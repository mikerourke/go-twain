package twain

// CallBack (TW_CALLBACK) is used to register callbacks.
type CallBack struct {
	CallBackProc MemRef
	RefCon       MemRef
	Message      Int16
}

// Identity (TW_IDENTITY) provides identification information about a TWAIN
// entity.
type Identity struct {
	ID              MemRef
	Version         Version
	ProtocolMajor   UInt16
	ProtocolMinor   UInt16
	SupportedGroups UInt32
	Manufacturer    Str32
	ProductFamily   Str32
	ProductName     Str32
}
