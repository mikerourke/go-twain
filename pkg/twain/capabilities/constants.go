package capabilities

import "github.com/mikerourke/go-twain/pkg/twain"

// AlarmType is an alias for the CAP_ALARMS constants.
type AlarmType twain.UInt16

const (
	// AlarmTypeAlarm is an alias for TWAL_ALARM.
	AlarmTypeAlarm AlarmType = iota

	// AlarmTypeFeedError is an alias for TWAL_FEEDERERROR.
	AlarmTypeFeedError

	// AlarmTypeFeederWarning is an alias for TWAL_FEEDERWARNING.
	AlarmTypeFeederWarning

	// AlarmTypeBarcode is an alias for TWAL_BARCODE.
	AlarmTypeBarcode

	// AlarmTypeDoubleFeed is an alias for TWAL_DOUBLEFEED.
	AlarmTypeDoubleFeed

	// AlarmTypeJam is an alias for TWAL_JAM.
	AlarmTypeJam

	// AlarmTypePatchCode is an alias for TWAL_PATCHCODE.
	AlarmTypePatchCode

	// AlarmTypePower is an alias for TWAL_POWER.
	AlarmTypePower

	// AlarmTypeSkew is an alias for TWAL_SKEW.
	AlarmTypeSkew
)

// AutoSizeType is an alias for the ICAP_AUTOSIZE values.
type AutoSizeType twain.UInt16

const (
	// AutoSizeTypeNone is an alias for TWAS_NONE.
	AutoSizeTypeNone AutoSizeType = iota

	// AutoSizeTypeAuto is an alias for TWAS_AUTO.
	AutoSizeTypeAuto

	// AutoSizeTypeCurrent is an alias for TWAS_CURRENT.
	AutoSizeTypeCurrent
)

// BarcodeRotation is an alias for the TWEI_BARCODEROTATION values.
type BarcodeRotation twain.UInt16

const (
	// BarcodeRotation0 is an alias for TWBCOR_ROT0.
	BarcodeRotation0 BarcodeRotation = iota

	// BarcodeRotation90 is an alias for TWBCOR_ROT90.
	BarcodeRotation90

	// BarcodeRotation180 is an alias for TWBCOR_ROT180.
	BarcodeRotation180

	// BarcodeRotation270 is an alias for TWBCOR_ROT270.
	BarcodeRotation270

	// BarcodeRotationX is an alias for TWBCOR_ROTX.
	BarcodeRotationX
)

// BarcodeSearchMode is an alias for ICAP_BARCODESEARCHMODE values.
type BarcodeSearchMode twain.UInt16

const (
	// BarcodeSearchModeHoriz is an alias for TWBD_HORZ.
	BarcodeSearchModeHoriz BarcodeSearchMode = iota

	// BarcodeSearchModeVert is an alias for TWBD_VERT.
	BarcodeSearchModeVert

	// BarcodeSearchModeHorizVert is an alias for TWBD_HORZVERT.
	BarcodeSearchModeHorizVert

	// BarcodeSearchModeVertHoriz is an alias for TWBD_VERTHORZ.
	BarcodeSearchModeVertHoriz
)

// BitOrder is an alias for ICAP_BITORDER values.
type BitOrder twain.UInt16

const (
	// BitOrderLSBFirst is an alias for TWBO_LSBFIRST.
	BitOrderLSBFirst BitOrder = iota

	// BitOrderMSBFirst is an alias for TWBO_MSBFIRST.
	BitOrderMSBFirst
)

// AutoDiscardBlankPages is an alias for ICAP_AUTODISCARDBLANKPAGES values.
type AutoDiscardBlankPages twain.UInt16

const (
	// AutoDiscardBlankPagesDisable is an alias for TWBP_DISABLE.
	AutoDiscardBlankPagesDisable AutoDiscardBlankPages = -2

	// AutoDiscardBlankPagesAuto is an alias for TWBP_AUTO.
	AutoDiscardBlankPagesAuto AutoDiscardBlankPages = -1
)

// BitDepthReduction is an alias for ICAP_BITDEPTHREDUCTION values.
type BitDepthReduction twain.UInt16

const (
	// BitDepthReductionThreshold is an alias for TWBR_THRESHOLD.
	BitDepthReductionThreshold BitDepthReduction = iota

	// BitDepthReductionHalfTone is an alias for TWBR_HALFTONE.
	BitDepthReductionHalfTone

	// BitDepthReductionCustomHalfTone is an alias for TWBR_CUSTHALFTONE.
	BitDepthReductionCustomHalfTone

	// BitDepthReductionDiffusion is an alias for TWBR_DIFFUSION.
	BitDepthReductionDiffusion

	// BitDepthReductionDynamicThreshold is an alias for TWBR_DYNAMICTHRESHOLD.
	BitDepthReductionDynamicThreshold
)

// SupportedBarcodeType is an alias for ICAP_SUPPORTEDBARCODETYPES and
// TWEI_BARCODETYPE values.
type SupportedBarcodeType twain.UInt16

