package twain

// CapabilityType is an alias for CAP_ and ICAP_ values.
type CapabilityType uint16

// CapabilityCustomBase is the base of custom capabilities. It is an alias for
// CAP_CUSTOMBASE.
const CapabilityCustomBase CapabilityType = 0x8000

// CapabilityTransferCount is an alias for CAP_XFERCOUNT. All data sources are
// required to support this capability.
const CapabilityTransferCount CapabilityType = 0x8000

// Image data sources are _required_ to support these capabilities.
const (
	// CapabilityCompression is an alias for ICAP_COMPRESSION.
	CapabilityCompression CapabilityType = 0x0100

	// CapabilityPixelType is an alias for ICAP_PIXELTYPE.
	CapabilityPixelType CapabilityType = 0x0101

	// CapabilityUnits is an alias for ICAP_UNITS.
	CapabilityUnits CapabilityType = 0x0102

	// CapabilityTransferMechanism is an alias for ICAP_XFERMECH.
	CapabilityTransferMechanism CapabilityType = 0x0103
)

// All data sources _may_ support these capabilities.
const (
	// CapabilityAuthor is an alias for CAP_AUTHOR.
	CapabilityAuthor CapabilityType = 0x1000

	// CapabilityCaption is an alias for CAP_CAPTION.
	CapabilityCaption CapabilityType = 0x1001

	// CapabilityFeederEnabled is an alias for CAP_FEEDERENABLED.
	CapabilityFeederEnabled CapabilityType = 0x1002

	// CapabilityFeederLoaded is an alias for CAP_FEEDERLOADED.
	CapabilityFeederLoaded CapabilityType = 0x1003

	// CapabilityTimeDate is an alias for CAP_TIMEDATE.
	CapabilityTimeDate CapabilityType = 0x1004

	// CapabilitySupportedCaps is an alias for CAP_SUPPORTEDCAPS.
	CapabilitySupportedCaps CapabilityType = 0x1005

	// CapabilityExtendedCaps is an alias for CAP_EXTENDEDCAPS.
	CapabilityExtendedCaps CapabilityType = 0x1006

	// CapabilityAutoFeed is an alias for CAP_AUTOFEED.
	CapabilityAutoFeed CapabilityType = 0x1007

	// CapabilityClearPage is an alias for CAP_CLEARPAGE.
	CapabilityClearPage CapabilityType = 0x1008

	// CapabilityFeedPage is an alias for CAP_FEEDPAGE.
	CapabilityFeedPage CapabilityType = 0x1009

	// CapabilityRewindPage is an alias for CAP_REWINDPAGE.
	CapabilityRewindPage CapabilityType = 0x100a

	// CapabilityIndicators is an alias for CAP_INDICATORS.
	CapabilityIndicators CapabilityType = 0x100b

	// CapabilityPaperDetectable is an alias for CAP_PAPERDETECTABLE.
	CapabilityPaperDetectable CapabilityType = 0x100d

	// CapabilityUIControllable is an alias for CAP_UICONTROLLABLE.
	CapabilityUIControllable CapabilityType = 0x100e

	// CapabilityDeviceOnline is an alias for CAP_DEVICEONLINE.
	CapabilityDeviceOnline CapabilityType = 0x100f

	// CapabilityAutoScan is an alias for CAP_AUTOSCAN.
	CapabilityAutoScan CapabilityType = 0x1010

	// CapabilityThumbnailsEnabled is an alias for CAP_THUMBNAILSENABLED.
	CapabilityThumbnailsEnabled CapabilityType = 0x1011

	// CapabilityDuplex is an alias for CAP_DUPLEX.
	CapabilityDuplex CapabilityType = 0x1012

	// CapabilityDuplexEnabled is an alias for CAP_DUPLEXENABLED.
	CapabilityDuplexEnabled CapabilityType = 0x1013

	// CapabilityEnabledSUIOnly is an alias for CAP_ENABLEDSUIONLY.
	CapabilityEnabledSUIOnly CapabilityType = 0x1014

	// CapabilityCustomdsData is an alias for CAP_CUSTOMDSDATA.
	CapabilityCustomdsData CapabilityType = 0x1015

	// CapabilityEndorser is an alias for CAP_ENDORSER.
	CapabilityEndorser CapabilityType = 0x1016

	// CapabilityJobControl is an alias for CAP_JOBCONTROL.
	CapabilityJobControl CapabilityType = 0x1017

	// CapabilityAlarms is an alias for CAP_ALARMS.
	CapabilityAlarms CapabilityType = 0x1018

	// CapabilityAlarmVolume is an alias for CAP_ALARMVOLUME.
	CapabilityAlarmVolume CapabilityType = 0x1019

	// CapabilityAutomaticCapture is an alias for CAP_AUTOMATICCAPTURE.
	CapabilityAutomaticCapture CapabilityType = 0x101a

	// CapabilityTimeBeforeFirstCapture is an alias for CAP_TIMEBEFOREFIRSTCAPTURE.
	CapabilityTimeBeforeFirstCapture CapabilityType = 0x101b

	// CapabilityTimeBetweenCaptures is an alias for CAP_TIMEBETWEENCAPTURES.
	CapabilityTimeBetweenCaptures CapabilityType = 0x101c

	// CapabilityMaxBatchBuffers is an alias for CAP_MAXBATCHBUFFERS.
	CapabilityMaxBatchBuffers CapabilityType = 0x101e

	// CapabilityDeviceTimeDate is an alias for CAP_DEVICETIMEDATE.
	CapabilityDeviceTimeDate CapabilityType = 0x101f

	// CapabilityPowerSupply is an alias for CAP_POWERSUPPLY.
	CapabilityPowerSupply CapabilityType = 0x1020

	// CapabilityCameraPreviewUI is an alias for CAP_CAMERAPREVIEWUI.
	CapabilityCameraPreviewUI CapabilityType = 0x1021

	// CapabilityDeviceEvent is an alias for CAP_DEVICEEVENT.
	CapabilityDeviceEvent CapabilityType = 0x1022

	// CapabilitySerialNumber is an alias for CAP_SERIALNUMBER.
	CapabilitySerialNumber CapabilityType = 0x1024

	// CapabilityPrinter is an alias for CAP_PRINTER.
	CapabilityPrinter CapabilityType = 0x1026

	// CapabilityPrinterEnabled is an alias for CAP_PRINTERENABLED.
	CapabilityPrinterEnabled CapabilityType = 0x1027

	// CapabilityPrinterIndex is an alias for CAP_PRINTERINDEX.
	CapabilityPrinterIndex CapabilityType = 0x1028

	// CapabilityPrinterMode is an alias for CAP_PRINTERMODE.
	CapabilityPrinterMode CapabilityType = 0x1029

	// CapabilityPrinterString is an alias for CAP_PRINTERSTRING.
	CapabilityPrinterString CapabilityType = 0x102a

	// CapabilityPrinterSuffix is an alias for CAP_PRINTERSUFFIX.
	CapabilityPrinterSuffix CapabilityType = 0x102b

	// CapabilityLanguage is an alias for CAP_LANGUAGE.
	CapabilityLanguage CapabilityType = 0x102c

	// CapabilityFeederAlignment is an alias for CAP_FEEDERALIGNMENT.
	CapabilityFeederAlignment CapabilityType = 0x102d

	// CapabilityFeederOrder is an alias for CAP_FEEDERORDER.
	CapabilityFeederOrder CapabilityType = 0x102e

	// CapabilityReacquireAllowed is an alias for CAP_REACQUIREALLOWED.
	CapabilityReacquireAllowed CapabilityType = 0x1030

	// CapabilityBatteryMinutes is an alias for CAP_BATTERYMINUTES.
	CapabilityBatteryMinutes CapabilityType = 0x1032

	// CapabilityBatteryPercentage is an alias for CAP_BATTERYPERCENTAGE.
	CapabilityBatteryPercentage CapabilityType = 0x1033

	// CapabilityCameraSide is an alias for CAP_CAMERASIDE.
	CapabilityCameraSide CapabilityType = 0x1034

	// CapabilitySegmented is an alias for CAP_SEGMENTED.
	CapabilitySegmented CapabilityType = 0x1035

	// CapabilityCameraEnabled is an alias for CAP_CAMERAENABLED.
	CapabilityCameraEnabled CapabilityType = 0x1036

	// CapabilityCameraOrder is an alias for CAP_CAMERAORDER.
	CapabilityCameraOrder CapabilityType = 0x1037

	// CapabilityMICREnabled is an alias for CAP_MICRENABLED.
	CapabilityMICREnabled CapabilityType = 0x1038

	// CapabilityFeederPrep is an alias for CAP_FEEDERPREP.
	CapabilityFeederPrep CapabilityType = 0x1039

	// CapabilityFeederPocket is an alias for CAP_FEEDERPOCKET.
	CapabilityFeederPocket CapabilityType = 0x103a

	// CapabilityAutomaticSenseMedium is an alias for CAP_AUTOMATICSENSEMEDIUM.
	CapabilityAutomaticSenseMedium CapabilityType = 0x103b

	// CapabilityCustomInterfaceGUID is an alias for CAP_CUSTOMINTERFACEGUID.
	CapabilityCustomInterfaceGUID CapabilityType = 0x103c

	// CapabilitySupportedCapsSegmentUnique is an alias for CAP_SUPPORTEDCAPSSEGMENTUNIQUE.
	CapabilitySupportedCapsSegmentUnique CapabilityType = 0x103d

	// CapabilitySupportedDATs is an alias for CAP_SUPPORTEDDATS.
	CapabilitySupportedDATs CapabilityType = 0x103e

	// CapabilityDoubleFeedDetection is an alias for CAP_DOUBLEFEEDDETECTION.
	CapabilityDoubleFeedDetection CapabilityType = 0x103f

	// CapabilityDoubleFeedDetectionLength is an alias for CAP_DOUBLEFEEDDETECTIONLENGTH.
	CapabilityDoubleFeedDetectionLength CapabilityType = 0x1040

	// CapabilityDoubleFeedDetectionSensitivity is an alias for CAP_DOUBLEFEEDDETECTIONSENSITIVITY.
	CapabilityDoubleFeedDetectionSensitivity CapabilityType = 0x1041

	// CapabilityDoubleFeedDetectionResponse is an alias for CAP_DOUBLEFEEDDETECTIONRESPONSE.
	CapabilityDoubleFeedDetectionResponse CapabilityType = 0x1042

	// CapabilityPaperHandling is an alias for CAP_PAPERHANDLING.
	CapabilityPaperHandling CapabilityType = 0x1043

	// CapabilityIndicatorsMode is an alias for CAP_INDICATORSMODE.
	CapabilityIndicatorsMode CapabilityType = 0x1044

	// CapabilityPrinterVerticalOffset is an alias for CAP_PRINTERVERTICALOFFSET.
	CapabilityPrinterVerticalOffset CapabilityType = 0x1045

	// CapabilityPowerSaveTime is an alias for CAP_POWERSAVETIME.
	CapabilityPowerSaveTime CapabilityType = 0x1046

	// CapabilityPrinterCharRotation is an alias for CAP_PRINTERCHARROTATION.
	CapabilityPrinterCharRotation CapabilityType = 0x1047

	// CapabilityPrinterFontStyle is an alias for CAP_PRINTERFONTSTYLE.
	CapabilityPrinterFontStyle CapabilityType = 0x1048

	// CapabilityPrinterIndexLeadChar is an alias for CAP_PRINTERINDEXLEADCHAR.
	CapabilityPrinterIndexLeadChar CapabilityType = 0x1049

	// CapabilityPrinterIndexMaxValue is an alias for CAP_PRINTERINDEXMAXVALUE.
	CapabilityPrinterIndexMaxValue CapabilityType = 0x104A

	// CapabilityPrinterIndexNumDigits is an alias for CAP_PRINTERINDEXNUMDIGITS.
	CapabilityPrinterIndexNumDigits CapabilityType = 0x104B

	// CapabilityPrinterIndexStep is an alias for CAP_PRINTERINDEXSTEP.
	CapabilityPrinterIndexStep CapabilityType = 0x104C

	// CapabilityPrinterIndexTrigger is an alias for CAP_PRINTERINDEXTRIGGER.
	CapabilityPrinterIndexTrigger CapabilityType = 0x104D

	// CapabilityPrinterStringPreview is an alias for CAP_PRINTERSTRINGPREVIEW.
	CapabilityPrinterStringPreview CapabilityType = 0x104E

	// CapabilitySheetCount is an alias for CAP_SHEETCOUNT.
	CapabilitySheetCount CapabilityType = 0x104F
)

