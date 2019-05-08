// Code generated by 'go generate'; DO NOT EDIT.

package service

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modwtsapi32 = windows.NewLazySystemDLL("wtsapi32.dll")
	modadvapi32 = windows.NewLazySystemDLL("advapi32.dll")
	moduserenv  = windows.NewLazySystemDLL("userenv.dll")
	modkernel32 = windows.NewLazySystemDLL("kernel32.dll")

	procWTSQueryUserToken          = modwtsapi32.NewProc("WTSQueryUserToken")
	procWTSEnumerateSessionsW      = modwtsapi32.NewProc("WTSEnumerateSessionsW")
	procWTSFreeMemory              = modwtsapi32.NewProc("WTSFreeMemory")
	procGetSecurityInfo            = modadvapi32.NewProc("GetSecurityInfo")
	procAddAccessAllowedAce        = modadvapi32.NewProc("AddAccessAllowedAce")
	procSetSecurityDescriptorDacl  = modadvapi32.NewProc("SetSecurityDescriptorDacl")
	procSetSecurityDescriptorSacl  = modadvapi32.NewProc("SetSecurityDescriptorSacl")
	procGetAclInformation          = modadvapi32.NewProc("GetAclInformation")
	procGetAce                     = modadvapi32.NewProc("GetAce")
	procAddAce                     = modadvapi32.NewProc("AddAce")
	procInitializeAcl              = modadvapi32.NewProc("InitializeAcl")
	procMakeAbsoluteSD             = modadvapi32.NewProc("MakeAbsoluteSD")
	procMakeSelfRelativeSD         = modadvapi32.NewProc("MakeSelfRelativeSD")
	procCreateEnvironmentBlock     = moduserenv.NewProc("CreateEnvironmentBlock")
	procDestroyEnvironmentBlock    = moduserenv.NewProc("DestroyEnvironmentBlock")
	procNotifyServiceStatusChangeW = modadvapi32.NewProc("NotifyServiceStatusChangeW")
	procSleepEx                    = modkernel32.NewProc("SleepEx")
)

