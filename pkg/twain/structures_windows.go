package twain

// TWCallBack (TW_CALLBACK) is used to register callbacks.
type TWCallBack struct {
	CallBackProc TWMemRef
	RefCon       TWUInt32
	Message      TWInt16
}

// TWIdentity (TW_IDENTITY) provides identification information about a TWAIN
// entity.
type TWIdentity struct {
	ID              TWUInt32
	Version         TWVersion
	ProtocolMajor   TWUInt16
	ProtocolMinor   TWUInt16
	SupportedGroups TWUInt32
	Manufacturer    TWStr32
	ProductFamily   TWStr32
	ProductName     TWStr32
}