// Image data sources _may_ support these capabilities.
const (
	// CapabilityImageAutoBright is an alias for ICAP_AUTOBRIGHT.
	CapabilityImageAutoBright CapabilityType = 0x1100

	// CapabilityImageBrightness is an alias for ICAP_BRIGHTNESS.
	CapabilityImageBrightness CapabilityType = 0x1101

	// CapabilityImageContrast is an alias for ICAP_CONTRAST.
	CapabilityImageContrast CapabilityType = 0x1103

	// CapabilityImageCustHalftone is an alias for ICAP_CUSTHALFTONE.
	CapabilityImageCustHalftone CapabilityType = 0x1104

	// CapabilityImageExposureTime is an alias for ICAP_EXPOSURETIME.
	CapabilityImageExposureTime CapabilityType = 0x1105

	// CapabilityImageFilter is an alias for ICAP_FILTER.
	CapabilityImageFilter CapabilityType = 0x1106

	// CapabilityImageFlashUsed is an alias for ICAP_FLASHUSED.
	CapabilityImageFlashUsed CapabilityType = 0x1107

	// CapabilityImageGamma is an alias for ICAP_GAMMA.
	CapabilityImageGamma CapabilityType = 0x1108

	// CapabilityImageHalftones is an alias for ICAP_HALFTONES.
	CapabilityImageHalftones CapabilityType = 0x1109

	// CapabilityImageHighlight is an alias for ICAP_HIGHLIGHT.
	CapabilityImageHighlight CapabilityType = 0x110a

	// CapabilityImageImageFileFormat is an alias for ICAP_IMAGEFILEFORMAT.
	CapabilityImageImageFileFormat CapabilityType = 0x110c

	// CapabilityImageLampState is an alias for ICAP_LAMPSTATE.
	CapabilityImageLampState CapabilityType = 0x110d

	// CapabilityImageLightSource is an alias for ICAP_LIGHTSOURCE.
	CapabilityImageLightSource CapabilityType = 0x110e

	// CapabilityImageOrientation is an alias for ICAP_ORIENTATION.
	CapabilityImageOrientation CapabilityType = 0x1110

	// CapabilityImagePhysicalWidth is an alias for ICAP_PHYSICALWIDTH.
	CapabilityImagePhysicalWidth CapabilityType = 0x1111

	// CapabilityImagePhysicalHeight is an alias for ICAP_PHYSICALHEIGHT.
	CapabilityImagePhysicalHeight CapabilityType = 0x1112

	// CapabilityImageShadow is an alias for ICAP_SHADOW.
	CapabilityImageShadow CapabilityType = 0x1113

	// CapabilityImageFrames is an alias for ICAP_FRAMES.
	CapabilityImageFrames CapabilityType = 0x1114

	// CapabilityImageXNativeResolution is an alias for ICAP_XNATIVERESOLUTION.
	CapabilityImageXNativeResolution CapabilityType = 0x1116

	// CapabilityImageYNativeResolution is an alias for ICAP_YNATIVERESOLUTION.
	CapabilityImageYNativeResolution CapabilityType = 0x1117

	// CapabilityImageXResolution is an alias for ICAP_XRESOLUTION.
	CapabilityImageXResolution CapabilityType = 0x1118

	// CapabilityImageYResolution is an alias for ICAP_YRESOLUTION.
	CapabilityImageYResolution CapabilityType = 0x1119

	// CapabilityImageMaxFrames is an alias for ICAP_MAXFRAMES.
	CapabilityImageMaxFrames CapabilityType = 0x111a

	// CapabilityImageTiles is an alias for ICAP_TILES.
	CapabilityImageTiles CapabilityType = 0x111b

	// CapabilityImageBitOrder is an alias for ICAP_BITORDER.
	CapabilityImageBitOrder CapabilityType = 0x111c

	// CapabilityImageCCITTKFactor is an alias for ICAP_CCITTKFACTOR.
	CapabilityImageCCITTKFactor CapabilityType = 0x111d

	// CapabilityImageLightPath is an alias for ICAP_LIGHTPATH.
	CapabilityImageLightPath CapabilityType = 0x111e

	// CapabilityImagePixelFlavor is an alias for ICAP_PIXELFLAVOR.
	CapabilityImagePixelFlavor CapabilityType = 0x111f

	// CapabilityImagePlanarChunky is an alias for ICAP_PLANARCHUNKY.
	CapabilityImagePlanarChunky CapabilityType = 0x1120

	// CapabilityImageRotation is an alias for ICAP_ROTATION.
	CapabilityImageRotation CapabilityType = 0x1121

	// CapabilityImageSupportedSizes is an alias for ICAP_SUPPORTEDSIZES.
	CapabilityImageSupportedSizes CapabilityType = 0x1122

	// CapabilityImageThreshold is an alias for ICAP_THRESHOLD.
	CapabilityImageThreshold CapabilityType = 0x1123

	// CapabilityImageXScaling is an alias for ICAP_XSCALING.
	CapabilityImageXScaling CapabilityType = 0x1124

	// CapabilityImageYScaling is an alias for ICAP_YSCALING.
	CapabilityImageYScaling CapabilityType = 0x1125

	// CapabilityImageBitOrderCodes is an alias for ICAP_BITORDERCODES.
	CapabilityImageBitOrderCodes CapabilityType = 0x1126

	// CapabilityImagePixelFlavorCodes is an alias for ICAP_PIXELFLAVORCODES.
	CapabilityImagePixelFlavorCodes CapabilityType = 0x1127

	// CapabilityImageJPEGPixelType is an alias for ICAP_JPEGPIXELTYPE.
	CapabilityImageJPEGPixelType CapabilityType = 0x1128

	// CapabilityImageTimeFill is an alias for ICAP_TIMEFILL.
	CapabilityImageTimeFill CapabilityType = 0x112a

	// CapabilityImageBitDepth is an alias for ICAP_BITDEPTH.
	CapabilityImageBitDepth CapabilityType = 0x112b

	// CapabilityImageBitDepthReduction is an alias for ICAP_BITDEPTHREDUCTION.
	CapabilityImageBitDepthReduction CapabilityType = 0x112c

	// CapabilityImageUndefinedImageSize is an alias for ICAP_UNDEFINEDIMAGESIZE.
	CapabilityImageUndefinedImageSize CapabilityType = 0x112d

	// CapabilityImageImageDataset is an alias for ICAP_IMAGEDATASET.
	CapabilityImageImageDataset CapabilityType = 0x112e

	// CapabilityImageExtImageInfo is an alias for ICAP_EXTIMAGEINFO.
	CapabilityImageExtImageInfo CapabilityType = 0x112f

	// CapabilityImageMinimumHeight is an alias for ICAP_MINIMUMHEIGHT.
	CapabilityImageMinimumHeight CapabilityType = 0x1130

	// CapabilityImageMinimumWidth is an alias for ICAP_MINIMUMWIDTH.
	CapabilityImageMinimumWidth CapabilityType = 0x1131

	// CapabilityImageAutoDiscardBlankPages is an alias for ICAP_AUTODISCARDBLANKPAGES.
	CapabilityImageAutoDiscardBlankPages CapabilityType = 0x1134

	// CapabilityImageFlipRotation is an alias for ICAP_FLIPROTATION.
	CapabilityImageFlipRotation CapabilityType = 0x1136

	// CapabilityImageBarcodeDetectionEnabled is an alias for ICAP_BARCODEDETECTIONENABLED.
	CapabilityImageBarcodeDetectionEnabled CapabilityType = 0x1137

	// CapabilityImageSupportedBarcodeTypes is an alias for ICAP_SUPPORTEDBARCODETYPES.
	CapabilityImageSupportedBarcodeTypes CapabilityType = 0x1138

	// CapabilityImageBarcodeMaxSearchPriorities is an alias for ICAP_BARCODEMAXSEARCHPRIORITIES.
	CapabilityImageBarcodeMaxSearchPriorities CapabilityType = 0x1139

	// CapabilityImageBarcodeSearchPriorities is an alias for ICAP_BARCODESEARCHPRIORITIES.
	CapabilityImageBarcodeSearchPriorities CapabilityType = 0x113a

	// CapabilityImageBarcodeSearchMode is an alias for ICAP_BARCODESEARCHMODE.
	CapabilityImageBarcodeSearchMode CapabilityType = 0x113b

	// CapabilityImageBarcodeMaxRetries is an alias for ICAP_BARCODEMAXRETRIES.
	CapabilityImageBarcodeMaxRetries CapabilityType = 0x113c

	// CapabilityImageBarcodeTimeout is an alias for ICAP_BARCODETIMEOUT.
	CapabilityImageBarcodeTimeout CapabilityType = 0x113d

	// CapabilityImageZoomFactor is an alias for ICAP_ZOOMFACTOR.
	CapabilityImageZoomFactor CapabilityType = 0x113e

	// CapabilityImagePatchCodeDetectionEnabled is an alias for ICAP_PATCHCODEDETECTIONENABLED.
	CapabilityImagePatchCodeDetectionEnabled CapabilityType = 0x113f

	// CapabilityImageSupportedPatchCodeTypes is an alias for ICAP_SUPPORTEDPATCHCODETYPES.
	CapabilityImageSupportedPatchCodeTypes CapabilityType = 0x1140

	// CapabilityImagePatchCodeMaxSearchPriorities is an alias for ICAP_PATCHCODEMAXSEARCHPRIORITIES.
	CapabilityImagePatchCodeMaxSearchPriorities CapabilityType = 0x1141

	// CapabilityImagePatchCodeSearchPriorities is an alias for ICAP_PATCHCODESEARCHPRIORITIES.
	CapabilityImagePatchCodeSearchPriorities CapabilityType = 0x1142

	// CapabilityImagePatchCodeSearchMode is an alias for ICAP_PATCHCODESEARCHMODE.
	CapabilityImagePatchCodeSearchMode CapabilityType = 0x1143

	// CapabilityImagePatchCodeMaxRetries is an alias for ICAP_PATCHCODEMAXRETRIES.
	CapabilityImagePatchCodeMaxRetries CapabilityType = 0x1144

	// CapabilityImagePatchCodeTimeout is an alias for ICAP_PATCHCODETIMEOUT.
	CapabilityImagePatchCodeTimeout CapabilityType = 0x1145

	// CapabilityImageFlashUsed2 is an alias for ICAP_FLASHUSED2.
	CapabilityImageFlashUsed2 CapabilityType = 0x1146

	// CapabilityImageImageFilter is an alias for ICAP_IMAGEFILTER.
	CapabilityImageImageFilter CapabilityType = 0x1147

	// CapabilityImageNoiseFilter is an alias for ICAP_NOISEFILTER.
	CapabilityImageNoiseFilter CapabilityType = 0x1148

	// CapabilityImageOverScan is an alias for ICAP_OVERSCAN.
	CapabilityImageOverScan CapabilityType = 0x1149

	// CapabilityImageAutomaticBorderDetection is an alias for ICAP_AUTOMATICBORDERDETECTION.
	CapabilityImageAutomaticBorderDetection CapabilityType = 0x1150

	// CapabilityImageAutomaticDESkew is an alias for ICAP_AUTOMATICDESKEW.
	CapabilityImageAutomaticDESkew CapabilityType = 0x1151

	// CapabilityImageAutomaticRotate is an alias for ICAP_AUTOMATICROTATE.
	CapabilityImageAutomaticRotate CapabilityType = 0x1152

	// CapabilityImageJPEGQuality is an alias for ICAP_JPEGQUALITY.
	CapabilityImageJPEGQuality CapabilityType = 0x1153

	// CapabilityImageFeederType is an alias for ICAP_FEEDERTYPE.
	CapabilityImageFeederType CapabilityType = 0x1154

	// CapabilityImageICCProfile is an alias for ICAP_ICCPROFILE.
	CapabilityImageICCProfile CapabilityType = 0x1155

	// CapabilityImageAutoSize is an alias for ICAP_AUTOSIZE.
	CapabilityImageAutoSize CapabilityType = 0x1156

	// CapabilityImageAutomaticCropUsesFrame is an alias for ICAP_AUTOMATICCROPUSESFRAME.
	CapabilityImageAutomaticCropUsesFrame CapabilityType = 0x1157

	// CapabilityImageAutomaticLengthDetection is an alias for ICAP_AUTOMATICLENGTHDETECTION.
	CapabilityImageAutomaticLengthDetection CapabilityType = 0x1158

	// CapabilityImageAutomaticColorEnabled is an alias for ICAP_AUTOMATICCOLORENABLED.
	CapabilityImageAutomaticColorEnabled CapabilityType = 0x1159

	// CapabilityImageAutomaticColorNonColorPixelType is an alias for ICAP_AUTOMATICCOLORNONCOLORPIXELTYPE.
	CapabilityImageAutomaticColorNonColorPixelType CapabilityType = 0x115a

	// CapabilityImageColorManagementEnabled is an alias for ICAP_COLORMANAGEMENTENABLED.
	CapabilityImageColorManagementEnabled CapabilityType = 0x115b

	// CapabilityImageImageMerge is an alias for ICAP_IMAGEMERGE.
	CapabilityImageImageMerge CapabilityType = 0x115c

	// CapabilityImageImageMergeHeightThreshold is an alias for ICAP_IMAGEMERGEHEIGHTTHRESHOLD.
	CapabilityImageImageMergeHeightThreshold CapabilityType = 0x115d

	// CapabilityImageSupportedExtImageInfo is an alias for ICAP_SUPPORTEDEXTIMAGEINFO.
	CapabilityImageSupportedExtImageInfo CapabilityType = 0x115e

	// CapabilityImageFilmType is an alias for ICAP_FILMTYPE.
	CapabilityImageFilmType CapabilityType = 0x115f

	// CapabilityImageMirror is an alias for ICAP_MIRROR.
	CapabilityImageMirror CapabilityType = 0x1160

	// CapabilityImageJPEGSubsampling is an alias for ICAP_JPEGSUBSAMPLING.
	CapabilityImageJPEGSubsampling CapabilityType = 0x1161
)

