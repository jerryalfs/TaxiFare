package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Delivery is an interface wrapping http router to passing parameter from body or query string
type Delivery interface {
	GetTaxiFareHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params)
}