const (
	// SupportedBarcodeType3OF9 is an alias for TWBT_3OF9.
	SupportedBarcodeType3OF9 SupportedBarcodeType = iota

	// SupportedBarcodeType2OF5Interleaved is an alias for TWBT_2OF5INTERLEAVED.
	SupportedBarcodeType2OF5Interleaved

	// SupportedBarcodeType2OF5NonInterleaved is an alias for TWBT_2OF5NONINTERLEAVED.
	SupportedBarcodeType2OF5NonInterleaved

	// SupportedBarcodeTypeCode93 is an alias for TWBT_CODE93.
	SupportedBarcodeTypeCode93

	// SupportedBarcodeTypeCode128 is an alias for TWBT_CODE128.
	SupportedBarcodeTypeCode128

	// SupportedBarcodeTypeUCC128 is an alias for TWBT_UCC128.
	SupportedBarcodeTypeUCC128

	// SupportedBarcodeTypeCodaBar is an alias for TWBT_CODABAR.
	SupportedBarcodeTypeCodaBar

	// SupportedBarcodeTypeUPCA is an alias for TWBT_UPCA.
	SupportedBarcodeTypeUPCA

	// SupportedBarcodeTypeUPCE is an alias for TWBT_UPCE.
	SupportedBarcodeTypeUPCE

	// SupportedBarcodeTypeEAN8 is an alias for TWBT_EAN8.
	SupportedBarcodeTypeEAN8

	// SupportedBarcodeTypeEAN13 is an alias for TWBT_EAN13.
	SupportedBarcodeTypeEAN13

	// SupportedBarcodeTypePostNet is an alias for TWBT_POSTNET.
	SupportedBarcodeTypePostNet

	// SupportedBarcodeTypePDF417 is an alias for TWBT_PDF417.
	SupportedBarcodeTypePDF417

	// SupportedBarcodeType2OF5Industrial is an alias for TWBT_2OF5INDUSTRIAL.
	SupportedBarcodeType2OF5Industrial

	// SupportedBarcodeType2OF5Matrix is an alias for TWBT_2OF5MATRIX.
	SupportedBarcodeType2OF5Matrix

	// SupportedBarcodeType2OF5DataLogic is an alias for TWBT_2OF5DATALOGIC.
	SupportedBarcodeType2OF5DataLogic

	// SupportedBarcodeType2OF5IATA is an alias for TWBT_2OF5IATA.
	SupportedBarcodeType2OF5IATA

	// SupportedBarcodeType3OF9FullASCII is an alias for TWBT_3OF9FULLASCII.
	SupportedBarcodeType3OF9FullASCII

	// SupportedBarcodeTypeCodaBarWithStartStop is an alias for TWBT_CODABARWITHSTARTSTOP.
	SupportedBarcodeTypeCodaBarWithStartStop

	// SupportedBarcodeTypeMAXICode is an alias for TWBT_MAXICODE.
	SupportedBarcodeTypeMAXICode

	// SupportedBarcodeTypeQRCode is an alias for TWBT_QRCODE.
	SupportedBarcodeTypeQRCode
)

// CompressionType is an alias for ICAP_COMPRESSION values.
type CompressionType twain.UInt16

const (
	// CompressionTypeNone is an alias for TWCP_NONE.
	CompressionTypeNone CompressionType = iota

	// CompressionTypePackBits is an alias for TWCP_PACKBITS.
	CompressionTypePackBits

	// CompressionTypeGroup31D is an alias for TWCP_GROUP31D.
	CompressionTypeGroup31D

	// CompressionTypeGroup31DEOL is an alias for TWCP_GROUP31DEOL.
	CompressionTypeGroup31DEOL

	// CompressionTypeGroup32D is an alias for TWCP_GROUP32D.
	CompressionTypeGroup32D

	// CompressionTypeGroup4 is an alias for TWCP_GROUP4.
	CompressionTypeGroup4

	// CompressionTypeJPEG is an alias for TWCP_JPEG.
	CompressionTypeJPEG

	// CompressionTypeLZW is an alias for TWCP_LZW.
	CompressionTypeLZW

	// CompressionTypeJBIG is an alias for TWCP_JBIG.
	CompressionTypeJBIG

	// CompressionTypePNG is an alias for TWCP_PNG.
	CompressionTypePNG

	// CompressionTypeRLE4 is an alias for TWCP_RLE4.
	CompressionTypeRLE4

	// CompressionTypeRLE8 is an alias for TWCP_RLE8.
	CompressionTypeRLE8

	// CompressionTypeBitFields is an alias for TWCP_BITFIELDS.
	CompressionTypeBitFields

	// CompressionTypeZIP is an alias for TWCP_ZIP.
	CompressionTypeZIP

	// CompressionTypeJPEG2000 is an alias for TWCP_JPEG2000.
	CompressionTypeJPEG2000
)

// CameraSide is an alias for CAP_CAMERASIDE and TWEI_PAGESIDE values.
type CameraSide twain.UInt16

const (
	// CameraSideBoth is an alias for TWCS_BOTH.
	CameraSideBoth CameraSide = iota

	// CameraSideTop is an alias for TWCS_TOP.
	CameraSideTop

	// CameraSideBottom is an alias for TWCS_BOTTOM.
	CameraSideBottom
)

// DeviceEventType is an alias for CAP_DEVICEEVENT values.
type DeviceEventType twain.UInt16

const DeviceEventTypeCustomEvents DeviceEventType = 0x8000

const (
	// DeviceEventTypeCheckAutomaticCapture is an alias for TWDE_CHECKAUTOMATICCAPTURE.
	DeviceEventTypeCheckAutomaticCapture DeviceEventType = iota

	// DeviceEventTypeCheckBattery is an alias for TWDE_CHECKBATTERY.
	DeviceEventTypeCheckBattery

	// DeviceEventTypeCheckDeviceOnline is an alias for TWDE_CHECKDEVICEONLINE.
	DeviceEventTypeCheckDeviceOnline

	// DeviceEventTypeCheckFlash is an alias for TWDE_CHECKFLASH.
	DeviceEventTypeCheckFlash

	// DeviceEventTypeCheckPowerSupply is an alias for TWDE_CHECKPOWERSUPPLY.
	DeviceEventTypeCheckPowerSupply

	// DeviceEventTypeCheckResolution is an alias for TWDE_CHECKRESOLUTION.
	DeviceEventTypeCheckResolution

	// DeviceEventTypeDeviceAdded is an alias for TWDE_DEVICEADDED.
	DeviceEventTypeDeviceAdded

	// DeviceEventTypeDeviceOffline is an alias for TWDE_DEVICEOFFLINE.
	DeviceEventTypeDeviceOffline

	// DeviceEventTypeDeviceReady is an alias for TWDE_DEVICEREADY.
	DeviceEventTypeDeviceReady

	// DeviceEventTypeDeviceRemoved is an alias for TWDE_DEVICEREMOVED.
	DeviceEventTypeDeviceRemoved

	// DeviceEventTypeImageCaptured is an alias for TWDE_IMAGECAPTURED.
	DeviceEventTypeImageCaptured

	// DeviceEventTypeImageDeleted is an alias for TWDE_IMAGEDELETED.
	DeviceEventTypeImageDeleted

	// DeviceEventTypePaperDoubleFeed is an alias for TWDE_PAPERDOUBLEFEED.
	DeviceEventTypePaperDoubleFeed

	// DeviceEventTypePaperJam is an alias for TWDE_PAPERJAM.
	DeviceEventTypePaperJam

	// DeviceEventTypeLampFailure is an alias for TWDE_LAMPFAILURE.
	DeviceEventTypeLampFailure

	// DeviceEventTypePowerSave is an alias for TWDE_POWERSAVE.
	DeviceEventTypePowerSave

	// DeviceEventTypePowerSaveNotify is an alias for TWDE_POWERSAVENOTIFY.
	DeviceEventTypePowerSaveNotify
)

