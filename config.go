package goss

import (
	"fmt"
	"time"
)

const (
	unsupportedOptionsStructFmt = "unsupported struct specified for options (%T)"
)

// DialConfig configures a socket dial attempt.
type DialConfig struct {
	// Path is the path to use when connecting to the socket.
	// Generally speaking, this can be a Unix file path.
	//
	// When compiled for Windows, the Windows pipe prefix is
	// automatically applied if it is not already present.
	// Programs not using this library can connect to the
	// pipe by specifying WindowsPipePrefix + Path.
	Path string

	// Timeout is the maximum amount of time to wait for
	// the dial operation to succeed. A zero time.Duration
	// means no timeout (with the exception of Windows,
	// where a timeout is chosen by go-winio if a zero
	// value is specified).
	Timeout time.Duration
}

func (o DialConfig) Validate() error {
	if len(o.Path) == 0 {
		return fmt.Errorf("path cannot be empty")
	}

	return nil
}

// ListenConfig configures a listener socket.
type ListenConfig struct {
	// Path is the path to use when creating a new listener socket.
	//
	// When compiled for Windows, the Windows pipe prefix is
	// automatically appended to the path if it is not present.
	Path string

	// SystemOptions is an optional pointer to an operating system
	// specific struct. If non-nil, the struct's values are used
	// instead of the defaults.
	//
	// The supported types are:
	//	- Unix systems: *UnixListenerOptions
	//	- Windows: *winio.PipeConfig
	SystemOptions interface{}
}

func (o ListenConfig) Validate() error {
	if len(o.Path) == 0 {
		return fmt.Errorf("path cannot be empty")
	}

	return nil
}
