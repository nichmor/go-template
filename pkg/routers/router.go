package routers

import (
	"github.com/valyala/fasthttp"
)

func Health(ctx *fasthttp.RequestCtx) {
	ctx.Response.AppendBody([]byte("healthy"))
}