// CapabilityAudioTransferMechanism is an alias for ACAP_XFERMECH. Image data
// sources _may_ support this audio capability.
const CapabilityAudioTransferMechanism CapabilityType = 0x1202

// Alarm is an alias for CAP_ALARMS values.
type Alarm uint16

const (
	// AlarmAlarm is an alias for TWAL_ALARM.
	AlarmAlarm Alarm = iota

	// AlarmFeedError is an alias for TWAL_FEEDERERROR.
	AlarmFeedError

	// AlarmFeederWarning is an alias for TWAL_FEEDERWARNING.
	AlarmFeederWarning

	// AlarmBarcode is an alias for TWAL_BARCODE.
	AlarmBarcode

	// AlarmDoubleFeed is an alias for TWAL_DOUBLEFEED.
	AlarmDoubleFeed

	// AlarmJam is an alias for TWAL_JAM.
	AlarmJam

	// AlarmPatchCode is an alias for TWAL_PATCHCODE.
	AlarmPatchCode

	// AlarmPower is an alias for TWAL_POWER.
	AlarmPower

	// AlarmSkew is an alias for TWAL_SKEW.
	AlarmSkew
)

// AutoSize is an alias for the ICAP_AUTOSIZE values.
type AutoSize uint16

const (
	// AutoSizeNone is an alias for TWAS_NONE.
	AutoSizeNone AutoSize = iota

	// AutoSizeAuto is an alias for TWAS_AUTO.
	AutoSizeAuto

	// AutoSizeCurrent is an alias for TWAS_CURRENT.
	AutoSizeCurrent
)

