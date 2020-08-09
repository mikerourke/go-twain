package capabilities

import "github.com/mikerourke/go-twain/pkg/twain"

// AlarmValue is an alias for the CAP_ALARMS constants.
type AlarmValue twain.UInt16

const (
	// AlarmValueAlarm is an alias for TWAL_ALARM.
	AlarmValueAlarm AlarmValue = iota

	// AlarmValueFeedError is an alias for TWAL_FEEDERERROR.
	AlarmValueFeedError

	// AlarmValueFeederWarning is an alias for TWAL_FEEDERWARNING.
	AlarmValueFeederWarning

	// AlarmValueBarcode is an alias for TWAL_BARCODE.
	AlarmValueBarcode

	// AlarmValueDoubleFeed is an alias for TWAL_DOUBLEFEED.
	AlarmValueDoubleFeed

	// AlarmValueJam is an alias for TWAL_JAM.
	AlarmValueJam

	// AlarmValuePatchCode is an alias for TWAL_PATCHCODE.
	AlarmValuePatchCode

	// AlarmValuePower is an alias for TWAL_POWER.
	AlarmValuePower

	// AlarmValueSkew is an alias for TWAL_SKEW.
	AlarmValueSkew
)

// AutoSizeValue is an alias for the ICAP_AUTOSIZE values.
type AutoSizeValue twain.UInt16

const (
	// AutoSizeValueNone is an alias for TWAS_NONE.
	AutoSizeValueNone AutoSizeValue = iota

	// AutoSizeValueAuto is an alias for TWAS_AUTO.
	AutoSizeValueAuto

	// AutoSizeValueCurrent is an alias for TWAS_CURRENT.
	AutoSizeValueCurrent
)

// BarcodeRotationValue is an alias for the TWEI_BARCODEROTATION values.
type BarcodeRotationValue twain.UInt16

const (
	// BarcodeRotationValue0 is an alias for TWBCOR_ROT0.
	BarcodeRotationValue0 BarcodeRotationValue = iota

	// BarcodeRotationValue90 is an alias for TWBCOR_ROT90.
	BarcodeRotationValue90

	// BarcodeRotationValue180 is an alias for TWBCOR_ROT180.
	BarcodeRotationValue180

	// BarcodeRotationValue270 is an alias for TWBCOR_ROT270.
	BarcodeRotationValue270

	// BarcodeRotationValueX is an alias for TWBCOR_ROTX.
	BarcodeRotationValueX
)

// BarcodeSearchModeValue is an alias for ICAP_BARCODESEARCHMODE values.
type BarcodeSearchModeValue twain.UInt16

const (
	// BarcodeSearchModeValueHoriz is an alias for TWBD_HORZ.
	BarcodeSearchModeValueHoriz BarcodeSearchModeValue = iota

	// BarcodeSearchModeValueVert is an alias for TWBD_VERT.
	BarcodeSearchModeValueVert

	// BarcodeSearchModeValueHorizVert is an alias for TWBD_HORZVERT.
	BarcodeSearchModeValueHorizVert

	// BarcodeSearchModeValueVertHoriz is an alias for TWBD_VERTHORZ.
	BarcodeSearchModeValueVertHoriz
)

// BitOrderValue is an alias for ICAP_BITORDER values.
type BitOrderValue twain.UInt16

const (
	// BitOrderValueLSBFirst is an alias for TWBO_LSBFIRST.
	BitOrderValueLSBFirst BitOrderValue = iota

	// BitOrderValueMSBFirst is an alias for TWBO_MSBFIRST.
	BitOrderValueMSBFirst
)

// AutoDiscardBlankPagesValue is an alias for ICAP_AUTODISCARDBLANKPAGES values.
type AutoDiscardBlankPagesValue twain.UInt16

const (
	// AutoDiscardBlankPagesValueDisable is an alias for TWBP_DISABLE.
	AutoDiscardBlankPagesValueDisable AutoDiscardBlankPagesValue = -2

	// AutoDiscardBlankPagesValueAuto is an alias for TWBP_AUTO.
	AutoDiscardBlankPagesValueAuto AutoDiscardBlankPagesValue = -1
)

// BitDepthReductionValue is an alias for ICAP_BITDEPTHREDUCTION values.
type BitDepthReductionValue twain.UInt16

const (
	// BitDepthReductionValueThreshold is an alias for TWBR_THRESHOLD.
	BitDepthReductionValueThreshold BitDepthReductionValue = iota

	// BitDepthReductionValueHalfTone is an alias for TWBR_HALFTONE.
	BitDepthReductionValueHalfTone

	// BitDepthReductionValueCustomHalfTone is an alias for TWBR_CUSTHALFTONE.
	BitDepthReductionValueCustomHalfTone

	// BitDepthReductionValueDiffusion is an alias for TWBR_DIFFUSION.
	BitDepthReductionValueDiffusion

	// BitDepthReductionValueDynamicThreshold is an alias for TWBR_DYNAMICTHRESHOLD.
	BitDepthReductionValueDynamicThreshold
)

// SupportedBarcodeTypeValue is an alias for ICAP_SUPPORTEDBARCODETYPES and
// TWEI_BARCODETYPE values.
type SupportedBarcodeTypeValue twain.UInt16

