package twain

// String types that map to TW_STR** types from the TWAIN specification.
// These include room for the strings and a NULL char.

// TWStr32 (TW_STR32) represents a string that can hold up to 32 characters.
type TWStr32 [34]int8

// TWStr64 (TW_STR64) represents a string that can hold up to 64 characters.
type TWStr64 [66]int8

// TWStr128 (TW_STR128) represents a string that can hold up to 128 characters.
type TWStr128 [130]int8

// TWStr255 (TW_STR255) represents a string that can hold up to 255 characters.
type TWStr255 [256]int8
