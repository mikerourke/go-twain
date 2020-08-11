package twain

import "unsafe"

// Note that we don't have TW_INT8, TW_INT16, TW_UINT8, etc. defined here because
// Go already takes care of those types automatically based on the target
// platform. The TWAIN header file assigns them to accommodate for which
// compiler/platform is being used.

// Bool is an alias for TW_BOOL.
type Bool uint16

// TODO: Check if these types are valid.
// Handle is an alias for TW_HANDLE.
type Handle unsafe.Pointer

// MemRef is an alias for TW_MEMREF.
type MemRef *uint8

// Byte is an alias for BYTE.
type Byte uint8

// Fix32 (TW_FIX32) is a fixed point structure type.
type Fix32 struct {
	Whole int16
	Frac  uint16
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
	ItemType uint16
	NumItems uint32
	ItemList [1]uint8
}

// AudioInfo (TW_AUDIO_INFO) is information about audio data.
type AudioInfo struct {
	Name     Str255
	Reserved uint32
}

// Callback2 (TW_CALLBACK2) is used to register callbacks.
type Callback2 struct {
	CallbackProc MemRef
	RefCon       uintptr
	Message      int16
}

// Capability (TW_CAPABILITY) is used by application to get/set capability
// from/in a data source.
type Capability struct {
	// Cap represents the capability ID to get or set.
	Cap uint16

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
	ColorSpace      uint16
	LowEndian       int16
	DeviceDependent int16
	VersionNumber   int32
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
	InfoLength uint32
	DataHandle Handle
}

// DeviceEvent (TW_DEVICEEVENT) provides information about the Event that was
// raised by the Source.
type DeviceEvent struct {
	Event                  uint32
	DeviceName             Str255
	BatteryMinutes         uint32
	BatteryPercentage      int16
	PowerSupply            int32
	XResolution            Fix32
	YResolution            Fix32
	FlashUsed2             uint32
	AutomaticCapture       uint32
	TimeBeforeFirstCapture uint32
	TimeBetweenCaptures    uint32
}

// Element8 (TW_ELEMENT8) holds the tri-stimulus color palette information
// for Palette8 structures.
type Element8 struct {
	Index    uint8
	Channel1 uint8
	Channel2 uint8
	Channel3 uint8
}

// Enumeration (TW_ENUMERATION) stores a group of individual values describing
// a capability.
type Enumeration struct {
	ItemType     uint16
	NumItems     uint32
	CurrentIndex uint32
	DefaultIndex uint32
	ItemList     [1]uint8
}

// Event (TW_EVENT) is used to pass application events/messages from the
// application to the Source.
type Event struct {
	Event   MemRef
	Message uint16
}

// Info (TW_INFO) is used to pass specific information between the data source
// and the application.
// TODO: Find out if this will break with union. See https://stackoverflow.com/questions/31557539/golang-how-to-simulate-union-type-efficiently
type Info struct {
	InfoID   uint16
	ItemType uint16
	NumItems uint16

	// ReturnCode represents a union type that contains a ReturnCode or CondCode,
	// but the CondCode has been deprecated.
	ReturnCode [2]byte // uint16
	Item       uintptr
}

// ExtImageInfo (TW_EXTIMAGEINFO) provides extended image information.
type ExtImageInfo struct {
	NumInfos uint32
	Info     [1]Info
}

// FileSystem (TW_FILESYSTEM) provides information about the currently
// selected device.
// TODO: Find out if unions break this struct.
type FileSystem struct {
	InputName  Str255
	OutputName Str255
	Context    MemRef

	// RecursiveOrSubdirectories is a union that can be either an int (for
	// Recursive) or a Bool for Subdirectories.
	RecursiveOrSubdirectories [4]byte // int or Bool

	// FileTypeOrFileSystemType is a union that can be either a int32 for
	// FileType or a uint32 for FileSystemType.
	FileTypeOrFileSystemType [4]byte // int32 or uint32
	Size                     uint32
	CreateTimeDate           Str32
	ModifiedTimeDate         Str32
	FreeSpace                uint32
	NewImageSize             int32
	NumberOfFiles            uint32
	NumberOfSnippets         uint32
	DeviceGroupMask          uint32
	Reserved                 [508]int8
}

// GrayResponse (TW_GRAYRESPONSE) is used by the application to specify a set
// of mapping values to be applied to grayscale data.
type GrayResponse struct {
	Response [1]Element8
}

// Version (TW_VERSION) represents a general way to describe the version of
// software that is running.
type Version struct {
	MajorNum uint16
	MinorNum uint16
	Language uint16
	Country  uint16
	Info     Str32
}

