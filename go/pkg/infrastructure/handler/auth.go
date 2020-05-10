package handler

type AuthHandler interface {
	// Login() http.HandlerFunc
	// Logout() http.HandlerFunc
}

type authHandler struct {
}

func NewAuthHandler() {

}

// func (ah *authHandler) Login() http.HandlerFunc {
// 	return func()
// }

// func (ah *authHandler) Logout() http.HandlerFunc {

// }
