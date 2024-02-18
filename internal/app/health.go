package app

import (
	"fmt"
	"html"
	"net/http"
)

type HealthHandler struct{}

func (c *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}
