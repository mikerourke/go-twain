package twain

// TWFix32 (TW_FIX32) is a fixed point structure type.
type TWFix32 struct {
	Whole TWInt16
	Frac  TWUInt16
}

// TWFrame (TW_FRAME) defines a frame rectangle in ICAP_UNITS coordinates.
type TWFrame struct {
	Left   TWFix32
	Top    TWFix32
	Right  TWFix32
	Bottom TWFix32
}

// TWDecodeFunction (TW_DECODEFUNCTION) defines the parameters used for
// channel-specific transformation.
type TWDecodeFunction struct {
	StartIn     TWFix32
	BreakIn     TWFix32
	EndIn       TWFix32
	StartOut    TWFix32
	BreakOut    TWFix32
	EndOut      TWFix32
	Gamma       TWFix32
	SampleCount TWFix32
}

// TWTransformStage (TW_TRANSFORMSTAGE) stores a Fixed point number in two
// parts, a whole and a fractional part.
// TODO: Find out if array brackets are valid for struct.
type TWTransformStage struct {
	Decode [3]TWDecodeFunction
	Mix    [3][3]TWFix32
}

// TWArray (TW_ARRAY) is a container for array of values.
type TWArray struct {
	ItemType TWUInt16
	NumItems TWUInt32
	ItemList [1]TWUInt8
}

// TWAudioInfo (TW_AUDIO_INFO) is information about audio data.
type TWAudioInfo struct {
	Name     TWStr255
	Reserved TWUInt32
}

// TWCallBack2 (TW_CALLBACK2) is used to register callbacks.
type TWCallBack2 struct {
	CallBackProc TWMemRef
	RefCon       TWUIntPtr
	Message      TWInt16
}

// TWCapability (TW_CAPABILITY) is used by application to get/set capability
// from/in a data source.
type TWCapability struct {
	// Cap represents the capability ID to get or set.
	Cap TWUInt16

	// ConType is the type of container to use.
	ConType TWUInt16

	// ConHandle is the handle to container of type ConType.
	ConHandle TWHandle
}

// TWCIEPoint (TW_CIEPOINT) defines a CIE XYZ space tri-stimulus value.
type TWCIEPoint struct {
	X TWFix32
	Y TWFix32
	Z TWFix32
}

// TWCIEColor (TW_CIECOLOR) defines the mapping from an RGB color space device
// into CIE 1931 (XYZ) color space.
type TWCIEColor struct {
	ColorSpace      TWUInt16
	LowEndian       TWInt16
	DeviceDependent TWInt16
	VersionNumber   TWInt32
	StageABC        TWTransformStage
	StageLMN        TWTransformStage
	WhitePoint      TWCIEPoint
	BlackPoint      TWCIEPoint
	WhitePaper      TWCIEPoint
	BlackInk        TWCIEPoint
	Samples         [1]TWFix32
}

// TWCustomDSData (TW_CUSTOMDSDATA) allows for a data source and application to
// pass custom data to each other.
type TWCustomDSData struct {
	InfoLength TWUInt32
	DataHandle TWHandle
}

// TWDeviceEvent (TW_DEVICEEVENT) provides information about the Event that was
// raised by the Source.
type TWDeviceEvent struct {
	Event                  TWUInt32
	DeviceName             TWStr255
	BatteryMinutes         TWUInt32
	BatteryPercentage      TWInt16
	PowerSupply            TWInt32
	XResolution            TWFix32
	YResolution            TWFix32
	FlashUsed2             TWUInt32
	AutomaticCapture       TWUInt32
	TimeBeforeFirstCapture TWUInt32
	TimeBetweenCaptures    TWUInt32
}

// TWElement8 (TW_ELEMENT8) holds the tri-stimulus color palette information
// for TWPalette8 structures.
type TWElement8 struct {
	Index    TWUInt8
	Channel1 TWUInt8
	Channel2 TWUInt8
	Channel3 TWUInt8
}