const (
	// SupportedBarcodeTypeValue3OF9 is an alias for TWBT_3OF9.
	SupportedBarcodeTypeValue3OF9 SupportedBarcodeTypeValue = iota

	// SupportedBarcodeTypeValue2OF5Interleaved is an alias for TWBT_2OF5INTERLEAVED.
	SupportedBarcodeTypeValue2OF5Interleaved

	// SupportedBarcodeTypeValue2OF5NonInterleaved is an alias for TWBT_2OF5NONINTERLEAVED.
	SupportedBarcodeTypeValue2OF5NonInterleaved

	// SupportedBarcodeTypeValueCode93 is an alias for TWBT_CODE93.
	SupportedBarcodeTypeValueCode93

	// SupportedBarcodeTypeValueCode128 is an alias for TWBT_CODE128.
	SupportedBarcodeTypeValueCode128

	// SupportedBarcodeTypeValueUCC128 is an alias for TWBT_UCC128.
	SupportedBarcodeTypeValueUCC128

	// SupportedBarcodeTypeValueCodaBar is an alias for TWBT_CODABAR.
	SupportedBarcodeTypeValueCodaBar

	// SupportedBarcodeTypeValueUPCA is an alias for TWBT_UPCA.
	SupportedBarcodeTypeValueUPCA

	// SupportedBarcodeTypeValueUPCE is an alias for TWBT_UPCE.
	SupportedBarcodeTypeValueUPCE

	// SupportedBarcodeTypeValueEAN8 is an alias for TWBT_EAN8.
	SupportedBarcodeTypeValueEAN8

	// SupportedBarcodeTypeValueEAN13 is an alias for TWBT_EAN13.
	SupportedBarcodeTypeValueEAN13

	// SupportedBarcodeTypeValuePostNet is an alias for TWBT_POSTNET.
	SupportedBarcodeTypeValuePostNet

	// SupportedBarcodeTypeValuePDF417 is an alias for TWBT_PDF417.
	SupportedBarcodeTypeValuePDF417

	// SupportedBarcodeTypeValue2OF5Industrial is an alias for TWBT_2OF5INDUSTRIAL.
	SupportedBarcodeTypeValue2OF5Industrial

	// SupportedBarcodeTypeValue2OF5Matrix is an alias for TWBT_2OF5MATRIX.
	SupportedBarcodeTypeValue2OF5Matrix

	// SupportedBarcodeTypeValue2OF5DataLogic is an alias for TWBT_2OF5DATALOGIC.
	SupportedBarcodeTypeValue2OF5DataLogic

	// SupportedBarcodeTypeValue2OF5IATA is an alias for TWBT_2OF5IATA.
	SupportedBarcodeTypeValue2OF5IATA

	// SupportedBarcodeTypeValue3OF9FullASCII is an alias for TWBT_3OF9FULLASCII.
	SupportedBarcodeTypeValue3OF9FullASCII

	// SupportedBarcodeTypeValueCodaBarWithStartStop is an alias for TWBT_CODABARWITHSTARTSTOP.
	SupportedBarcodeTypeValueCodaBarWithStartStop

	// SupportedBarcodeTypeValueMAXICode is an alias for TWBT_MAXICODE.
	SupportedBarcodeTypeValueMAXICode

	// SupportedBarcodeTypeValueQRCode is an alias for TWBT_QRCODE.
	SupportedBarcodeTypeValueQRCode
)

// CompressionValue is an alias for ICAP_COMPRESSION values.
type CompressionValue twain.UInt16

const (
	// CompressionValueNone is an alias for TWCP_NONE.
	CompressionValueNone CompressionValue = iota

	// CompressionValuePackBits is an alias for TWCP_PACKBITS.
	CompressionValuePackBits

	// CompressionValueGroup31D is an alias for TWCP_GROUP31D.
	CompressionValueGroup31D

	// CompressionValueGroup31DEOL is an alias for TWCP_GROUP31DEOL.
	CompressionValueGroup31DEOL

	// CompressionValueGroup32D is an alias for TWCP_GROUP32D.
	CompressionValueGroup32D

	// CompressionValueGroup4 is an alias for TWCP_GROUP4.
	CompressionValueGroup4

	// CompressionValueJPEG is an alias for TWCP_JPEG.
	CompressionValueJPEG

	// CompressionValueLZW is an alias for TWCP_LZW.
	CompressionValueLZW

	// CompressionValueJBIG is an alias for TWCP_JBIG.
	CompressionValueJBIG

	// CompressionValuePNG is an alias for TWCP_PNG.
	CompressionValuePNG

	// CompressionValueRLE4 is an alias for TWCP_RLE4.
	CompressionValueRLE4

	// CompressionValueRLE8 is an alias for TWCP_RLE8.
	CompressionValueRLE8

	// CompressionValueBitFields is an alias for TWCP_BITFIELDS.
	CompressionValueBitFields

	// CompressionValueZIP is an alias for TWCP_ZIP.
	CompressionValueZIP

	// CompressionValueJPEG2000 is an alias for TWCP_JPEG2000.
	CompressionValueJPEG2000
)

// CameraSideValue is an alias for CAP_CAMERASIDE and TWEI_PAGESIDE values.
type CameraSideValue twain.UInt16

const (
	// CameraSideValueBoth is an alias for TWCS_BOTH.
	CameraSideValueBoth CameraSideValue = iota

	// CameraSideValueTop is an alias for TWCS_TOP.
	CameraSideValueTop

	// CameraSideValueBottom is an alias for TWCS_BOTTOM.
	CameraSideValueBottom
)

// DeviceEventValue is an alias for CAP_DEVICEEVENT values.
type DeviceEventValue twain.UInt16

const DeviceEventValueCustomEvents DeviceEventValue = 0x8000

const (
	// DeviceEventValueCheckAutomaticCapture is an alias for TWDE_CHECKAUTOMATICCAPTURE.
	DeviceEventValueCheckAutomaticCapture DeviceEventValue = iota

	// DeviceEventValueCheckBattery is an alias for TWDE_CHECKBATTERY.
	DeviceEventValueCheckBattery

	// DeviceEventValueCheckDeviceOnline is an alias for TWDE_CHECKDEVICEONLINE.
	DeviceEventValueCheckDeviceOnline

	// DeviceEventValueCheckFlash is an alias for TWDE_CHECKFLASH.
	DeviceEventValueCheckFlash

	// DeviceEventValueCheckPowerSupplyValue is an alias for TWDE_CHECKPOWERSUPPLY.
	DeviceEventValueCheckPowerSupplyValue

	// DeviceEventValueCheckResolution is an alias for TWDE_CHECKRESOLUTION.
	DeviceEventValueCheckResolution

	// DeviceEventValueDeviceAdded is an alias for TWDE_DEVICEADDED.
	DeviceEventValueDeviceAdded

	// DeviceEventValueDeviceOffline is an alias for TWDE_DEVICEOFFLINE.
	DeviceEventValueDeviceOffline

	// DeviceEventValueDeviceReady is an alias for TWDE_DEVICEREADY.
	DeviceEventValueDeviceReady

	// DeviceEventValueDeviceRemoved is an alias for TWDE_DEVICEREMOVED.
	DeviceEventValueDeviceRemoved

	// DeviceEventValueImageCaptured is an alias for TWDE_IMAGECAPTURED.
	DeviceEventValueImageCaptured

	// DeviceEventValueImageDeleted is an alias for TWDE_IMAGEDELETED.
	DeviceEventValueImageDeleted

	// DeviceEventValuePaperDoubleFeed is an alias for TWDE_PAPERDOUBLEFEED.
	DeviceEventValuePaperDoubleFeed

	// DeviceEventValuePaperJam is an alias for TWDE_PAPERJAM.
	DeviceEventValuePaperJam

	// DeviceEventValueLampFailure is an alias for TWDE_LAMPFAILURE.
	DeviceEventValueLampFailure

	// DeviceEventValuePowerSave is an alias for TWDE_POWERSAVE.
	DeviceEventValuePowerSave

	// DeviceEventValuePowerSaveNotify is an alias for TWDE_POWERSAVENOTIFY.
	DeviceEventValuePowerSaveNotify
)

