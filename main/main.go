package main

import (
	"context"
	"fmt"
	"github.com/urfave/negroni"
	"jiong"
	"net/http"
)

func main() {
	ctx := context.Background()
	mux := http.NewServeMux()
	mux.Handle("/", jiong.Handler(ctx, test))
	n := negroni.New()
	n.UseHandler(mux)
	http.ListenAndServe(":8080", n)
}

func test(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	reqID := ctx.Value(jiong.RequestIDKey).(string)
	fmt.Fprintf(w, "Hello request ID %s\n", reqID)
}
