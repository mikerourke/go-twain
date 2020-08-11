package twain

// String types that map to TW_STR** types from the TWAIN specification.
// These include room for the strings and a NULL char.

// Str32 is an alias for TW_STR32. It represents a string that can hold up to
// 32 characters.
type Str32 [34]int8

// Str64 is an alias for TW_STR64. It represents a string that can hold up to
// 64 characters.
type Str64 [66]int8

// Str128 is an alias for TW_STR128. It represents a string that can hold up to
// 128 characters.
type Str128 [130]int8

// Str255 is an alias for TW_STR255. It represents a string that can hold up to
// 255 characters.
type Str255 [256]int8

// Callback (TW_CALLBACK) is used to register callbacks.
type Callback struct {
	CallbackProc MemRef
	RefCon       uint32
	Message      int16
}

// Identity (TW_IDENTITY) provides identification information about a TWAIN
// entity.
type Identity struct {
	ID              uint32
	Version         Version
	ProtocolMajor   uint16
	ProtocolMinor   uint16
	SupportedGroups uint32
	Manufacturer    Str32
	ProductFamily   Str32
	ProductName     Str32
}
