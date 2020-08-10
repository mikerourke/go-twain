// Package codes wraps the return, condition, and query codes constants (prefixed
// with TWRC_, TWCC_, and TWQC_) from the TWAIN header file.
package codes

import "github.com/mikerourke/go-twain/pkg/twain"

// ReturnCode is an alias for TWRC values.
type ReturnCode twain.UInt16

// ReturnCustomBase is an alias for TWRC_CUSTOMBASE.
const ReturnCustomBase ReturnCode = 0x8000

const (
	// ReturnSuccess is an alias for TWRC_SUCCESS.
	ReturnSuccess ReturnCode = iota

	// ReturnFailure is an alias for TWRC_FAILURE.
	ReturnFailure

	// ReturnCheckStatus is an alias for TWRC_CHECKSTATUS.
	ReturnCheckStatus

	// ReturnCancel is an alias for TWRC_CANCEL.
	ReturnCancel

	// ReturnDSEvent is an alias for TWRC_DSEVENT.
	ReturnDSEvent

	// ReturnNotDSEvent is an alias for TWRC_NOTDSEVENT.
	ReturnNotDSEvent

	// ReturnTransferDone is an alias for TWRC_XFERDONE.
	ReturnTransferDone

	// ReturnEndOfList is an alias for TWRC_ENDOFLIST.
	ReturnEndOfList

	// ReturnInfoNotSupported is an alias for TWRC_INFONOTSUPPORTED.
	ReturnInfoNotSupported

	// ReturnDataNotAvailable is an alias for TWRC_DATANOTAVAILABLE.
	ReturnDataNotAvailable

	// ReturnBusy is an alias for TWRC_BUSY.
	ReturnBusy

	// ReturnScannerLocked is an alias for TWRC_SCANNERLOCKED.
	ReturnScannerLocked
)

// ConditionCode is an alias for TWCC values. The application gets these by
// doing data.ControlArgumentTypeStatus (DG_CONTROL/DAT_STATUS) and
// messages.Get (MSG_GET).
type ConditionCode twain.UInt16

// ConditionCustomBase is an alias for TWCC_CUSTOMBASE.
const ConditionCustomBase = 0x8000

const (
	// ConditionSuccess is an alias for TWCC_SUCCESS.
	ConditionSuccess ConditionCode = 0

	// ConditionBummer is an alias for TWCC_BUMMER.
	ConditionBummer ConditionCode = 1

	// ConditionLowMemory is an alias for TWCC_LOWMEMORY.
	ConditionLowMemory ConditionCode = 2

	// ConditionNoDS is an alias for TWCC_NODS.
	ConditionNoDS ConditionCode = 3

	// ConditionMaxConnections is an alias for TWCC_MAXCONNECTIONS.
	ConditionMaxConnections ConditionCode = 4

	// ConditionOperationError is an alias for TWCC_OPERATIONERROR.
	ConditionOperationError ConditionCode = 5

	// ConditionBadCap is an alias for TWCC_BADCAP.
	ConditionBadCap ConditionCode = 6

	// ConditionBadProtocol is an alias for TWCC_BADPROTOCOL.
	ConditionBadProtocol ConditionCode = 9

	// ConditionBadValue is an alias for TWCC_BADVALUE.
	ConditionBadValue ConditionCode = 10

	// ConditionSeqError is an alias for TWCC_SEQERROR.
	ConditionSeqError ConditionCode = 11

	// ConditionBaddest is an alias for TWCC_BADDEST.
	ConditionBaddest ConditionCode = 12

	// ConditionCapUnsupported is an alias for TWCC_CAPUNSUPPORTED.
	ConditionCapUnsupported ConditionCode = 13

	// ConditionCapBadOperation is an alias for TWCC_CAPBADOPERATION.
	ConditionCapBadOperation ConditionCode = 14

	// ConditionCapSeqError is an alias for TWCC_CAPSEQERROR.
	ConditionCapSeqError ConditionCode = 15

	// ConditionDenied is an alias for TWCC_DENIED.
	ConditionDenied ConditionCode = 16

	// ConditionFileExists is an alias for TWCC_FILEEXISTS.
	ConditionFileExists ConditionCode = 17

	// ConditionFileNotFound is an alias for TWCC_FILENOTFOUND.
	ConditionFileNotFound ConditionCode = 18

	// ConditionNotEmpty is an alias for TWCC_NOTEMPTY.
	ConditionNotEmpty ConditionCode = 19

	// ConditionPaperJam is an alias for TWCC_PAPERJAM.
	ConditionPaperJam ConditionCode = 20

	// ConditionPaperDoubleFeed is an alias for TWCC_PAPERDOUBLEFEED.
	ConditionPaperDoubleFeed ConditionCode = 21

	// ConditionFileWriteError is an alias for TWCC_FILEWRITEERROR.
	ConditionFileWriteError ConditionCode = 22

	// ConditionCheckDeviceOnline is an alias for TWCC_CHECKDEVICEONLINE.
	ConditionCheckDeviceOnline ConditionCode = 23

	// ConditionInterlock is an alias for TWCC_INTERLOCK.
	ConditionInterlock ConditionCode = 24

	// ConditionDamagedCorner is an alias for TWCC_DAMAGEDCORNER.
	ConditionDamagedCorner ConditionCode = 25

	// ConditionFocusError is an alias for TWCC_FOCUSERROR.
	ConditionFocusError ConditionCode = 26

	// ConditionDocTooLight is an alias for TWCC_DOCTOOLIGHT.
	ConditionDocTooLight ConditionCode = 27

	// ConditionDocTooDark is an alias for TWCC_DOCTOODARK.
	ConditionDocTooDark ConditionCode = 28

	// ConditionNoMedia is an alias for TWCC_NOMEDIA.
	ConditionNoMedia ConditionCode = 29
)

// QueryCode is an alias for TWQC values. These are bit patterns to query the
// operations that are supported by the data source on a capability. The
// application gets these through data.ControlArgumentTypeCapability
// (DG_CONTROL/DAT_CAPABILITY) and messages.QuerySupport (MSG_QUERYSUPPORT).
type QueryCode twain.UInt16

const (
	// QueryGet is an alias for TWQC_GET
	QueryGet QueryCode = 0x0001

	// QuerySet is an alias for TWQC_SET
	QuerySet QueryCode = 0x0002

	// QueryGetDefault is an alias for TWQC_GETDEFAULT
	QueryGetDefault QueryCode = 0x0004

	// QueryGetCurrent is an alias for TWQC_GETCURRENT
	QueryGetCurrent QueryCode = 0x0008

	// QueryReset is an alias for TWQC_RESET
	QueryReset QueryCode = 0x0010

	// QuerySetConstraint is an alias for TWQC_SETCONSTRAINT
	QuerySetConstraint QueryCode = 0x0020

	// QueryGetHelp is an alias for TWQC_GETHELP
	QueryGetHelp QueryCode = 0x0100

	// QueryGetLabel is an alias for TWQC_GETLABEL
	QueryGetLabel QueryCode = 0x0200

	// QueryGetLabelEnum is an alias for TWQC_GETLABELENUM
	QueryGetLabelEnum QueryCode = 0x0400
)
