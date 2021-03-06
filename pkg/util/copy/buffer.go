// Copyright (c) 2018, Sylabs Inc. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE.md file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package copy

import (
	"bytes"
	"sync"
)

// TerminalBuffer captures the last line displayed on terminal.
type TerminalBuffer struct {
	data []byte
	sync.Mutex
}

// NewTerminalBuffer returns an instantiated TerminalBuffer.
func NewTerminalBuffer() *TerminalBuffer {
	b := &TerminalBuffer{}
	b.data = make([]byte, 0)
	return b
}

// Write implements the write interface to store last terminal line.
func (b *TerminalBuffer) Write(p []byte) (n int, err error) {
	b.Lock()
	defer b.Unlock()

	if bytes.IndexByte(p, '\n') >= 0 {
		b.data = nil
	} else {
		b.data = append(b.data, p...)
	}

	return len(p), nil
}

// Line returns the last terminal line.
func (b *TerminalBuffer) Line() []byte {
	return b.data
}
