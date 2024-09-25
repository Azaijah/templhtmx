package types

type key string

const UserCtxKey key = "user"

type AuthenticatedUser struct {
	Email    string
	LoggedIn bool
}
