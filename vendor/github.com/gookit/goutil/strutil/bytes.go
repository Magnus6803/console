package strutil

import (
	"github.com/gookit/goutil/byteutil"
)

// Buffer wrap and extends the bytes.Buffer
type Buffer = byteutil.Buffer

// NewBuffer instance
func NewBuffer() *Buffer {
	return &Buffer{}
}

// ByteChanPool struct
//
// Usage:
//
//	bp := strutil.NewByteChanPool(500, 1024, 1024)
//	buf:=bp.Get()
//	defer bp.Put(buf)
//	// use buf do something ...
type ByteChanPool = byteutil.ChanPool

// NewByteChanPool instance
func NewByteChanPool(maxSize, width, capWidth int) *ByteChanPool {
	return byteutil.NewChanPool(maxSize, width, capWidth)
}
