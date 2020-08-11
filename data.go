package twain

// Group represents the Data Group.
type DataGroup uint32

const (
	// DataGroupControl is an alias for DG_CONTROL.
	DataGroupControl DataGroup = 0x0001

	// DataGroupImage is an alias for DG_IMAGE.
	DataGroupImage DataGroup = 0x0002

	// DataGroupAudio is an alias for DG_AUDIO.
	DataGroupAudio DataGroup = 0x0004
)

// SupportedFunctionality represents the functions supported.
// Supported Functionality constants must be powers of 2 as they are used as
// bitflags when Application asks DSM to present a list of DSs.
// To support backward capability the App and DS will not use the fields
type SupportedFunctionality uint32

const (
	// SupportedFunctionalityDSM2 is an alias for DF_DSM2.
	SupportedFunctionalityDSM2 = 0x10000000

	// SupportedFunctionalityApp2 is an alias for DF_APP2.
	SupportedFunctionalityApp2 = 0x20000000

	// SupportedFunctionalityDS2 is an alias for DF_DS2.
	SupportedFunctionalityDS2 = 0x40000000

	// SupportedFunctionalityMask is an alias for DG_MASK.
	SupportedFunctionalityMask = 0xFFFF
)

// DataArgumentType is an alias for Data Argument Types. They can represent the
// DG_CONTROL, DG_IMAGE, or DG_AUDIO Data Groups.
type DataArgumentType uint16

const (
	// DataArgumentCapability is an alias for DAT_CAPABILITY.
	DataArgumentCapability DataArgumentType = 0x0001

	// DataArgumentEvent is an alias for DAT_EVENT.
	DataArgumentEvent DataArgumentType = 0x0002

	// DataArgumentIdentity is an alias for DAT_IDENTITY.
	DataArgumentIdentity DataArgumentType = 0x0003

	// DataArgumentParent is an alias for DAT_PARENT.
	DataArgumentParent DataArgumentType = 0x0004

	// DataArgumentPendingTransfers is an alias for DAT_PENDINGXFERS.
	DataArgumentPendingTransfers DataArgumentType = 0x0005

	// DataArgumentSetupMemTransfer is an alias for DAT_SETUPMEMXFER.
	DataArgumentSetupMemTransfer DataArgumentType = 0x0006

	// DataArgumentSetupFileTransfer is an alias for DAT_SETUPFILEXFER.
	DataArgumentSetupFileTransfer DataArgumentType = 0x0007

	// DataArgumentStatus is an alias for DAT_STATUS.
	DataArgumentStatus DataArgumentType = 0x0008

	// DataArgumentUserInterface is an alias for DAT_USERINTERFACE.
	DataArgumentUserInterface DataArgumentType = 0x0009

	// DataArgumentTransferGroup is an alias for DAT_XFERGROUP.
	DataArgumentTransferGroup DataArgumentType = 0x000a

	// DataArgumentCustomDSData is an alias for DAT_CUSTOMDSDATA.
	DataArgumentCustomDSData DataArgumentType = 0x000c

	// DataArgumentDeviceEvent is an alias for DAT_DEVICEEVENT.
	DataArgumentDeviceEvent DataArgumentType = 0x000d

	// DataArgumentFileSystem is an alias for DAT_FILESYSTEM.
	DataArgumentFileSystem DataArgumentType = 0x000e

	// DataArgumentPassThru is an alias for DAT_PASSTHRU.
	DataArgumentPassThru DataArgumentType = 0x000f

	// DataArgumentCallback is an alias for DAT_CALLBACK.
	DataArgumentCallback DataArgumentType = 0x0010

	// DataArgumentStatusUTF8 is an alias for DAT_STATUSUTF8.
	DataArgumentStatusUTF8 DataArgumentType = 0x0011

	// DataArgumentCallback2 is an alias for DAT_CALLBACK2.
	DataArgumentCallback2 DataArgumentType = 0x0012

	// DataArgumentMetrics is an alias for DAT_METRICS.
	DataArgumentMetrics DataArgumentType = 0x0013

	// DataArgumentTwainDirect is an alias for DAT_TWAINDIRECT.
	DataArgumentTwainDirect DataArgumentType = 0x0014
)

const (
	// DataArgumentImageInfo is an alias for DAT_IMAGEINFO.
	DataArgumentImageInfo DataArgumentType = 0x0101

	// DataArgumentImageLayout is an alias for DAT_IMAGELAYOUT.
	DataArgumentImageLayout DataArgumentType = 0x0102

	// DataArgumentImageMemTransfer is an alias for DAT_IMAGEMEMXFER.
	DataArgumentImageMemTransfer DataArgumentType = 0x0103

	// DataArgumentImageNativeTransfer is an alias for DAT_IMAGENATIVEXFER.
	DataArgumentImageNativeTransfer DataArgumentType = 0x0104

	// DataArgumentImageFileTransfer is an alias for DAT_IMAGEFILEXFER.
	DataArgumentImageFileTransfer DataArgumentType = 0x0105

	// DataArgumentCIEColor is an alias for DAT_CIECOLOR.
	DataArgumentCIEColor DataArgumentType = 0x0106

	// DataArgumentGrayResponse is an alias for DAT_GRAYRESPONSE.
	DataArgumentGrayResponse DataArgumentType = 0x0107

	// DataArgumentRGBResponse is an alias for DAT_RGBRESPONSE.
	DataArgumentRGBResponse DataArgumentType = 0x0108

	// DataArgumentJPEGCompression is an alias for DAT_JPEGCOMPRESSION.
	DataArgumentJPEGCompression DataArgumentType = 0x0109

	// DataArgumentPalette8 is an alias for DAT_PALETTE8.
	DataArgumentPalette8 DataArgumentType = 0x010a

	// DataArgumentExtImageInfo is an alias for DAT_EXTIMAGEINFO.
	DataArgumentExtImageInfo DataArgumentType = 0x010b

	// DataArgumentFilter is an alias for DAT_FILTER.
	DataArgumentFilter DataArgumentType = 0x010c
)

const (
	// DataArgumentAudioFileTransfer is an alias for DAT_AUDIOFILEXFER.
	DataArgumentAudioFileTransfer DataArgumentType = 0x0201

	// DataArgumentAudioInfo is an alias for DAT_AUDIOINFO.
	DataArgumentAudioInfo DataArgumentType = 0x0202

	// DataArgumentAudioNativeTransfer is an alias for DAT_AUDIONATIVEXFER.
	DataArgumentAudioNativeTransfer DataArgumentType = 0x0203
)