// PassThruDirection is an alias for TW_PASSTHRU.Direction values.
type PassThruDirection twain.UInt16

const (
	// PassThruDirectionGet is an alias for TWDR_GET.
	PassThruDirectionGet = 1

	// PassThruDirectionSet is an alias for TWDR_SET.
	PassThruDirectionSet = 2
)

// DeskEWStatus is an alias for TWEI_DESKEWSTATUS values.
type DeskEWStatus twain.UInt16

const (
	// DeskEWStatusSuccess is an alias for TWDSK_SUCCESS.
	DeskEWStatusSuccess DeskEWStatus = iota

	// DeskEWStatusReportOnly is an alias for TWDSK_REPORTONLY.
	DeskEWStatusReportOnly

	// DeskEWStatusFail is an alias for TWDSK_FAIL.
	DeskEWStatusFail

	// DeskEWStatusDisabled is an alias for TWDSK_DISABLED.
	DeskEWStatusDisabled
)

// DuplexType is an alias for CAP_DUPLEX values.
type DuplexType twain.UInt16

const (
	// DuplexTypeNone is an alias for TWDX_NONE.
	DuplexTypeNone DuplexType = iota

	// DuplexType1PassDuplex is an alias for TWDX_1PASSDUPLEX.
	DuplexType1PassDuplex

	// DuplexType2PassDuplex is an alias for TWDX_2PASSDUPLEX.
	DuplexType2PassDuplex
)

// FeederAlignment is an alias for CAP_FEEDERALIGNMENT values.
type FeederAlignment twain.UInt16

const (
	// FeederAlignmentNone is an alias for TWFA_NONE.
	FeederAlignmentNone FeederAlignment = iota

	// FeederAlignmentLeft is an alias for TWFA_LEFT.
	FeederAlignmentLeft

	// FeederAlignmentCenter is an alias for TWFA_CENTER.
	FeederAlignmentCenter

	// FeederAlignmentRight is an alias for TWFA_RIGHT.
	FeederAlignmentRight
)

// FeederType is an alias for ICAP_FEEDERTYPE values.
type FeederType twain.UInt16

const (
	// FeederTypeGeneral is an alias for TWFE_GENERAL.
	FeederTypeGeneral FeederType = iota

	// FeederTypePhoto is an alias for TWFE_PHOTO.
	FeederTypePhoto
)

// ImageFileFormat is an alias for ICAP_IMAGEFILEFORMAT values.
type ImageFileFormat twain.UInt16

const (
	// ImageFileFormatTIFF is an alias for TWFF_TIFF.
	ImageFileFormatTIFF ImageFileFormat = iota

	// ImageFileFormatPICT is an alias for TWFF_PICT.
	ImageFileFormatPICT

	// ImageFileFormatBMP is an alias for TWFF_BMP.
	ImageFileFormatBMP

	// ImageFileFormatXBM is an alias for TWFF_XBM.
	ImageFileFormatXBM

	// ImageFileFormatJFIF is an alias for TWFF_JFIF.
	ImageFileFormatJFIF

	// ImageFileFormatFPX is an alias for TWFF_FPX.
	ImageFileFormatFPX

	// ImageFileFormatTIFFMulti is an alias for TWFF_TIFFMULTI.
	ImageFileFormatTIFFMulti

	// ImageFileFormatPNG is an alias for TWFF_PNG.
	ImageFileFormatPNG

	// ImageFileFormatSPIFF is an alias for TWFF_SPIFF.
	ImageFileFormatSPIFF

	// ImageFileFormatEXIF is an alias for TWFF_EXIF.
	ImageFileFormatEXIF

	// ImageFileFormatPDF is an alias for TWFF_PDF.
	ImageFileFormatPDF

	// ImageFileFormatJP2 is an alias for TWFF_JP2.
	ImageFileFormatJP2

	// ImageFileFormatJPX is an alias for TWFF_JPX.
	ImageFileFormatJPX

	// ImageFileFormatDejaVu is an alias for TWFF_DEJAVU.
	ImageFileFormatDejaVu

	// ImageFileFormatPDFA is an alias for TWFF_PDFA.
	ImageFileFormatPDFA

	// ImageFileFormatPDFA2 is an alias for TWFF_PDFA2.
	ImageFileFormatPDFA2

	// ImageFileFormatPDFRaster is an alias for TWFF_PDFRASTER.
	ImageFileFormatPDFRaster
)

// FlashUsed2 is an alias for ICAP_FLASHUSED2 values.
type FlashUsed2 twain.UInt16

const (
	// FlashUsed2None is an alias for TWFL_NONE.
	FlashUsed2None FlashUsed2 = iota

	// FlashUsed2Off is an alias for TWFL_OFF.
	FlashUsed2Off

	// FlashUsed2On is an alias for TWFL_ON.
	FlashUsed2On

	// FlashUsed2Auto is an alias for TWFL_AUTO.
	FlashUsed2Auto

	// FlashUsed2RedEye is an alias for TWFL_REDEYE.
	FlashUsed2RedEye
)

