package wrapper

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"go-template/internals/model"
)

func (w Wrapper) Ping(ctx context.Context, input model.Ping) (*model.Pong, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "Interface.User.AcceptConsent")
	defer sp.Finish()

	sp.LogFields(
		log.String("event", "ping"),
	)

	return w.Ping(ctx, input)
}