// PassThruDirectionValue is an alias for TW_PASSTHRU.Direction values.
type PassThruDirectionValue twain.UInt16

const (
	// PassThruDirectionValueGet is an alias for TWDR_GET.
	PassThruDirectionValueGet = 1

	// PassThruDirectionValueSet is an alias for TWDR_SET.
	PassThruDirectionValueSet = 2
)

// DeskEWStatusValue is an alias for TWEI_DESKEWSTATUS values.
type DeskEWStatusValue twain.UInt16

const (
	// DeskEWStatusValueSuccess is an alias for TWDSK_SUCCESS.
	DeskEWStatusValueSuccess DeskEWStatusValue = iota

	// DeskEWStatusValueReportOnly is an alias for TWDSK_REPORTONLY.
	DeskEWStatusValueReportOnly

	// DeskEWStatusValueFail is an alias for TWDSK_FAIL.
	DeskEWStatusValueFail

	// DeskEWStatusValueDisabled is an alias for TWDSK_DISABLED.
	DeskEWStatusValueDisabled
)

// DuplexValue is an alias for CAP_DUPLEX values.
type DuplexValue twain.UInt16

const (
	// DuplexValueNone is an alias for TWDX_NONE.
	DuplexValueNone DuplexValue = iota

	// DuplexValue1PassDuplex is an alias for TWDX_1PASSDUPLEX.
	DuplexValue1PassDuplex

	// DuplexValue2PassDuplex is an alias for TWDX_2PASSDUPLEX.
	DuplexValue2PassDuplex
)

// FeederAlignmentValue is an alias for CAP_FEEDERALIGNMENT values.
type FeederAlignmentValue twain.UInt16

const (
	// FeederAlignmentValueNone is an alias for TWFA_NONE.
	FeederAlignmentValueNone FeederAlignmentValue = iota

	// FeederAlignmentValueLeft is an alias for TWFA_LEFT.
	FeederAlignmentValueLeft

	// FeederAlignmentValueCenter is an alias for TWFA_CENTER.
	FeederAlignmentValueCenter

	// FeederAlignmentValueRight is an alias for TWFA_RIGHT.
	FeederAlignmentValueRight
)

// FeederTypeValue is an alias for ICAP_FEEDERTYPE values.
type FeederTypeValue twain.UInt16

const (
	// FeederTypeValueGeneral is an alias for TWFE_GENERAL.
	FeederTypeValueGeneral FeederTypeValue = iota

	// FeederTypeValuePhoto is an alias for TWFE_PHOTO.
	FeederTypeValuePhoto
)

// ImageFileFormatValue is an alias for ICAP_IMAGEFILEFORMAT values.
type ImageFileFormatValue twain.UInt16

const (
	// ImageFileFormatValueTIFF is an alias for TWFF_TIFF.
	ImageFileFormatValueTIFF ImageFileFormatValue = iota

	// ImageFileFormatValuePICT is an alias for TWFF_PICT.
	ImageFileFormatValuePICT

	// ImageFileFormatValueBMP is an alias for TWFF_BMP.
	ImageFileFormatValueBMP

	// ImageFileFormatValueXBM is an alias for TWFF_XBM.
	ImageFileFormatValueXBM

	// ImageFileFormatValueJFIF is an alias for TWFF_JFIF.
	ImageFileFormatValueJFIF

	// ImageFileFormatValueFPX is an alias for TWFF_FPX.
	ImageFileFormatValueFPX

	// ImageFileFormatValueTIFFMulti is an alias for TWFF_TIFFMULTI.
	ImageFileFormatValueTIFFMulti

	// ImageFileFormatValuePNG is an alias for TWFF_PNG.
	ImageFileFormatValuePNG

	// ImageFileFormatValueSPIFF is an alias for TWFF_SPIFF.
	ImageFileFormatValueSPIFF

	// ImageFileFormatValueEXIF is an alias for TWFF_EXIF.
	ImageFileFormatValueEXIF

	// ImageFileFormatValuePDF is an alias for TWFF_PDF.
	ImageFileFormatValuePDF

	// ImageFileFormatValueJP2 is an alias for TWFF_JP2.
	ImageFileFormatValueJP2

	// ImageFileFormatValueJPX is an alias for TWFF_JPX.
	ImageFileFormatValueJPX

	// ImageFileFormatValueDejaVu is an alias for TWFF_DEJAVU.
	ImageFileFormatValueDejaVu

	// ImageFileFormatValuePDFA is an alias for TWFF_PDFA.
	ImageFileFormatValuePDFA

	// ImageFileFormatValuePDFA2 is an alias for TWFF_PDFA2.
	ImageFileFormatValuePDFA2

	// ImageFileFormatValuePDFRaster is an alias for TWFF_PDFRASTER.
	ImageFileFormatValuePDFRaster
)

// FlashUsed2Value is an alias for ICAP_FLASHUSED2 values.
type FlashUsed2Value twain.UInt16

const (
	// FlashUsed2ValueNone is an alias for TWFL_NONE.
	FlashUsed2ValueNone FlashUsed2Value = iota

	// FlashUsed2ValueOff is an alias for TWFL_OFF.
	FlashUsed2ValueOff

	// FlashUsed2ValueOn is an alias for TWFL_ON.
	FlashUsed2ValueOn

	// FlashUsed2ValueAuto is an alias for TWFL_AUTO.
	FlashUsed2ValueAuto

	// FlashUsed2ValueRedEye is an alias for TWFL_REDEYE.
	FlashUsed2ValueRedEye
)

