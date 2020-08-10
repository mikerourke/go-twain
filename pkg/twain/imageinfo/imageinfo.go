package imageinfo

import "github.com/mikerourke/go-twain/pkg/twain"

type ExtImageInfoAttr twain.UInt16

const (
	// BarcodeX is an alias for TWEI_BARCODEX.
	BarcodeX ExtImageInfoAttr = 0x1200

	// BarcodeY is an alias for TWEI_BARCODEY.
	BarcodeY ExtImageInfoAttr = 0x1201

	// BarcodeText is an alias for TWEI_BARCODETEXT.
	BarcodeText ExtImageInfoAttr = 0x1202

	// BarcodeType is an alias for TWEI_BARCODETYPE.
	BarcodeType ExtImageInfoAttr = 0x1203

	// DeShadeTop is an alias for TWEI_DESHADETOP.
	DeShadeTop ExtImageInfoAttr = 0x1204

	// DeShadeLeft is an alias for TWEI_DESHADELEFT.
	DeShadeLeft ExtImageInfoAttr = 0x1205

	// DeShadeHeight is an alias for TWEI_DESHADEHEIGHT.
	DeShadeHeight ExtImageInfoAttr = 0x1206

	// DeShadeWidth is an alias for TWEI_DESHADEWIDTH.
	DeShadeWidth ExtImageInfoAttr = 0x1207

	// DeShadeSize is an alias for TWEI_DESHADESIZE.
	DeShadeSize ExtImageInfoAttr = 0x1208

	// SpecklesRemoved is an alias for TWEI_SPECKLESREMOVED.
	SpecklesRemoved ExtImageInfoAttr = 0x1209

	// HorzLineXCoord is an alias for TWEI_HORZLINEXCOORD.
	HorzLineXCoord ExtImageInfoAttr = 0x120A

	// HorzLineYCoord is an alias for TWEI_HORZLINEYCOORD.
	HorzLineYCoord ExtImageInfoAttr = 0x120B

	// HorzLineLength is an alias for TWEI_HORZLINELENGTH.
	HorzLineLength ExtImageInfoAttr = 0x120C

	// HorzLineThickness is an alias for TWEI_HORZLINETHICKNESS.
	HorzLineThickness ExtImageInfoAttr = 0x120D

	// VertLineXCoord is an alias for TWEI_VERTLINEXCOORD.
	VertLineXCoord ExtImageInfoAttr = 0x120E

	// VertLineYCoord is an alias for TWEI_VERTLINEYCOORD.
	VertLineYCoord ExtImageInfoAttr = 0x120F

	// VertLineLength is an alias for TWEI_VERTLINELENGTH.
	VertLineLength ExtImageInfoAttr = 0x1210

	// VertLineThickness is an alias for TWEI_VERTLINETHICKNESS.
	VertLineThickness ExtImageInfoAttr = 0x1211

	// PatchCode is an alias for TWEI_PATCHCODE.
	PatchCode ExtImageInfoAttr = 0x1212

	// EndorsedText is an alias for TWEI_ENDORSEDTEXT.
	EndorsedText ExtImageInfoAttr = 0x1213

	// FormConfidence is an alias for TWEI_FORMCONFIDENCE.
	FormConfidence ExtImageInfoAttr = 0x1214

	// FormTemplateMatch is an alias for TWEI_FORMTEMPLATEMATCH.
	FormTemplateMatch ExtImageInfoAttr = 0x1215

	// FormTemplatePageMatch is an alias for TWEI_FORMTEMPLATEPAGEMATCH.
	FormTemplatePageMatch ExtImageInfoAttr = 0x1216

	// FormHorzDocOffset is an alias for TWEI_FORMHORZDOCOFFSET.
	FormHorzDocOffset ExtImageInfoAttr = 0x1217

	// FormVertDocOffset is an alias for TWEI_FORMVERTDOCOFFSET.
	FormVertDocOffset ExtImageInfoAttr = 0x1218

	// BarcodeCount is an alias for TWEI_BARCODECOUNT.
	BarcodeCount ExtImageInfoAttr = 0x1219

	// BarcodeConfidence is an alias for TWEI_BARCODECONFIDENCE.
	BarcodeConfidence ExtImageInfoAttr = 0x121A

	// BarcodeRotation is an alias for TWEI_BARCODEROTATION.
	BarcodeRotation ExtImageInfoAttr = 0x121B

	// BarcodeTextLength is an alias for TWEI_BARCODETEXTLENGTH.
	BarcodeTextLength ExtImageInfoAttr = 0x121C

	// DeShadeCount is an alias for TWEI_DESHADECOUNT.
	DeShadeCount ExtImageInfoAttr = 0x121D

	// DeShadeBlackCountOld is an alias for TWEI_DESHADEBLACKCOUNTOLD.
	DeShadeBlackCountOld ExtImageInfoAttr = 0x121E

	// DeShadeBlackCountNew is an alias for TWEI_DESHADEBLACKCOUNTNEW.
	DeShadeBlackCountNew ExtImageInfoAttr = 0x121F

	// DeShadeBlackRLMin is an alias for TWEI_DESHADEBLACKRLMIN.
	DeShadeBlackRLMin ExtImageInfoAttr = 0x1220

	// DeShadeBlackRLMax is an alias for TWEI_DESHADEBLACKRLMAX.
	DeShadeBlackRLMax ExtImageInfoAttr = 0x1221

	// DeShadeWhiteCountOld is an alias for TWEI_DESHADEWHITECOUNTOLD.
	DeShadeWhiteCountOld ExtImageInfoAttr = 0x1222

	// DeShadeWhiteCountNew is an alias for TWEI_DESHADEWHITECOUNTNEW.
	DeShadeWhiteCountNew ExtImageInfoAttr = 0x1223

	// DeShadeWhiteRLMin is an alias for TWEI_DESHADEWHITERLMIN.
	DeShadeWhiteRLMin ExtImageInfoAttr = 0x1224

	// DeShadeWhiteRLAve is an alias for TWEI_DESHADEWHITERLAVE.
	DeShadeWhiteRLAve ExtImageInfoAttr = 0x1225

	// DeShadeWhiteRLMax is an alias for TWEI_DESHADEWHITERLMAX.
	DeShadeWhiteRLMax ExtImageInfoAttr = 0x1226

	// BlackSpecklesRemoved is an alias for TWEI_BLACKSPECKLESREMOVED.
	BlackSpecklesRemoved ExtImageInfoAttr = 0x1227

	// WhiteSpecklesRemoved is an alias for TWEI_WHITESPECKLESREMOVED.
	WhiteSpecklesRemoved ExtImageInfoAttr = 0x1228

	// HorzLineCount is an alias for TWEI_HORZLINECOUNT.
	HorzLineCount ExtImageInfoAttr = 0x1229

	// VertLineCount is an alias for TWEI_VERTLINECOUNT.
	VertLineCount ExtImageInfoAttr = 0x122A

	// DeSkewStatus is an alias for TWEI_DESKEWSTATUS.
	DeSkewStatus ExtImageInfoAttr = 0x122B

	// SkewOriginalAngle is an alias for TWEI_SKEWORIGINALANGLE.
	SkewOriginalAngle ExtImageInfoAttr = 0x122C

	// SkewFinalAngle is an alias for TWEI_SKEWFINALANGLE.
	SkewFinalAngle ExtImageInfoAttr = 0x122D

	// SkewConfidence is an alias for TWEI_SKEWCONFIDENCE.
	SkewConfidence ExtImageInfoAttr = 0x122E

	// SkewWindowX1 is an alias for TWEI_SKEWWINDOWX1.
	SkewWindowX1 ExtImageInfoAttr = 0x122F

	// SkewWindowY1 is an alias for TWEI_SKEWWINDOWY1.
	SkewWindowY1 ExtImageInfoAttr = 0x1230

	// SkewWindowX2 is an alias for TWEI_SKEWWINDOWX2.
	SkewWindowX2 ExtImageInfoAttr = 0x1231

	// SkewWindowY2 is an alias for TWEI_SKEWWINDOWY2.
	SkewWindowY2 ExtImageInfoAttr = 0x1232

	// SkewWindowX3 is an alias for TWEI_SKEWWINDOWX3.
	SkewWindowX3 ExtImageInfoAttr = 0x1233

	// SkewWindowY3 is an alias for TWEI_SKEWWINDOWY3.
	SkewWindowY3 ExtImageInfoAttr = 0x1234

	// SkewWindowX4 is an alias for TWEI_SKEWWINDOWX4.
	SkewWindowX4 ExtImageInfoAttr = 0x1235

	// SkewWindowY4 is an alias for TWEI_SKEWWINDOWY4.
	SkewWindowY4 ExtImageInfoAttr = 0x1236

	// BookName is an alias for TWEI_BOOKNAME.
	BookName ExtImageInfoAttr = 0x1238

	// ChapterNumber is an alias for TWEI_CHAPTERNUMBER.
	ChapterNumber ExtImageInfoAttr = 0x1239

	// DocumentNumber is an alias for TWEI_DOCUMENTNUMBER.
	DocumentNumber ExtImageInfoAttr = 0x123A

	// PageNumber is an alias for TWEI_PAGENUMBER.
	PageNumber ExtImageInfoAttr = 0x123B

	// Camera is an alias for TWEI_CAMERA.
	Camera ExtImageInfoAttr = 0x123C

	// FrameNumber is an alias for TWEI_FRAMENUMBER.
	FrameNumber ExtImageInfoAttr = 0x123D

	// Frame is an alias for TWEI_FRAME.
	Frame ExtImageInfoAttr = 0x123E

	// PixelFlavor is an alias for TWEI_PIXELFLAVOR.
	PixelFlavor ExtImageInfoAttr = 0x123F

	// ICCProfile is an alias for TWEI_ICCPROFILE.
	ICCProfile ExtImageInfoAttr = 0x1240

	// LastSegment is an alias for TWEI_LASTSEGMENT.
	LastSegment ExtImageInfoAttr = 0x1241

	// SegmentNumber is an alias for TWEI_SEGMENTNUMBER.
	SegmentNumber ExtImageInfoAttr = 0x1242

	// MagData is an alias for TWEI_MAGDATA.
	MagData ExtImageInfoAttr = 0x1243

	// MagType is an alias for TWEI_MAGTYPE.
	MagType ExtImageInfoAttr = 0x1244

	// PageSide is an alias for TWEI_PAGESIDE.
	PageSide ExtImageInfoAttr = 0x1245

	// FileSystemSource is an alias for TWEI_FILESYSTEMSOURCE.
	FileSystemSource ExtImageInfoAttr = 0x1246

	// ImageMerged is an alias for TWEI_IMAGEMERGED.
	ImageMerged ExtImageInfoAttr = 0x1247

	// MagDataLength is an alias for TWEI_MAGDATALENGTH.
	MagDataLength ExtImageInfoAttr = 0x1248

	// PaperCount is an alias for TWEI_PAPERCOUNT.
	PaperCount ExtImageInfoAttr = 0x1249

	// PrinterText is an alias for TWEI_PRINTERTEXT.
	PrinterText ExtImageInfoAttr = 0x124A

	// TwainDirectMetadata is an alias for TWEI_TWAINDIRECTMETADATA.
	TwainDirectMetadata ExtImageInfoAttr = 0x124B
)