// ImageInfo (TW_IMAGEINFO) describes the "real" image data, that is, the
// complete image being transferred between the Source and application.
type ImageInfo struct {
	XResolution     Fix32
	YResolution     Fix32
	ImageWidth      int32
	ImageLength     int32
	SamplesPerPixel int16
	BitsPerSample   [8]int16
	BitsPerPixel    int16
	Planar          Bool
	PixelType       int16
	Compression     uint16
}

// ImageLayout (TW_IMAGELAYOUT) contains information about the original size
// of the acquired image.
type ImageLayout struct {
	Frame          Frame
	DocumentNumber uint32
	PageNumber     uint32
	FrameNumber    uint32
}

// Memory (TW_MEMORY) provides information for managing memory buffers.
type Memory struct {
	Flags  MemoryFlag
	Length uint32
	TheMem MemRef
}

// ImageMemTransfer (TW_IMAGEMEMXFER) describes the form of the acquired data
// being passed from the Source to the application.
type ImageMemTransfer struct {
	Compression  uint16
	BytesPerRow  uint32
	Columns      uint32
	Rows         uint32
	XOffset      uint32
	YOffset      uint32
	BytesWritten uint32
	Memory       Memory
}

// JPEGCompression (TW_JPEGCOMPRESSION) describes the information necessary
// to transfer a JPEG-compressed image.
type JPEGCompression struct {
	ColorSpace       uint16
	SubSampling      uint32
	NumComponents    uint16
	RestartFrequency uint16
	QuantMap         [4]uint16
	QuantTable       [4]Memory
	HuffmanMap       [4]uint16
	HuffmanDC        [2]Memory
	HuffmanAC        [2]Memory
}

// Metrics (TW_METRICS) collects scanning metrics after returning to state 4.
type Metrics struct {
	SizeOf     uint32
	ImageCount uint32
	SheetCount uint32
}

// OneValue (TW_ONEVALUE) stores a single value (item) which describes a
// capability.
type OneValue struct {
	ItemType uint16
	Item     uint32
}

// Palette8 (TW_PALETTE8) holds the color palette information.
type Palette8 struct {
	NumColors   uint16
	PaletteType uint16
	Colors      [256]Element8
}

// PassThru (TW_PASSTHRU) is used to bypass the TWAIN protocol when
// communicating with a device.
type PassThru struct {
	Command              MemRef
	CommandBytes         uint32
	Direction            int32
	DataHandle           MemRef
	DataBytes            uint32
	DataBytesTransferred uint32
}

// PendingTransfers (TW_PENDINGXFERS) tells the application how many more
// complete transfers the Source currently has available.
// TODO: Find out if union breaks this struct.
type PendingTransfers struct {
	Count         uint16
	EOJOrReserved [4]byte // uint32
}

// Range (TW_RANGE) stores a range of individual values describing a
// capability.
type Range struct {
	ItemType     uint16
	MinValue     uint32
	MaxValue     uint32
	StepSize     uint32
	DefaultValue uint32
	CurrentValue uint32
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
	Format   uint16
	VRefNum  int16
}

// SetupMemTransfer (TW_SETUPMEMXFER) provides the application information
// about the Source's requirements and preferences regarding allocation of
// transfer buffer(s).
type SetupMemTransfer struct {
	MinBufSize uint32
	MaxBufSize uint32
	Preferred  uint32
}

// Status (TW_STATUS) describes the status of a source.
type Status struct {
	ConditionCode uint16
	Data          uint16
}

// StatusUTF8 (TW_STATUSUTF8) translates the contents of Status into a
// localized UTF-8 string.
type StatusUTF8 struct {
	Status     Status
	Size       uint32
	UTF8String Handle
}

// TwainDirect is an alias for TW_TWAINDIRECT.
type TwainDirect struct {
	SizeOf               uint32
	CommunicationManager uint16
	Send                 Handle
	SendSize             uint32
	Receive              Handle
	ReceiveSize          uint32
}

// UserInterface (TW_USERINTERFACE) is used to handle the user interface
// coordination between an application and a Source.
type UserInterface struct {
	ShowUI       Bool
	ModalUI      Bool
	ParentHandle Handle
}

// FilterDescriptor (TW_FILTER_DESCRIPTOR) is used with data.Filter (DAT_FILTER).
type FilterDescriptor struct {
	Size            uint32
	HueStart        uint32
	HueEnd          uint32
	SaturationStart uint32
	SaturationEnd   uint32
	ValueStart      uint32
	ValueEnd        uint32
	Replacement     uint32
}

// Filter (TW_FILTER) is used with data.Filter (DAT_FILTER).
type Filter struct {
	Size               uint32
	DescriptorCount    uint32
	MaxDescriptorCount uint32
	Condition          uint32
	DescriptorsHandle  Handle
}
