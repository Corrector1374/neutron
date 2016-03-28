package chunksplit

import (
	"io"
)

type writer struct {
	Sep string
	Len int

	w io.Writer
	i int
}

func (w *writer) Write(b []byte) (N int, err error) {
	to := w.Len - w.i

	for len(b) > to {
		var n int
		n, err = w.w.Write(b[:to])
		if err != nil {
			return
		}
		N += n
		b = b[to:]

		_, err = w.w.Write([]byte(w.Sep))
		if err != nil {
			return
		}

		w.i = 0
		to = w.Len
	}

	w.i += len(b)

	n, err := w.w.Write(b)
	if err != nil {
		return
	}
	N += n

	return
}

func New(sep string, l int, w io.Writer) io.Writer {
	return &writer{
		Sep: sep,
		Len: l,
		w: w,
	}
}