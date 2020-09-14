package main

import (
    "fmt"
    "github.com/valyala/fasthttp"
    "sync"
)


func HandelHeaderCalls(ctx *fasthttp.RequestCtx) {
    // notice that we may access MyHandler properties here - see h.foobar.


    ctx.Response.SetBody(ctx.Request.Header.RawHeaders())
}

func main() {
    fmt.Print("Listening to 8080..\n")
    wg := sync.WaitGroup{}
    wg.Add(1)
    fasthttp.ListenAndServe(":8080", HandelHeaderCalls)
    wg.Wait()
}
