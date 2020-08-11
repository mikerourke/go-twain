package twain

// ReturnCode is an alias for TWRC_ values.
type ReturnCode uint16

// ReturnCodeCustomBase is an alias for TWRC_CUSTOMBASE.
const ReturnCodeCustomBase ReturnCode = 0x8000

const (
	// ReturnCodeSuccess is an alias for TWRC_SUCCESS.
	ReturnCodeSuccess ReturnCode = iota

	// ReturnCodeFailure is an alias for TWRC_FAILURE.
	ReturnCodeFailure

	// ReturnCodeCheckStatus is an alias for TWRC_CHECKSTATUS.
	ReturnCodeCheckStatus

	// ReturnCodeCancel is an alias for TWRC_CANCEL.
	ReturnCodeCancel

	// ReturnCodeDSEvent is an alias for TWRC_DSEVENT.
	ReturnCodeDSEvent

	// ReturnCodeNotDSEvent is an alias for TWRC_NOTDSEVENT.
	ReturnCodeNotDSEvent

	// ReturnCodeTransferDone is an alias for TWRC_XFERDONE.
	ReturnCodeTransferDone

	// ReturnCodeEndOfList is an alias for TWRC_ENDOFLIST.
	ReturnCodeEndOfList

	// ReturnCodeInfoNotSupported is an alias for TWRC_INFONOTSUPPORTED.
	ReturnCodeInfoNotSupported

	// ReturnCodeDataNotAvailable is an alias for TWRC_DATANOTAVAILABLE.
	ReturnCodeDataNotAvailable

	// ReturnCodeBusy is an alias for TWRC_BUSY.
	ReturnCodeBusy

	// ReturnCodeScannerLocked is an alias for TWRC_SCANNERLOCKED.
	ReturnCodeScannerLocked
)

// ConditionCode is an alias for TWCC_ values. The application gets these by
// doing data.ControlArgumentTypeStatus (DG_CONTROL/DAT_STATUS) and
// messages.Get (MSG_GET).
type ConditionCode uint16

// ConditionCodeCustomBase is an alias for TWCC_CUSTOMBASE.
const ConditionCodeCustomBase = 0x8000

