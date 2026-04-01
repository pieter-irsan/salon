package router

import (
	"net/http"
	"time"
)

//go:generate go tool oapi-codegen -config ./oapi-codegen.yaml ../../oapi/spec.dev.yaml

// Request Lifecycle
//
// 	Client connects
// 	│
// 	├──[ReadHeaderTimeout]──► headers fully received
// 	│
// 	├──[ReadTimeout]────────► entire request read (headers + body)
// 	│
// 	│  (handler runs)
// 	│
// 	├──[WriteTimeout]───────► response fully sent
// 	│
// 	│  (connection kept alive)
// 	│
// 	└──[IdleTimeout]────────► next request arrives, or connection closed

func New(port string, handler ServerInterface) *http.Server {
	return &http.Server{
		Addr:              ":" + port,
		Handler:           Handler(handler),
		ReadHeaderTimeout: 2 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
}