// FeedOrder is an alias for CAP_FEEDERORDER values.
type FeedOrder twain.UInt16

const (
	// FeedOrderFirstPageFirst is an alias for TWFO_FIRSTPAGEFIRST.
	FeedOrderFirstPageFirst FeedOrder = iota

	// FeedOrderLastPageFirst is an alias for TWFO_LASTPAGEFIRST.
	FeedOrderLastPageFirst
)

// FeederPocket is an alias for CAP_FEEDERPOCKET values.
type FeederPocket twain.UInt16

const (
	// FeederPocketError is an alias for TWFP_POCKETERROR.
	FeederPocketError FeederPocket = iota

	// FeederPocket1 is an alias for TWFP_POCKET1.
	FeederPocket1

	// FeederPocket2 is an alias for TWFP_POCKET2.
	FeederPocket2

	// FeederPocket3 is an alias for TWFP_POCKET3.
	FeederPocket3

	// FeederPocket4 is an alias for TWFP_POCKET4.
	FeederPocket4

	// FeederPocket5 is an alias for TWFP_POCKET5.
	FeederPocket5

	// FeederPocket6 is an alias for TWFP_POCKET6.
	FeederPocket6

	// FeederPocket7 is an alias for TWFP_POCKET7.
	FeederPocket7

	// FeederPocket8 is an alias for TWFP_POCKET8.
	FeederPocket8

	// FeederPocket9 is an alias for TWFP_POCKET9.
	FeederPocket9

	// FeederPocket10 is an alias for TWFP_POCKET10.
	FeederPocket10

	// FeederPocket11 is an alias for TWFP_POCKET11.
	FeederPocket11

	// FeederPocket12 is an alias for TWFP_POCKET12.
	FeederPocket12

	// FeederPocket13 is an alias for TWFP_POCKET13.
	FeederPocket13

	// FeederPocket14 is an alias for TWFP_POCKET14.
	FeederPocket14

	// FeederPocket15 is an alias for TWFP_POCKET15.
	FeederPocket15

	// FeederPocket16 is an alias for TWFP_POCKET16.
	FeederPocket16
)

// FlipRotation is an alias for ICAP_FLIPROTATION values.
type FlipRotation twain.UInt16

const (
	// FlipRotationBook is an alias for TWFR_BOOK.
	FlipRotationBook FlipRotation = iota

	// FlipRotationFanFold is an alias for TWFR_FANFOLD.
	FlipRotationFanFold
)

// FilterColor is an alias for ICAP_FILTER values.
type FilterColor twain.UInt16

const (
	// FilterColorRed is an alias for TWFT_RED.
	FilterColorRed FilterColor = iota

	// FilterColorGreen is an alias for TWFT_GREEN.
	FilterColorGreen

	// FilterColorBlue is an alias for TWFT_BLUE.
	FilterColorBlue

	// FilterColorNone is an alias for TWFT_NONE.
	FilterColorNone

	// FilterColorWhite is an alias for TWFT_WHITE.
	FilterColorWhite

	// FilterColorCyan is an alias for TWFT_CYAN.
	FilterColorCyan

	// FilterColorMagenta is an alias for TWFT_MAGENTA.
	FilterColorMagenta

	// FilterColorYellow is an alias for TWFT_YELLOW.
	FilterColorYellow

	// FilterColorBlack is an alias for TWFT_BLACK.
	FilterColorBlack
)

// ICCProfile is an alias for ICAP_ICCPROFILE values.
type ICCProfile twain.UInt16

const (
	// ICCProfileNone is an alias for TWIC_NONE.
	ICCProfileNone ICCProfile = iota

	// ICCProfileLink is an alias for TWIC_LINK.
	ICCProfileLink

	// ICCProfileEmbed is an alias for TWIC_EMBED.
	ICCProfileEmbed
)

// ImageFilterType is an alias for ICAP_IMAGEFILTER values.
type ImageFilterType twain.UInt16

const (
	// ImageFilterTypeNone is an alias for TWIF_NONE.
	ImageFilterTypeNone ImageFilterType = iota

	// ImageFilterTypeAuto is an alias for TWIF_AUTO.
	ImageFilterTypeAuto

	// ImageFilterTypeLowPass is an alias for TWIF_LOWPASS.
	ImageFilterTypeLowPass

	// ImageFilterTypeBandPass is an alias for TWIF_BANDPASS.
	ImageFilterTypeBandPass

	// ImageFilterTypeHighPass is an alias for TWIF_HIGHPASS.
	ImageFilterTypeHighPass

	// ImageFilterTypeText is an alias for TWIF_TEXT.
	ImageFilterTypeText = ImageFilterTypeBandPass

	// ImageFilterTypeFineLine is an alias for TWIF_FINELINE.
	ImageFilterTypeFineLine = ImageFilterTypeHighPass
)

// ImageMergeType is an alias for ICAP_IMAGEMERGE values.
type ImageMergeType twain.UInt16

const (
	// ImageMergeTypeNone is an alias for TWIM_NONE.
	ImageMergeTypeNone ImageMergeType = iota

	// ImageMergeTypeFrontOnTop is an alias for TWIM_FRONTONTOP.
	ImageMergeTypeFrontOnTop

	// ImageMergeTypeFrontOnBottom is an alias for TWIM_FRONTONBOTTOM.
	ImageMergeTypeFrontOnBottom

	// ImageMergeTypeFrontOnLeft is an alias for TWIM_FRONTONLEFT.
	ImageMergeTypeFrontOnLeft

	// ImageMergeTypeFrontOnRight is an alias for TWIM_FRONTONRIGHT.
	ImageMergeTypeFrontOnRight
)

