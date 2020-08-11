package twain

// Message is an alias for MSG_ values.
type Message uint32

// Generic messages may be used with any of several DATs.
const (
	// MessageNull is an alias for MSG_NULL.
	MessageNull Message = 0x0000

	// MessageCustomBase is an alias for MSG_CUSTOMBASE.
	MessageCustomBase Message = 0x8000

	// MessageGet is an alias for MSG_GET.
	MessageGet Message = 0x0001

	// MessageGetCurrent is an alias for MSG_GETCURRENT.
	MessageGetCurrent Message = 0x0002

	// MessageGetDefault is an alias for MSG_GETDEFAULT.
	MessageGetDefault Message = 0x0003

	// MessageGetFirst is an alias for MSG_GETFIRST.
	MessageGetFirst Message = 0x0004

	// MessageGetNext is an alias for MSG_GETNEXT.
	MessageGetNext Message = 0x0005

	// MessageSet is an alias for MSG_SET.
	MessageSet Message = 0x0006

	// MessageReset is an alias for MSG_RESET.
	MessageReset Message = 0x0007

	// MessageQuerySupport is an alias for MSG_QUERYSUPPORT.
	MessageQuerySupport Message = 0x0008

	// MessageGetHelp is an alias for MSG_GETHELP.
	MessageGetHelp Message = 0x0009

	// MessageGetLabel is an alias for MSG_GETLABEL.
	MessageGetLabel Message = 0x000a

	// MessageGetLabelEnum is an alias for MSG_GETLABELENUM.
	MessageGetLabelEnum Message = 0x000b

	// MessageSetConstraint is an alias for MSG_SETCONSTRAINT.
	MessageSetConstraint Message = 0x000c
)

// Messages used with DAT_NULL.
const (
	// MessageTransferReady is an alias for MSG_XFERREADY.
	MessageTransferReady Message = 0x0101

	// MessageCloseDSReq is an alias for MSG_CLOSEDSREQ.
	MessageCloseDSReq Message = 0x0102

	// MessageCloseDSOK is an alias for MSG_CLOSEDSOK.
	MessageCloseDSOK Message = 0x0103

	// MessageDeviceEvent is an alias for MSG_DEVICEEVENT.
	MessageDeviceEvent Message = 0x0104
)

// Messages used with a pointer to DataParent (DAT_PARENT) data.
const (
	// MessageOpenDSM is an alias for MSG_OPENDSM.
	MessageOpenDSM Message = 0x0301

	// MessageCloseDSM is an alias for MSG_CLOSEDSM.
	MessageCloseDSM Message = 0x0302
)

// Messages used with a pointer to a DataIdentity (DAT_IDENTITY) structure.
const (
	// OpenDS is an alias for MSG_OPENDS.
	MessageOpenDS Message = 0x0401

	// CloseDS is an alias for MSG_CLOSEDS.
	MessageCloseDS Message = 0x0402

	// UserSelect is an alias for MSG_USERSELECT.
	MessageUserSelect Message = 0x0403
)

// Messages used with a pointer to a DataUserInterface (DAT_USERINTERFACE)
// structure.
const (
	// MessageDisableDS is an alias for MSG_DISABLEDS.
	MessageDisableDS Message = 0x0501

	// MessageEnableDS is an alias for MSG_ENABLEDS.
	MessageEnableDS Message = 0x0502

	// MessageEnableDSUIOnly is an alias for MSG_ENABLEDSUIONLY.
	MessageEnableDSUIOnly Message = 0x0503
)

// MessageProcessEvent is an alias for MSG_PROCESSEVENT. It is used with a
// pointer to a DataEvent (DAT_EVENT) structure
const MessageProcessEvent Message = 0x0601

// Messages used with a pointer to a DataPendingTransfers (DAT_PENDINGXFERS)
// structure.
const (
	// MessageEndTransfer is an alias for MSG_ENDXFER.
	MessageEndTransfer Message = 0x0701

	// MessageStopFeeder is an alias for MSG_STOPFEEDER.
	MessageStopFeeder Message = 0x0702
)

// Messages used with a pointer to a DataFileSystem (DAT_FILESYSTEM) structure.
const (
	// MessageChangeDirectory is an alias for MSG_CHANGEDIRECTORY.
	MessageChangeDirectory Message = 0x0801

	// MessageCreateDirectory is an alias for MSG_CREATEDIRECTORY.
	MessageCreateDirectory Message = 0x0802

	// MessageDelete is an alias for MSG_DELETE.
	MessageDelete Message = 0x0803

	// MessageFormatMedia is an alias for MSG_FORMATMEDIA.
	MessageFormatMedia Message = 0x0804

	// MessageGetClose is an alias for MSG_GETCLOSE.
	MessageGetClose Message = 0x0805

	// MessageGetFirstFile is an alias for MSG_GETFIRSTFILE.
	MessageGetFirstFile Message = 0x0806

	// MessageGetInfo is an alias for MSG_GETINFO.
	MessageGetInfo Message = 0x0807

	// MessageGetNextFile is an alias for MSG_GETNEXTFILE.
	MessageGetNextFile Message = 0x0808

	// MessageRename is an alias for MSG_RENAME.
	MessageRename Message = 0x0809

	// MessageCopy is an alias for MSG_COPY.
	MessageCopy Message = 0x080A

	// MessageAutomaticCaptureDirectory is an alias for MSG_AUTOMATICCAPTUREDIRECTORY.
	MessageAutomaticCaptureDirectory Message = 0x080B
)

// MessagePassThru is an alias for MSG_PASSTHRU. It is used with a pointer to a
// DataPassThru (DAT_PASSTHRU) structure.
const MessagePassThru Message = 0x0901

// MessageRegisterCallback is an alias for MSG_REGISTER_CALLBACK. It is used with
// DataCallback (DAT_CALLBACK).
const MessageRegisterCallback Message = 0x0902

// MessageResetAll is an alias for MSG_RESETALL.  It is used with DataCapability
// (DAT_CAPABILITY).
const MessageResetAll Message = 0x0A01

// MessageSetTask is an alias for MSG_SETTASK. It is used with used with
// DataTwainDirect (DAT_TWAINDIRECT).
const MessageSetTask Message = 0x0B01
