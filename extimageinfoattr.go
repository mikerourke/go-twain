package twain

// ExtImageInfoAttr is an alias for TWEI_ values.
type ExtImageInfoAttr uint16

const (
	// ExtImageBarcodeX is an alias for TWEI_BARCODEX.
	ExtImageBarcodeX ExtImageInfoAttr = 0x1200

	// ExtImageBarcodeY is an alias for TWEI_BARCODEY.
	ExtImageBarcodeY ExtImageInfoAttr = 0x1201

	// ExtImageBarcodeText is an alias for TWEI_BARCODETEXT.
	ExtImageBarcodeText ExtImageInfoAttr = 0x1202

	// ExtImageBarcodeType is an alias for TWEI_BARCODETYPE.
	ExtImageBarcodeType ExtImageInfoAttr = 0x1203

	// ExtImageDeShadeTop is an alias for TWEI_DESHADETOP.
	ExtImageDeShadeTop ExtImageInfoAttr = 0x1204

	// ExtImageDeShadeLeft is an alias for TWEI_DESHADELEFT.
	ExtImageDeShadeLeft ExtImageInfoAttr = 0x1205

	// ExtImageDeShadeHeight is an alias for TWEI_DESHADEHEIGHT.
	ExtImageDeShadeHeight ExtImageInfoAttr = 0x1206

	// ExtImageDeShadeWidth is an alias for TWEI_DESHADEWIDTH.
	ExtImageDeShadeWidth ExtImageInfoAttr = 0x1207

	// ExtImageDeShadeSize is an alias for TWEI_DESHADESIZE.
	ExtImageDeShadeSize ExtImageInfoAttr = 0x1208

	// ExtImageSpecklesRemoved is an alias for TWEI_SPECKLESREMOVED.
	ExtImageSpecklesRemoved ExtImageInfoAttr = 0x1209

	// ExtImageHorzLineXCoord is an alias for TWEI_HORZLINEXCOORD.
	ExtImageHorzLineXCoord ExtImageInfoAttr = 0x120A

	// ExtImageHorzLineYCoord is an alias for TWEI_HORZLINEYCOORD.
	ExtImageHorzLineYCoord ExtImageInfoAttr = 0x120B

	// ExtImageHorzLineLength is an alias for TWEI_HORZLINELENGTH.
	ExtImageHorzLineLength ExtImageInfoAttr = 0x120C

	// ExtImageHorzLineThickness is an alias for TWEI_HORZLINETHICKNESS.
	ExtImageHorzLineThickness ExtImageInfoAttr = 0x120D

	// ExtImageVertLineXCoord is an alias for TWEI_VERTLINEXCOORD.
	ExtImageVertLineXCoord ExtImageInfoAttr = 0x120E

	// ExtImageVertLineYCoord is an alias for TWEI_VERTLINEYCOORD.
	ExtImageVertLineYCoord ExtImageInfoAttr = 0x120F

	// ExtImageVertLineLength is an alias for TWEI_VERTLINELENGTH.
	ExtImageVertLineLength ExtImageInfoAttr = 0x1210

	// ExtImageVertLineThickness is an alias for TWEI_VERTLINETHICKNESS.
	ExtImageVertLineThickness ExtImageInfoAttr = 0x1211

	// ExtImagePatchCode is an alias for TWEI_PATCHCODE.
	ExtImagePatchCode ExtImageInfoAttr = 0x1212

	// ExtImageEndorsedText is an alias for TWEI_ENDORSEDTEXT.
	ExtImageEndorsedText ExtImageInfoAttr = 0x1213

	// ExtImageFormConfidence is an alias for TWEI_FORMCONFIDENCE.
	ExtImageFormConfidence ExtImageInfoAttr = 0x1214

	// ExtImageFormTemplateMatch is an alias for TWEI_FORMTEMPLATEMATCH.
	ExtImageFormTemplateMatch ExtImageInfoAttr = 0x1215

	// ExtImageFormTemplatePageMatch is an alias for TWEI_FORMTEMPLATEPAGEMATCH.
	ExtImageFormTemplatePageMatch ExtImageInfoAttr = 0x1216

	// ExtImageFormHorzDocOffset is an alias for TWEI_FORMHORZDOCOFFSET.
	ExtImageFormHorzDocOffset ExtImageInfoAttr = 0x1217

	// ExtImageFormVertDocOffset is an alias for TWEI_FORMVERTDOCOFFSET.
	ExtImageFormVertDocOffset ExtImageInfoAttr = 0x1218

	// ExtImageBarcodeCount is an alias for TWEI_BARCODECOUNT.
	ExtImageBarcodeCount ExtImageInfoAttr = 0x1219

	// ExtImageBarcodeConfidence is an alias for TWEI_BARCODECONFIDENCE.
	ExtImageBarcodeConfidence ExtImageInfoAttr = 0x121A

	// ExtImageBarcodeRotation is an alias for TWEI_BARCODEROTATION.
	ExtImageBarcodeRotation ExtImageInfoAttr = 0x121B

	// ExtImageBarcodeTextLength is an alias for TWEI_BARCODETEXTLENGTH.
	ExtImageBarcodeTextLength ExtImageInfoAttr = 0x121C

	// ExtImageDeShadeCount is an alias for TWEI_DESHADECOUNT.
	ExtImageDeShadeCount ExtImageInfoAttr = 0x121D

	// ExtImageDeShadeBlackCountOld is an alias for TWEI_DESHADEBLACKCOUNTOLD.
	ExtImageDeShadeBlackCountOld ExtImageInfoAttr = 0x121E

	// ExtImageDeShadeBlackCountNew is an alias for TWEI_DESHADEBLACKCOUNTNEW.
	ExtImageDeShadeBlackCountNew ExtImageInfoAttr = 0x121F

	// ExtImageDeShadeBlackRLMin is an alias for TWEI_DESHADEBLACKRLMIN.
	ExtImageDeShadeBlackRLMin ExtImageInfoAttr = 0x1220

	// ExtImageDeShadeBlackRLMax is an alias for TWEI_DESHADEBLACKRLMAX.
	ExtImageDeShadeBlackRLMax ExtImageInfoAttr = 0x1221

	// ExtImageDeShadeWhiteCountOld is an alias for TWEI_DESHADEWHITECOUNTOLD.
	ExtImageDeShadeWhiteCountOld ExtImageInfoAttr = 0x1222

	// ExtImageDeShadeWhiteCountNew is an alias for TWEI_DESHADEWHITECOUNTNEW.
	ExtImageDeShadeWhiteCountNew ExtImageInfoAttr = 0x1223

	// ExtImageDeShadeWhiteRLMin is an alias for TWEI_DESHADEWHITERLMIN.
	ExtImageDeShadeWhiteRLMin ExtImageInfoAttr = 0x1224

	// ExtImageDeShadeWhiteRLAve is an alias for TWEI_DESHADEWHITERLAVE.
	ExtImageDeShadeWhiteRLAve ExtImageInfoAttr = 0x1225

	// ExtImageDeShadeWhiteRLMax is an alias for TWEI_DESHADEWHITERLMAX.
	ExtImageDeShadeWhiteRLMax ExtImageInfoAttr = 0x1226

	// ExtImageBlackSpecklesRemoved is an alias for TWEI_BLACKSPECKLESREMOVED.
	ExtImageBlackSpecklesRemoved ExtImageInfoAttr = 0x1227

	// ExtImageWhiteSpecklesRemoved is an alias for TWEI_WHITESPECKLESREMOVED.
	ExtImageWhiteSpecklesRemoved ExtImageInfoAttr = 0x1228

	// ExtImageHorzLineCount is an alias for TWEI_HORZLINECOUNT.
	ExtImageHorzLineCount ExtImageInfoAttr = 0x1229

	// ExtImageVertLineCount is an alias for TWEI_VERTLINECOUNT.
	ExtImageVertLineCount ExtImageInfoAttr = 0x122A

	// ExtImageDeSkewStatus is an alias for TWEI_DESKEWSTATUS.
	ExtImageDeSkewStatus ExtImageInfoAttr = 0x122B

	// ExtImageSkewOriginalAngle is an alias for TWEI_SKEWORIGINALANGLE.
	ExtImageSkewOriginalAngle ExtImageInfoAttr = 0x122C

	// ExtImageSkewFinalAngle is an alias for TWEI_SKEWFINALANGLE.
	ExtImageSkewFinalAngle ExtImageInfoAttr = 0x122D

	// ExtImageSkewConfidence is an alias for TWEI_SKEWCONFIDENCE.
	ExtImageSkewConfidence ExtImageInfoAttr = 0x122E

	// ExtImageSkewWindowX1 is an alias for TWEI_SKEWWINDOWX1.
	ExtImageSkewWindowX1 ExtImageInfoAttr = 0x122F

	// ExtImageSkewWindowY1 is an alias for TWEI_SKEWWINDOWY1.
	ExtImageSkewWindowY1 ExtImageInfoAttr = 0x1230

	// ExtImageSkewWindowX2 is an alias for TWEI_SKEWWINDOWX2.
	ExtImageSkewWindowX2 ExtImageInfoAttr = 0x1231

	// ExtImageSkewWindowY2 is an alias for TWEI_SKEWWINDOWY2.
	ExtImageSkewWindowY2 ExtImageInfoAttr = 0x1232

	// ExtImageSkewWindowX3 is an alias for TWEI_SKEWWINDOWX3.
	ExtImageSkewWindowX3 ExtImageInfoAttr = 0x1233

	// ExtImageSkewWindowY3 is an alias for TWEI_SKEWWINDOWY3.
	ExtImageSkewWindowY3 ExtImageInfoAttr = 0x1234

	// ExtImageSkewWindowX4 is an alias for TWEI_SKEWWINDOWX4.
	ExtImageSkewWindowX4 ExtImageInfoAttr = 0x1235

	// ExtImageSkewWindowY4 is an alias for TWEI_SKEWWINDOWY4.
	ExtImageSkewWindowY4 ExtImageInfoAttr = 0x1236

	// ExtImageBookName is an alias for TWEI_BOOKNAME.
	ExtImageBookName ExtImageInfoAttr = 0x1238

	// ExtImageChapterNumber is an alias for TWEI_CHAPTERNUMBER.
	ExtImageChapterNumber ExtImageInfoAttr = 0x1239

	// ExtImageDocumentNumber is an alias for TWEI_DOCUMENTNUMBER.
	ExtImageDocumentNumber ExtImageInfoAttr = 0x123A

	// ExtImagePageNumber is an alias for TWEI_PAGENUMBER.
	ExtImagePageNumber ExtImageInfoAttr = 0x123B

	// ExtImageCamera is an alias for TWEI_CAMERA.
	ExtImageCamera ExtImageInfoAttr = 0x123C

	// ExtImageFrameNumber is an alias for TWEI_FRAMENUMBER.
	ExtImageFrameNumber ExtImageInfoAttr = 0x123D

	// ExtImageFrame is an alias for TWEI_FRAME.
	ExtImageFrame ExtImageInfoAttr = 0x123E

	// ExtImagePixelFlavor is an alias for TWEI_PIXELFLAVOR.
	ExtImagePixelFlavor ExtImageInfoAttr = 0x123F

	// ExtImageICCProfile is an alias for TWEI_ICCPROFILE.
	ExtImageICCProfile ExtImageInfoAttr = 0x1240

	// ExtImageLastSegment is an alias for TWEI_LASTSEGMENT.
	ExtImageLastSegment ExtImageInfoAttr = 0x1241

	// ExtImageSegmentNumber is an alias for TWEI_SEGMENTNUMBER.
	ExtImageSegmentNumber ExtImageInfoAttr = 0x1242

	// ExtImageMagData is an alias for TWEI_MAGDATA.
	ExtImageMagData ExtImageInfoAttr = 0x1243

	// ExtImageMagType is an alias for TWEI_MAGTYPE.
	ExtImageMagType ExtImageInfoAttr = 0x1244

	// ExtImagePageSide is an alias for TWEI_PAGESIDE.
	ExtImagePageSide ExtImageInfoAttr = 0x1245

	// ExtImageFileSystemSource is an alias for TWEI_FILESYSTEMSOURCE.
	ExtImageFileSystemSource ExtImageInfoAttr = 0x1246

	// ExtImageImageMerged is an alias for TWEI_IMAGEMERGED.
	ExtImageImageMerged ExtImageInfoAttr = 0x1247

	// ExtImageMagDataLength is an alias for TWEI_MAGDATALENGTH.
	ExtImageMagDataLength ExtImageInfoAttr = 0x1248

	// ExtImagePaperCount is an alias for TWEI_PAPERCOUNT.
	ExtImagePaperCount ExtImageInfoAttr = 0x1249

	// ExtImagePrinterText is an alias for TWEI_PRINTERTEXT.
	ExtImagePrinterText ExtImageInfoAttr = 0x124A

	// ExtImageTwainDirectMetadata is an alias for TWEI_TWAINDIRECTMETADATA.
	ExtImageTwainDirectMetadata ExtImageInfoAttr = 0x124B
)
