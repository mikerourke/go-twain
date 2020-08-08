package twain

import "unsafe"

// TWInt8 (TW_INT8) represents a signed 8-bit integer.
type TWInt8 int8

// TWInt16 (TW_INT16) represents a signed 16-bit integer.
type TWInt16 int16

// TWInt32 (TW_INT32) represents a signed 32-bit integer.
type TWInt32 int32

// TWUInt8 (TW_UINT8) represents an unsigned 8-bit integer.
type TWUInt8 uint8

// TWUInt16 (TW_UINT16) represents an unsigned 16-bit integer.
type TWUInt16 uint16

// TWUInt32 (TW_UINT32) represents an unsigned 32-bit integer.
type TWUInt32 uint32

// TWBool (TW_BOOL) represents a boolean value.
type TWBool uint16

// TODO: Check if these types are valid.
// TWHandle (TW_HANDLE) represents a pointer to a handle.
type TWHandle unsafe.Pointer

// TWMemRef (TW_MEMREF) represents a reference in memory.
type TWMemRef *uint8

// TWUIntPtr (TW_UINTPTR) represents an unsigned integer pointer.
type TWUIntPtr uintptr
