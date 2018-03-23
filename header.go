// Copyright 2012 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package win

const (
	HDS_NOSIZING = 0x0800
)

type HDITEM struct {
	Mask       uint32
	Cxy        int32
	PszText    *uint16
	Hbm        HBITMAP
	CchTextMax int32
	Fmt        int32
	LParam     uintptr
	IImage     int32
	IOrder     int32
	Type       uint32
	PvFilter   uintptr
}

const (
	HDI_WIDTH      = 0x0001
	HDI_HEIGHT     = HDI_WIDTH
	HDI_TEXT       = 0x0002
	HDI_FORMAT     = 0x0004
	HDI_LPARAM     = 0x0008
	HDI_BITMAP     = 0x0010
	HDI_IMAGE      = 0x0020
	HDI_DI_SETITEM = 0x0040
	HDI_ORDER      = 0x0080
	HDI_FILTER     = 0x0100
	HDI_STATE      = 0x0200
)

const (
	HDF_LEFT            = 0x0000
	HDF_RIGHT           = 0x0001
	HDF_CENTER          = 0x0002
	HDF_JUSTIFYMASK     = 0x0003
	HDF_RTLREADING      = 0x0004
	HDF_CHECKBOX        = 0x0040
	HDF_CHECKED         = 0x0080
	HDF_FIXEDWIDTH      = 0x0100
	HDF_SORTDOWN        = 0x0200
	HDF_SORTUP          = 0x0400
	HDF_IMAGE           = 0x0800
	HDF_BITMAP_ON_RIGHT = 0x1000
	HDF_BITMAP          = 0x2000
	HDF_STRING          = 0x4000
	HDF_OWNERDRAW       = 0x8000
	HDF_SPLITBUTTON     = 0x1000000
)

const (
	HDIS_FOCUSED = 0x00000001
)

const (
	HDM_FIRST                  = 0x1200
	HDM_GETITEMCOUNT           = HDM_FIRST + 0
	HDM_DELETEITEM             = HDM_FIRST + 2
	HDM_LAYOUT                 = HDM_FIRST + 5
	HDM_HITTEST                = HDM_FIRST + 6
	HDM_GETITEMRECT            = HDM_FIRST + 7
	HDM_SETIMAGELIST           = HDM_FIRST + 8
	HDM_GETIMAGELIST           = HDM_FIRST + 9
	HDM_INSERTITEM             = HDM_FIRST + 10
	HDM_GETITEM                = HDM_FIRST + 11
	HDM_SETITEM                = HDM_FIRST + 12
	HDM_ORDERTOINDEX           = HDM_FIRST + 15
	HDM_CREATEDRAGIMAGE        = HDM_FIRST + 16
	HDM_GETORDERARRAY          = HDM_FIRST + 17
	HDM_SETORDERARRAY          = HDM_FIRST + 18
	HDM_SETHOTDIVIDER          = HDM_FIRST + 19
	HDM_SETBITMAPMARGIN        = HDM_FIRST + 20
	HDM_GETBITMAPMARGIN        = HDM_FIRST + 21
	HDM_SETFILTERCHANGETIMEOUT = HDM_FIRST + 22
	HDM_EDITFILTER             = HDM_FIRST + 23
	HDM_CLEARFILTER            = HDM_FIRST + 24
	HDM_GETITEMDROPDOWNRECT    = HDM_FIRST + 25
	HDM_GETOVERFLOWRECT        = HDM_FIRST + 26
	HDM_GETFOCUSEDITEM         = HDM_FIRST + 27
	HDM_SETFOCUSEDITEM         = HDM_FIRST + 28
	HDM_SETUNICODEFORMAT       = CCM_SETUNICODEFORMAT
	HDM_GETUNICODEFORMAT       = CCM_GETUNICODEFORMAT
)

const (
	HHT_NOWHERE         = 0x0001
	HHT_ONHEADER        = 0x0002
	HHT_ONDIVIDER       = 0x0004
	HHT_ONDIVOPEN       = 0x0008
	HHT_ONFILTER        = 0x0010
	HHT_ONFILTERBUTTON  = 0x0020
	HHT_ABOVE           = 0x0100
	HHT_BELOW           = 0x0200
	HHT_TORIGHT         = 0x0400
	HHT_TOLEFT          = 0x0800
	HHT_ONITEMSTATEICON = 0x1000
	HHT_ONDROPDOWN      = 0x2000
	HHT_ONOVERFLOW      = 0x4000
)

const (
	HDN_FIRST              = ^uint32(300)
	HDN_BEGINDRAG          = HDN_FIRST - 10
	HDN_ENDDRAG            = HDN_FIRST - 11
	HDN_FILTERCHANGE       = HDN_FIRST - 12
	HDN_FILTERBTNCLICK     = HDN_FIRST - 13
	HDN_BEGINFILTEREDIT    = HDN_FIRST - 14
	HDN_ENDFILTEREDIT      = HDN_FIRST - 15
	HDN_ITEMSTATEICONCLICK = HDN_FIRST - 16
	HDN_ITEMKEYDOWN        = HDN_FIRST - 17
	HDN_DROPDOWN           = HDN_FIRST - 18
	HDN_OVERFLOWCLICK      = HDN_FIRST - 19
	HDN_ITEMCHANGING       = HDN_FIRST - 20
	HDN_ITEMCHANGED        = HDN_FIRST - 21
	HDN_ITEMCLICK          = HDN_FIRST - 22
	HDN_ITEMDBLCLICK       = HDN_FIRST - 23
	HDN_DIVIDERDBLCLICK    = HDN_FIRST - 25
	HDN_BEGINTRACK         = HDN_FIRST - 26
	HDN_ENDTRACK           = HDN_FIRST - 27
	HDN_TRACK              = HDN_FIRST - 28
	HDN_GETDISPINFO        = HDN_FIRST - 29
)
