package filepicker

import "strconv"

// FPError represents an error which could be produced by filepicker client.
type FPError int

// Error satisfies builtin.error interface. It prints the error string with
// the reason of failure.
func (e FPError) Error() string {
	var prefix = "filepicker: " + strconv.Itoa(int(e)) + " - "
	if msg, ok := errmsgs[e]; ok {
		return prefix + msg
	}
	return prefix + "unknown error"
}

const (
	ErrFileNotFound         FPError = 115
	ErrGeneralReadError     FPError = 118
	ErrFileStoreUnreachable FPError = 151
	ErrRemoteUrlUnreachable FPError = 152
	ErrBadParameters        FPError = 400
	ErrInvalidRequest       FPError = 403
)

var errmsgs = map[FPError]string{
	ErrFileNotFound:         "file not found",
	ErrGeneralReadError:     "general read error",
	ErrFileStoreUnreachable: "the file store could not be reached",
	ErrRemoteUrlUnreachable: "the remote URL could not be reached",
	ErrBadParameters:        "bad parameters were passed to the server",
	ErrInvalidRequest:       "invalid request",
}
