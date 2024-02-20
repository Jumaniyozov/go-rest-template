package response

import (
	"encoding/json"
	"github.com/rs/zerolog"
	"net/http"
)

type Response struct {
	logger *zerolog.Logger
}

func New(log *zerolog.Logger) *Response {
	return &Response{
		logger: log,
	}
}

func (r *Response) InternalErrorMessage(w http.ResponseWriter, err error) {
	if r.logger != nil {
		r.logger.Error().Err(err).Msg("Error while writing response")
	}
	http.Error(w, "error trying to send json", http.StatusInternalServerError)
}

func (r *Response) WriteJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(js)
	if err != nil {
		return err
	}
	return nil
}

func (r *Response) FetchError(w http.ResponseWriter, msg string) {
	resp := Message{
		Message: msg,
		Code:    http.StatusInternalServerError,
		Success: false,
		Data:    nil,
	}

	err := r.WriteJSON(w, http.StatusInternalServerError, resp, nil)
	if err != nil {
		r.InternalErrorMessage(w, err)
	}
}

func (r *Response) BadRequest(w http.ResponseWriter, msg string) {
	rsp := Message{
		Message: msg,
		Code:    http.StatusBadRequest,
		Success: false,
		Data:    nil,
	}
	err := r.WriteJSON(w, http.StatusBadRequest, rsp, nil)
	if err != nil {
		r.InternalErrorMessage(w, err)
	}
}

func (r *Response) PermissionDenied(w http.ResponseWriter, msg string) {
	rsp := Message{
		Message: msg,
		Code:    http.StatusUnauthorized,
		Success: false,
		Data:    nil,
	}
	err := r.WriteJSON(w, http.StatusUnauthorized, rsp, http.Header{
		"WWW-Authenticate": []string{"Basic realm=Restricted"},
	})
	if err != nil {
		r.InternalErrorMessage(w, err)
	}
}

func (r *Response) InternalServerError(w http.ResponseWriter, msg string) {
	rsp := Message{
		Message: msg,
		Code:    http.StatusInternalServerError,
		Success: false,
		Data:    nil,
	}
	err := r.WriteJSON(w, http.StatusInternalServerError, rsp, nil)
	if err != nil {
		r.InternalErrorMessage(w, err)
	}
}
