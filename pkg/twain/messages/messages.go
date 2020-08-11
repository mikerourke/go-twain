// Package messages wraps the messages constants (prefixed with MSG_) from the
// TWAIN header file.
package messages

import "github.com/mikerourke/go-twain/pkg/twain"

type Message twain.UInt32

// Generic messages may be used with any of several DATs.
const (
	// Null is an alias for MSG_NULL.
	Null Message = 0x0000

	// CustomBase is an alias for MSG_CUSTOMBASE.
	CustomBase Message = 0x8000

	// Get is an alias for MSG_GET.
	Get Message = 0x0001

	// GetCurrent is an alias for MSG_GETCURRENT.
	GetCurrent Message = 0x0002

	// GetDefault is an alias for MSG_GETDEFAULT.
	GetDefault Message = 0x0003

	// GetFirst is an alias for MSG_GETFIRST.
	GetFirst Message = 0x0004

	// GetNext is an alias for MSG_GETNEXT.
	GetNext Message = 0x0005

	// Set is an alias for MSG_SET.
	Set Message = 0x0006

	// Reset is an alias for MSG_RESET.
	Reset Message = 0x0007

	// QuerySupport is an alias for MSG_QUERYSUPPORT.
	QuerySupport Message = 0x0008

	// GetHelp is an alias for MSG_GETHELP.
	GetHelp Message = 0x0009

	// GetLabel is an alias for MSG_GETLABEL.
	GetLabel Message = 0x000a

	// GetLabelEnum is an alias for MSG_GETLABELENUM.
	GetLabelEnum Message = 0x000b

	// SetConstraint is an alias for MSG_SETCONSTRAINT.
	SetConstraint Message = 0x000c
)

// Messages used with DAT_NULL.
const (
	// TransferReady is an alias for MSG_XFERREADY.
	TransferReady Message = 0x0101

	// CloseDSReq is an alias for MSG_CLOSEDSREQ.
	CloseDSReq Message = 0x0102

	// CloseDSOK is an alias for MSG_CLOSEDSOK.
	CloseDSOK Message = 0x0103

	// DeviceEvent is an alias for MSG_DEVICEEVENT.
	DeviceEvent Message = 0x0104
)

// Messages used with a pointer to DataParent (DAT_PARENT) data.
const (
	// OpenDSM is an alias for MSG_OPENDSM.
	OpenDSM Message = 0x0301

	// CloseDSM is an alias for MSG_CLOSEDSM.
	CloseDSM Message = 0x0302
)

// Messages used with a pointer to a DataIdentity (DAT_IDENTITY) structure.
const (
	// OpenDS is an alias for MSG_OPENDS.
	OpenDS Message = 0x0401

	// CloseDS is an alias for MSG_CLOSEDS.
	CloseDS Message = 0x0402

	// UserSelect is an alias for MSG_USERSELECT.
	UserSelect Message = 0x0403
)

// Messages used with a pointer to a DataUserInterface (DAT_USERINTERFACE)
// structure.
const (
	// DisableDS is an alias for MSG_DISABLEDS.
	DisableDS Message = 0x0501

	// EnableDS is an alias for MSG_ENABLEDS.
	EnableDS Message = 0x0502

	// EnableDSUIOnly is an alias for MSG_ENABLEDSUIONLY.
	EnableDSUIOnly Message = 0x0503
)

// ProcessEvent is an alias for MSG_PROCESSEVENT. It is used with a pointer to
// a DataEvent (DAT_EVENT) structure
const ProcessEvent Message = 0x0601

// Messages used with a pointer to a DataPendingTransfers (DAT_PENDINGXFERS)
// structure.
const (
	// EndTransfer is an alias for MSG_ENDXFER.
	EndTransfer Message = 0x0701

	// StopFeeder is an alias for MSG_STOPFEEDER.
	StopFeeder Message = 0x0702
)

// Messages used with a pointer to a DataFileSystem (DAT_FILESYSTEM) structure.
const (
	// ChangeDirectory is an alias for MSG_CHANGEDIRECTORY.
	ChangeDirectory Message = 0x0801

	// CreateDirectory is an alias for MSG_CREATEDIRECTORY.
	CreateDirectory Message = 0x0802

	// Delete is an alias for MSG_DELETE.
	Delete Message = 0x0803

	// FormatMedia is an alias for MSG_FORMATMEDIA.
	FormatMedia Message = 0x0804

	// GetClose is an alias for MSG_GETCLOSE.
	GetClose Message = 0x0805

	// GetFirstFile is an alias for MSG_GETFIRSTFILE.
	GetFirstFile Message = 0x0806

	// GetInfo is an alias for MSG_GETINFO.
	GetInfo Message = 0x0807

	// GetNextFile is an alias for MSG_GETNEXTFILE.
	GetNextFile Message = 0x0808

	// Rename is an alias for MSG_RENAME.
	Rename Message = 0x0809

	// Copy is an alias for MSG_COPY.
	Copy Message = 0x080A

	// AutomaticCaptureDirectory is an alias for MSG_AUTOMATICCAPTUREDIRECTORY.
	AutomaticCaptureDirectory Message = 0x080B
)

// PassThru is an alias for MSG_PASSTHRU. It is used with a pointer to a
// DataPassThru (DAT_PASSTHRU) structure.
const PassThru Message = 0x0901

// RegisterCallback is an alias for MSG_REGISTER_CALLBACK. It is used with
// DataCallback (DAT_CALLBACK).
const RegisterCallback Message = 0x0902

// ResetAll is an alias for MSG_RESETALL.  It is used with DataCapability
// (DAT_CAPABILITY).
const ResetAll Message = 0x0A01

// SetTask is an alias for MSG_SETTASK. It is used with used with
// DataTwainDirect (DAT_TWAINDIRECT).
const SetTask Message = 0x0B01
