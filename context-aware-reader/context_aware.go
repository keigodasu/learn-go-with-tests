package context_aware_reader

import (
	"context"
	"io"
)

func NewCancellableReader(ctx context.Context, rdr io.Reader) io.Reader {
	return &readCtx{
		ctx:      ctx,
		delegate: rdr,
	}
}

func (r *readCtx) Read(p []byte) (n int, err error) {
	if err := r.ctx.Err(); err != nil {
		return 0, err
	}
	return r.delegate.Read(p)
}

type readCtx struct {
	ctx      context.Context
	delegate io.Reader
}