// FeedOrderValue is an alias for CAP_FEEDERORDER values.
type FeedOrderValue twain.UInt16

const (
	// FeedOrderValueFirstPageFirst is an alias for TWFO_FIRSTPAGEFIRST.
	FeedOrderValueFirstPageFirst FeedOrderValue = iota

	// FeedOrderValueLastPageFirst is an alias for TWFO_LASTPAGEFIRST.
	FeedOrderValueLastPageFirst
)

// FeederPocketValue is an alias for CAP_FEEDERPOCKET values.
type FeederPocketValue twain.UInt16

const (
	// FeederPocketValueError is an alias for TWFP_POCKETERROR.
	FeederPocketValueError FeederPocketValue = iota

	// FeederPocketValue1 is an alias for TWFP_POCKET1.
	FeederPocketValue1

	// FeederPocketValue2 is an alias for TWFP_POCKET2.
	FeederPocketValue2

	// FeederPocketValue3 is an alias for TWFP_POCKET3.
	FeederPocketValue3

	// FeederPocketValue4 is an alias for TWFP_POCKET4.
	FeederPocketValue4

	// FeederPocketValue5 is an alias for TWFP_POCKET5.
	FeederPocketValue5

	// FeederPocketValue6 is an alias for TWFP_POCKET6.
	FeederPocketValue6

	// FeederPocketValue7 is an alias for TWFP_POCKET7.
	FeederPocketValue7

	// FeederPocketValue8 is an alias for TWFP_POCKET8.
	FeederPocketValue8

	// FeederPocketValue9 is an alias for TWFP_POCKET9.
	FeederPocketValue9

	// FeederPocketValue10 is an alias for TWFP_POCKET10.
	FeederPocketValue10

	// FeederPocketValue11 is an alias for TWFP_POCKET11.
	FeederPocketValue11

	// FeederPocketValue12 is an alias for TWFP_POCKET12.
	FeederPocketValue12

	// FeederPocketValue13 is an alias for TWFP_POCKET13.
	FeederPocketValue13

	// FeederPocketValue14 is an alias for TWFP_POCKET14.
	FeederPocketValue14

	// FeederPocketValue15 is an alias for TWFP_POCKET15.
	FeederPocketValue15

	// FeederPocketValue16 is an alias for TWFP_POCKET16.
	FeederPocketValue16
)

// FlipRotationValue is an alias for ICAP_FLIPROTATION values.
type FlipRotationValue twain.UInt16

const (
	// FlipRotationValueBook is an alias for TWFR_BOOK.
	FlipRotationValueBook FlipRotationValue = iota

	// FlipRotationValueFanFold is an alias for TWFR_FANFOLD.
	FlipRotationValueFanFold
)

// FilterValue is an alias for ICAP_FILTER values.
type FilterValue twain.UInt16

const (
	// FilterValueRed is an alias for TWFT_RED.
	FilterValueRed FilterValue = iota

	// FilterValueGreen is an alias for TWFT_GREEN.
	FilterValueGreen

	// FilterValueBlue is an alias for TWFT_BLUE.
	FilterValueBlue

	// FilterValueNone is an alias for TWFT_NONE.
	FilterValueNone

	// FilterValueWhite is an alias for TWFT_WHITE.
	FilterValueWhite

	// FilterValueCyan is an alias for TWFT_CYAN.
	FilterValueCyan

	// FilterValueMagenta is an alias for TWFT_MAGENTA.
	FilterValueMagenta

	// FilterValueYellow is an alias for TWFT_YELLOW.
	FilterValueYellow

	// FilterValueBlack is an alias for TWFT_BLACK.
	FilterValueBlack
)

// ICCProfileValue is an alias for ICAP_ICCPROFILE values.
type ICCProfileValue twain.UInt16

const (
	// ICCProfileValueNone is an alias for TWIC_NONE.
	ICCProfileValueNone ICCProfileValue = iota

	// ICCProfileValueLink is an alias for TWIC_LINK.
	ICCProfileValueLink

	// ICCProfileValueEmbed is an alias for TWIC_EMBED.
	ICCProfileValueEmbed
)

// ImageFilterValue is an alias for ICAP_IMAGEFILTER values.
type ImageFilterValue twain.UInt16

const (
	// ImageFilterValueNone is an alias for TWIF_NONE.
	ImageFilterValueNone ImageFilterValue = iota

	// ImageFilterValueAuto is an alias for TWIF_AUTO.
	ImageFilterValueAuto

	// ImageFilterValueLowPass is an alias for TWIF_LOWPASS.
	ImageFilterValueLowPass

	// ImageFilterValueBandPass is an alias for TWIF_BANDPASS.
	ImageFilterValueBandPass

	// ImageFilterValueHighPass is an alias for TWIF_HIGHPASS.
	ImageFilterValueHighPass

	// ImageFilterValueText is an alias for TWIF_TEXT.
	ImageFilterValueText = ImageFilterValueBandPass

	// ImageFilterValueFineLine is an alias for TWIF_FINELINE.
	ImageFilterValueFineLine = ImageFilterValueHighPass
)

// ImageMergeValue is an alias for ICAP_IMAGEMERGE values.
type ImageMergeValue twain.UInt16

const (
	// ImageMergeValueNone is an alias for TWIM_NONE.
	ImageMergeValueNone ImageMergeValue = iota

	// ImageMergeValueFrontOnTop is an alias for TWIM_FRONTONTOP.
	ImageMergeValueFrontOnTop

	// ImageMergeValueFrontOnBottom is an alias for TWIM_FRONTONBOTTOM.
	ImageMergeValueFrontOnBottom

	// ImageMergeValueFrontOnLeft is an alias for TWIM_FRONTONLEFT.
	ImageMergeValueFrontOnLeft

	// ImageMergeValueFrontOnRight is an alias for TWIM_FRONTONRIGHT.
	ImageMergeValueFrontOnRight
)

// JobControlValue is an alias for CAP_JOBCONTROL values.
type JobControlValue twain.UInt16

