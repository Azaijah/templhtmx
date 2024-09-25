package handler

import (
	"context"
	"net/http"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/templhtmx/types"
)

func WithUser(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/public") {
			next.ServeHTTP(w, r)
			return
		}

		user := types.AuthenticatedUser{Email: "testemail@gmail.com", LoggedIn: true}
		ctx := context.WithValue(r.Context(), types.UserCtxKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))

		log.Debug().Any("auth user", user).Msg("with user action")

	}

	return http.HandlerFunc(fn)

}
