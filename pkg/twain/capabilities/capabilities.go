package capabilities

import "github.com/mikerourke/go-twain/pkg/twain"

type Capability twain.UInt16

// CustomBase is the base of custom capabilities. It is an alias for
// CAP_CUSTOMBASE.
const CustomBase Capability = 0x8000

// TransferCount is an alias for CAP_XFERCOUNT. All data sources are required
// to support this capability.
const TransferCount Capability = 0x8000

// Image data sources are _required_ to support these capabilities.
const (
	// Compression is an alias for ICAP_COMPRESSION.
	Compression Capability = 0x0100

	// PixelType is an alias for ICAP_PIXELTYPE.
	PixelType Capability = 0x0101

	// Units is an alias for ICAP_UNITS.
	Units Capability = 0x0102

	// TransferMechanism is an alias for ICAP_XFERMECH.
	TransferMechanism Capability = 0x0103
)

// All data sources _may_ support these capabilities.
const (
	// Author is an alias for CAP_AUTHOR.
	Author Capability = 0x1000

	// Caption is an alias for CAP_CAPTION.
	Caption Capability = 0x1001

	// FeederEnabled is an alias for CAP_FEEDERENABLED.
	FeederEnabled Capability = 0x1002

	// FeederLoaded is an alias for CAP_FEEDERLOADED.
	FeederLoaded Capability = 0x1003

	// TimeDate is an alias for CAP_TIMEDATE.
	TimeDate Capability = 0x1004

	// SupportedCaps is an alias for CAP_SUPPORTEDCAPS.
	SupportedCaps Capability = 0x1005

	// ExtendedCaps is an alias for CAP_EXTENDEDCAPS.
	ExtendedCaps Capability = 0x1006

	// AutoFeed is an alias for CAP_AUTOFEED.
	AutoFeed Capability = 0x1007

	// ClearPage is an alias for CAP_CLEARPAGE.
	ClearPage Capability = 0x1008

	// FeedPage is an alias for CAP_FEEDPAGE.
	FeedPage Capability = 0x1009

	// RewindPage is an alias for CAP_REWINDPAGE.
	RewindPage Capability = 0x100a

	// Indicators is an alias for CAP_INDICATORS.
	Indicators Capability = 0x100b

	// PaperDetectable is an alias for CAP_PAPERDETECTABLE.
	PaperDetectable Capability = 0x100d

	// UIControllable is an alias for CAP_UICONTROLLABLE.
	UIControllable Capability = 0x100e

	// DeviceOnline is an alias for CAP_DEVICEONLINE.
	DeviceOnline Capability = 0x100f

	// AutoScan is an alias for CAP_AUTOSCAN.
	AutoScan Capability = 0x1010

	// ThumbnailsEnabled is an alias for CAP_THUMBNAILSENABLED.
	ThumbnailsEnabled Capability = 0x1011

	// Duplex is an alias for CAP_DUPLEX.
	Duplex Capability = 0x1012

	// DuplexEnabled is an alias for CAP_DUPLEXENABLED.
	DuplexEnabled Capability = 0x1013

	// EnabledSUIOnly is an alias for CAP_ENABLEDSUIONLY.
	EnabledSUIOnly Capability = 0x1014

	// CustomdsData is an alias for CAP_CUSTOMDSDATA.
	CustomdsData Capability = 0x1015

	// Endorser is an alias for CAP_ENDORSER.
	Endorser Capability = 0x1016

	// JobControl is an alias for CAP_JOBCONTROL.
	JobControl Capability = 0x1017

	// Alarms is an alias for CAP_ALARMS.
	Alarms Capability = 0x1018

	// AlarmVolume is an alias for CAP_ALARMVOLUME.
	AlarmVolume Capability = 0x1019

	// AutomaticCapture is an alias for CAP_AUTOMATICCAPTURE.
	AutomaticCapture Capability = 0x101a

	// TimeBeforeFirstCapture is an alias for CAP_TIMEBEFOREFIRSTCAPTURE.
	TimeBeforeFirstCapture Capability = 0x101b

	// TimeBetweenCaptures is an alias for CAP_TIMEBETWEENCAPTURES.
	TimeBetweenCaptures Capability = 0x101c

	// MaxBatchBuffers is an alias for CAP_MAXBATCHBUFFERS.
	MaxBatchBuffers Capability = 0x101e

	// DeviceTimeDate is an alias for CAP_DEVICETIMEDATE.
	DeviceTimeDate Capability = 0x101f

	// PowerSupply is an alias for CAP_POWERSUPPLY.
	PowerSupply Capability = 0x1020

	// CameraPreviewUI is an alias for CAP_CAMERAPREVIEWUI.
	CameraPreviewUI Capability = 0x1021

	// DeviceEvent is an alias for CAP_DEVICEEVENT.
	DeviceEvent Capability = 0x1022

	// SerialNumber is an alias for CAP_SERIALNUMBER.
	SerialNumber Capability = 0x1024

	// Printer is an alias for CAP_PRINTER.
	Printer Capability = 0x1026

	// PrinterEnabled is an alias for CAP_PRINTERENABLED.
	PrinterEnabled Capability = 0x1027

	// PrinterIndex is an alias for CAP_PRINTERINDEX.
	PrinterIndex Capability = 0x1028

	// PrinterMode is an alias for CAP_PRINTERMODE.
	PrinterMode Capability = 0x1029

	// PrinterString is an alias for CAP_PRINTERSTRING.
	PrinterString Capability = 0x102a

	// PrinterSuffix is an alias for CAP_PRINTERSUFFIX.
	PrinterSuffix Capability = 0x102b

	// Language is an alias for CAP_LANGUAGE.
	Language Capability = 0x102c

	// FeederAlignment is an alias for CAP_FEEDERALIGNMENT.
	FeederAlignment Capability = 0x102d

	// FeederOrder is an alias for CAP_FEEDERORDER.
	FeederOrder Capability = 0x102e

	// ReacquireAllowed is an alias for CAP_REACQUIREALLOWED.
	ReacquireAllowed Capability = 0x1030

	// BatteryMinutes is an alias for CAP_BATTERYMINUTES.
	BatteryMinutes Capability = 0x1032

	// BatteryPercentage is an alias for CAP_BATTERYPERCENTAGE.
	BatteryPercentage Capability = 0x1033

	// CameraSide is an alias for CAP_CAMERASIDE.
	CameraSide Capability = 0x1034

	// Segmented is an alias for CAP_SEGMENTED.
	Segmented Capability = 0x1035

	// CameraEnabled is an alias for CAP_CAMERAENABLED.
	CameraEnabled Capability = 0x1036

	// CameraOrder is an alias for CAP_CAMERAORDER.
	CameraOrder Capability = 0x1037

	// MICREnabled is an alias for CAP_MICRENABLED.
	MICREnabled Capability = 0x1038

	// FeederPrep is an alias for CAP_FEEDERPREP.
	FeederPrep Capability = 0x1039

	// FeederPocket is an alias for CAP_FEEDERPOCKET.
	FeederPocket Capability = 0x103a

	// AutomaticSenseMedium is an alias for CAP_AUTOMATICSENSEMEDIUM.
	AutomaticSenseMedium Capability = 0x103b

	// CustomInterfaceGUID is an alias for CAP_CUSTOMINTERFACEGUID.
	CustomInterfaceGUID Capability = 0x103c

	// SupportedCapsSegmentUnique is an alias for CAP_SUPPORTEDCAPSSEGMENTUNIQUE.
	SupportedCapsSegmentUnique Capability = 0x103d

	// SupportedDATs is an alias for CAP_SUPPORTEDDATS.
	SupportedDATs Capability = 0x103e

	// DoubleFeedDetection is an alias for CAP_DOUBLEFEEDDETECTION.
	DoubleFeedDetection Capability = 0x103f

	// DoubleFeedDetectionLength is an alias for CAP_DOUBLEFEEDDETECTIONLENGTH.
	DoubleFeedDetectionLength Capability = 0x1040

	// DoubleFeedDetectionSensitivity is an alias for CAP_DOUBLEFEEDDETECTIONSENSITIVITY.
	DoubleFeedDetectionSensitivity Capability = 0x1041

	// DoubleFeedDetectionResponse is an alias for CAP_DOUBLEFEEDDETECTIONRESPONSE.
	DoubleFeedDetectionResponse Capability = 0x1042

	// PaperHandling is an alias for CAP_PAPERHANDLING.
	PaperHandling Capability = 0x1043

	// IndicatorsMode is an alias for CAP_INDICATORSMODE.
	IndicatorsMode Capability = 0x1044

	// PrinterVerticalOffset is an alias for CAP_PRINTERVERTICALOFFSET.
	PrinterVerticalOffset Capability = 0x1045

	// PowerSaveTime is an alias for CAP_POWERSAVETIME.
	PowerSaveTime Capability = 0x1046

	// PrinterCharRotation is an alias for CAP_PRINTERCHARROTATION.
	PrinterCharRotation Capability = 0x1047

	// PrinterFontStyle is an alias for CAP_PRINTERFONTSTYLE.
	PrinterFontStyle Capability = 0x1048

	// PrinterIndexLeadChar is an alias for CAP_PRINTERINDEXLEADCHAR.
	PrinterIndexLeadChar Capability = 0x1049

	// PrinterIndexMaxValue is an alias for CAP_PRINTERINDEXMAXVALUE.
	PrinterIndexMaxValue Capability = 0x104A

	// PrinterIndexNumDigits is an alias for CAP_PRINTERINDEXNUMDIGITS.
	PrinterIndexNumDigits Capability = 0x104B

	// PrinterIndexStep is an alias for CAP_PRINTERINDEXSTEP.
	PrinterIndexStep Capability = 0x104C

	// PrinterIndexTrigger is an alias for CAP_PRINTERINDEXTRIGGER.
	PrinterIndexTrigger Capability = 0x104D

	// PrinterStringPreview is an alias for CAP_PRINTERSTRINGPREVIEW.
	PrinterStringPreview Capability = 0x104E

	// SheetCount is an alias for CAP_SHEETCOUNT.
	SheetCount Capability = 0x104F
)

