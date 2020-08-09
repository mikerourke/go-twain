package twain

import "unsafe"

// Int8 (TW_INT8) represents a signed 8-bit integer.
type Int8 int8

// Int16 (TW_INT16) represents a signed 16-bit integer.
type Int16 int16

// Int32 (TW_INT32) represents a signed 32-bit integer.
type Int32 int32

// UInt8 (TW_UINT8) represents an unsigned 8-bit integer.
type UInt8 uint8

// UInt16 (TW_UINT16) represents an unsigned 16-bit integer.
type UInt16 uint16

// UInt32 (TW_UINT32) represents an unsigned 32-bit integer.
type UInt32 uint32

// Bool (TW_BOOL) represents a boolean value.
type Bool uint16

// TODO: Check if these types are valid.
// Handle (TW_HANDLE) represents a pointer to a handle.
type Handle unsafe.Pointer

// MemRef (TW_MEMREF) represents a reference in memory.
type MemRef *uint8

// UIntPtr (TW_UINTPTR) represents an unsigned integer pointer.
type UIntPtr uintptr
