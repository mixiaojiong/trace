package middleware

import (
	"fmt"
	"github.com/satori/go.uuid"
	"net/http"
)

type TraceId struct {
}

func NewTraceId() *TraceId {
	return &TraceId{}
}

func (this *TraceId) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var traceId string
	traceIds := r.Header["Traceid"]
	if traceIds != nil && len(traceIds) > 0 && traceIds[0] != "" {
		traceId = traceIds[0]
	} else {
		u1 := uuid.NewV4()
		traceId = u1.String()
	}
	fmt.Printf("traceId: %s\n", traceId)
	w.Header().Set("Traceid", traceId)
	next(w, r)
}