// TWEnumeration (TW_ENUMERATION) stores a group of individual values describing
// a capability.
type TWEnumeration struct {
	ItemType     TWUInt16
	NumItems     TWUInt32
	CurrentIndex TWUInt32
	DefaultIndex TWUInt32
	ItemList     [1]TWUInt8
}

// TWEvent (TW_EVENT) is used to pass application events/messages from the
// application to the Source.
type TWEvent struct {
	Event   TWMemRef
	Message TWUInt16
}

// TWInfo (TW_INFO) is used to pass specific information between the data source
// and the application.
// TODO: Find out if this will break with union. See https://stackoverflow.com/questions/31557539/golang-how-to-simulate-union-type-efficiently
type TWInfo struct {
	InfoID   TWUInt16
	ItemType TWUInt16
	NumItems TWUInt16

	// ReturnCode represents a union type that contains a ReturnCode or CondCode,
	// but the CondCode has been deprecated.
	ReturnCode TWUInt16
	Item       TWUIntPtr
}

// TWExtImageInfo (TW_EXTIMAGEINFO) provides extended image information.
type TWExtImageInfo struct {
	NumInfos TWUInt32
	Info     [1]TWInfo
}

// TWFileSystem (TW_FILESYSTEM) provides information about the currently
// selected device.
type TWFileSystem struct {
	InputName  TWStr255
	OutputName TWStr255
	Context    TWMemRef

	// RecursiveOrSubdirectories is a union that can be either an int (for
	// Recursive) or a TWBool for Subdirectories.
	RecursiveOrSubdirectories interface{}

	// FileTypeOrFileSystemType is a union that can be either a TWInt32 for
	// FileType or a TWUInt32 for FileSystemType.
	FileTypeOrFileSystemType interface{}
	Size                     TWUInt32
	CreateTimeDate           TWStr32
	ModifiedTimeDate         TWStr32
	FreeSpace                TWUInt32
	NewImageSize             TWInt32
	NumberOfFiles            TWUInt32
	NumberOfSnippets         TWUInt32
	DeviceGroupMask          TWUInt32
	Reserved                 [508]TWInt8
}

// TWGrayResponse (TW_GRAYRESPONSE) is used by the application to specify a set
// of mapping values to be applied to grayscale data.
type TWGrayResponse struct {
	Response [1]TWElement8
}

// TWVersion (TW_VERSION) represents a general way to describe the version of
// software that is running.
type TWVersion struct {
	MajorNum TWUInt16
	MinorNum TWUInt16
	Language TWUInt16
	Country  TWUInt16
	Info     TWStr32
}

// TWImageInfo (TW_IMAGEINFO) describes the "real" image data, that is, the
// complete image being transferred between the Source and application.
type TWImageInfo struct {
	XResolution     TWFix32
	YResolution     TWFix32
	ImageWidth      TWInt32
	ImageLength     TWInt32
	SamplesPerPixel TWInt16
	BitsPerSample   [8]TWInt16
	BitsPerPixel    TWInt16
	Planar          TWBool
	PixelType       TWInt16
	Compression     TWUInt16
}

// TWImageLayout (TW_IMAGELAYOUT) contains information about the original size
// of the acquired image.
type TWImageLayout struct {
	Frame          TWFrame
	DocumentNumber TWUInt32
	PageNumber     TWUInt32
	FrameNumber    TWUInt32
}

// TWMemory (TW_MEMORY) provides information for managing memory buffers.
type TWMemory struct {
	Flags  TWUInt32
	Length TWUInt32
	TheMem TWMemRef
}

// TWImageMemTransfer (TW_IMAGEMEMXFER) describes the form of the acquired data
// being passed from the Source to the application.
type TWImageMemTransfer struct {
	Compression  TWUInt16
	BytesPerRow  TWUInt32
	Columns      TWUInt32
	Rows         TWUInt32
	XOffset      TWUInt32
	YOffset      TWUInt32
	BytesWritten TWUInt32
	Memory       TWMemory
}