// Image data sources _may_ support these capabilities.
const (
	// ImageAutoBright is an alias for ICAP_AUTOBRIGHT.
	ImageAutoBright Capability = 0x1100

	// ImageBrightness is an alias for ICAP_BRIGHTNESS.
	ImageBrightness Capability = 0x1101

	// ImageContrast is an alias for ICAP_CONTRAST.
	ImageContrast Capability = 0x1103

	// ImageCustHalftone is an alias for ICAP_CUSTHALFTONE.
	ImageCustHalftone Capability = 0x1104

	// ImageExposureTime is an alias for ICAP_EXPOSURETIME.
	ImageExposureTime Capability = 0x1105

	// ImageFilter is an alias for ICAP_FILTER.
	ImageFilter Capability = 0x1106

	// ImageFlashUsed is an alias for ICAP_FLASHUSED.
	ImageFlashUsed Capability = 0x1107

	// ImageGamma is an alias for ICAP_GAMMA.
	ImageGamma Capability = 0x1108

	// ImageHalftones is an alias for ICAP_HALFTONES.
	ImageHalftones Capability = 0x1109

	// ImageHighlight is an alias for ICAP_HIGHLIGHT.
	ImageHighlight Capability = 0x110a

	// ImageImageFileFormat is an alias for ICAP_IMAGEFILEFORMAT.
	ImageImageFileFormat Capability = 0x110c

	// ImageLampState is an alias for ICAP_LAMPSTATE.
	ImageLampState Capability = 0x110d

	// ImageLightSource is an alias for ICAP_LIGHTSOURCE.
	ImageLightSource Capability = 0x110e

	// ImageOrientation is an alias for ICAP_ORIENTATION.
	ImageOrientation Capability = 0x1110

	// ImagePhysicalWidth is an alias for ICAP_PHYSICALWIDTH.
	ImagePhysicalWidth Capability = 0x1111

	// ImagePhysicalHeight is an alias for ICAP_PHYSICALHEIGHT.
	ImagePhysicalHeight Capability = 0x1112

	// ImageShadow is an alias for ICAP_SHADOW.
	ImageShadow Capability = 0x1113

	// ImageFrames is an alias for ICAP_FRAMES.
	ImageFrames Capability = 0x1114

	// ImageXNativeResolution is an alias for ICAP_XNATIVERESOLUTION.
	ImageXNativeResolution Capability = 0x1116

	// ImageYNativeResolution is an alias for ICAP_YNATIVERESOLUTION.
	ImageYNativeResolution Capability = 0x1117

	// ImageXResolution is an alias for ICAP_XRESOLUTION.
	ImageXResolution Capability = 0x1118

	// ImageYResolution is an alias for ICAP_YRESOLUTION.
	ImageYResolution Capability = 0x1119

	// ImageMaxFrames is an alias for ICAP_MAXFRAMES.
	ImageMaxFrames Capability = 0x111a

	// ImageTiles is an alias for ICAP_TILES.
	ImageTiles Capability = 0x111b

	// ImageBitOrder is an alias for ICAP_BITORDER.
	ImageBitOrder Capability = 0x111c

	// ImageCCITTKFactor is an alias for ICAP_CCITTKFACTOR.
	ImageCCITTKFactor Capability = 0x111d

	// ImageLightPath is an alias for ICAP_LIGHTPATH.
	ImageLightPath Capability = 0x111e

	// ImagePixelFlavor is an alias for ICAP_PIXELFLAVOR.
	ImagePixelFlavor Capability = 0x111f

	// ImagePlanarChunky is an alias for ICAP_PLANARCHUNKY.
	ImagePlanarChunky Capability = 0x1120

	// ImageRotation is an alias for ICAP_ROTATION.
	ImageRotation Capability = 0x1121

	// ImageSupportedSizes is an alias for ICAP_SUPPORTEDSIZES.
	ImageSupportedSizes Capability = 0x1122

	// ImageThreshold is an alias for ICAP_THRESHOLD.
	ImageThreshold Capability = 0x1123

	// ImageXScaling is an alias for ICAP_XSCALING.
	ImageXScaling Capability = 0x1124

	// ImageYScaling is an alias for ICAP_YSCALING.
	ImageYScaling Capability = 0x1125

	// ImageBitOrderCodes is an alias for ICAP_BITORDERCODES.
	ImageBitOrderCodes Capability = 0x1126

	// ImagePixelFlavorCodes is an alias for ICAP_PIXELFLAVORCODES.
	ImagePixelFlavorCodes Capability = 0x1127

	// ImageJPEGPixelType is an alias for ICAP_JPEGPIXELTYPE.
	ImageJPEGPixelType Capability = 0x1128

	// ImageTimeFill is an alias for ICAP_TIMEFILL.
	ImageTimeFill Capability = 0x112a

	// ImageBitDepth is an alias for ICAP_BITDEPTH.
	ImageBitDepth Capability = 0x112b

	// ImageBitDepthReduction is an alias for ICAP_BITDEPTHREDUCTION.
	ImageBitDepthReduction Capability = 0x112c

	// ImageUndefinedImageSize is an alias for ICAP_UNDEFINEDIMAGESIZE.
	ImageUndefinedImageSize Capability = 0x112d

	// ImageImageDataset is an alias for ICAP_IMAGEDATASET.
	ImageImageDataset Capability = 0x112e

	// ImageExtImageInfo is an alias for ICAP_EXTIMAGEINFO.
	ImageExtImageInfo Capability = 0x112f

	// ImageMinimumHeight is an alias for ICAP_MINIMUMHEIGHT.
	ImageMinimumHeight Capability = 0x1130

	// ImageMinimumWidth is an alias for ICAP_MINIMUMWIDTH.
	ImageMinimumWidth Capability = 0x1131

	// ImageAutoDiscardBlankPages is an alias for ICAP_AUTODISCARDBLANKPAGES.
	ImageAutoDiscardBlankPages Capability = 0x1134

	// ImageFlipRotation is an alias for ICAP_FLIPROTATION.
	ImageFlipRotation Capability = 0x1136

	// ImageBarcodeDetectionEnabled is an alias for ICAP_BARCODEDETECTIONENABLED.
	ImageBarcodeDetectionEnabled Capability = 0x1137

	// ImageSupportedBarcodeTypes is an alias for ICAP_SUPPORTEDBARCODETYPES.
	ImageSupportedBarcodeTypes Capability = 0x1138

	// ImageBarcodeMaxSearchPriorities is an alias for ICAP_BARCODEMAXSEARCHPRIORITIES.
	ImageBarcodeMaxSearchPriorities Capability = 0x1139

	// ImageBarcodeSearchPriorities is an alias for ICAP_BARCODESEARCHPRIORITIES.
	ImageBarcodeSearchPriorities Capability = 0x113a

	// ImageBarcodeSearchMode is an alias for ICAP_BARCODESEARCHMODE.
	ImageBarcodeSearchMode Capability = 0x113b

	// ImageBarcodeMaxRetries is an alias for ICAP_BARCODEMAXRETRIES.
	ImageBarcodeMaxRetries Capability = 0x113c

	// ImageBarcodeTimeout is an alias for ICAP_BARCODETIMEOUT.
	ImageBarcodeTimeout Capability = 0x113d

	// ImageZoomFactor is an alias for ICAP_ZOOMFACTOR.
	ImageZoomFactor Capability = 0x113e

	// ImagePatchCodeDetectionEnabled is an alias for ICAP_PATCHCODEDETECTIONENABLED.
	ImagePatchCodeDetectionEnabled Capability = 0x113f

	// ImageSupportedPatchCodeTypes is an alias for ICAP_SUPPORTEDPATCHCODETYPES.
	ImageSupportedPatchCodeTypes Capability = 0x1140

	// ImagePatchCodeMaxSearchPriorities is an alias for ICAP_PATCHCODEMAXSEARCHPRIORITIES.
	ImagePatchCodeMaxSearchPriorities Capability = 0x1141

	// ImagePatchCodeSearchPriorities is an alias for ICAP_PATCHCODESEARCHPRIORITIES.
	ImagePatchCodeSearchPriorities Capability = 0x1142

	// ImagePatchCodeSearchMode is an alias for ICAP_PATCHCODESEARCHMODE.
	ImagePatchCodeSearchMode Capability = 0x1143

	// ImagePatchCodeMaxRetries is an alias for ICAP_PATCHCODEMAXRETRIES.
	ImagePatchCodeMaxRetries Capability = 0x1144

	// ImagePatchCodeTimeout is an alias for ICAP_PATCHCODETIMEOUT.
	ImagePatchCodeTimeout Capability = 0x1145

	// ImageFlashUsed2 is an alias for ICAP_FLASHUSED2.
	ImageFlashUsed2 Capability = 0x1146

	// ImageImageFilter is an alias for ICAP_IMAGEFILTER.
	ImageImageFilter Capability = 0x1147

	// ImageNoiseFilter is an alias for ICAP_NOISEFILTER.
	ImageNoiseFilter Capability = 0x1148

	// ImageOverScan is an alias for ICAP_OVERSCAN.
	ImageOverScan Capability = 0x1149

	// ImageAutomaticBorderDetection is an alias for ICAP_AUTOMATICBORDERDETECTION.
	ImageAutomaticBorderDetection Capability = 0x1150

	// ImageAutomaticDESkew is an alias for ICAP_AUTOMATICDESKEW.
	ImageAutomaticDESkew Capability = 0x1151

	// ImageAutomaticRotate is an alias for ICAP_AUTOMATICROTATE.
	ImageAutomaticRotate Capability = 0x1152

	// ImageJPEGQuality is an alias for ICAP_JPEGQUALITY.
	ImageJPEGQuality Capability = 0x1153

	// ImageFeederType is an alias for ICAP_FEEDERTYPE.
	ImageFeederType Capability = 0x1154

	// ImageICCProfile is an alias for ICAP_ICCPROFILE.
	ImageICCProfile Capability = 0x1155

	// ImageAutoSize is an alias for ICAP_AUTOSIZE.
	ImageAutoSize Capability = 0x1156

	// ImageAutomaticCropUsesFrame is an alias for ICAP_AUTOMATICCROPUSESFRAME.
	ImageAutomaticCropUsesFrame Capability = 0x1157

	// ImageAutomaticLengthDetection is an alias for ICAP_AUTOMATICLENGTHDETECTION.
	ImageAutomaticLengthDetection Capability = 0x1158

	// ImageAutomaticColorEnabled is an alias for ICAP_AUTOMATICCOLORENABLED.
	ImageAutomaticColorEnabled Capability = 0x1159

	// ImageAutomaticColorNonColorPixelType is an alias for ICAP_AUTOMATICCOLORNONCOLORPIXELTYPE.
	ImageAutomaticColorNonColorPixelType Capability = 0x115a

	// ImageColorManagementEnabled is an alias for ICAP_COLORMANAGEMENTENABLED.
	ImageColorManagementEnabled Capability = 0x115b

	// ImageImageMerge is an alias for ICAP_IMAGEMERGE.
	ImageImageMerge Capability = 0x115c

	// ImageImageMergeHeightThreshold is an alias for ICAP_IMAGEMERGEHEIGHTTHRESHOLD.
	ImageImageMergeHeightThreshold Capability = 0x115d

	// ImageSupportedExtImageInfo is an alias for ICAP_SUPPORTEDEXTIMAGEINFO.
	ImageSupportedExtImageInfo Capability = 0x115e

	// ImageFilmType is an alias for ICAP_FILMTYPE.
	ImageFilmType Capability = 0x115f

	// ImageMirror is an alias for ICAP_MIRROR.
	ImageMirror Capability = 0x1160

	// ImageJPEGSubsampling is an alias for ICAP_JPEGSUBSAMPLING.
	ImageJPEGSubsampling Capability = 0x1161
)

// AudioTransferMechanism is an alias for ACAP_XFERMECH. Image data sources
// _may_ support this audio capability.
const AudioTransferMechanism Capability = 0x1202