// BarcodeRotation is an alias for the TWEI_BARCODEROTATION values.
type BarcodeRotation uint16

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
type BarcodeSearchMode uint16

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
type BitOrder uint16

const (
	// BitOrderLSBFirst is an alias for TWBO_LSBFIRST.
	BitOrderLSBFirst BitOrder = iota

	// BitOrderMSBFirst is an alias for TWBO_MSBFIRST.
	BitOrderMSBFirst
)

// AutoDiscardBlankPages is an alias for ICAP_AUTODISCARDBLANKPAGES values.
type AutoDiscardBlankPages int16

const (
	// AutoDiscardBlankPagesDisable is an alias for TWBP_DISABLE.
	AutoDiscardBlankPagesDisable AutoDiscardBlankPages = -2

	// AutoDiscardBlankPagesAuto is an alias for TWBP_AUTO.
	AutoDiscardBlankPagesAuto AutoDiscardBlankPages = -1
)

// BitDepthReduction is an alias for ICAP_BITDEPTHREDUCTION values.
type BitDepthReduction uint16

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
type SupportedBarcodeType uint16

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

// Compression is an alias for ICAP_COMPRESSION values.
type Compression uint16

const (
	// CompressionNone is an alias for TWCP_NONE.
	CompressionNone Compression = iota

	// CompressionPackBits is an alias for TWCP_PACKBITS.
	CompressionPackBits

	// CompressionGroup31D is an alias for TWCP_GROUP31D.
	CompressionGroup31D

	// CompressionGroup31DEOL is an alias for TWCP_GROUP31DEOL.
	CompressionGroup31DEOL

	// CompressionGroup32D is an alias for TWCP_GROUP32D.
	CompressionGroup32D

	// CompressionGroup4 is an alias for TWCP_GROUP4.
	CompressionGroup4

	// CompressionJPEG is an alias for TWCP_JPEG.
	CompressionJPEG

	// CompressionLZW is an alias for TWCP_LZW.
	CompressionLZW

	// CompressionJBIG is an alias for TWCP_JBIG.
	CompressionJBIG

	// CompressionPNG is an alias for TWCP_PNG.
	CompressionPNG

	// CompressionRLE4 is an alias for TWCP_RLE4.
	CompressionRLE4

	// CompressionRLE8 is an alias for TWCP_RLE8.
	CompressionRLE8

	// CompressionBitFields is an alias for TWCP_BITFIELDS.
	CompressionBitFields

	// CompressionZIP is an alias for TWCP_ZIP.
	CompressionZIP

	// CompressionJPEG2000 is an alias for TWCP_JPEG2000.
	CompressionJPEG2000
)