const (
	// JobControlValueNone is an alias for TWJC_NONE.
	JobControlValueNone JobControlValue = iota

	// JobControlValueJSIC is an alias for TWJC_JSIC.
	JobControlValueJSIC

	// JobControlValueJSIS is an alias for TWJC_JSIS.
	JobControlValueJSIS

	// JobControlValueJSXC is an alias for TWJC_JSXC.
	JobControlValueJSXC

	// JobControlValueJSXS is an alias for TWJC_JSXS.
	JobControlValueJSXS
)

// JPEGQualityValue is an alias for ICAP_JPEGQUALITY values.
type JPEGQualityValue twain.UInt16

const (
	// JPEGQualityValueUnknown is an alias for TWJQ_UNKNOWN.
	JPEGQualityValueUnknown JPEGQualityValue = -4

	// JPEGQualityValueLow is an alias for TWJQ_LOW.
	JPEGQualityValueLow JPEGQualityValue = -3

	// JPEGQualityValueMedium is an alias for TWJQ_MEDIUM.
	JPEGQualityValueMedium JPEGQualityValue = -2

	// JPEGQualityValueHigh is an alias for TWJQ_HIGH.
	JPEGQualityValueHigh JPEGQualityValue = -1
)

// LightPathValue is an alias for ICAP_LIGHTPATH values.
type LightPathValue twain.UInt16

const (
	// LightPathValueReflective is an alias for TWLP_REFLECTIVE.
	LightPathValueReflective LightPathValue = iota

	// LightPathValueTransmissive is an alias for TWLP_TRANSMISSIVE.
	LightPathValueTransmissive
)

// LightSourceValue is an alias for ICAP_LIGHTSOURCE values.
type LightSourceValue twain.UInt16

const (
	// LightSourceValueRed is an alias for TWLS_RED.
	LightSourceValueRed LightSourceValue = iota

	// LightSourceValueGreen is an alias for TWLS_GREEN.
	LightSourceValueGreen

	// LightSourceValueBlue is an alias for TWLS_BLUE.
	LightSourceValueBlue

	// LightSourceValueNone is an alias for TWLS_NONE.
	LightSourceValueNone

	// LightSourceValueWhite is an alias for TWLS_WHITE.
	LightSourceValueWhite

	// LightSourceValueUV is an alias for TWLS_UV.
	LightSourceValueUV

	// LightSourceValueIR is an alias for TWLS_IR.
	LightSourceValueIR
)

// NoiseFilterValue is an alias for ICAP_NOISEFILTER values.
type NoiseFilterValue twain.UInt16

const (
	// NoiseFilterValueNone is an alias for TWNF_NONE.
	NoiseFilterValueNone NoiseFilterValue = iota

	// NoiseFilterValueAuto is an alias for TWNF_AUTO.
	NoiseFilterValueAuto

	// NoiseFilterValueLonePixel is an alias for TWNF_LONEPIXEL.
	NoiseFilterValueLonePixel

	// NoiseFilterValueMajorityRule is an alias for TWNF_MAJORITYRULE.
	NoiseFilterValueMajorityRule
)

// OrientationValue is an alias for ICAP_ORIENTATION values.
type OrientationValue twain.UInt16

const (
	// OrientationValueRot0 is an alias for TWOR_ROT0.
	OrientationValueRot0 OrientationValue = iota

	// OrientationValueRot90 is an alias for TWOR_ROT90.
	OrientationValueRot90

	// OrientationValueRot180 is an alias for TWOR_ROT180.
	OrientationValueRot180

	// OrientationValueRot270 is an alias for TWOR_ROT270.
	OrientationValueRot270

	// OrientationValueAuto is an alias for TWOR_AUTO.
	OrientationValueAuto

	// OrientationValueAutoText is an alias for TWOR_AUTOTEXT.
	OrientationValueAutoText

	// OrientationValueAutoPicture is an alias for TWOR_AUTOPICTURE.
	OrientationValueAutoPicture

	// OrientationValuePortrait is an alias for TWOR_PORTRAIT.
	OrientationValuePortrait = OrientationValueRot0

	// OrientationValueLandscape is an alias for TWOR_LANDSCAPE.
	OrientationValueLandscape = OrientationValueRot270
)

// OverScanValue is an alias for ICAP_OVERSCAN values.
type OverScanValue twain.UInt16

const (
	// OverScanValueNone is an alias for TWOV_NONE.
	OverScanValueNone OverScanValue = iota

	// OverScanValueAuto is an alias for TWOV_AUTO.
	OverScanValueAuto

	// OverScanValueTopBottom is an alias for TWOV_TOPBOTTOM.
	OverScanValueTopBottom

	// OverScanValueLeftRight is an alias for TWOV_LEFTRIGHT.
	OverScanValueLeftRight

	// OverScanValueAll is an alias for TWOV_ALL.
	OverScanValueAll
)

// PlanarChunkyValue is an alias for ICAP_PLANARCHUNKY values.
type PlanarChunkyValue twain.UInt16

const (
	// PlanarChunkyValueChunky is an alias for TWPC_CHUNKY.
	PlanarChunkyValueChunky PlanarChunkyValue = iota

	// PlanarChunkyValuePlanar is an alias for TWPC_PLANAR.
	PlanarChunkyValuePlanar
)

// PixelFlavorValue is an alias for ICAP_PIXELFLAVOR values.
type PixelFlavorValue twain.UInt16

const (
	// PixelFlavorValueChocolate is an alias for TWPF_CHOCOLATE.
	PixelFlavorValueChocolate PixelFlavorValue = iota

	// PixelFlavorValueVanilla is an alias for TWPF_VANILLA.
	PixelFlavorValueVanilla
)

// PrinterModeValue is an alias for CAP_PRINTERMODE values.
type PrinterModeValue twain.UInt16

const (
	// PrinterModeValueSingleString is an alias for TWPM_SINGLESTRING.
	PrinterModeValueSingleString PrinterModeValue = iota

	// PrinterModeValueMultiString is an alias for TWPM_MULTISTRING.
	PrinterModeValueMultiString

	// PrinterModeValueCompoundString is an alias for TWPM_COMPOUNDSTRING.
	PrinterModeValueCompoundString
)

// PrinterValue is an alias for CAP_PRINTER values.
type PrinterValue twain.UInt16

