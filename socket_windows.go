package goss

import (
	"fmt"
	"math"
	"net"
	"strings"
	"time"

	"github.com/Microsoft/go-winio"
)

const (
	WindowsPipePrefix = `\\.\pipe\`
)

// Dial dials an existing socket.
func Dial(config DialConfig) (net.Conn, error) {
	err := config.Validate()
	if err != nil {
		return nil, err
	}

	var timeout *time.Duration
	if config.Timeout > 0 {
		timeout = &config.Timeout
	}

	return winio.DialPipe(ensurePipePrefix(config.Path), timeout)
}

// Listen creates a new listening socket.
func Listen(config ListenConfig) (net.Listener, error) {
	err := config.Validate()
	if err != nil {
		return nil, err
	}

	var pipeConfig *winio.PipeConfig
	switch asserted := config.SystemOptions.(type) {
	case *winio.PipeConfig:
		pipeConfig = asserted
	case nil:
		pipeConfig = &winio.PipeConfig{
			SecurityDescriptor: "",
			MessageMode:        false,
			InputBufferSize:    math.MaxInt32,
			OutputBufferSize:   math.MaxInt32,
		}
	default:
		return nil, fmt.Errorf(unsupportedOptionsStructFmt, config.SystemOptions)
	}

	return winio.ListenPipe(ensurePipePrefix(config.Path), pipeConfig)
}

func ensurePipePrefix(pipePath string) string {
	if strings.HasPrefix(pipePath, WindowsPipePrefix) {
		return pipePath
	}

	return WindowsPipePrefix + pipePath
}