// CameraSide is an alias for CAP_CAMERASIDE and TWEI_PAGESIDE values.
type CameraSide uint16

const (
	// CameraSideBoth is an alias for TWCS_BOTH.
	CameraSideBoth CameraSide = iota

	// CameraSideTop is an alias for TWCS_TOP.
	CameraSideTop

	// CameraSideBottom is an alias for TWCS_BOTTOM.
	CameraSideBottom
)

// DeviceEventType is an alias for CAP_DEVICEEVENT values.
type DeviceEventType uint32

// DeviceEventTypeCustomEvents is an alias for TWDE_CUSTOMEVENTS.
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

	// DeviceEventTypeCheckPowerSupplyValue is an alias for TWDE_CHECKPOWERSUPPLY.
	DeviceEventTypeCheckPowerSupplyValue

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
type PassThruDirection int32

const (
	// PassThruDirectionGet is an alias for TWDR_GET.
	PassThruDirectionGet PassThruDirection = iota + 1

	// PassThruDirectionSet is an alias for TWDR_SET.
	PassThruDirectionSet
)

// DeskEWStatus is an alias for TWEI_DESKEWSTATUS values.
type DeskEWStatus uint16

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

// Duplex is an alias for CAP_DUPLEX values.
type Duplex uint16

const (
	// DuplexNone is an alias for TWDX_NONE.
	DuplexNone Duplex = iota

	// Duplex1PassDuplex is an alias for TWDX_1PASSDUPLEX.
	Duplex1PassDuplex

	// Duplex2PassDuplex is an alias for TWDX_2PASSDUPLEX.
	Duplex2PassDuplex
)

