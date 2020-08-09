package twain

// CallBack (TW_CALLBACK) is used to register callbacks.
type CallBack struct {
	CallBackProc MemRef
	RefCon       UInt32
	Message      Int16
}

// Identity (TW_IDENTITY) provides identification information about a TWAIN
// entity.
type Identity struct {
	ID              UInt32
	Version         Version
	ProtocolMajor   UInt16
	ProtocolMinor   UInt16
	SupportedGroups UInt32
	Manufacturer    Str32
	ProductFamily   Str32
	ProductName     Str32
}
