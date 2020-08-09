package twain

import "github.com/mikerourke/go-twain/pkg/twain/countries"

// Fix32 (TW_FIX32) is a fixed point structure type.
type Fix32 struct {
	Whole Int16
	Frac  UInt16
}

// Frame (TW_FRAME) defines a frame rectangle in ICAP_UNITS coordinates.
type Frame struct {
	Left   Fix32
	Top    Fix32
	Right  Fix32
	Bottom Fix32
}

// DecodeFunction (TW_DECODEFUNCTION) defines the parameters used for
// channel-specific transformation.
type DecodeFunction struct {
	StartIn     Fix32
	BreakIn     Fix32
	EndIn       Fix32
	StartOut    Fix32
	BreakOut    Fix32
	EndOut      Fix32
	Gamma       Fix32
	SampleCount Fix32
}

// TransformStage (TW_TRANSFORMSTAGE) stores a Fixed point number in two
// parts, a whole and a fractional part.
// TODO: Find out if array brackets are valid for struct.
type TransformStage struct {
	Decode [3]DecodeFunction
	Mix    [3][3]Fix32
}

// Array (TW_ARRAY) is a container for array of values.
type Array struct {
	ItemType UInt16
	NumItems UInt32
	ItemList [1]UInt8
}

// AudioInfo (TW_AUDIO_INFO) is information about audio data.
type AudioInfo struct {
	Name     Str255
	Reserved UInt32
}

// CallBack2 (TW_CALLBACK2) is used to register callbacks.
type CallBack2 struct {
	CallBackProc MemRef
	RefCon       UIntPtr
	Message      Int16
}

// Capability (TW_CAPABILITY) is used by application to get/set capability
// from/in a data source.
type Capability struct {
	// Cap represents the capability ID to get or set.
	Cap UInt16

	// ConType is the type of container to use.
	ConType ContainerType

	// ConHandle is the handle to container of type ConType.
	ConHandle Handle
}

// CIEPoint (TW_CIEPOINT) defines a CIE XYZ space tri-stimulus value.
type CIEPoint struct {
	X Fix32
	Y Fix32
	Z Fix32
}

// CIEColor (TW_CIECOLOR) defines the mapping from an RGB color space device
// into CIE 1931 (XYZ) color space.
type CIEColor struct {
	ColorSpace      UInt16
	LowEndian       Int16
	DeviceDependent Int16
	VersionNumber   Int32
	StageABC        TransformStage
	StageLMN        TransformStage
	WhitePoint      CIEPoint
	BlackPoint      CIEPoint
	WhitePaper      CIEPoint
	BlackInk        CIEPoint
	Samples         [1]Fix32
}

// CustomDSData (TW_CUSTOMDSDATA) allows for a data source and application to
// pass custom data to each other.
type CustomDSData struct {
	InfoLength UInt32
	DataHandle Handle
}

// DeviceEvent (TW_DEVICEEVENT) provides information about the Event that was
// raised by the Source.
type DeviceEvent struct {
	Event                  UInt32
	DeviceName             Str255
	BatteryMinutes         UInt32
	BatteryPercentage      Int16
	PowerSupply            Int32
	XResolution            Fix32
	YResolution            Fix32
	FlashUsed2             UInt32
	AutomaticCapture       UInt32
	TimeBeforeFirstCapture UInt32
	TimeBetweenCaptures    UInt32
}

// Element8 (TW_ELEMENT8) holds the tri-stimulus color palette information
// for Palette8 structures.
type Element8 struct {
	Index    UInt8
	Channel1 UInt8
	Channel2 UInt8
	Channel3 UInt8
}

// Enumeration (TW_ENUMERATION) stores a group of individual values describing
// a capability.
type Enumeration struct {
	ItemType     UInt16
	NumItems     UInt32
	CurrentIndex UInt32
	DefaultIndex UInt32
	ItemList     [1]UInt8
}

// Event (TW_EVENT) is used to pass application events/messages from the
// application to the Source.
type Event struct {
	Event   MemRef
	Message UInt16
}

// Info (TW_INFO) is used to pass specific information between the data source
// and the application.
// TODO: Find out if this will break with union. See https://stackoverflow.com/questions/31557539/golang-how-to-simulate-union-type-efficiently
type Info struct {
	InfoID   UInt16
	ItemType UInt16
	NumItems UInt16

	// ReturnCode represents a union type that contains a ReturnCode or CondCode,
	// but the CondCode has been deprecated.
	ReturnCode UInt16
	Item       UIntPtr
}

// ExtImageInfo (TW_EXTIMAGEINFO) provides extended image information.
type ExtImageInfo struct {
	NumInfos UInt32
	Info     [1]Info
}

// FileSystem (TW_FILESYSTEM) provides information about the currently
// selected device.
type FileSystem struct {
	InputName  Str255
	OutputName Str255
	Context    MemRef

	// RecursiveOrSubdirectories is a union that can be either an int (for
	// Recursive) or a Bool for Subdirectories.
	RecursiveOrSubdirectories interface{}

	// FileTypeOrFileSystemType is a union that can be either a Int32 for
	// FileType or a UInt32 for FileSystemType.
	FileTypeOrFileSystemType interface{}
	Size                     UInt32
	CreateTimeDate           Str32
	ModifiedTimeDate         Str32
	FreeSpace                UInt32
	NewImageSize             Int32
	NumberOfFiles            UInt32
	NumberOfSnippets         UInt32
	DeviceGroupMask          UInt32
	Reserved                 [508]Int8
}