// FeederAlignment is an alias for CAP_FEEDERALIGNMENT values.
type FeederAlignment uint16

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
type FeederType uint16

const (
	// FeederTypeGeneral is an alias for TWFE_GENERAL.
	FeederTypeGeneral FeederType = iota

	// FeederTypePhoto is an alias for TWFE_PHOTO.
	FeederTypePhoto
)

// ImageFileFormat is an alias for ICAP_IMAGEFILEFORMAT values.
type ImageFileFormat uint16

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
type FlashUsed2 uint32

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
type FeedOrder uint16

const (
	// FeedOrderFirstPageFirst is an alias for TWFO_FIRSTPAGEFIRST.
	FeedOrderFirstPageFirst FeedOrder = iota

	// FeedOrderLastPageFirst is an alias for TWFO_LASTPAGEFIRST.
	FeedOrderLastPageFirst
)

// FeederPocket is an alias for CAP_FEEDERPOCKET values.
type FeederPocket uint16

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
type FlipRotation uint16

const (
	// FlipRotationBook is an alias for TWFR_BOOK.
	FlipRotationBook FlipRotation = iota

	// FlipRotationFanFold is an alias for TWFR_FANFOLD.
	FlipRotationFanFold
)

// FilterType is an alias for ICAP_FILTER values.
type FilterType uint16

const (
	// FilterTypeRed is an alias for TWFT_RED.
	FilterTypeRed FilterType = iota

	// FilterTypeGreen is an alias for TWFT_GREEN.
	FilterTypeGreen

	// FilterTypeBlue is an alias for TWFT_BLUE.
	FilterTypeBlue

	// FilterTypeNone is an alias for TWFT_NONE.
	FilterTypeNone

	// FilterTypeWhite is an alias for TWFT_WHITE.
	FilterTypeWhite

	// FilterTypeCyan is an alias for TWFT_CYAN.
	FilterTypeCyan

	// FilterTypeMagenta is an alias for TWFT_MAGENTA.
	FilterTypeMagenta

	// FilterTypeYellow is an alias for TWFT_YELLOW.
	FilterTypeYellow

	// FilterTypeBlack is an alias for TWFT_BLACK.
	FilterTypeBlack
)

