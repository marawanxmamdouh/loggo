/*
Copyright © 2022 Aurelio Calegari, et al.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package pkg

import (
	"github.com/marawanxmamdouh/loggo/reader"
)

// Reader is an interface for reading log streams.
type Reader interface {
	// StreamInto feeds the strChan channel for every streamed line.
	StreamInto() error
	// Close finalises and invalidates this stream reader.
	Close()
	// ChanReader returns the outbound channel reader
	ChanReader() <-chan string
	// ErrorNotifier registers a callback func that's called upon fatal streaming log.
	ErrorNotifier(onError func(err error))
}

// ReaderType represents the type of reader.
type ReaderType = reader.Type

const (
	// TypeFile represents a file reader.
	TypeFile = reader.TypeFile
	// TypePipe represents a pipe reader.
	TypePipe = reader.TypePipe
)

// NewReader creates a new reader for the given file or stdin if fileName is empty.
func NewReader(fileName string, strChan chan string) Reader {
	return reader.MakeReader(fileName, strChan)
}