const (
	// PrinterValueImprinterTopBefore is an alias for TWPR_IMPRINTERTOPBEFORE.
	PrinterValueImprinterTopBefore PrinterValue = iota

	// PrinterValueImprinterTopAfter is an alias for TWPR_IMPRINTERTOPAFTER.
	PrinterValueImprinterTopAfter

	// PrinterValueImprinterBottomBefore is an alias for TWPR_IMPRINTERBOTTOMBEFORE.
	PrinterValueImprinterBottomBefore

	// PrinterValueImprinterBottomAfter is an alias for TWPR_IMPRINTERBOTTOMAFTER.
	PrinterValueImprinterBottomAfter

	// PrinterValueEndorserTopBefore is an alias for TWPR_ENDORSERTOPBEFORE.
	PrinterValueEndorserTopBefore

	// PrinterValueEndorserTopAfter is an alias for TWPR_ENDORSERTOPAFTER.
	PrinterValueEndorserTopAfter

	// PrinterValueEndorserBottomBefore is an alias for TWPR_ENDORSERBOTTOMBEFORE.
	PrinterValueEndorserBottomBefore

	// PrinterValueEndorserBottomAfter is an alias for TWPR_ENDORSERBOTTOMAFTER.
	PrinterValueEndorserBottomAfter
)

// PrinterFontStyleValue is an alias for CAP_PRINTERFONTSTYLE values.
type PrinterFontStyleValue twain.UInt16

const (
	// PrinterFontStyleValueNormal is an alias for TWPF_NORMAL.
	PrinterFontStyleValueNormal PrinterFontStyleValue = iota

	// PrinterFontStyleValueBold is an alias for TWPF_BOLD.
	PrinterFontStyleValueBold

	// PrinterFontStyleValueItalic is an alias for TWPF_ITALIC.
	PrinterFontStyleValueItalic

	// PrinterFontStyleValueLargeSize is an alias for TWPF_LARGESIZE.
	PrinterFontStyleValueLargeSize

	// PrinterFontStyleValueSmallSize is an alias for TWPF_SMALLSIZE.
	PrinterFontStyleValueSmallSize
)

// PrinterIndexTriggerValue is an alias for CAP_PRINTERINDEXTRIGGER values.
type PrinterIndexTriggerValue twain.UInt16

const (
	// PrinterIndexTriggerValuePage is an alias for TWCT_PAGE.
	PrinterIndexTriggerValuePage PrinterIndexTriggerValue = iota

	// PrinterIndexTriggerValuePatch1 is an alias for TWCT_PATCH1.
	PrinterIndexTriggerValuePatch1

	// PrinterIndexTriggerValuePatch2 is an alias for TWCT_PATCH2.
	PrinterIndexTriggerValuePatch2

	// PrinterIndexTriggerValuePatch3 is an alias for TWCT_PATCH3.
	PrinterIndexTriggerValuePatch3

	// PrinterIndexTriggerValuePatch4 is an alias for TWCT_PATCH4.
	PrinterIndexTriggerValuePatch4

	// PrinterIndexTriggerValuePatchT is an alias for TWCT_PATCHT.
	PrinterIndexTriggerValuePatchT

	// PrinterIndexTriggerValuePatch6 is an alias for TWCT_PATCH6.
	PrinterIndexTriggerValuePatch6
)

// PowerSupplyValue is an alias for CAP_POWERSUPPLY values.
type PowerSupplyValue twain.UInt16

const (
	// PowerSupplyValueExternal is an alias for TWPS_EXTERNAL.
	PowerSupplyValueExternal PowerSupplyValue = iota

	// PowerSupplyValueBattery is an alias for TWPS_BATTERY.
	PowerSupplyValueBattery
)

// PixelTypeValue is an alias for ICAP_PIXELTYPE values.
type PixelTypeValue twain.UInt16

const (
	// PixelTypeValueBW is an alias for TWPT_BW.
	PixelTypeValueBW PixelTypeValue = iota

	// PixelTypeValueGray is an alias for TWPT_GRAY.
	PixelTypeValueGray

	// PixelTypeValueRGB is an alias for TWPT_RGB.
	PixelTypeValueRGB

	// PixelTypeValuePalette is an alias for TWPT_PALETTE.
	PixelTypeValuePalette

	// PixelTypeValueCMY is an alias for TWPT_CMY.
	PixelTypeValueCMY

	// PixelTypeValueCMYK is an alias for TWPT_CMYK.
	PixelTypeValueCMYK

	// PixelTypeValueYUV is an alias for TWPT_YUV.
	PixelTypeValueYUV

	// PixelTypeValueYUVK is an alias for TWPT_YUVK.
	PixelTypeValueYUVK

	// PixelTypeValueCIEXYZ is an alias for TWPT_CIEXYZ.
	PixelTypeValueCIEXYZ

	// PixelTypeValueLAB is an alias for TWPT_LAB.
	PixelTypeValueLAB

	// PixelTypeValueSRGB is an alias for TWPT_SRGB.
	PixelTypeValueSRGB

	// PixelTypeValueSCRGB is an alias for TWPT_SCRGB.
	PixelTypeValueSCRGB

	// PixelTypeValueInfrared is an alias for TWPT_INFRARED.
	PixelTypeValueInfrared PixelTypeValue = 16
)

// SegmentedValue is an alias for CAP_SEGMENTED values.
type SegmentedValue twain.UInt16

const (
	// SegmentedValueNone is an alias for TWSG_NONE.
	SegmentedValueNone SegmentedValue = iota

	// SegmentedValueAuto is an alias for TWSG_AUTO.
	SegmentedValueAuto

	// SegmentedValueManual is an alias for TWSG_MANUAL.
	SegmentedValueManual
)

// FilmTypeValue is an alias for ICAP_FILMTYPE values.
type FilmTypeValue twain.UInt16

const (
	// FilmTypeValuePositive is an alias for TWFM_POSITIVE.
	FilmTypeValuePositive FilmTypeValue = iota

	// FilmTypeValueNegative is an alias for TWFM_NEGATIVE.
	FilmTypeValueNegative
)

// DoubleFeedDetectionValue is an alias for CAP_DOUBLEFEEDDETECTION values.
type DoubleFeedDetectionValue twain.UInt16

const (
	// DoubleFeedDetectionValueUltrasonic is an alias for TWDF_ULTRASONIC.
	DoubleFeedDetectionValueUltrasonic DoubleFeedDetectionValue = iota

	// DoubleFeedDetectionValueByLength is an alias for TWDF_BYLENGTH.
	DoubleFeedDetectionValueByLength

	// DoubleFeedDetectionValueInfrared is an alias for TWDF_INFRARED.
	DoubleFeedDetectionValueInfrared
)

