package response

import (
	"github.com/rs/zerolog"
	"net/http"
)

func InternalErrorMessage(w http.ResponseWriter, err error, logger *zerolog.Logger) {
	logger.Error().Err(err).Msg("Error while writing response")
	http.Error(w, "error trying to send json", http.StatusInternalServerError)
}