// JobControlType is an alias for CAP_JOBCONTROL values.
type JobControlType twain.UInt16

const (
	// JobControlTypeNone is an alias for TWJC_NONE.
	JobControlTypeNone JobControlType = iota

	// JobControlTypeJSIC is an alias for TWJC_JSIC.
	JobControlTypeJSIC

	// JobControlTypeJSIS is an alias for TWJC_JSIS.
	JobControlTypeJSIS

	// JobControlTypeJSXC is an alias for TWJC_JSXC.
	JobControlTypeJSXC

	// JobControlTypeJSXS is an alias for TWJC_JSXS.
	JobControlTypeJSXS
)

// JPEGQuality is an alias for ICAP_JPEGQUALITY values.
type JPEGQuality twain.UInt16

const (
	// JPEGQualityUnknown is an alias for TWJQ_UNKNOWN.
	JPEGQualityUnknown JPEGQuality = -4

	// JPEGQualityLow is an alias for TWJQ_LOW.
	JPEGQualityLow JPEGQuality = -3

	// JPEGQualityMedium is an alias for TWJQ_MEDIUM.
	JPEGQualityMedium JPEGQuality = -2

	// JPEGQualityHigh is an alias for TWJQ_HIGH.
	JPEGQualityHigh JPEGQuality = -1
)

// LightPath is an alias for ICAP_LIGHTPATH values.
type LightPath twain.UInt16

const (
	// LightPathReflective is an alias for TWLP_REFLECTIVE.
	LightPathReflective LightPath = iota

	// LightPathTransmissive is an alias for TWLP_TRANSMISSIVE.
	LightPathTransmissive
)

// LightSource is an alias for ICAP_LIGHTSOURCE values.
type LightSource twain.UInt16

const (
	// LightSourceRed is an alias for TWLS_RED.
	LightSourceRed LightSource = iota

	// LightSourceGreen is an alias for TWLS_GREEN.
	LightSourceGreen

	// LightSourceBlue is an alias for TWLS_BLUE.
	LightSourceBlue

	// LightSourceNone is an alias for TWLS_NONE.
	LightSourceNone

	// LightSourceWhite is an alias for TWLS_WHITE.
	LightSourceWhite

	// LightSourceUV is an alias for TWLS_UV.
	LightSourceUV

	// LightSourceIR is an alias for TWLS_IR.
	LightSourceIR
)

// NoiseFilter is an alias for ICAP_NOISEFILTER values.
type NoiseFilter twain.UInt16

const (
	// NoiseFilterNone is an alias for TWNF_NONE.
	NoiseFilterNone NoiseFilter = iota

	// NoiseFilterAuto is an alias for TWNF_AUTO.
	NoiseFilterAuto

	// NoiseFilterLonePixel is an alias for TWNF_LONEPIXEL.
	NoiseFilterLonePixel

	// NoiseFilterMajorityRule is an alias for TWNF_MAJORITYRULE.
	NoiseFilterMajorityRule
)

type Orientation twain.UInt16

const (
	// OrientationRot0 is an alias for TWOR_ROT0.
	OrientationRot0 Orientation = iota

	// OrientationRot90 is an alias for TWOR_ROT90.
	OrientationRot90

	// OrientationRot180 is an alias for TWOR_ROT180.
	OrientationRot180

	// OrientationRot270 is an alias for TWOR_ROT270.
	OrientationRot270

	// OrientationAuto is an alias for TWOR_AUTO.
	OrientationAuto

	// OrientationAutoText is an alias for TWOR_AUTOTEXT.
	OrientationAutoText

	// OrientationAutoPicture is an alias for TWOR_AUTOPICTURE.
	OrientationAutoPicture

	// OrientationPortrait is an alias for TWOR_PORTRAIT.
	OrientationPortrait = OrientationRot0

	// OrientationLandscape is an alias for TWOR_LANDSCAPE.
	OrientationLandscape = OrientationRot270
)

// OverScan is an alias for ICAP_OVERSCAN values.
type OverScan twain.UInt16

const (
	// OverScanNone is an alias for TWOV_NONE.
	OverScanNone OverScan = iota

	// OverScanAuto is an alias for TWOV_AUTO.
	OverScanAuto

	// OverScanTopBottom is an alias for TWOV_TOPBOTTOM.
	OverScanTopBottom

	// OverScanLeftRight is an alias for TWOV_LEFTRIGHT.
	OverScanLeftRight

	// OverScanAll is an alias for TWOV_ALL.
	OverScanAll
)

// PlanarChunky is an alias for ICAP_PLANARCHUNKY values.
type PlanarChunky twain.UInt16

const (
	// PlanarChunkyChunky is an alias for TWPC_CHUNKY.
	PlanarChunkyChunky PlanarChunky = iota

	// PlanarChunkyPlanar is an alias for TWPC_PLANAR.
	PlanarChunkyPlanar
)

// PixelFlavor is an alias for ICAP_PIXELFLAVOR values.
type PixelFlavor twain.UInt16

const (
	// PixelFlavorChocolate is an alias for TWPF_CHOCOLATE.
	PixelFlavorChocolate PixelFlavor = iota

	// PixelFlavorVanilla is an alias for TWPF_VANILLA.
	PixelFlavorVanilla
)

// PrinterMode is an alias for CAP_PRINTERMODE values.
type PrinterMode twain.UInt16

const (
	// PrinterModeSingleString is an alias for TWPM_SINGLESTRING.
	PrinterModeSingleString PrinterMode = iota

	// PrinterModeMultiString is an alias for TWPM_MULTISTRING.
	PrinterModeMultiString

	// PrinterModeCompoundString is an alias for TWPM_COMPOUNDSTRING.
	PrinterModeCompoundString
)

// PrinterFlag is an alias for CAP_PRINTER values.
type PrinterFlag twain.UInt16

