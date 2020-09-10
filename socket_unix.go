// +build !windows

package ossocket

import (
	"fmt"
	"net"
	"os"
)

// Dial dials an existing socket.
func Dial(config DialConfig) (net.Conn, error) {
	err := config.Validate()
	if err != nil {
		return nil, err
	}

	return net.DialTimeout("unix", config.Path, config.Timeout)
}

// UnixListenerOptions are Unix system specific options to use when creating
// a listening Unix socket.
type UnixListenerOptions struct {
	// TryRemove, if true, will cause the library to attempt removal
	// of the socket before creating the listener.
	TryRemove bool

	// FileMode is the file mode to chmod the socket to.
	FileMode os.FileMode
}

// Listen creates a new listening socket.
func Listen(config ListenConfig) (net.Listener, error) {
	var options *UnixListenerOptions
	switch asserted := config.SystemOptions.(type) {
	case *UnixListenerOptions:
		options = asserted
	case nil:
		// Do nothing.
	default:
		return nil, fmt.Errorf(unsupportedOptionsStructFmt, config.SystemOptions)
	}

	if options != nil && options.TryRemove {
		os.Remove(config.Path)
	}

	l, err := net.Listen("unix", config.Path)
	if err != nil {
		return nil, err
	}

	if options != nil {
		err := os.Chmod(config.Path, options.FileMode)
		if err != nil {
			l.Close()
			return nil, err
		}
	}

	return l, nil
}
