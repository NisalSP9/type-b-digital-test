package router

import (
	"net/http"

	"github.com/NisalSP9/type-b-digital-test/internal/handler"
	"github.com/NisalSP9/type-b-digital-test/internal/middleware"
)

func New() http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handler.ServerHealth)
	mux.HandleFunc("GET /hello-world", handler.CheckName)

	return middleware.EnableCORS(mux)

}