const (
	// PrinterFlagImprinterTopBefore is an alias for TWPR_IMPRINTERTOPBEFORE.
	PrinterFlagImprinterTopBefore PrinterFlag = iota

	// PrinterFlagImprinterTopAfter is an alias for TWPR_IMPRINTERTOPAFTER.
	PrinterFlagImprinterTopAfter

	// PrinterFlagImprinterBottomBefore is an alias for TWPR_IMPRINTERBOTTOMBEFORE.
	PrinterFlagImprinterBottomBefore

	// PrinterFlagImprinterBottomAfter is an alias for TWPR_IMPRINTERBOTTOMAFTER.
	PrinterFlagImprinterBottomAfter

	// PrinterFlagEndorserTopBefore is an alias for TWPR_ENDORSERTOPBEFORE.
	PrinterFlagEndorserTopBefore

	// PrinterFlagEndorserTopAfter is an alias for TWPR_ENDORSERTOPAFTER.
	PrinterFlagEndorserTopAfter

	// PrinterFlagEndorserBottomBefore is an alias for TWPR_ENDORSERBOTTOMBEFORE.
	PrinterFlagEndorserBottomBefore

	// PrinterFlagEndorserBottomAfter is an alias for TWPR_ENDORSERBOTTOMAFTER.
	PrinterFlagEndorserBottomAfter
)

// PrinterFontStyle is an alias for CAP_PRINTERFONTSTYLE values.
type PrinterFontStyle twain.UInt16

const (
	// PrinterFontStyleNormal is an alias for TWPF_NORMAL.
	PrinterFontStyleNormal PrinterFontStyle = iota

	// PrinterFontStyleBold is an alias for TWPF_BOLD.
	PrinterFontStyleBold

	// PrinterFontStyleItalic is an alias for TWPF_ITALIC.
	PrinterFontStyleItalic

	// PrinterFontStyleLargeSize is an alias for TWPF_LARGESIZE.
	PrinterFontStyleLargeSize

	// PrinterFontStyleSmallSize is an alias for TWPF_SMALLSIZE.
	PrinterFontStyleSmallSize
)

// PrinterIndexTrigger is an alias for CAP_PRINTERINDEXTRIGGER values.
type PrinterIndexTrigger twain.UInt16

const (
	// PrinterIndexTriggerPage is an alias for TWCT_PAGE.
	PrinterIndexTriggerPage PrinterIndexTrigger = iota

	// PrinterIndexTriggerPatch1 is an alias for TWCT_PATCH1.
	PrinterIndexTriggerPatch1

	// PrinterIndexTriggerPatch2 is an alias for TWCT_PATCH2.
	PrinterIndexTriggerPatch2

	// PrinterIndexTriggerPatch3 is an alias for TWCT_PATCH3.
	PrinterIndexTriggerPatch3

	// PrinterIndexTriggerPatch4 is an alias for TWCT_PATCH4.
	PrinterIndexTriggerPatch4

	// PrinterIndexTriggerPatchT is an alias for TWCT_PATCHT.
	PrinterIndexTriggerPatchT

	// PrinterIndexTriggerPatch6 is an alias for TWCT_PATCH6.
	PrinterIndexTriggerPatch6
)

// PowerSupply is an alias for CAP_POWERSUPPLY values.
type PowerSupply twain.UInt16

const (
	// PowerSupplyExternal is an alias for TWPS_EXTERNAL.
	PowerSupplyExternal PowerSupply = iota

	// PowerSupplyBattery is an alias for TWPS_BATTERY.
	PowerSupplyBattery
)

// PixelType is an alias for ICAP_PIXELTYPE values.
type PixelType twain.UInt16

const (
	// PixelTypeBW is an alias for TWPT_BW.
	PixelTypeBW PixelType = iota

	// PixelTypeGray is an alias for TWPT_GRAY.
	PixelTypeGray

	// PixelTypeRGB is an alias for TWPT_RGB.
	PixelTypeRGB

	// PixelTypePalette is an alias for TWPT_PALETTE.
	PixelTypePalette

	// PixelTypeCMY is an alias for TWPT_CMY.
	PixelTypeCMY

	// PixelTypeCMYK is an alias for TWPT_CMYK.
	PixelTypeCMYK

	// PixelTypeYUV is an alias for TWPT_YUV.
	PixelTypeYUV

	// PixelTypeYUVK is an alias for TWPT_YUVK.
	PixelTypeYUVK

	// PixelTypeCIEXYZ is an alias for TWPT_CIEXYZ.
	PixelTypeCIEXYZ

	// PixelTypeLAB is an alias for TWPT_LAB.
	PixelTypeLAB

	// PixelTypeSRGB is an alias for TWPT_SRGB.
	PixelTypeSRGB

	// PixelTypeSCRGB is an alias for TWPT_SCRGB.
	PixelTypeSCRGB

	// PixelTypeInfrared is an alias for TWPT_INFRARED.
	PixelTypeInfrared PixelType = 16
)

// Segmented is an alias for CAP_SEGMENTED values.
type Segmented twain.UInt16

const (
	// SegmentedNone is an alias for TWSG_NONE.
	SegmentedNone Segmented = iota

	// SegmentedAuto is an alias for TWSG_AUTO.
	SegmentedAuto

	// SegmentedManual is an alias for TWSG_MANUAL.
	SegmentedManual
)

// FilmType is an alias for ICAP_FILMTYPE values.
type FilmType twain.UInt16

const (
	// FilmTypePositive is an alias for TWFM_POSITIVE.
	FilmTypePositive FilmType = iota

	// FilmTypeNegative is an alias for TWFM_NEGATIVE.
	FilmTypeNegative
)

// DoubleFeedDetection is an alias for CAP_DOUBLEFEEDDETECTION values.
type DoubleFeedDetection twain.UInt16

const (
	// DoubleFeedDetectionUltrasonic is an alias for TWDF_ULTRASONIC.
	DoubleFeedDetectionUltrasonic DoubleFeedDetection = iota

	// DoubleFeedDetectionByLength is an alias for TWDF_BYLENGTH.
	DoubleFeedDetectionByLength

	// DoubleFeedDetectionInfrared is an alias for TWDF_INFRARED.
	DoubleFeedDetectionInfrared
)

