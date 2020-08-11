// Package capabilities provides the names and values associated with
// capabilities (prefixed with CAP_ and ICAP_) from the TWAIN header file.
package capabilities

import "github.com/mikerourke/go-twain/pkg/twain"

type Capability twain.UInt16

// UseCustomBase is the base of custom capabilities. It is an alias for
// CAP_CUSTOMBASE.
const UseCustomBase Capability = 0x8000

// UseTransferCount is an alias for CAP_XFERCOUNT. All data sources are required
// to support this capability.
const UseTransferCount Capability = 0x8000

// Image data sources are _required_ to support these capabilities.
const (
	// UseCompression is an alias for ICAP_COMPRESSION.
	UseCompression Capability = 0x0100

	// UsePixelType is an alias for ICAP_PIXELTYPE.
	UsePixelType Capability = 0x0101

	// UseUnits is an alias for ICAP_UNITS.
	UseUnits Capability = 0x0102

	// UseTransferMechanism is an alias for ICAP_XFERMECH.
	UseTransferMechanism Capability = 0x0103
)

// All data sources _may_ support these capabilities.
const (
	// UseAuthor is an alias for CAP_AUTHOR.
	UseAuthor Capability = 0x1000

	// UseCaption is an alias for CAP_CAPTION.
	UseCaption Capability = 0x1001

	// UseFeederEnabled is an alias for CAP_FEEDERENABLED.
	UseFeederEnabled Capability = 0x1002

	// UseFeederLoaded is an alias for CAP_FEEDERLOADED.
	UseFeederLoaded Capability = 0x1003

	// UseTimeDate is an alias for CAP_TIMEDATE.
	UseTimeDate Capability = 0x1004

	// UseSupportedCaps is an alias for CAP_SUPPORTEDCAPS.
	UseSupportedCaps Capability = 0x1005

	// UseExtendedCaps is an alias for CAP_EXTENDEDCAPS.
	UseExtendedCaps Capability = 0x1006

	// UseAutoFeed is an alias for CAP_AUTOFEED.
	UseAutoFeed Capability = 0x1007

	// UseClearPage is an alias for CAP_CLEARPAGE.
	UseClearPage Capability = 0x1008

	// UseFeedPage is an alias for CAP_FEEDPAGE.
	UseFeedPage Capability = 0x1009

	// UseRewindPage is an alias for CAP_REWINDPAGE.
	UseRewindPage Capability = 0x100a

	// UseIndicators is an alias for CAP_INDICATORS.
	UseIndicators Capability = 0x100b

	// UsePaperDetectable is an alias for CAP_PAPERDETECTABLE.
	UsePaperDetectable Capability = 0x100d

	// UseUIControllable is an alias for CAP_UICONTROLLABLE.
	UseUIControllable Capability = 0x100e

	// UseDeviceOnline is an alias for CAP_DEVICEONLINE.
	UseDeviceOnline Capability = 0x100f

	// UseAutoScan is an alias for CAP_AUTOSCAN.
	UseAutoScan Capability = 0x1010

	// UseThumbnailsEnabled is an alias for CAP_THUMBNAILSENABLED.
	UseThumbnailsEnabled Capability = 0x1011

	// UseDuplex is an alias for CAP_DUPLEX.
	UseDuplex Capability = 0x1012

	// UseDuplexEnabled is an alias for CAP_DUPLEXENABLED.
	UseDuplexEnabled Capability = 0x1013

	// UseEnabledSUIOnly is an alias for CAP_ENABLEDSUIONLY.
	UseEnabledSUIOnly Capability = 0x1014

	// UseCustomdsData is an alias for CAP_CUSTOMDSDATA.
	UseCustomdsData Capability = 0x1015

	// UseEndorser is an alias for CAP_ENDORSER.
	UseEndorser Capability = 0x1016

	// UseJobControl is an alias for CAP_JOBCONTROL.
	UseJobControl Capability = 0x1017

	// UseAlarms is an alias for CAP_ALARMS.
	UseAlarms Capability = 0x1018

	// UseAlarmVolume is an alias for CAP_ALARMVOLUME.
	UseAlarmVolume Capability = 0x1019

	// UseAutomaticCapture is an alias for CAP_AUTOMATICCAPTURE.
	UseAutomaticCapture Capability = 0x101a

	// UseTimeBeforeFirstCapture is an alias for CAP_TIMEBEFOREFIRSTCAPTURE.
	UseTimeBeforeFirstCapture Capability = 0x101b

	// UseTimeBetweenCaptures is an alias for CAP_TIMEBETWEENCAPTURES.
	UseTimeBetweenCaptures Capability = 0x101c

	// UseMaxBatchBuffers is an alias for CAP_MAXBATCHBUFFERS.
	UseMaxBatchBuffers Capability = 0x101e

	// UseDeviceTimeDate is an alias for CAP_DEVICETIMEDATE.
	UseDeviceTimeDate Capability = 0x101f

	// UsePowerSupply is an alias for CAP_POWERSUPPLY.
	UsePowerSupply Capability = 0x1020

	// UseCameraPreviewUI is an alias for CAP_CAMERAPREVIEWUI.
	UseCameraPreviewUI Capability = 0x1021

	// UseDeviceEvent is an alias for CAP_DEVICEEVENT.
	UseDeviceEvent Capability = 0x1022

	// UseSerialNumber is an alias for CAP_SERIALNUMBER.
	UseSerialNumber Capability = 0x1024

	// UsePrinter is an alias for CAP_PRINTER.
	UsePrinter Capability = 0x1026

	// UsePrinterEnabled is an alias for CAP_PRINTERENABLED.
	UsePrinterEnabled Capability = 0x1027

	// UsePrinterIndex is an alias for CAP_PRINTERINDEX.
	UsePrinterIndex Capability = 0x1028

	// UsePrinterMode is an alias for CAP_PRINTERMODE.
	UsePrinterMode Capability = 0x1029

	// UsePrinterString is an alias for CAP_PRINTERSTRING.
	UsePrinterString Capability = 0x102a

	// UsePrinterSuffix is an alias for CAP_PRINTERSUFFIX.
	UsePrinterSuffix Capability = 0x102b

	// UseLanguage is an alias for CAP_LANGUAGE.
	UseLanguage Capability = 0x102c

	// UseFeederAlignment is an alias for CAP_FEEDERALIGNMENT.
	UseFeederAlignment Capability = 0x102d

	// UseFeederOrder is an alias for CAP_FEEDERORDER.
	UseFeederOrder Capability = 0x102e

	// UseReacquireAllowed is an alias for CAP_REACQUIREALLOWED.
	UseReacquireAllowed Capability = 0x1030

	// UseBatteryMinutes is an alias for CAP_BATTERYMINUTES.
	UseBatteryMinutes Capability = 0x1032

	// UseBatteryPercentage is an alias for CAP_BATTERYPERCENTAGE.
	UseBatteryPercentage Capability = 0x1033

	// UseCameraSide is an alias for CAP_CAMERASIDE.
	UseCameraSide Capability = 0x1034

	// UseSegmented is an alias for CAP_SEGMENTED.
	UseSegmented Capability = 0x1035

	// UseCameraEnabled is an alias for CAP_CAMERAENABLED.
	UseCameraEnabled Capability = 0x1036

	// UseCameraOrder is an alias for CAP_CAMERAORDER.
	UseCameraOrder Capability = 0x1037

	// UseMICREnabled is an alias for CAP_MICRENABLED.
	UseMICREnabled Capability = 0x1038

	// UseFeederPrep is an alias for CAP_FEEDERPREP.
	UseFeederPrep Capability = 0x1039

	// UseFeederPocket is an alias for CAP_FEEDERPOCKET.
	UseFeederPocket Capability = 0x103a

	// UseAutomaticSenseMedium is an alias for CAP_AUTOMATICSENSEMEDIUM.
	UseAutomaticSenseMedium Capability = 0x103b

	// UseCustomInterfaceGUID is an alias for CAP_CUSTOMINTERFACEGUID.
	UseCustomInterfaceGUID Capability = 0x103c

	// UseSupportedCapsSegmentUnique is an alias for CAP_SUPPORTEDCAPSSEGMENTUNIQUE.
	UseSupportedCapsSegmentUnique Capability = 0x103d

	// UseSupportedDATs is an alias for CAP_SUPPORTEDDATS.
	UseSupportedDATs Capability = 0x103e

	// UseDoubleFeedDetection is an alias for CAP_DOUBLEFEEDDETECTION.
	UseDoubleFeedDetection Capability = 0x103f

	// UseDoubleFeedDetectionLength is an alias for CAP_DOUBLEFEEDDETECTIONLENGTH.
	UseDoubleFeedDetectionLength Capability = 0x1040

	// UseDoubleFeedDetectionSensitivity is an alias for CAP_DOUBLEFEEDDETECTIONSENSITIVITY.
	UseDoubleFeedDetectionSensitivity Capability = 0x1041

	// UseDoubleFeedDetectionResponse is an alias for CAP_DOUBLEFEEDDETECTIONRESPONSE.
	UseDoubleFeedDetectionResponse Capability = 0x1042

	// UsePaperHandling is an alias for CAP_PAPERHANDLING.
	UsePaperHandling Capability = 0x1043

	// UseIndicatorsMode is an alias for CAP_INDICATORSMODE.
	UseIndicatorsMode Capability = 0x1044

	// UsePrinterVerticalOffset is an alias for CAP_PRINTERVERTICALOFFSET.
	UsePrinterVerticalOffset Capability = 0x1045

	// UsePowerSaveTime is an alias for CAP_POWERSAVETIME.
	UsePowerSaveTime Capability = 0x1046

	// UsePrinterCharRotation is an alias for CAP_PRINTERCHARROTATION.
	UsePrinterCharRotation Capability = 0x1047

	// UsePrinterFontStyle is an alias for CAP_PRINTERFONTSTYLE.
	UsePrinterFontStyle Capability = 0x1048

	// UsePrinterIndexLeadChar is an alias for CAP_PRINTERINDEXLEADCHAR.
	UsePrinterIndexLeadChar Capability = 0x1049

	// UsePrinterIndexMaxValue is an alias for CAP_PRINTERINDEXMAXVALUE.
	UsePrinterIndexMaxValue Capability = 0x104A

	// UsePrinterIndexNumDigits is an alias for CAP_PRINTERINDEXNUMDIGITS.
	UsePrinterIndexNumDigits Capability = 0x104B

	// UsePrinterIndexStep is an alias for CAP_PRINTERINDEXSTEP.
	UsePrinterIndexStep Capability = 0x104C

	// UsePrinterIndexTrigger is an alias for CAP_PRINTERINDEXTRIGGER.
	UsePrinterIndexTrigger Capability = 0x104D

	// UsePrinterStringPreview is an alias for CAP_PRINTERSTRINGPREVIEW.
	UsePrinterStringPreview Capability = 0x104E

	// UseSheetCount is an alias for CAP_SHEETCOUNT.
	UseSheetCount Capability = 0x104F
)

