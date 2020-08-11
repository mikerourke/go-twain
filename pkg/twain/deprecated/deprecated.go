// Package deprecated contains the deprecated types/values from the TWAIN header
// file.
package deprecated

import (
	"github.com/mikerourke/go-twain/pkg/twain"
	"github.com/mikerourke/go-twain/pkg/twain/capabilities"
	"github.com/mikerourke/go-twain/pkg/twain/codes"
	"github.com/mikerourke/go-twain/pkg/twain/data"
	"github.com/mikerourke/go-twain/pkg/twain/messages"
)

// TODO: Validate the Huge types to ensure they work correctly.

// HugeByte is an alias for HPBYTE.
type HugeByte = twain.Byte

// HugeVoid is an alias for HPVOID.
type HugeVoid func()

// TODO: Validate the Uni512 type to ensure the syntax is valid.
// Uni512 is an alias for TW_UNI512.
type Uni512 [512]string

const (
	// TypeFlagStr1024 is an alias for TWTY_STR1024.
	TypeFlagStr1024 twain.TypeFlag = 0x000d

	// TypeFlagUni512 is an alias for TWTY_UNI512.
	TypeFlagUni512 twain.TypeFlag = 0x000e
)

// ImageFileFormatJPN is an alias for TWFF_JPN.
const ImageFileFormatJPN capabilities.ImageFileFormat = 12

const (
	// ControlTwunkIdentity is an alias for DAT_TWUNKIDENTITY.
	ControlTwunkIdentity data.ArgumentType = 0x000b

	// ControlSetupFileTransfer2 is an alias for DAT_SETUPFILEXFER2.
	ControlSetupFileTransfer2 data.ArgumentType = 0x0301
)

const (
	// UseClearBuffers is an alias for CAP_CLEARBUFFERS.
	UseClearBuffers capabilities.Capability = 0x101d

	// UseSupportedCapsExt is an alias for CAP_SUPPORTEDCAPSEXT.
	UseSupportedCapsExt capabilities.Capability = 0x100c

	// UsePageMultipleAcquire is an alias for CAP_PAGEMULTIPLEACQUIRE.
	UsePageMultipleAcquire capabilities.Capability = 0x1023

	// UsePaperBinding is an alias for CAP_PAPERBINDING.
	UsePaperBinding capabilities.Capability = 0x102f

	// UsePassThru is an alias for CAP_PASSTHRU.
	UsePassThru capabilities.Capability = 0x1031

	// UsePowerDownTime is an alias for CAP_POWERDOWNTIME.
	UsePowerDownTime capabilities.Capability = 0x1034

	// UseAudioFileFormat is an alias for ACAP_AUDIOFILEFORMAT.
	UseAudioFileFormat capabilities.Capability = 0x1201
)

const (
	// CheckStatus is an alias for MSG_CHECKSTATUS.
	CheckStatus messages.Message = 0x0201

	// InvokeCallback is an alias for MSG_INVOKE_CALLBACK. Use data.Null
	// (DAT_NULL) and messages.* (MSG_xxx) instead.
	InvokeCallback messages.Message = 0x0903
)

// QueryConstrainable is an alias for TWQC_CONSTRAINABLE.
const QueryConstrainable codes.QueryCode = 0x0040

// TransferMechanismFile2 is an alias for TWSX_FILE2.
const TransferMechanismFile2 capabilities.TransferMechanism = 3

// FileSystem is an alias for TWFS values.
type FileSystem twain.UInt16

const (
	// FileSystemFileSystem is an alias for TWFS_FILESYSTEM.
	FileSystemFileSystem FileSystem = iota

	// FileSystemRecursiveDelete is an alias for TWFS_RECURSIVEDELETE.
	FileSystemRecursiveDelete
)