// TWJPEGCompression (TW_JPEGCOMPRESSION) describes the information necessary
// to transfer a JPEG-compressed image.
type TWJPEGCompression struct {
	ColorSpace       TWUInt16
	SubSampling      TWUInt32
	NumComponents    TWUInt16
	RestartFrequency TWUInt16
	QuantMap         [4]TWUInt16
	QuantTable       [4]TWMemory
	HuffmanMap       [4]TWUInt16
	HuffmanDC        [2]TWMemory
	HuffmanAC        [2]TWMemory
}

// TWMetrics (TW_METRICS) collects scanning metrics after returning to state 4.
type TWMetrics struct {
	SizeOf     TWUInt32
	ImageCount TWUInt32
	SheetCount TWUInt32
}

// TWOneValue (TW_ONEVALUE) stores a single value (item) which describes a
// capability.
type TWOneValue struct {
	ItemType TWUInt16
	Item     TWUInt32
}

// TWPalette8 (TW_PALETTE8) holds the color palette information.
type TWPalette8 struct {
	NumColors   TWUInt16
	PaletteType TWUInt16
	Colors      [256]TWElement8
}

// TWPassThru (TW_PASSTHRU) is used to bypass the TWAIN protocol when
// communicating with a device.
type TWPassThru struct {
	pCommand             TWMemRef
	CommandBytes         TWUInt32
	Direction            TWInt32
	pData                TWMemRef
	DataBytes            TWUInt32
	DataBytesTransferred TWUInt32
}

// TWPendingTransfers (TW_PENDINGXFERS) tells the application how many more
// complete transfers the Source currently has available.
// TODO: Find out if union breaks this struct.
type TWPendingTransfers struct {
	Count         TWUInt16
	EOJOrReserved TWUInt32
}

// TWRange (TW_RANGE) stores a range of individual values describing a
// capability.
type TWRange struct {
	ItemType     TWUInt16
	MinValue     TWUInt32
	MaxValue     TWUInt32
	StepSize     TWUInt32
	DefaultValue TWUInt32
	CurrentValue TWUInt32
}

// TWRGBResponse (TW_RGBRESPONSE) is used by the application to specify a set
// of mapping values to be applied to RGB color data.
type TWRGBResponse struct {
	Response [1]TWElement8
}

// TWSetupFileTransfer (TW_SETUPFILEXFER) describes the file format and file
// specification information for a transfer through a disk file.
type TWSetupFileTransfer struct {
	FileName TWStr255
	Format   TWUInt16
	VRefNum  TWInt16
}

// TWSetupMemTransfer (TW_SETUPMEMXFER) provides the application information
// about the Source's requirements and preferences regarding allocation of
// transfer buffer(s).
type TWSetupMemTransfer struct {
	MinBufSize TWUInt32
	MaxBufSize TWUInt32
	Preferred  TWUInt32
}

// TWStatus (TW_STATUS) describes the status of a source.
type TWStatus struct {
	ConditionCode TWUInt16
	Data          TWUInt16
}

// TWStatusUTF8 (TW_STATUSUTF8) translates the contents of Status into a
// localized UTF-8 string.
type TWStatusUTF8 struct {
	Status     TWStatus
	Size       TWUInt32
	UTF8String TWHandle
}

// TWTwainDirect (TW_TWAINDIRECT) is not documented.
type TWTwainDirect struct {
	SizeOf               TWUInt32
	CommunicationManager TWUInt16
	Send                 TWHandle
	SendSize             TWUInt32
	Receive              TWHandle
	ReceiveSize          TWUInt32
}

// TWUserInterface (TW_USERINTERFACE) is used to handle the user interface
// coordination between an application and a Source.
type TWUserInterface struct {
	ShowUI       TWBool
	ModalUI      TWBool
	ParentHandle TWHandle
}