// Image data sources _may_ support these capabilities.
const (
	// UseImageAutoBright is an alias for ICAP_AUTOBRIGHT.
	UseImageAutoBright Capability = 0x1100

	// UseImageBrightness is an alias for ICAP_BRIGHTNESS.
	UseImageBrightness Capability = 0x1101

	// UseImageContrast is an alias for ICAP_CONTRAST.
	UseImageContrast Capability = 0x1103

	// UseImageCustHalftone is an alias for ICAP_CUSTHALFTONE.
	UseImageCustHalftone Capability = 0x1104

	// UseImageExposureTime is an alias for ICAP_EXPOSURETIME.
	UseImageExposureTime Capability = 0x1105

	// UseImageFilter is an alias for ICAP_FILTER.
	UseImageFilter Capability = 0x1106

	// UseImageFlashUsed is an alias for ICAP_FLASHUSED.
	UseImageFlashUsed Capability = 0x1107

	// UseImageGamma is an alias for ICAP_GAMMA.
	UseImageGamma Capability = 0x1108

	// UseImageHalftones is an alias for ICAP_HALFTONES.
	UseImageHalftones Capability = 0x1109

	// UseImageHighlight is an alias for ICAP_HIGHLIGHT.
	UseImageHighlight Capability = 0x110a

	// UseImageImageFileFormat is an alias for ICAP_IMAGEFILEFORMAT.
	UseImageImageFileFormat Capability = 0x110c

	// UseImageLampState is an alias for ICAP_LAMPSTATE.
	UseImageLampState Capability = 0x110d

	// UseImageLightSource is an alias for ICAP_LIGHTSOURCE.
	UseImageLightSource Capability = 0x110e

	// UseImageOrientation is an alias for ICAP_ORIENTATION.
	UseImageOrientation Capability = 0x1110

	// UseImagePhysicalWidth is an alias for ICAP_PHYSICALWIDTH.
	UseImagePhysicalWidth Capability = 0x1111

	// UseImagePhysicalHeight is an alias for ICAP_PHYSICALHEIGHT.
	UseImagePhysicalHeight Capability = 0x1112

	// UseImageShadow is an alias for ICAP_SHADOW.
	UseImageShadow Capability = 0x1113

	// UseImageFrames is an alias for ICAP_FRAMES.
	UseImageFrames Capability = 0x1114

	// UseImageXNativeResolution is an alias for ICAP_XNATIVERESOLUTION.
	UseImageXNativeResolution Capability = 0x1116

	// UseImageYNativeResolution is an alias for ICAP_YNATIVERESOLUTION.
	UseImageYNativeResolution Capability = 0x1117

	// UseImageXResolution is an alias for ICAP_XRESOLUTION.
	UseImageXResolution Capability = 0x1118

	// UseImageYResolution is an alias for ICAP_YRESOLUTION.
	UseImageYResolution Capability = 0x1119

	// UseImageMaxFrames is an alias for ICAP_MAXFRAMES.
	UseImageMaxFrames Capability = 0x111a

	// UseImageTiles is an alias for ICAP_TILES.
	UseImageTiles Capability = 0x111b

	// UseImageBitOrder is an alias for ICAP_BITORDER.
	UseImageBitOrder Capability = 0x111c

	// UseImageCCITTKFactor is an alias for ICAP_CCITTKFACTOR.
	UseImageCCITTKFactor Capability = 0x111d

	// UseImageLightPath is an alias for ICAP_LIGHTPATH.
	UseImageLightPath Capability = 0x111e

	// UseImagePixelFlavor is an alias for ICAP_PIXELFLAVOR.
	UseImagePixelFlavor Capability = 0x111f

	// UseImagePlanarChunky is an alias for ICAP_PLANARCHUNKY.
	UseImagePlanarChunky Capability = 0x1120

	// UseImageRotation is an alias for ICAP_ROTATION.
	UseImageRotation Capability = 0x1121

	// UseImageSupportedSizes is an alias for ICAP_SUPPORTEDSIZES.
	UseImageSupportedSizes Capability = 0x1122

	// UseImageThreshold is an alias for ICAP_THRESHOLD.
	UseImageThreshold Capability = 0x1123

	// UseImageXScaling is an alias for ICAP_XSCALING.
	UseImageXScaling Capability = 0x1124

	// UseImageYScaling is an alias for ICAP_YSCALING.
	UseImageYScaling Capability = 0x1125

	// UseImageBitOrderCodes is an alias for ICAP_BITORDERCODES.
	UseImageBitOrderCodes Capability = 0x1126

	// UseImagePixelFlavorCodes is an alias for ICAP_PIXELFLAVORCODES.
	UseImagePixelFlavorCodes Capability = 0x1127

	// UseImageJPEGPixelType is an alias for ICAP_JPEGPIXELTYPE.
	UseImageJPEGPixelType Capability = 0x1128

	// UseImageTimeFill is an alias for ICAP_TIMEFILL.
	UseImageTimeFill Capability = 0x112a

	// UseImageBitDepth is an alias for ICAP_BITDEPTH.
	UseImageBitDepth Capability = 0x112b

	// UseImageBitDepthReduction is an alias for ICAP_BITDEPTHREDUCTION.
	UseImageBitDepthReduction Capability = 0x112c

	// UseImageUndefinedImageSize is an alias for ICAP_UNDEFINEDIMAGESIZE.
	UseImageUndefinedImageSize Capability = 0x112d

	// UseImageImageDataset is an alias for ICAP_IMAGEDATASET.
	UseImageImageDataset Capability = 0x112e

	// UseImageExtImageInfo is an alias for ICAP_EXTIMAGEINFO.
	UseImageExtImageInfo Capability = 0x112f

	// UseImageMinimumHeight is an alias for ICAP_MINIMUMHEIGHT.
	UseImageMinimumHeight Capability = 0x1130

	// UseImageMinimumWidth is an alias for ICAP_MINIMUMWIDTH.
	UseImageMinimumWidth Capability = 0x1131

	// UseImageAutoDiscardBlankPages is an alias for ICAP_AUTODISCARDBLANKPAGES.
	UseImageAutoDiscardBlankPages Capability = 0x1134

	// UseImageFlipRotation is an alias for ICAP_FLIPROTATION.
	UseImageFlipRotation Capability = 0x1136

	// UseImageBarcodeDetectionEnabled is an alias for ICAP_BARCODEDETECTIONENABLED.
	UseImageBarcodeDetectionEnabled Capability = 0x1137

	// UseImageSupportedBarcodeTypes is an alias for ICAP_SUPPORTEDBARCODETYPES.
	UseImageSupportedBarcodeTypes Capability = 0x1138

	// UseImageBarcodeMaxSearchPriorities is an alias for ICAP_BARCODEMAXSEARCHPRIORITIES.
	UseImageBarcodeMaxSearchPriorities Capability = 0x1139

	// UseImageBarcodeSearchPriorities is an alias for ICAP_BARCODESEARCHPRIORITIES.
	UseImageBarcodeSearchPriorities Capability = 0x113a

	// UseImageBarcodeSearchMode is an alias for ICAP_BARCODESEARCHMODE.
	UseImageBarcodeSearchMode Capability = 0x113b

	// UseImageBarcodeMaxRetries is an alias for ICAP_BARCODEMAXRETRIES.
	UseImageBarcodeMaxRetries Capability = 0x113c

	// UseImageBarcodeTimeout is an alias for ICAP_BARCODETIMEOUT.
	UseImageBarcodeTimeout Capability = 0x113d

	// UseImageZoomFactor is an alias for ICAP_ZOOMFACTOR.
	UseImageZoomFactor Capability = 0x113e

	// UseImagePatchCodeDetectionEnabled is an alias for ICAP_PATCHCODEDETECTIONENABLED.
	UseImagePatchCodeDetectionEnabled Capability = 0x113f

	// UseImageSupportedPatchCodeTypes is an alias for ICAP_SUPPORTEDPATCHCODETYPES.
	UseImageSupportedPatchCodeTypes Capability = 0x1140

	// UseImagePatchCodeMaxSearchPriorities is an alias for ICAP_PATCHCODEMAXSEARCHPRIORITIES.
	UseImagePatchCodeMaxSearchPriorities Capability = 0x1141

	// UseImagePatchCodeSearchPriorities is an alias for ICAP_PATCHCODESEARCHPRIORITIES.
	UseImagePatchCodeSearchPriorities Capability = 0x1142

	// UseImagePatchCodeSearchMode is an alias for ICAP_PATCHCODESEARCHMODE.
	UseImagePatchCodeSearchMode Capability = 0x1143

	// UseImagePatchCodeMaxRetries is an alias for ICAP_PATCHCODEMAXRETRIES.
	UseImagePatchCodeMaxRetries Capability = 0x1144

	// UseImagePatchCodeTimeout is an alias for ICAP_PATCHCODETIMEOUT.
	UseImagePatchCodeTimeout Capability = 0x1145

	// UseImageFlashUsed2 is an alias for ICAP_FLASHUSED2.
	UseImageFlashUsed2 Capability = 0x1146

	// UseImageImageFilter is an alias for ICAP_IMAGEFILTER.
	UseImageImageFilter Capability = 0x1147

	// UseImageNoiseFilter is an alias for ICAP_NOISEFILTER.
	UseImageNoiseFilter Capability = 0x1148

	// UseImageOverScan is an alias for ICAP_OVERSCAN.
	UseImageOverScan Capability = 0x1149

	// UseImageAutomaticBorderDetection is an alias for ICAP_AUTOMATICBORDERDETECTION.
	UseImageAutomaticBorderDetection Capability = 0x1150

	// UseImageAutomaticDESkew is an alias for ICAP_AUTOMATICDESKEW.
	UseImageAutomaticDESkew Capability = 0x1151

	// UseImageAutomaticRotate is an alias for ICAP_AUTOMATICROTATE.
	UseImageAutomaticRotate Capability = 0x1152

	// UseImageJPEGQuality is an alias for ICAP_JPEGQUALITY.
	UseImageJPEGQuality Capability = 0x1153

	// UseImageFeederType is an alias for ICAP_FEEDERTYPE.
	UseImageFeederType Capability = 0x1154

	// UseImageICCProfile is an alias for ICAP_ICCPROFILE.
	UseImageICCProfile Capability = 0x1155

	// UseImageAutoSize is an alias for ICAP_AUTOSIZE.
	UseImageAutoSize Capability = 0x1156

	// UseImageAutomaticCropUsesFrame is an alias for ICAP_AUTOMATICCROPUSESFRAME.
	UseImageAutomaticCropUsesFrame Capability = 0x1157

	// UseImageAutomaticLengthDetection is an alias for ICAP_AUTOMATICLENGTHDETECTION.
	UseImageAutomaticLengthDetection Capability = 0x1158

	// UseImageAutomaticColorEnabled is an alias for ICAP_AUTOMATICCOLORENABLED.
	UseImageAutomaticColorEnabled Capability = 0x1159

	// UseImageAutomaticColorNonColorPixelType is an alias for ICAP_AUTOMATICCOLORNONCOLORPIXELTYPE.
	UseImageAutomaticColorNonColorPixelType Capability = 0x115a

	// UseImageColorManagementEnabled is an alias for ICAP_COLORMANAGEMENTENABLED.
	UseImageColorManagementEnabled Capability = 0x115b

	// UseImageImageMerge is an alias for ICAP_IMAGEMERGE.
	UseImageImageMerge Capability = 0x115c

	// UseImageImageMergeHeightThreshold is an alias for ICAP_IMAGEMERGEHEIGHTTHRESHOLD.
	UseImageImageMergeHeightThreshold Capability = 0x115d

	// UseImageSupportedExtImageInfo is an alias for ICAP_SUPPORTEDEXTIMAGEINFO.
	UseImageSupportedExtImageInfo Capability = 0x115e

	// UseImageFilmType is an alias for ICAP_FILMTYPE.
	UseImageFilmType Capability = 0x115f

	// UseImageMirror is an alias for ICAP_MIRROR.
	UseImageMirror Capability = 0x1160

	// UseImageJPEGSubsampling is an alias for ICAP_JPEGSUBSAMPLING.
	UseImageJPEGSubsampling Capability = 0x1161
)

// AudioTransferMechanism is an alias for ACAP_XFERMECH. Image data sources
// _may_ support this audio capability.
const UseAudioTransferMechanism Capability = 0x1202