// GrayResponse (TW_GRAYRESPONSE) is used by the application to specify a set
// of mapping values to be applied to grayscale data.
type GrayResponse struct {
	Response [1]Element8
}

// Version (TW_VERSION) represents a general way to describe the version of
// software that is running.
type Version struct {
	MajorNum UInt16
	MinorNum UInt16
	Language UInt16
	Country  countries.Country
	Info     Str32
}

// ImageInfo (TW_IMAGEINFO) describes the "real" image data, that is, the
// complete image being transferred between the Source and application.
type ImageInfo struct {
	XResolution     Fix32
	YResolution     Fix32
	ImageWidth      Int32
	ImageLength     Int32
	SamplesPerPixel Int16
	BitsPerSample   [8]Int16
	BitsPerPixel    Int16
	Planar          Bool
	PixelType       Int16
	Compression     UInt16
}

// ImageLayout (TW_IMAGELAYOUT) contains information about the original size
// of the acquired image.
type ImageLayout struct {
	Frame          Frame
	DocumentNumber UInt32
	PageNumber     UInt32
	FrameNumber    UInt32
}

// Memory (TW_MEMORY) provides information for managing memory buffers.
type Memory struct {
	Flags  MemoryFlag
	Length UInt32
	TheMem MemRef
}

// ImageMemTransfer (TW_IMAGEMEMXFER) describes the form of the acquired data
// being passed from the Source to the application.
type ImageMemTransfer struct {
	Compression  UInt16
	BytesPerRow  UInt32
	Columns      UInt32
	Rows         UInt32
	XOffset      UInt32
	YOffset      UInt32
	BytesWritten UInt32
	Memory       Memory
}

// JPEGCompression (TW_JPEGCOMPRESSION) describes the information necessary
// to transfer a JPEG-compressed image.
type JPEGCompression struct {
	ColorSpace       UInt16
	SubSampling      UInt32
	NumComponents    UInt16
	RestartFrequency UInt16
	QuantMap         [4]UInt16
	QuantTable       [4]Memory
	HuffmanMap       [4]UInt16
	HuffmanDC        [2]Memory
	HuffmanAC        [2]Memory
}

// Metrics (TW_METRICS) collects scanning metrics after returning to state 4.
type Metrics struct {
	SizeOf     UInt32
	ImageCount UInt32
	SheetCount UInt32
}

// OneValue (TW_ONEVALUE) stores a single value (item) which describes a
// capability.
type OneValue struct {
	ItemType UInt16
	Item     UInt32
}

// Palette8 (TW_PALETTE8) holds the color palette information.
type Palette8 struct {
	NumColors   UInt16
	PaletteType UInt16
	Colors      [256]Element8
}

// PassThru (TW_PASSTHRU) is used to bypass the TWAIN protocol when
// communicating with a device.
type PassThru struct {
	Command              MemRef
	CommandBytes         UInt32
	Direction            Int32
	DataHandle           MemRef
	DataBytes            UInt32
	DataBytesTransferred UInt32
}

// PendingTransfers (TW_PENDINGXFERS) tells the application how many more
// complete transfers the Source currently has available.
// TODO: Find out if union breaks this struct.
type PendingTransfers struct {
	Count         UInt16
	EOJOrReserved UInt32
}

// Range (TW_RANGE) stores a range of individual values describing a
// capability.
type Range struct {
	ItemType     UInt16
	MinValue     UInt32
	MaxValue     UInt32
	StepSize     UInt32
	DefaultValue UInt32
	CurrentValue UInt32
}

// RGBResponse (TW_RGBRESPONSE) is used by the application to specify a set
// of mapping values to be applied to RGB color data.
type RGBResponse struct {
	Response [1]Element8
}

// SetupFileTransfer (TW_SETUPFILEXFER) describes the file format and file
// specification information for a transfer through a disk file.
type SetupFileTransfer struct {
	FileName Str255
	Format   UInt16
	VRefNum  Int16
}

// SetupMemTransfer (TW_SETUPMEMXFER) provides the application information
// about the Source's requirements and preferences regarding allocation of
// transfer buffer(s).
type SetupMemTransfer struct {
	MinBufSize UInt32
	MaxBufSize UInt32
	Preferred  UInt32
}

// Status (TW_STATUS) describes the status of a source.
type Status struct {
	ConditionCode UInt16
	Data          UInt16
}

// StatusUTF8 (TW_STATUSUTF8) translates the contents of Status into a
// localized UTF-8 string.
type StatusUTF8 struct {
	Status     Status
	Size       UInt32
	UTF8String Handle
}

// TwainDirect (TW_TWAINDIRECT) is not documented.
type TwainDirect struct {
	SizeOf               UInt32
	CommunicationManager UInt16
	Send                 Handle
	SendSize             UInt32
	Receive              Handle
	ReceiveSize          UInt32
}

// UserInterface (TW_USERINTERFACE) is used to handle the user interface
// coordination between an application and a Source.
type UserInterface struct {
	ShowUI       Bool
	ModalUI      Bool
	ParentHandle Handle
}
