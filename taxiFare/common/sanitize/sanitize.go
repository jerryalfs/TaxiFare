package sanitize

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
	"taxiFare/common"
)

// BindError remove unused parameter binding error
func BindError(errs error) (result error) {
	if errs == nil {
		return
	}

	result = errs

	var errorMessages []string
	messages := strings.Split(errs.Error(), ",")
	for _, m := range messages {
		trim := strings.Trim(m, " ")
		// ignore empty string
		if trim == "" {
			continue
		}
		// ignore parse error
		if strings.Contains(trim, "strconv.Parse") {
			continue
		}
		errorMessages = append(errorMessages, trim)
	}
	if len(errorMessages) == 0 {
		return nil
	}

	return errors.New(strings.Join(errorMessages, ", "))
}

// RenderError write error response into http writer
func RenderError(w http.ResponseWriter, err error, status int) {
	if err == nil {
		return
	}

	response := common.Error{
		Status: status,
		Errors: []string{err.Error()},
	}
	Render(w, response, "")
}

// Render write response into http writer
func Render(w http.ResponseWriter, response interface{}, callback string) {
	jsonResponse, _ := jsoniter.ConfigFastest.Marshal(response)
	expires := time.Now().Add(5 * time.Minute)
	now := time.Now()

	w.Header().Set("Cache-Control", "public, max-age=300")
	w.Header().Set("Expires", expires.Format(time.RFC1123))
	w.Header().Set("Date", now.Format(time.RFC1123))

	if callback != "" {
		w.Header().Set("Content-Type", "text/javascript")
		fmt.Fprintf(w, "%s(%s)", callback, jsonResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Write(jsonResponse)
}


