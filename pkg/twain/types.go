package twain

import "unsafe"

// Int8 is an alias for TW_INT8.
type Int8 int8

// Int16 is an alias for TW_INT16.
type Int16 int16

// Int32 is an alias for TW_INT32.
type Int32 int32

// UInt8 is an alias for TW_UINT8.
type UInt8 uint8

// UInt16 is an alias for TW_UINT16.
type UInt16 uint16

// UInt32 is an alias for TW_UINT32.
type UInt32 uint32

// Bool is an alias for TW_BOOL.
type Bool uint16

// TODO: Check if these types are valid.
// Handle is an alias for TW_HANDLE.
type Handle unsafe.Pointer

// MemRef is an alias for TW_MEMREF.
type MemRef *uint8

// UIntPtr is an alias for TW_UINTPTR.
type UIntPtr uintptr

// Byte is an alias for BYTE.
type Byte uint8
