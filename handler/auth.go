package handler

import (
	"net/http"

	"github.com/nedpals/supabase-go"
	"github.com/rs/zerolog/log"
	"github.com/templhtmx/view/auth"
)

func HandleSigninIndex(w http.ResponseWriter, r *http.Request) error {

	return render(r, w, auth.Signin())
}

func HandleSigninCreate(w http.ResponseWriter, r *http.Request) error {

	creds := supabase.UserCredentials{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	log.Debug().Any("creds", creds).Msg("user entered creds")

	return render(r, w, auth.LoginForm(creds, auth.LoginErrors{
		InvalidCreds: "The username or password you provided is incorrect",
	}))

}