// DoubleFeedDetectionSensitivityValue is an alias for CAP_DOUBLEFEEDDETECTIONSENSITIVITY values.
type DoubleFeedDetectionSensitivityValue twain.UInt16

const (
	// DoubleFeedDetectionSensitivityValueLow is an alias for TWUS_LOW.
	DoubleFeedDetectionSensitivityValueLow DoubleFeedDetectionSensitivityValue = iota

	// DoubleFeedDetectionSensitivityValueMedium is an alias for TWUS_MEDIUM.
	DoubleFeedDetectionSensitivityValueMedium

	// DoubleFeedDetectionSensitivityValueHigh is an alias for TWUS_HIGH.
	DoubleFeedDetectionSensitivityValueHigh
)

// DoubleFeedDetectionResponseValue is an alias for CAP_DOUBLEFEEDDETECTIONRESPONSE values.
type DoubleFeedDetectionResponseValue twain.UInt16

const (
	// DoubleFeedDetectionResponseValueStop is an alias for TWDP_STOP.
	DoubleFeedDetectionResponseValueStop DoubleFeedDetectionResponseValue = iota

	// DoubleFeedDetectionResponseValueStopAndWait is an alias for TWDP_STOPANDWAIT.
	DoubleFeedDetectionResponseValueStopAndWait

	// DoubleFeedDetectionResponseValueSound is an alias for TWDP_SOUND.
	DoubleFeedDetectionResponseValueSound

	// DoubleFeedDetectionResponseValueDoNotImprint is an alias for TWDP_DONOTIMPRINT.
	DoubleFeedDetectionResponseValueDoNotImprint
)

// MirrorValue is an alias for ICAP_MIRROR values.
type MirrorValue twain.UInt16

const (
	// MirrorValueNone is an alias for TWMR_NONE.
	MirrorValueNone MirrorValue = iota

	// MirrorValueVertical is an alias for TWMR_VERTICAL.
	MirrorValueVertical

	// MirrorValueHorizontal is an alias for TWMR_HORIZONTAL.
	MirrorValueHorizontal
)

// JPEGSubSamplingValue is an alias for ICAP_JPEGSUBSAMPLING values.
type JPEGSubSamplingValue twain.UInt16

const (
	// JPEGSubSamplingValue444YCBCR is an alias for TWJS_444YCBCR.
	JPEGSubSamplingValue444YCBCR JPEGSubSamplingValue = iota

	// JPEGSubSamplingValue444RGB is an alias for TWJS_444RGB.
	JPEGSubSamplingValue444RGB

	// JPEGSubSamplingValue422 is an alias for TWJS_422.
	JPEGSubSamplingValue422

	// JPEGSubSamplingValue421 is an alias for TWJS_421.
	JPEGSubSamplingValue421

	// JPEGSubSamplingValue411 is an alias for TWJS_411.
	JPEGSubSamplingValue411

	// JPEGSubSamplingValue420 is an alias for TWJS_420.
	JPEGSubSamplingValue420

	// JPEGSubSamplingValue410 is an alias for TWJS_410.
	JPEGSubSamplingValue410

	// JPEGSubSamplingValue311 is an alias for TWJS_311.
	JPEGSubSamplingValue311
)

// PaperHandlingValue is an alias for CAP_PAPERHANDLING values.
type PaperHandlingValue twain.UInt16

const (
	// PaperHandlingValueNormal is an alias for TWPH_NORMAL.
	PaperHandlingValueNormal PaperHandlingValue = iota

	// PaperHandlingValueFragile is an alias for TWPH_FRAGILE.
	PaperHandlingValueFragile

	// PaperHandlingValueThick is an alias for TWPH_THICK.
	PaperHandlingValueThick

	// PaperHandlingValueTrifold is an alias for TWPH_TRIFOLD.
	PaperHandlingValueTrifold

	// PaperHandlingValuePhotograph is an alias for TWPH_PHOTOGRAPH.
	PaperHandlingValuePhotograph
)

// IndicatorsModeValue is an alias for CAP_INDICATORSMODE values.
type IndicatorsModeValue twain.UInt16

const (
	// IndicatorsModeValueInfo is an alias for TWCI_INFO.
	IndicatorsModeValueInfo IndicatorsModeValue = iota

	// IndicatorsModeValueWarning is an alias for TWCI_WARNING.
	IndicatorsModeValueWarning

	// IndicatorsModeValueError is an alias for TWCI_ERROR.
	IndicatorsModeValueError

	// IndicatorsModeValueWarmup is an alias for TWCI_WARMUP.
	IndicatorsModeValueWarmup
)

// SupportedSizeValue is an alias for ICAP_SUPPORTEDSIZES values.
type SupportedSizeValue twain.UInt16