const (
	// ConditionCodeSuccess is an alias for TWCC_SUCCESS.
	ConditionCodeSuccess ConditionCode = 0

	// ConditionCodeBummer is an alias for TWCC_BUMMER.
	ConditionCodeBummer ConditionCode = 1

	// ConditionCodeLowMemory is an alias for TWCC_LOWMEMORY.
	ConditionCodeLowMemory ConditionCode = 2

	// ConditionCodeNoDS is an alias for TWCC_NODS.
	ConditionCodeNoDS ConditionCode = 3

	// ConditionCodeMaxConnections is an alias for TWCC_MAXCONNECTIONS.
	ConditionCodeMaxConnections ConditionCode = 4

	// ConditionCodeOperationError is an alias for TWCC_OPERATIONERROR.
	ConditionCodeOperationError ConditionCode = 5

	// ConditionCodeBadCap is an alias for TWCC_BADCAP.
	ConditionCodeBadCap ConditionCode = 6

	// ConditionCodeBadProtocol is an alias for TWCC_BADPROTOCOL.
	ConditionCodeBadProtocol ConditionCode = 9

	// ConditionCodeBadValue is an alias for TWCC_BADVALUE.
	ConditionCodeBadValue ConditionCode = 10

	// ConditionCodeSeqError is an alias for TWCC_SEQERROR.
	ConditionCodeSeqError ConditionCode = 11

	// ConditionCodeBaddest is an alias for TWCC_BADDEST.
	ConditionCodeBaddest ConditionCode = 12

	// ConditionCodeCapUnsupported is an alias for TWCC_CAPUNSUPPORTED.
	ConditionCodeCapUnsupported ConditionCode = 13

	// ConditionCodeCapBadOperation is an alias for TWCC_CAPBADOPERATION.
	ConditionCodeCapBadOperation ConditionCode = 14

	// ConditionCodeCapSeqError is an alias for TWCC_CAPSEQERROR.
	ConditionCodeCapSeqError ConditionCode = 15

	// ConditionCodeDenied is an alias for TWCC_DENIED.
	ConditionCodeDenied ConditionCode = 16

	// ConditionCodeFileExists is an alias for TWCC_FILEEXISTS.
	ConditionCodeFileExists ConditionCode = 17

	// ConditionCodeFileNotFound is an alias for TWCC_FILENOTFOUND.
	ConditionCodeFileNotFound ConditionCode = 18

	// ConditionCodeNotEmpty is an alias for TWCC_NOTEMPTY.
	ConditionCodeNotEmpty ConditionCode = 19

	// ConditionCodePaperJam is an alias for TWCC_PAPERJAM.
	ConditionCodePaperJam ConditionCode = 20

	// ConditionCodePaperDoubleFeed is an alias for TWCC_PAPERDOUBLEFEED.
	ConditionCodePaperDoubleFeed ConditionCode = 21

	// ConditionCodeFileWriteError is an alias for TWCC_FILEWRITEERROR.
	ConditionCodeFileWriteError ConditionCode = 22

	// ConditionCodeCheckDeviceOnline is an alias for TWCC_CHECKDEVICEONLINE.
	ConditionCodeCheckDeviceOnline ConditionCode = 23

	// ConditionCodeInterlock is an alias for TWCC_INTERLOCK.
	ConditionCodeInterlock ConditionCode = 24

	// ConditionCodeDamagedCorner is an alias for TWCC_DAMAGEDCORNER.
	ConditionCodeDamagedCorner ConditionCode = 25

	// ConditionCodeFocusError is an alias for TWCC_FOCUSERROR.
	ConditionCodeFocusError ConditionCode = 26

	// ConditionCodeDocTooLight is an alias for TWCC_DOCTOOLIGHT.
	ConditionCodeDocTooLight ConditionCode = 27

	// ConditionCodeDocTooDark is an alias for TWCC_DOCTOODARK.
	ConditionCodeDocTooDark ConditionCode = 28

	// ConditionCodeNoMedia is an alias for TWCC_NOMEDIA.
	ConditionCodeNoMedia ConditionCode = 29
)

// QueryCode is an alias for TWQC_ values. These are bit patterns to query the
// operations that are supported by the data source on a capability. The
// application gets these through data.ControlArgumentTypeCapability
// (DG_CONTROL/DAT_CAPABILITY) and messages.QuerySupport (MSG_QUERYSUPPORT).
type QueryCode uint16

const (
	// QueryCodeGet is an alias for TWQC_GET
	QueryCodeGet QueryCode = 0x0001

	// QueryCodeSet is an alias for TWQC_SET
	QueryCodeSet QueryCode = 0x0002

	// QueryCodeGetDefault is an alias for TWQC_GETDEFAULT
	QueryCodeGetDefault QueryCode = 0x0004

	// QueryCodeGetCurrent is an alias for TWQC_GETCURRENT
	QueryCodeGetCurrent QueryCode = 0x0008

	// QueryCodeReset is an alias for TWQC_RESET
	QueryCodeReset QueryCode = 0x0010

	// QueryCodeSetConstraint is an alias for TWQC_SETCONSTRAINT
	QueryCodeSetConstraint QueryCode = 0x0020

	// QueryCodeGetHelp is an alias for TWQC_GETHELP
	QueryCodeGetHelp QueryCode = 0x0100

	// QueryCodeGetLabel is an alias for TWQC_GETLABEL
	QueryCodeGetLabel QueryCode = 0x0200

	// QueryCodeGetLabelEnum is an alias for TWQC_GETLABELENUM
	QueryCodeGetLabelEnum QueryCode = 0x0400
)