const (
	// PixelTypeSRGB64 is an alias for TWPT_SRGB64.
	PixelTypeSRGB64 capabilities.PixelType = 11

	// PixelTypeBGR is an alias for TWPT_BGR.
	PixelTypeBGR capabilities.PixelType = 12

	// PixelTypeCIELAB is an alias for TWPT_CIELAB.
	PixelTypeCIELAB capabilities.PixelType = 13

	// PixelTypeCIELUV is an alias for TWPT_CIELUV.
	PixelTypeCIELUV capabilities.PixelType = 14

	// PixelTypeYCBCR is an alias for TWPT_YCBCR.
	PixelTypeYCBCR capabilities.PixelType = 15
)

const (
	// SupportedSizeB is an alias for TWSS_B.
	SupportedSizeB capabilities.SupportedSize = 8

	// SupportedSizeA4Letter is an alias for TWSS_A4LETTER.
	SupportedSizeA4Letter = capabilities.SupportedSizeA4

	// SupportedSizeB3 is an alias for TWSS_B3.
	SupportedSizeB3 = capabilities.SupportedSizeISOB3

	// SupportedSizeB4 is an alias for TWSS_B4.
	SupportedSizeB4 = capabilities.SupportedSizeISOB4

	// SupportedSizeB6 is an alias for TWSS_B6.
	SupportedSizeB6 = capabilities.SupportedSizeISOB6

	// SupportedSizeB5Letter is an alias for TWSS_B5LETTER.
	SupportedSizeB5Letter = capabilities.SupportedSizeJISB5
)

// AudioFileFormat is an alias for ACAP_AUDIOFILEFORMAT values.
type AudioFileFormat twain.UInt16

const (
	// AudioFileFormatWAV is an alias for TWAF_WAV.
	AudioFileFormatWAV AudioFileFormat = iota

	// AudioFileFormatAIFF is an alias for TWAF_AIFF.
	AudioFileFormatAIFF

	// AudioFileFormatAU is an alias for TWAF_AU.
	AudioFileFormatAU

	// AudioFileFormatSND is an alias for TWAF_SND.
	AudioFileFormatSND
)

// ClearBuffers is an alias for CAP_CLEARBUFFERS values.
type ClearBuffers twain.UInt16

const (
	// ClearBuffersAuto is an alias for TWCB_AUTO.
	ClearBuffersAuto ClearBuffers = iota

	// ClearBuffersClear is an alias for TWCB_CLEAR.
	ClearBuffersClear

	// ClearBuffersNoClear is an alias for TWCB_NOCLEAR.
	ClearBuffersNoClear
)

// SetupFileTransfer2 (DAT_SETUPFILEXFER2) sets up DS to application data
// transfer via a file.
type SetupFileTransfer2 struct {
	FileName     twain.MemRef
	FileNameType twain.UInt16
	Format       twain.UInt16
	VRefNum      twain.Int16
	ParID        twain.UInt32
}

// TwunkIdentity (DAT_TWUNKIDENTITY) provides DS identity and 'other'
// information necessary across thunk link.
type TwunkIdentity struct {
	Identity twain.Identity
	DSPath   twain.Str255
}

// TwunkDSEntryParams (TW_TWUNKDSENTRYPARAMS) provides DS_Entry parameters over
// thunk link.
type TwunkDSEntryParams struct {
	DestFlag    twain.Int8
	Dest        twain.Identity
	DataGroup   twain.Int32
	DataArgType twain.Int16
	Message     twain.Int16
	DataSize    twain.Int32
}

// TwunkDSEntryReturn (TW_TWUNKDSENTRYRETURN) provides DS_Entry results over
// thunk link.
type TwunkDSEntryReturn struct {
	ReturnCode    twain.UInt16
	ConditionCode twain.UInt16
	DataSize      twain.Int32
}

// CapExt is an alias for TW_CAPEXT.
type CapExt struct {
	Cap        capabilities.Capability
	Properties twain.UInt16
}

// SetupAudioFileTransfer (DAT_SETUPAUDIOFILEXFER) contains information required
// to setup an audio file transfer.
type SetupAudioFileTransfer struct {
	FileName twain.Str255
	Format   AudioFileFormat
	VRefNum  twain.Int16
}
