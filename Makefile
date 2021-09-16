GOFLAGS := -tags load_wgnt_from_rsrc -ldflags="-H windowsgui -s -w" -v -trimpath
export GOOS := windows
export PATH := $(CURDIR)/.deps/go/bin:$(PATH)

VERSION := $(shell sed -n 's/^\s*Number\s*=\s*"\([0-9.]\+\)"$$/\1/p' version/version.go)
empty :=
space := $(empty) $(empty)
comma := ,
RCFLAGS := -DWIREGUARD_VERSION_ARRAY=$(subst $(space),$(comma),$(wordlist 1,4,$(subst .,$(space),$(VERSION)) 0 0 0 0)) -DWIREGUARD_VERSION_STR=$(VERSION) -O coff -c 65001

rwildcard=$(foreach d,$(filter-out .deps,$(wildcard $1*)),$(call rwildcard,$d/,$2) $(filter $(subst *,%,$2),$d))
SOURCE_FILES := $(call rwildcard,,*.go) .deps/go/prepared go.mod go.sum
RESOURCE_FILES := resources.rc version/version.go manifest.xml $(patsubst %.svg,%.ico,$(wildcard ui/icon/*.svg)) .deps/wireguard-nt/prepared

DEPLOYMENT_HOST ?= winvm
DEPLOYMENT_PATH ?= Desktop

all: amd64/wireguard.exe x86/wireguard.exe arm64/wireguard.exe arm/wireguard.exe

define download =
.distfiles/$(1):
	mkdir -p .distfiles
	if ! curl -L#o $$@.unverified $(2); then rm -f $$@.unverified; exit 1; fi
	if ! echo "$(3)  $$@.unverified" | sha256sum -c; then rm -f $$@.unverified; exit 1; fi
	if ! mv $$@.unverified $$@; then rm -f $$@.unverified; exit 1; fi
endef

$(eval $(call download,go.tar.gz,https://golang.org/dl/go1.17.1.linux-amd64.tar.gz,dab7d9c34361dc21ec237d584590d72500652e7c909bf082758fb63064fca0ef))
$(eval $(call download,wireguard-nt.zip,https://download.wireguard.com/wireguard-nt/wireguard-nt-0.8.zip,a5841a345723734b00df432bdf1265ee0ffdea0b55ccf26f361217f4fe95f5a2))

.deps/go/prepared: .distfiles/go.tar.gz $(wildcard go-patches/*.patch)
	mkdir -p .deps
	rm -rf .deps/go
	bsdtar -C .deps -xf .distfiles/go.tar.gz
	chmod -R +w .deps/go
	cat $(filter %.patch,$^) | patch -f -N -r- -p1 -d .deps/go
	touch $@

.deps/wireguard-nt/prepared: .distfiles/wireguard-nt.zip
	mkdir -p .deps
	rm -rf .deps/wireguard-nt
	bsdtar -C .deps -xf .distfiles/wireguard-nt.zip
	touch $@

%.ico: %.svg
	convert -background none $< -define icon:auto-resize="256,192,128,96,64,48,40,32,24,20,16" -compress zip $@

resources_amd64.syso: $(RESOURCE_FILES)
	x86_64-w64-mingw32-windres $(RCFLAGS) -I .deps/wireguard-nt/bin/amd64 -i $< -o $@

resources_386.syso: $(RESOURCE_FILES)
	i686-w64-mingw32-windres $(RCFLAGS) -I .deps/wireguard-nt/bin/x86 -i $< -o $@

resources_arm.syso: $(RESOURCE_FILES)
	armv7-w64-mingw32-windres $(RCFLAGS) -I .deps/wireguard-nt/bin/arm -i $< -o $@

resources_arm64.syso: $(RESOURCE_FILES)
	aarch64-w64-mingw32-windres $(RCFLAGS) -I .deps/wireguard-nt/bin/arm64 -i $< -o $@

amd64/wireguard.exe: export GOARCH := amd64
amd64/wireguard.exe: resources_amd64.syso $(SOURCE_FILES)
	go build $(GOFLAGS) -o $@

x86/wireguard.exe: export GOARCH := 386
x86/wireguard.exe: resources_386.syso $(SOURCE_FILES)
	go build $(GOFLAGS) -o $@

arm/wireguard.exe: export GOARCH := arm
arm/wireguard.exe: export GOARM := 7
arm/wireguard.exe: resources_arm.syso $(SOURCE_FILES)
	go build $(GOFLAGS) -o $@

arm64/wireguard.exe: export GOARCH := arm64
arm64/wireguard.exe: resources_arm64.syso $(SOURCE_FILES)
	go build $(GOFLAGS) -o $@

remaster: export GOARCH := amd64
remaster: export GOPROXY := direct
remaster: .deps/go/prepared
	rm -f go.sum go.mod
	cp go.mod.master go.mod
	go get -d

fmt: export GOARCH := amd64
fmt: .deps/go/prepared
	go fmt ./...

generate: export GOOS :=
generate: .deps/go/prepared
	go generate -mod=mod ./...

crowdin:
	find locales -maxdepth 1 -mindepth 1 -type d \! -name en -exec rm -rf {} +
	curl -Lo - https://crowdin.com/backend/download/project/wireguard.zip | bsdtar -C locales -x -f - --strip-components 2 wireguard-windows
	find locales -name messages.gotext.json -exec bash -c '[[ $$(jq ".messages | length" {}) -ne 0 ]] || rm -rf "$$(dirname {})"' \;
	@$(MAKE) --no-print-directory generate

deploy: amd64/wireguard.exe
	-ssh $(DEPLOYMENT_HOST) -- 'taskkill /im wireguard.exe /f'
	scp $< $(DEPLOYMENT_HOST):$(DEPLOYMENT_PATH)

clean:
	rm -rf *.syso ui/icon/*.ico x86/ amd64/ arm/ arm64/ .deps

distclean: clean
	rm -rf .distfiles

.PHONY: deploy clean distclean fmt remaster generate all