func wtsQueryUserToken(session uint32, token *windows.Token) (err error) {
	r1, _, e1 := syscall.Syscall(procWTSQueryUserToken.Addr(), 2, uintptr(session), uintptr(unsafe.Pointer(token)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func wtsEnumerateSessions(handle windows.Handle, reserved uint32, version uint32, sessions **wtsSessionInfo, count *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procWTSEnumerateSessionsW.Addr(), 5, uintptr(handle), uintptr(reserved), uintptr(version), uintptr(unsafe.Pointer(sessions)), uintptr(unsafe.Pointer(count)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func wtsFreeMemory(ptr uintptr) {
	syscall.Syscall(procWTSFreeMemory.Addr(), 1, uintptr(ptr), 0, 0)
	return
}

func getSecurityInfo(handle windows.Handle, objectType uint32, si uint32, owner *uintptr, group *uintptr, dacl *uintptr, sacl *uintptr, securityDescriptor *uintptr) (err error) {
	r1, _, e1 := syscall.Syscall9(procGetSecurityInfo.Addr(), 8, uintptr(handle), uintptr(objectType), uintptr(si), uintptr(unsafe.Pointer(owner)), uintptr(unsafe.Pointer(group)), uintptr(unsafe.Pointer(dacl)), uintptr(unsafe.Pointer(sacl)), uintptr(unsafe.Pointer(securityDescriptor)), 0)
	if r1 != 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func addAccessAllowedAce(acl *byte, aceRevision uint32, accessmask uint32, sid *windows.SID) (err error) {
	r1, _, e1 := syscall.Syscall6(procAddAccessAllowedAce.Addr(), 4, uintptr(unsafe.Pointer(acl)), uintptr(aceRevision), uintptr(accessmask), uintptr(unsafe.Pointer(sid)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func setSecurityDescriptorDacl(securityDescriptor *byte, daclPresent bool, dacl *byte, defaulted bool) (err error) {
	var _p0 uint32
	if daclPresent {
		_p0 = 1
	} else {
		_p0 = 0
	}
	var _p1 uint32
	if defaulted {
		_p1 = 1
	} else {
		_p1 = 0
	}
	r1, _, e1 := syscall.Syscall6(procSetSecurityDescriptorDacl.Addr(), 4, uintptr(unsafe.Pointer(securityDescriptor)), uintptr(_p0), uintptr(unsafe.Pointer(dacl)), uintptr(_p1), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func setSecurityDescriptorSacl(securityDescriptor *byte, saclPresent bool, sacl *byte, defaulted bool) (err error) {
	var _p0 uint32
	if saclPresent {
		_p0 = 1
	} else {
		_p0 = 0
	}
	var _p1 uint32
	if defaulted {
		_p1 = 1
	} else {
		_p1 = 0
	}
	r1, _, e1 := syscall.Syscall6(procSetSecurityDescriptorSacl.Addr(), 4, uintptr(unsafe.Pointer(securityDescriptor)), uintptr(_p0), uintptr(unsafe.Pointer(sacl)), uintptr(_p1), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func getAclInformation(acl *byte, info *ACL_SIZE_INFORMATION, len uint32, infoclass uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetAclInformation.Addr(), 4, uintptr(unsafe.Pointer(acl)), uintptr(unsafe.Pointer(info)), uintptr(len), uintptr(infoclass), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func getAce(acl *byte, index uint32, ace **ACE_HEADER) (err error) {
	r1, _, e1 := syscall.Syscall(procGetAce.Addr(), 3, uintptr(unsafe.Pointer(acl)), uintptr(index), uintptr(unsafe.Pointer(ace)))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func addAce(acl *byte, revision uint32, index uint32, ace *ACE_HEADER, lenAce uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procAddAce.Addr(), 5, uintptr(unsafe.Pointer(acl)), uintptr(revision), uintptr(index), uintptr(unsafe.Pointer(ace)), uintptr(lenAce), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func initializeAcl(acl *byte, len uint32, revision uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procInitializeAcl.Addr(), 3, uintptr(unsafe.Pointer(acl)), uintptr(len), uintptr(revision))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func makeAbsoluteSd(selfRelativeSecurityDescriptor uintptr, absoluteSecurityDescriptor *byte, absoluteSecurityDescriptorSize *uint32, dacl *byte, daclSize *uint32, sacl *byte, saclSize *uint32, owner *byte, ownerSize *uint32, primaryGroup *byte, primaryGroupSize *uint32) (err error) {
	r1, _, e1 := syscall.Syscall12(procMakeAbsoluteSD.Addr(), 11, uintptr(selfRelativeSecurityDescriptor), uintptr(unsafe.Pointer(absoluteSecurityDescriptor)), uintptr(unsafe.Pointer(absoluteSecurityDescriptorSize)), uintptr(unsafe.Pointer(dacl)), uintptr(unsafe.Pointer(daclSize)), uintptr(unsafe.Pointer(sacl)), uintptr(unsafe.Pointer(saclSize)), uintptr(unsafe.Pointer(owner)), uintptr(unsafe.Pointer(ownerSize)), uintptr(unsafe.Pointer(primaryGroup)), uintptr(unsafe.Pointer(primaryGroupSize)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func makeSelfRelativeSd(absoluteSecurityDescriptor *byte, relativeSecurityDescriptor *byte, relativeSecurityDescriptorSize *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procMakeSelfRelativeSD.Addr(), 3, uintptr(unsafe.Pointer(absoluteSecurityDescriptor)), uintptr(unsafe.Pointer(relativeSecurityDescriptor)), uintptr(unsafe.Pointer(relativeSecurityDescriptorSize)))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func createEnvironmentBlock(block *uintptr, token windows.Token, inheritExisting bool) (err error) {
	var _p0 uint32
	if inheritExisting {
		_p0 = 1
	} else {
		_p0 = 0
	}
	r1, _, e1 := syscall.Syscall(procCreateEnvironmentBlock.Addr(), 3, uintptr(unsafe.Pointer(block)), uintptr(token), uintptr(_p0))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func destroyEnvironmentBlock(block uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procDestroyEnvironmentBlock.Addr(), 1, uintptr(block), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func notifyServiceStatusChange(service windows.Handle, notifyMask uint32, notifyBuffer uintptr) (status uint32) {
	r0, _, _ := syscall.Syscall(procNotifyServiceStatusChangeW.Addr(), 3, uintptr(service), uintptr(notifyMask), uintptr(notifyBuffer))
	status = uint32(r0)
	return
}

func sleepEx(milliseconds uint32, alertable bool) (ret uint32, err error) {
	var _p0 uint32
	if alertable {
		_p0 = 1
	} else {
		_p0 = 0
	}
	r0, _, e1 := syscall.Syscall(procSleepEx.Addr(), 2, uintptr(milliseconds), uintptr(_p0), 0)
	ret = uint32(r0)
	if ret == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}