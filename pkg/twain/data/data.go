// Package data wraps the data constants (prefixed with DF_, DG_, and DAT_) from
// the TWAIN header file.
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

// ArgumentType is an alias for Data Argument Types. They can represent the
// DG_CONTROL, DG_IMAGE, or DG_AUDIO Data Groups.
type ArgumentType twain.UInt16

const (
	// ControlCapability is an alias for DAT_CAPABILITY.
	ControlCapability ArgumentType = 0x0001

	// ControlEvent is an alias for DAT_EVENT.
	ControlEvent ArgumentType = 0x0002

	// ControlIdentity is an alias for DAT_IDENTITY.
	ControlIdentity ArgumentType = 0x0003

	// ControlParent is an alias for DAT_PARENT.
	ControlParent ArgumentType = 0x0004

	// ControlPendingTransfers is an alias for DAT_PENDINGXFERS.
	ControlPendingTransfers ArgumentType = 0x0005

	// ControlSetupMemTransfer is an alias for DAT_SETUPMEMXFER.
	ControlSetupMemTransfer ArgumentType = 0x0006

	// ControlSetupFileTransfer is an alias for DAT_SETUPFILEXFER.
	ControlSetupFileTransfer ArgumentType = 0x0007

	// ControlStatus is an alias for DAT_STATUS.
	ControlStatus ArgumentType = 0x0008

	// ControlUserInterface is an alias for DAT_USERINTERFACE.
	ControlUserInterface ArgumentType = 0x0009

	// ControlTransferGroup is an alias for DAT_XFERGROUP.
	ControlTransferGroup ArgumentType = 0x000a

	// ControlCustomDSData is an alias for DAT_CUSTOMDSDATA.
	ControlCustomDSData ArgumentType = 0x000c

	// ControlDeviceEvent is an alias for DAT_DEVICEEVENT.
	ControlDeviceEvent ArgumentType = 0x000d

	// ControlFileSystem is an alias for DAT_FILESYSTEM.
	ControlFileSystem ArgumentType = 0x000e

	// ControlPassThru is an alias for DAT_PASSTHRU.
	ControlPassThru ArgumentType = 0x000f

	// ControlCallBack is an alias for DAT_CALLBACK.
	ControlCallBack ArgumentType = 0x0010

	// ControlStatusUTF8 is an alias for DAT_STATUSUTF8.
	ControlStatusUTF8 ArgumentType = 0x0011

	// ControlCallBack2 is an alias for DAT_CALLBACK2.
	ControlCallBack2 ArgumentType = 0x0012

	// ControlMetrics is an alias for DAT_METRICS.
	ControlMetrics ArgumentType = 0x0013

	// ControlTwainDirect is an alias for DAT_TWAINDIRECT.
	ControlTwainDirect ArgumentType = 0x0014
)

const (
	// ImageImageInfo is an alias for DAT_IMAGEINFO.
	ImageImageInfo ArgumentType = 0x0101

	// ImageImageLayout is an alias for DAT_IMAGELAYOUT.
	ImageImageLayout ArgumentType = 0x0102

	// ImageImageMemTransfer is an alias for DAT_IMAGEMEMXFER.
	ImageImageMemTransfer ArgumentType = 0x0103

	// ImageImageNativeTransfer is an alias for DAT_IMAGENATIVEXFER.
	ImageImageNativeTransfer ArgumentType = 0x0104

	// ImageImageFileTransfer is an alias for DAT_IMAGEFILEXFER.
	ImageImageFileTransfer ArgumentType = 0x0105

	// ImageCIEColor is an alias for DAT_CIECOLOR.
	ImageCIEColor ArgumentType = 0x0106

	// ImageGrayResponse is an alias for DAT_GRAYRESPONSE.
	ImageGrayResponse ArgumentType = 0x0107

	// ImageRGBResponse is an alias for DAT_RGBRESPONSE.
	ImageRGBResponse ArgumentType = 0x0108

	// ImageJPEGCompression is an alias for DAT_JPEGCOMPRESSION.
	ImageJPEGCompression ArgumentType = 0x0109

	// ImagePalette8 is an alias for DAT_PALETTE8.
	ImagePalette8 ArgumentType = 0x010a

	// ImageExtImageInfo is an alias for DAT_EXTIMAGEINFO.
	ImageExtImageInfo ArgumentType = 0x010b

	// ImageFilter is an alias for DAT_FILTER.
	ImageFilter ArgumentType = 0x010c
)

const (
	// AudioFileTransfer is an alias for DAT_AUDIOFILEXFER.
	AudioFileTransfer ArgumentType = 0x0201

	// AudioInfo is an alias for DAT_AUDIOINFO.
	AudioInfo ArgumentType = 0x0202

	// AudioNativeTransfer is an alias for DAT_AUDIONATIVEXFER.
	AudioNativeTransfer ArgumentType = 0x0203
)
