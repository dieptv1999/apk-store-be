package lib

import (
	"github.com/gorilla/sessions"
)

func NewSessionStore(logger Logger) *sessions.CookieStore {
	store := sessions.NewCookieStore([]byte("techlens"))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 1, // 1 hour
		HttpOnly: true,
		Domain:   "localhost",
	}
	return store
}
