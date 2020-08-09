package data

import "github.com/mikerourke/go-twain/pkg/twain"

// Group represents the Data Group.
type Group twain.UInt32

const (
	// GroupControl is an alias for DG_CONTROL.
	GroupControl Group = 0x0001

	// GroupImage is an alias for DG_IMAGE.
	GroupImage Group = 0x0002

	// GroupAudio is an alias for DG_AUDIO.
	GroupAudio Group = 0x0004
)

// Functionality represents the functions supported.
// Supported Functionality constants must be powers of 2 as they are used as
// bitflags when Application asks DSM to present a list of DSs.
// To support backward capability the App and DS will not use the fields
type Functionality twain.UInt32

const (
	// FunctionalityDSM2 is an alias for DF_DSM2.
	FunctionalityDSM2 = 0x10000000

	// FunctionalityApp2 is an alias for DF_APP2.
	FunctionalityApp2 = 0x20000000

	// FunctionalityDS2 is an alias for DF_DS2.
	FunctionalityDS2 = 0x40000000

	// FunctionalityMask is an alias for DG_MASK.
	FunctionalityMask = 0xFFFF
)

// ControlArgumentType is an alias for Data Argument Types for the DG_CONTROL
// Data Group.
type ControlArgumentType twain.UInt16

const (
	// ControlArgumentTypeCapability is an alias for DAT_CAPABILITY.
	ControlArgumentTypeCapability ControlArgumentType = 0x0001

	// ControlArgumentTypeEvent is an alias for DAT_EVENT.
	ControlArgumentTypeEvent ControlArgumentType = 0x0002

	// ControlArgumentTypeIdentity is an alias for DAT_IDENTITY.
	ControlArgumentTypeIdentity ControlArgumentType = 0x0003

	// ControlArgumentTypeParent is an alias for DAT_PARENT.
	ControlArgumentTypeParent ControlArgumentType = 0x0004

	// ControlArgumentTypePendingTransfers is an alias for DAT_PENDINGXFERS.
	ControlArgumentTypePendingTransfers ControlArgumentType = 0x0005

	// ControlArgumentTypeSetupMemTransfer is an alias for DAT_SETUPMEMXFER.
	ControlArgumentTypeSetupMemTransfer ControlArgumentType = 0x0006

	// ControlArgumentTypeSetupFileTransfer is an alias for DAT_SETUPFILEXFER.
	ControlArgumentTypeSetupFileTransfer ControlArgumentType = 0x0007

	// ControlArgumentTypeStatus is an alias for DAT_STATUS.
	ControlArgumentTypeStatus ControlArgumentType = 0x0008

	// ControlArgumentTypeUserInterface is an alias for DAT_USERINTERFACE.
	ControlArgumentTypeUserInterface ControlArgumentType = 0x0009

	// ControlArgumentTypeTransferGroup is an alias for DAT_XFERGROUP.
	ControlArgumentTypeTransferGroup ControlArgumentType = 0x000a

	// ControlArgumentTypeCustomDSData is an alias for DAT_CUSTOMDSDATA.
	ControlArgumentTypeCustomDSData ControlArgumentType = 0x000c

	// ControlArgumentTypeDeviceEvent is an alias for DAT_DEVICEEVENT.
	ControlArgumentTypeDeviceEvent ControlArgumentType = 0x000d

	// ControlArgumentTypeFileSystem is an alias for DAT_FILESYSTEM.
	ControlArgumentTypeFileSystem ControlArgumentType = 0x000e

	// ControlArgumentTypePassThru is an alias for DAT_PASSTHRU.
	ControlArgumentTypePassThru ControlArgumentType = 0x000f

	// ControlArgumentTypeCallBack is an alias for DAT_CALLBACK.
	ControlArgumentTypeCallBack ControlArgumentType = 0x0010

	// ControlArgumentTypeStatusUTF8 is an alias for DAT_STATUSUTF8.
	ControlArgumentTypeStatusUTF8 ControlArgumentType = 0x0011

	// ControlArgumentTypeCallBack2 is an alias for DAT_CALLBACK2.
	ControlArgumentTypeCallBack2 ControlArgumentType = 0x0012

	// ControlArgumentTypeMetrics is an alias for DAT_METRICS.
	ControlArgumentTypeMetrics ControlArgumentType = 0x0013

	// ControlArgumentTypeTwainDirect is an alias for DAT_TWAINDIRECT.
	ControlArgumentTypeTwainDirect ControlArgumentType = 0x0014
)

// ImageArgumentType is an alias for Data Argument Types for the DG_IMAGE
// Data Group.
type ImageArgumentType twain.UInt16

const (
	// ImageArgumentTypeImageInfo is an alias for DAT_IMAGEINFO.
	ImageArgumentTypeImageInfo ImageArgumentType = 0x0101

	// ImageArgumentTypeImageLayout is an alias for DAT_IMAGELAYOUT.
	ImageArgumentTypeImageLayout ImageArgumentType = 0x0102

	// ImageArgumentTypeImageMemTransfer is an alias for DAT_IMAGEMEMXFER.
	ImageArgumentTypeImageMemTransfer ImageArgumentType = 0x0103

	// ImageArgumentTypeImageNativeTransfer is an alias for DAT_IMAGENATIVEXFER.
	ImageArgumentTypeImageNativeTransfer ImageArgumentType = 0x0104

	// ImageArgumentTypeImageFileTransfer is an alias for DAT_IMAGEFILEXFER.
	ImageArgumentTypeImageFileTransfer ImageArgumentType = 0x0105

	// ImageArgumentTypeCIEColor is an alias for DAT_CIECOLOR.
	ImageArgumentTypeCIEColor ImageArgumentType = 0x0106

	// ImageArgumentTypeGrayResponse is an alias for DAT_GRAYRESPONSE.
	ImageArgumentTypeGrayResponse ImageArgumentType = 0x0107

	// ImageArgumentTypeRGBResponse is an alias for DAT_RGBRESPONSE.
	ImageArgumentTypeRGBResponse ImageArgumentType = 0x0108

	// ImageArgumentTypeJPEGCompression is an alias for DAT_JPEGCOMPRESSION.
	ImageArgumentTypeJPEGCompression ImageArgumentType = 0x0109

	// ImageArgumentTypePalette8 is an alias for DAT_PALETTE8.
	ImageArgumentTypePalette8 ImageArgumentType = 0x010a

	// ImageArgumentTypeExtImageInfo is an alias for DAT_EXTIMAGEINFO.
	ImageArgumentTypeExtImageInfo ImageArgumentType = 0x010b

	// ImageArgumentTypeFilter is an alias for DAT_FILTER.
	ImageArgumentTypeFilter ImageArgumentType = 0x010c
)

// AudioArgumentType is an alias for Data Argument Types for the DG_AUDIO
// Data Group.
type AudioArgumentType twain.UInt16

const (
	// AudioArgumentTypeFileTransfer is an alias for DAT_AUDIOFILEXFER.
	AudioArgumentTypeFileTransfer AudioArgumentType = 0x0201

	// AudioArgumentTypeInfo is an alias for DAT_AUDIOINFO.
	AudioArgumentTypeInfo AudioArgumentType = 0x0202

	// AudioArgumentTypeNativeTransfer is an alias for DAT_AUDIONATIVEXFER.
	AudioArgumentTypeNativeTransfer AudioArgumentType = 0x0203
)
