package keys

import (
	"context"

	keysv1 "github.com/unkeyed/unkey/apps/agent/gen/proto/keys/v1"
	"github.com/unkeyed/unkey/apps/agent/pkg/tracing"
)

type tracingMiddleware struct {
	tracer tracing.Tracer
	next   KeyService
}

func WithTracing(tracer tracing.Tracer) Middleware {
	return func(svc KeyService) KeyService {
		return &tracingMiddleware{
			tracer: tracer,
			next:   svc,
		}
	}
}

func (mw *tracingMiddleware) CreateKey(ctx context.Context, req *keysv1.CreateKeyRequest) (*keysv1.CreateKeyResponse, error) {
	ctx, span := mw.tracer.Start(ctx, tracing.NewSpanName("workspaces", "CreateKey"))
	defer span.End()

	res, err := mw.next.CreateKey(ctx, req)
	if err != nil {
		span.RecordError(err)
	}
	return res, err
}

func (mw *tracingMiddleware) SoftDeleteKey(ctx context.Context, req *keysv1.SoftDeleteKeyRequest) (*keysv1.SoftDeleteKeyResponse, error) {
	ctx, span := mw.tracer.Start(ctx, tracing.NewSpanName("workspaces", "SoftDeleteKey"))
	defer span.End()

	res, err := mw.next.SoftDeleteKey(ctx, req)
	if err != nil {
		span.RecordError(err)
	}
	return res, err
}
