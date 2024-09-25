package handler

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/rs/zerolog/log"
	"github.com/templhtmx/types"
)

func render(r *http.Request, w http.ResponseWriter, component templ.Component) error {
	return component.Render(r.Context(), w)
}

func getAuthenticatedUser(r *http.Request) types.AuthenticatedUser {

	user, ok := r.Context().Value(types.UserCtxKey).(types.AuthenticatedUser)

	if !ok {
		return types.AuthenticatedUser{}
	}

	return user

}

func Make(h func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			log.Error().AnErr("error", err).Str("path", r.URL.Path).Msg("interal server error")
		}
	}

}