// ICCProfile is an alias for ICAP_ICCPROFILE values.
type ICCProfile uint16

const (
	// ICCProfileNone is an alias for TWIC_NONE.
	ICCProfileNone ICCProfile = iota

	// ICCProfileLink is an alias for TWIC_LINK.
	ICCProfileLink

	// ICCProfileEmbed is an alias for TWIC_EMBED.
	ICCProfileEmbed
)

// ImageFilter is an alias for ICAP_IMAGEFILTER values.
type ImageFilter uint16

const (
	// ImageFilterNone is an alias for TWIF_NONE.
	ImageFilterNone ImageFilter = iota

	// ImageFilterAuto is an alias for TWIF_AUTO.
	ImageFilterAuto

	// ImageFilterLowPass is an alias for TWIF_LOWPASS.
	ImageFilterLowPass

	// ImageFilterBandPass is an alias for TWIF_BANDPASS.
	ImageFilterBandPass

	// ImageFilterHighPass is an alias for TWIF_HIGHPASS.
	ImageFilterHighPass

	// ImageFilterText is an alias for TWIF_TEXT.
	ImageFilterText = ImageFilterBandPass

	// ImageFilterFineLine is an alias for TWIF_FINELINE.
	ImageFilterFineLine = ImageFilterHighPass
)

// ImageMerge is an alias for ICAP_IMAGEMERGE values.
type ImageMerge uint16

const (
	// ImageMergeNone is an alias for TWIM_NONE.
	ImageMergeNone ImageMerge = iota

	// ImageMergeFrontOnTop is an alias for TWIM_FRONTONTOP.
	ImageMergeFrontOnTop

	// ImageMergeFrontOnBottom is an alias for TWIM_FRONTONBOTTOM.
	ImageMergeFrontOnBottom

	// ImageMergeFrontOnLeft is an alias for TWIM_FRONTONLEFT.
	ImageMergeFrontOnLeft

	// ImageMergeFrontOnRight is an alias for TWIM_FRONTONRIGHT.
	ImageMergeFrontOnRight
)

// JobControl is an alias for CAP_JOBCONTROL values.
type JobControl uint16

const (
	// JobControlNone is an alias for TWJC_NONE.
	JobControlNone JobControl = iota

	// JobControlJSIC is an alias for TWJC_JSIC.
	JobControlJSIC

	// JobControlJSIS is an alias for TWJC_JSIS.
	JobControlJSIS

	// JobControlJSXC is an alias for TWJC_JSXC.
	JobControlJSXC

	// JobControlJSXS is an alias for TWJC_JSXS.
	JobControlJSXS
)

// JPEGQuality is an alias for ICAP_JPEGQUALITY values.
type JPEGQuality int16

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
type LightPath uint16

const (
	// LightPathReflective is an alias for TWLP_REFLECTIVE.
	LightPathReflective LightPath = iota

	// LightPathTransmissive is an alias for TWLP_TRANSMISSIVE.
	LightPathTransmissive
)

// LightSource is an alias for ICAP_LIGHTSOURCE values.
type LightSource uint16

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
type NoiseFilter uint16

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

// Orientation is an alias for ICAP_ORIENTATION values.
type Orientation uint16

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
type OverScan uint16

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
type PlanarChunky uint16

const (
	// PlanarChunkyChunky is an alias for TWPC_CHUNKY.
	PlanarChunkyChunky PlanarChunky = iota

	// PlanarChunkyPlanar is an alias for TWPC_PLANAR.
	PlanarChunkyPlanar
)

// PixelFlavor is an alias for ICAP_PIXELFLAVOR values.
type PixelFlavor uint16

const (
	// PixelFlavorChocolate is an alias for TWPF_CHOCOLATE.
	PixelFlavorChocolate PixelFlavor = iota

	// PixelFlavorVanilla is an alias for TWPF_VANILLA.
	PixelFlavorVanilla
)

// PrinterMode is an alias for CAP_PRINTERMODE values.
type PrinterMode uint16

const (
	// PrinterModeSingleString is an alias for TWPM_SINGLESTRING.
	PrinterModeSingleString PrinterMode = iota

	// PrinterModeMultiString is an alias for TWPM_MULTISTRING.
	PrinterModeMultiString

	// PrinterModeCompoundString is an alias for TWPM_COMPOUNDSTRING.
	PrinterModeCompoundString
)

// Printer is an alias for CAP_PRINTER values.
type Printer uint16

const (
	// PrinterImprinterTopBefore is an alias for TWPR_IMPRINTERTOPBEFORE.
	PrinterImprinterTopBefore Printer = iota

	// PrinterImprinterTopAfter is an alias for TWPR_IMPRINTERTOPAFTER.
	PrinterImprinterTopAfter

	// PrinterImprinterBottomBefore is an alias for TWPR_IMPRINTERBOTTOMBEFORE.
	PrinterImprinterBottomBefore

	// PrinterImprinterBottomAfter is an alias for TWPR_IMPRINTERBOTTOMAFTER.
	PrinterImprinterBottomAfter

	// PrinterEndorserTopBefore is an alias for TWPR_ENDORSERTOPBEFORE.
	PrinterEndorserTopBefore

	// PrinterEndorserTopAfter is an alias for TWPR_ENDORSERTOPAFTER.
	PrinterEndorserTopAfter

	// PrinterEndorserBottomBefore is an alias for TWPR_ENDORSERBOTTOMBEFORE.
	PrinterEndorserBottomBefore

	// PrinterEndorserBottomAfter is an alias for TWPR_ENDORSERBOTTOMAFTER.
	PrinterEndorserBottomAfter
)

// PrinterFontStyle is an alias for CAP_PRINTERFONTSTYLE values.
type PrinterFontStyle uint16

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
type PrinterIndexTrigger uint16

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
type PowerSupply int32

