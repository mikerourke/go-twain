package twain

// String types that map to TW_STR** types from the TWAIN specification.
// These include room for the strings and a NULL char.

// Str32 (TW_STR32) represents a string that can hold up to 32 characters.
type Str32 [34]int8

// Str64 (TW_STR64) represents a string that can hold up to 64 characters.
type Str64 [66]int8

// Str128 (TW_STR128) represents a string that can hold up to 128 characters.
type Str128 [130]int8

// Str255 (TW_STR255) represents a string that can hold up to 255 characters.
type Str255 [256]int8
