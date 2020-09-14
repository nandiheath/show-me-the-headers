package main

import (
    "github.com/valyala/fasthttp"
)


func HandelHeaderCalls(ctx *fasthttp.RequestCtx) {
    // notice that we may access MyHandler properties here - see h.foobar.


    ctx.Response.SetBody(ctx.Request.Header.RawHeaders())
}

func main() {
    fasthttp.ListenAndServe(":8080", HandelHeaderCalls)
}