// DoubleFeedDetectionSensitivity is an alias for CAP_DOUBLEFEEDDETECTIONSENSITIVITY values.
type DoubleFeedDetectionSensitivity twain.UInt16

const (
	// DoubleFeedDetectionSensitivityLow is an alias for TWUS_LOW.
	DoubleFeedDetectionSensitivityLow DoubleFeedDetectionSensitivity = iota

	// DoubleFeedDetectionSensitivityMedium is an alias for TWUS_MEDIUM.
	DoubleFeedDetectionSensitivityMedium

	// DoubleFeedDetectionSensitivityHigh is an alias for TWUS_HIGH.
	DoubleFeedDetectionSensitivityHigh
)

// DoubleFeedDetectionResponse is an alias for CAP_DOUBLEFEEDDETECTIONRESPONSE values.
type DoubleFeedDetectionResponse twain.UInt16

const (
	// DoubleFeedDetectionResponseStop is an alias for TWDP_STOP.
	DoubleFeedDetectionResponseStop DoubleFeedDetectionResponse = iota

	// DoubleFeedDetectionResponseStopAndWait is an alias for TWDP_STOPANDWAIT.
	DoubleFeedDetectionResponseStopAndWait

	// DoubleFeedDetectionResponseSound is an alias for TWDP_SOUND.
	DoubleFeedDetectionResponseSound

	// DoubleFeedDetectionResponseDoNotImprint is an alias for TWDP_DONOTIMPRINT.
	DoubleFeedDetectionResponseDoNotImprint
)

// MirrorType is an alias for ICAP_MIRROR values.
type MirrorType twain.UInt16

const (
	// MirrorTypeNone is an alias for TWMR_NONE.
	MirrorTypeNone MirrorType = iota

	// MirrorTypeVertical is an alias for TWMR_VERTICAL.
	MirrorTypeVertical

	// MirrorTypeHorizontal is an alias for TWMR_HORIZONTAL.
	MirrorTypeHorizontal
)

// JPEGSubSamplingType is an alias for ICAP_JPEGSUBSAMPLING values.
type JPEGSubSamplingType twain.UInt16

const (
	// JPEGSubSamplingType444YCBCR is an alias for TWJS_444YCBCR.
	JPEGSubSamplingType444YCBCR JPEGSubSamplingType = iota

	// JPEGSubSamplingType444RGB is an alias for TWJS_444RGB.
	JPEGSubSamplingType444RGB

	// JPEGSubSamplingType422 is an alias for TWJS_422.
	JPEGSubSamplingType422

	// JPEGSubSamplingType421 is an alias for TWJS_421.
	JPEGSubSamplingType421

	// JPEGSubSamplingType411 is an alias for TWJS_411.
	JPEGSubSamplingType411

	// JPEGSubSamplingType420 is an alias for TWJS_420.
	JPEGSubSamplingType420

	// JPEGSubSamplingType410 is an alias for TWJS_410.
	JPEGSubSamplingType410

	// JPEGSubSamplingType311 is an alias for TWJS_311.
	JPEGSubSamplingType311
)

// PaperHandlingType is an alias for CAP_PAPERHANDLING values.
type PaperHandlingType twain.UInt16

const (
	// PaperHandlingTypeNormal is an alias for TWPH_NORMAL.
	PaperHandlingTypeNormal PaperHandlingType = iota

	// PaperHandlingTypeFragile is an alias for TWPH_FRAGILE.
	PaperHandlingTypeFragile

	// PaperHandlingTypeThick is an alias for TWPH_THICK.
	PaperHandlingTypeThick

	// PaperHandlingTypeTrifold is an alias for TWPH_TRIFOLD.
	PaperHandlingTypeTrifold

	// PaperHandlingTypePhotograph is an alias for TWPH_PHOTOGRAPH.
	PaperHandlingTypePhotograph
)

// IndicatorsMode is an alias for CAP_INDICATORSMODE values.
type IndicatorsMode twain.UInt16

const (
	// IndicatorsModeInfo is an alias for TWCI_INFO.
	IndicatorsModeInfo IndicatorsMode = iota

	// IndicatorsModeWarning is an alias for TWCI_WARNING.
	IndicatorsModeWarning

	// IndicatorsModeError is an alias for TWCI_ERROR.
	IndicatorsModeError

	// IndicatorsModeWarmup is an alias for TWCI_WARMUP.
	IndicatorsModeWarmup
)

// SupportedSize is an alias for ICAP_SUPPORTEDSIZES values.
type SupportedSize twain.UInt16