const (
	// PowerSupplyExternal is an alias for TWPS_EXTERNAL.
	PowerSupplyExternal PowerSupply = iota

	// PowerSupplyBattery is an alias for TWPS_BATTERY.
	PowerSupplyBattery
)

// PixelType is an alias for ICAP_PIXELTYPE values.
type PixelType int16

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
type Segmented uint16

const (
	// SegmentedNone is an alias for TWSG_NONE.
	SegmentedNone Segmented = iota

	// SegmentedAuto is an alias for TWSG_AUTO.
	SegmentedAuto

	// SegmentedManual is an alias for TWSG_MANUAL.
	SegmentedManual
)

// FilmType is an alias for ICAP_FILMTYPE values.
type FilmType uint16

const (
	// FilmTypePositive is an alias for TWFM_POSITIVE.
	FilmTypePositive FilmType = iota

	// FilmTypeNegative is an alias for TWFM_NEGATIVE.
	FilmTypeNegative
)

// DoubleFeedDetection is an alias for CAP_DOUBLEFEEDDETECTION values.
type DoubleFeedDetection uint16

const (
	// DoubleFeedDetectionUltrasonic is an alias for TWDF_ULTRASONIC.
	DoubleFeedDetectionUltrasonic DoubleFeedDetection = iota

	// DoubleFeedDetectionByLength is an alias for TWDF_BYLENGTH.
	DoubleFeedDetectionByLength

	// DoubleFeedDetectionInfrared is an alias for TWDF_INFRARED.
	DoubleFeedDetectionInfrared
)

// DoubleFeedDetectionSensitivity is an alias for CAP_DOUBLEFEEDDETECTIONSENSITIVITY values.
type DoubleFeedDetectionSensitivity uint16

const (
	// DoubleFeedDetectionSensitivityLow is an alias for TWUS_LOW.
	DoubleFeedDetectionSensitivityLow DoubleFeedDetectionSensitivity = iota

	// DoubleFeedDetectionSensitivityMedium is an alias for TWUS_MEDIUM.
	DoubleFeedDetectionSensitivityMedium

	// DoubleFeedDetectionSensitivityHigh is an alias for TWUS_HIGH.
	DoubleFeedDetectionSensitivityHigh
)

// DoubleFeedDetectionResponse is an alias for CAP_DOUBLEFEEDDETECTIONRESPONSE values.
type DoubleFeedDetectionResponse uint16

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

// MirrorValue is an alias for ICAP_MIRROR values.
type MirrorValue uint16

const (
	// MirrorValueNone is an alias for TWMR_NONE.
	MirrorValueNone MirrorValue = iota

	// MirrorValueVertical is an alias for TWMR_VERTICAL.
	MirrorValueVertical

	// MirrorValueHorizontal is an alias for TWMR_HORIZONTAL.
	MirrorValueHorizontal
)

// JPEGSubSampling is an alias for ICAP_JPEGSUBSAMPLING values.
type JPEGSubSampling uint32

const (
	// JPEGSubSampling444YCBCR is an alias for TWJS_444YCBCR.
	JPEGSubSampling444YCBCR JPEGSubSampling = iota

	// JPEGSubSampling444RGB is an alias for TWJS_444RGB.
	JPEGSubSampling444RGB

	// JPEGSubSampling422 is an alias for TWJS_422.
	JPEGSubSampling422

	// JPEGSubSampling421 is an alias for TWJS_421.
	JPEGSubSampling421

	// JPEGSubSampling411 is an alias for TWJS_411.
	JPEGSubSampling411

	// JPEGSubSampling420 is an alias for TWJS_420.
	JPEGSubSampling420

	// JPEGSubSampling410 is an alias for TWJS_410.
	JPEGSubSampling410

	// JPEGSubSampling311 is an alias for TWJS_311.
	JPEGSubSampling311
)

// PaperHandling is an alias for CAP_PAPERHANDLING values.
type PaperHandling uint16

const (
	// PaperHandlingNormal is an alias for TWPH_NORMAL.
	PaperHandlingNormal PaperHandling = iota

	// PaperHandlingFragile is an alias for TWPH_FRAGILE.
	PaperHandlingFragile

	// PaperHandlingThick is an alias for TWPH_THICK.
	PaperHandlingThick

	// PaperHandlingTrifold is an alias for TWPH_TRIFOLD.
	PaperHandlingTrifold

	// PaperHandlingPhotograph is an alias for TWPH_PHOTOGRAPH.
	PaperHandlingValuePhotograph
)

// IndicatorsMode is an alias for CAP_INDICATORSMODE values.
type IndicatorsMode uint16

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
type SupportedSize uint16

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

// TransferMechanism is an alias for ICAP_XFERMECH values.
type TransferMechanism uint16

const (
	// TransferMechanismNative is an alias for TWSX_NATIVE.
	TransferMechanismNative TransferMechanism = iota

	// TransferMechanismFile is an alias for TWSX_FILE.
	TransferMechanismFile

	// TransferMechanismMemory is an alias for TWSX_MEMORY.
	TransferMechanismMemory

	// TransferMechanismMemFile is an alias for TWSX_MEMFILE.
	TransferMechanismMemFile
)

// UnitsValue is an alias for ICAP_UNITS values.
type Units uint16

const (
	// UnitsInches is an alias for TWUN_INCHES.
	UnitsInches Units = iota

	// UnitsCentimeters is an alias for TWUN_CENTIMETERS.
	UnitsCentimeters

	// UnitsPicas is an alias for TWUN_PICAS.
	UnitsPicas

	// UnitsPoints is an alias for TWUN_POINTS.
	UnitsPoints

	// UnitsTwips is an alias for TWUN_TWIPS.
	UnitsTwips

	// UnitsPixels is an alias for TWUN_PIXELS.
	UnitsPixels

	// UnitsMillimeters is an alias for TWUN_MILLIMETERS.
	UnitsMillimeters
)