const (
	// SupportedSizeValueNone is an alias for TWSS_NONE.
	SupportedSizeValueNone SupportedSizeValue = iota

	// SupportedSizeValueA4 is an alias for TWSS_A4.
	SupportedSizeValueA4

	// SupportedSizeValueJISB5 is an alias for TWSS_JISB5.
	SupportedSizeValueJISB5

	// SupportedSizeValueUSLetter is an alias for TWSS_USLETTER.
	SupportedSizeValueUSLetter

	// SupportedSizeValueUSLegal is an alias for TWSS_USLEGAL.
	SupportedSizeValueUSLegal

	// SupportedSizeValueA5 is an alias for TWSS_A5.
	SupportedSizeValueA5

	// SupportedSizeValueISOB4 is an alias for TWSS_ISOB4.
	SupportedSizeValueISOB4

	// SupportedSizeValueISOB6 is an alias for TWSS_ISOB6.
	SupportedSizeValueISOB6

	// SupportedSizeValueUnused is needed because there is no size for 8.
	// Do not use this.
	SupportedSizeValueUnused

	// SupportedSizeValueUSLedger is an alias for TWSS_USLEDGER.
	SupportedSizeValueUSLedger

	// SupportedSizeValueUSExecutive is an alias for TWSS_USEXECUTIVE.
	SupportedSizeValueUSExecutive

	// SupportedSizeValueA3 is an alias for TWSS_A3.
	SupportedSizeValueA3

	// SupportedSizeValueISOB3 is an alias for TWSS_ISOB3.
	SupportedSizeValueISOB3

	// SupportedSizeValueA6 is an alias for TWSS_A6.
	SupportedSizeValueA6

	// SupportedSizeValueC4 is an alias for TWSS_C4.
	SupportedSizeValueC4

	// SupportedSizeValueC5 is an alias for TWSS_C5.
	SupportedSizeValueC5

	// SupportedSizeValueC6 is an alias for TWSS_C6.
	SupportedSizeValueC6

	// SupportedSizeValue4A0 is an alias for TWSS_4A0.
	SupportedSizeValue4A0

	// SupportedSizeValue2A0 is an alias for TWSS_2A0.
	SupportedSizeValue2A0

	// SupportedSizeValueA0 is an alias for TWSS_A0.
	SupportedSizeValueA0

	// SupportedSizeValueA1 is an alias for TWSS_A1.
	SupportedSizeValueA1

	// SupportedSizeValueA2 is an alias for TWSS_A2.
	SupportedSizeValueA2

	// SupportedSizeValueA7 is an alias for TWSS_A7.
	SupportedSizeValueA7

	// SupportedSizeValueA8 is an alias for TWSS_A8.
	SupportedSizeValueA8

	// SupportedSizeValueA9 is an alias for TWSS_A9.
	SupportedSizeValueA9

	// SupportedSizeValueA10 is an alias for TWSS_A10.
	SupportedSizeValueA10

	// SupportedSizeValueISOB0 is an alias for TWSS_ISOB0.
	SupportedSizeValueISOB0

	// SupportedSizeValueISOB1 is an alias for TWSS_ISOB1.
	SupportedSizeValueISOB1

	// SupportedSizeValueISOB2 is an alias for TWSS_ISOB2.
	SupportedSizeValueISOB2

	// SupportedSizeValueISOB5 is an alias for TWSS_ISOB5.
	SupportedSizeValueISOB5

	// SupportedSizeValueISOB7 is an alias for TWSS_ISOB7.
	SupportedSizeValueISOB7

	// SupportedSizeValueISOB8 is an alias for TWSS_ISOB8.
	SupportedSizeValueISOB8

	// SupportedSizeValueISOB9 is an alias for TWSS_ISOB9.
	SupportedSizeValueISOB9

	// SupportedSizeValueISOB10 is an alias for TWSS_ISOB10.
	SupportedSizeValueISOB10

	// SupportedSizeValueJISB0 is an alias for TWSS_JISB0.
	SupportedSizeValueJISB0

	// SupportedSizeValueJISB1 is an alias for TWSS_JISB1.
	SupportedSizeValueJISB1

	// SupportedSizeValueJISB2 is an alias for TWSS_JISB2.
	SupportedSizeValueJISB2

	// SupportedSizeValueJISB3 is an alias for TWSS_JISB3.
	SupportedSizeValueJISB3

	// SupportedSizeValueJISB4 is an alias for TWSS_JISB4.
	SupportedSizeValueJISB4

	// SupportedSizeValueJISB6 is an alias for TWSS_JISB6.
	SupportedSizeValueJISB6

	// SupportedSizeValueJISB7 is an alias for TWSS_JISB7.
	SupportedSizeValueJISB7

	// SupportedSizeValueJISB8 is an alias for TWSS_JISB8.
	SupportedSizeValueJISB8

	// SupportedSizeValueJISB9 is an alias for TWSS_JISB9.
	SupportedSizeValueJISB9

	// SupportedSizeValueJISB10 is an alias for TWSS_JISB10.
	SupportedSizeValueJISB10

	// SupportedSizeValueC0 is an alias for TWSS_C0.
	SupportedSizeValueC0

	// SupportedSizeValueC1 is an alias for TWSS_C1.
	SupportedSizeValueC1

	// SupportedSizeValueC2 is an alias for TWSS_C2.
	SupportedSizeValueC2

	// SupportedSizeValueC3 is an alias for TWSS_C3.
	SupportedSizeValueC3

	// SupportedSizeValueC7 is an alias for TWSS_C7.
	SupportedSizeValueC7

	// SupportedSizeValueC8 is an alias for TWSS_C8.
	SupportedSizeValueC8

	// SupportedSizeValueC9 is an alias for TWSS_C9.
	SupportedSizeValueC9

	// SupportedSizeValueC10 is an alias for TWSS_C10.
	SupportedSizeValueC10

	// SupportedSizeValueUSStatement is an alias for TWSS_USSTATEMENT.
	SupportedSizeValueUSStatement

	// SupportedSizeValueBusinessCard is an alias for TWSS_BUSINESSCARD.
	SupportedSizeValueBusinessCard

	// SupportedSizeValueMaxSize is an alias for TWSS_MAXSIZE.
	SupportedSizeValueMaxSize
)

// TransferMechanismValue is an alias for ICAP_XFERMECH values.
type TransferMechanismValue twain.UInt16

const (
	// TransferMechanismValueNative is an alias for TWSX_NATIVE.
	TransferMechanismValueNative TransferMechanismValue = iota

	// TransferMechanismValueFile is an alias for TWSX_FILE.
	TransferMechanismValueFile

	// TransferMechanismValueMemory is an alias for TWSX_MEMORY.
	TransferMechanismValueMemory

	// TransferMechanismValueMemFile is an alias for TWSX_MEMFILE.
	TransferMechanismValueMemFile
)

// UnitsValue is an alias for ICAP_UNITS values.
type UnitsValue twain.UInt16

const (
	// UnitsValueInches is an alias for TWUN_INCHES.
	UnitsValueInches UnitsValue = iota

	// UnitsValueCentimeters is an alias for TWUN_CENTIMETERS.
	UnitsValueCentimeters

	// UnitsValuePicas is an alias for TWUN_PICAS.
	UnitsValuePicas

	// UnitsValuePoints is an alias for TWUN_POINTS.
	UnitsValuePoints

	// UnitsValueTwips is an alias for TWUN_TWIPS.
	UnitsValueTwips

	// UnitsValuePixels is an alias for TWUN_PIXELS.
	UnitsValuePixels

	// UnitsValueMillimeters is an alias for TWUN_MILLIMETERS.
	UnitsValueMillimeters
)