const (
	// SupportedSizeNone is an alias for TWSS_NONE.
	SupportedSizeNone SupportedSize = iota

	// SupportedSizeA4 is an alias for TWSS_A4.
	SupportedSizeA4

	// SupportedSizeJISB5 is an alias for TWSS_JISB5.
	SupportedSizeJISB5

	// SupportedSizeUSLetter is an alias for TWSS_USLETTER.
	SupportedSizeUSLetter

	// SupportedSizeUSLegal is an alias for TWSS_USLEGAL.
	SupportedSizeUSLegal

	// SupportedSizeA5 is an alias for TWSS_A5.
	SupportedSizeA5

	// SupportedSizeISOB4 is an alias for TWSS_ISOB4.
	SupportedSizeISOB4

	// SupportedSizeISOB6 is an alias for TWSS_ISOB6.
	SupportedSizeISOB6

	// SupportedSizeUnused is needed because there is no size for 8.
	// Do not use this.
	SupportedSizeUnused

	// SupportedSizeUSLedger is an alias for TWSS_USLEDGER.
	SupportedSizeUSLedger

	// SupportedSizeUSExecutive is an alias for TWSS_USEXECUTIVE.
	SupportedSizeUSExecutive

	// SupportedSizeA3 is an alias for TWSS_A3.
	SupportedSizeA3

	// SupportedSizeISOB3 is an alias for TWSS_ISOB3.
	SupportedSizeISOB3

	// SupportedSizeA6 is an alias for TWSS_A6.
	SupportedSizeA6

	// SupportedSizeC4 is an alias for TWSS_C4.
	SupportedSizeC4

	// SupportedSizeC5 is an alias for TWSS_C5.
	SupportedSizeC5

	// SupportedSizeC6 is an alias for TWSS_C6.
	SupportedSizeC6

	// SupportedSize4A0 is an alias for TWSS_4A0.
	SupportedSize4A0

	// SupportedSize2A0 is an alias for TWSS_2A0.
	SupportedSize2A0

	// SupportedSizeA0 is an alias for TWSS_A0.
	SupportedSizeA0

	// SupportedSizeA1 is an alias for TWSS_A1.
	SupportedSizeA1

	// SupportedSizeA2 is an alias for TWSS_A2.
	SupportedSizeA2

	// SupportedSizeA7 is an alias for TWSS_A7.
	SupportedSizeA7

	// SupportedSizeA8 is an alias for TWSS_A8.
	SupportedSizeA8

	// SupportedSizeA9 is an alias for TWSS_A9.
	SupportedSizeA9

	// SupportedSizeA10 is an alias for TWSS_A10.
	SupportedSizeA10

	// SupportedSizeISOB0 is an alias for TWSS_ISOB0.
	SupportedSizeISOB0

	// SupportedSizeISOB1 is an alias for TWSS_ISOB1.
	SupportedSizeISOB1

	// SupportedSizeISOB2 is an alias for TWSS_ISOB2.
	SupportedSizeISOB2

	// SupportedSizeISOB5 is an alias for TWSS_ISOB5.
	SupportedSizeISOB5

	// SupportedSizeISOB7 is an alias for TWSS_ISOB7.
	SupportedSizeISOB7

	// SupportedSizeISOB8 is an alias for TWSS_ISOB8.
	SupportedSizeISOB8

	// SupportedSizeISOB9 is an alias for TWSS_ISOB9.
	SupportedSizeISOB9

	// SupportedSizeISOB10 is an alias for TWSS_ISOB10.
	SupportedSizeISOB10

	// SupportedSizeJISB0 is an alias for TWSS_JISB0.
	SupportedSizeJISB0

	// SupportedSizeJISB1 is an alias for TWSS_JISB1.
	SupportedSizeJISB1

	// SupportedSizeJISB2 is an alias for TWSS_JISB2.
	SupportedSizeJISB2

	// SupportedSizeJISB3 is an alias for TWSS_JISB3.
	SupportedSizeJISB3

	// SupportedSizeJISB4 is an alias for TWSS_JISB4.
	SupportedSizeJISB4

	// SupportedSizeJISB6 is an alias for TWSS_JISB6.
	SupportedSizeJISB6

	// SupportedSizeJISB7 is an alias for TWSS_JISB7.
	SupportedSizeJISB7

	// SupportedSizeJISB8 is an alias for TWSS_JISB8.
	SupportedSizeJISB8

	// SupportedSizeJISB9 is an alias for TWSS_JISB9.
	SupportedSizeJISB9

	// SupportedSizeJISB10 is an alias for TWSS_JISB10.
	SupportedSizeJISB10

	// SupportedSizeC0 is an alias for TWSS_C0.
	SupportedSizeC0

	// SupportedSizeC1 is an alias for TWSS_C1.
	SupportedSizeC1

	// SupportedSizeC2 is an alias for TWSS_C2.
	SupportedSizeC2

	// SupportedSizeC3 is an alias for TWSS_C3.
	SupportedSizeC3

	// SupportedSizeC7 is an alias for TWSS_C7.
	SupportedSizeC7

	// SupportedSizeC8 is an alias for TWSS_C8.
	SupportedSizeC8

	// SupportedSizeC9 is an alias for TWSS_C9.
	SupportedSizeC9

	// SupportedSizeC10 is an alias for TWSS_C10.
	SupportedSizeC10

	// SupportedSizeUSStatement is an alias for TWSS_USSTATEMENT.
	SupportedSizeUSStatement

	// SupportedSizeBusinessCard is an alias for TWSS_BUSINESSCARD.
	SupportedSizeBusinessCard

	// SupportedSizeMaxSize is an alias for TWSS_MAXSIZE.
	SupportedSizeMaxSize
)

// TransferMechanismType is an alias for ICAP_XFERMECH values.
type TransferMechanismType twain.UInt16

const (
	// TransferMechanismTypeNative is an alias for TWSX_NATIVE.
	TransferMechanismTypeNative TransferMechanismType = iota

	// TransferMechanismTypeFile is an alias for TWSX_FILE.
	TransferMechanismTypeFile

	// TransferMechanismTypeMemory is an alias for TWSX_MEMORY.
	TransferMechanismTypeMemory

	// TransferMechanismTypeMemFile is an alias for TWSX_MEMFILE.
	TransferMechanismTypeMemFile
)

// UnitType is an alias for ICAP_UNITS values.
type UnitType twain.UInt16

const (
	// UnitTypeInches is an alias for TWUN_INCHES.
	UnitTypeInches UnitType = iota

	// UnitTypeCentimeters is an alias for TWUN_CENTIMETERS.
	UnitTypeCentimeters

	// UnitTypePicas is an alias for TWUN_PICAS.
	UnitTypePicas

	// UnitTypePoints is an alias for TWUN_POINTS.
	UnitTypePoints

	// UnitTypeTwips is an alias for TWUN_TWIPS.
	UnitTypeTwips

	// UnitTypePixels is an alias for TWUN_PIXELS.
	UnitTypePixels

	// UnitTypeMillimeters is an alias for TWUN_MILLIMETERS.
	UnitTypeMillimeters
)
