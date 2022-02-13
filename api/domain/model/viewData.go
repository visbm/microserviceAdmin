package model

import (
	"net/http"
)

type ViewData struct {
	w http.ResponseWriter
	r *http.Request
}
