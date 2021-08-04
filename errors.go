package goextron

import "errors"

var (
	ErrInvalidInputNumber      = errors.New("invalid input number")
	ErrInvalidSwitchAttempt    = errors.New("invalid switch attempt in this mode")
	ErrInvalidCommand          = errors.New("invalid command")
	ErrInvalidPresetNumber     = errors.New("invalid preset number")
	ErrInvalidPortNumber       = errors.New("invalid port number")
	ErrInvalidParameter        = errors.New("invalid parameter")
	ErrInvalidForConfig        = errors.New("not valid for this configuration")
	ErrInvalidCmdForSignalType = errors.New("invalid command for signal type")
	ErrCommandTimedOut         = errors.New("command timed out")
	ErrBusy                    = errors.New("busy")
	ErrPrivilegeViolation      = errors.New("privilege violation")
	ErrDeviceNotPresent        = errors.New("device not present")
	ErrMaxConnExceeded         = errors.New("maximum number of connections exceeded")
	ErrInvalidEventNumber      = errors.New("invalid event number")
	ErrBadFile                 = errors.New("bad filename/file not found")
	ErrUnknown                 = errors.New("an unspecified error has occurred while performing this operation")
)

func (g Goextron) getErr(errStr string) error {
	switch errStr {
	case "E01":
		return ErrInvalidInputNumber
	case "E06":
		return ErrInvalidSwitchAttempt
	case "E10":
		return ErrInvalidCommand
	case "E11":
		return ErrInvalidPresetNumber
	case "E12":
		return ErrInvalidPortNumber
	case "E13":
		return ErrInvalidParameter
	case "E14":
		return ErrInvalidForConfig
	case "E17":
		return ErrInvalidCmdForSignalType
	case "E18":
		return ErrCommandTimedOut
	case "E22":
		return ErrBusy
	case "E24":
		return ErrPrivilegeViolation
	case "E25":
		return ErrDeviceNotPresent
	case "E26":
		return ErrMaxConnExceeded
	case "E27":
		return ErrInvalidEventNumber
	case "E28":
		return ErrBadFile
	default:
		return ErrUnknown
	}
}
