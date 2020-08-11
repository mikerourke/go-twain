package twain

const (
	// IconID is an alias for TWON_ICONID.
	IconID = 962

	// DSMID is an alias for TWON_DSMID.
	DSMID = 461

	// DSMCodeID is an alias for TWON_DSMCODEID.
	DSMCodeID = 63

	// DontCare8 is an alias for TWON_DONTCARE8.
	DontCare8 = 0xff

	// DontCare16 is an alias for TWON_DONTCARE16.
	DontCare16 = 0xffff

	// DontCare32 is an alias for TWON_DONTCARE32.
	DontCare32 = 0xffffffff
)

// ContainerType is an alias for TWON values. It represents a generic constant
// used to represent the container type in the Capability struct.
type ContainerType uint16

const (
	// ContainerTypeArray is an alias for TWON_ARRAY.
	ContainerTypeArray ContainerType = 3

	// ContainerTypeEnumeration is an alias for TWON_ENUMERATION.
	ContainerTypeEnumeration ContainerType = 4

	// ContainerTypeOneValue is an alias for TWON_ONEVALUE.
	ContainerTypeOneValue ContainerType = 5

	// ContainerTypeRange is an alias for TWON_RANGE.
	ContainerTypeRange ContainerType = 6
)

// MemoryFlag is an alias for TWMF values. It represents flags used in the
// Memory (TW_MEMORY) structure.
type MemoryFlag uint32

const (
	// MemoryFlagAppOwns is an alias for TWMF_APPOWNS.
	MemoryFlagAppOwns MemoryFlag = 0x0001

	// MemoryFlagDSMOwns is an alias for TWMF_DSMOWNS.
	MemoryFlagDSMOwns MemoryFlag = 0x0002

	// MemoryFlagDSOwns is an alias for TWMF_DSOWNS.
	MemoryFlagDSOwns MemoryFlag = 0x0003

	// MemoryFlagPointer is an alias for TWMF_POINTER.
	MemoryFlagPointer MemoryFlag = 0x0008

	// MemoryFlagHandle is an alias for TWMF_HANDLE.
	MemoryFlagHandle MemoryFlag = 0x0010
)

// TypeFlag is an alias for TWTY values. It represents the type used in
// the Memory (TW_MEMORY) structure.
type TypeFlag uint16

const (
	// TypeFlagint8 is an alias for TWTY_INT8.
	TypeFlagint8 TypeFlag = 0x0000

	// TypeFlagint16 is an alias for TWTY_INT16.
	TypeFlagint16 TypeFlag = 0x0001

	// TypeFlagint32 is an alias for TWTY_INT32.
	TypeFlagint32 TypeFlag = 0x0002

	// TypeFlagUint8 is an alias for TWTY_UINT8.
	TypeFlagUint8 TypeFlag = 0x0003

	// TypeFlaguint16 is an alias for TWTY_UINT16.
	TypeFlaguint16 TypeFlag = 0x0004

	// TypeFlaguint32 is an alias for TWTY_UINT32.
	TypeFlaguint32 TypeFlag = 0x0005

	// TypeFlagBool is an alias for TWTY_BOOL.
	TypeFlagBool TypeFlag = 0x0006

	// TypeFlagFix32 is an alias for TWTY_FIX32.
	TypeFlagFix32 TypeFlag = 0x0007

	// TypeFlagFrame is an alias for TWTY_FRAME.
	TypeFlagFrame TypeFlag = 0x0008

	// TypeFlagStr32 is an alias for TWTY_STR32.
	TypeFlagStr32 TypeFlag = 0x0009

	// TypeFlagStr64 is an alias for TWTY_STR64.
	TypeFlagStr64 TypeFlag = 0x000a

	// TypeFlagStr128 is an alias for TWTY_STR128.
	TypeFlagStr128 TypeFlag = 0x000b

	// TypeFlagStr255 is an alias for TWTY_STR255.
	TypeFlagStr255 TypeFlag = 0x000c

	// TypeFlagHandle is an alias for TWTY_HANDLE.
	TypeFlagHandle TypeFlag = 0x000f
)

// FileSystemFileType is an alias for TW_FILESYSTEM.FileType values.
type FileSystemFileType uint16

const (
	// FileSystemFileTypeCamera is an alias for TWFY_CAMERA.
	FileSystemFileTypeCamera FileSystemFileType = iota

	// FileSystemFileTypeCameraTop is an alias for TWFY_CAMERATOP.
	FileSystemFileTypeCameraTop

	// FileSystemFileTypeCameraBottom is an alias for TWFY_CAMERABOTTOM.
	FileSystemFileTypeCameraBottom

	// FileSystemFileTypeCameraPreview is an alias for TWFY_CAMERAPREVIEW.
	FileSystemFileTypeCameraPreview

	// FileSystemFileTypeDomain is an alias for TWFY_DOMAIN.
	FileSystemFileTypeDomain

	// FileSystemFileTypeHost is an alias for TWFY_HOST.
	FileSystemFileTypeHost

	// FileSystemFileTypeDirectory is an alias for TWFY_DIRECTORY.
	FileSystemFileTypeDirectory

	// FileSystemFileTypeImage is an alias for TWFY_IMAGE.
	FileSystemFileTypeImage

	// FileSystemFileTypeUnknown is an alias for TWFY_UNKNOWN.
	FileSystemFileTypeUnknown
)

// MagType is an alias for TWEI_MAGTYPE values.
type MagType uint16

const (
	// MagTypeMICR is an alias for TWMD_MICR.
	MagTypeMICR MagType = iota

	// MagTypeRaw is an alias for TWMD_RAW.
	MagTypeRaw

	// MagTypeInvalid is an alias for TWMD_INVALID.
	MagTypeInvalid
)

// Palette8Type is an alias for TWPA values. It represents palette types for
// Palette8 struct.
type Palette8Type uint16

const (
	// Palette8TypeRGB is an alias for TWPA_RGB.
	Palette8TypeRGB Palette8Type = iota

	// Palette8TypeGray is an alias for TWPA_GRAY.
	Palette8TypeGray

	// Palette8TypeCMY is an alias for TWPA_CMY.
	Palette8TypeCMY
)

// PatchCode is an alias for TWEI_PATCHCODE values.
type PatchCode uint16

const (
	// PatchCode1 is an alias for TWPCH_PATCH1.
	PatchCode1 PatchCode = iota

	// PatchCode2 is an alias for TWPCH_PATCH2.
	PatchCode2

	// PatchCode3 is an alias for TWPCH_PATCH3.
	PatchCode3

	// PatchCode4 is an alias for TWPCH_PATCH4.
	PatchCode4

	// PatchCode6 is an alias for TWPCH_PATCH6.
	PatchCode6

	// PatchCodeT is an alias for TWPCH_PATCHT.
	PatchCodeT
)
