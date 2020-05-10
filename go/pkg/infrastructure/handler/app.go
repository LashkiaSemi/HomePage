package handler

type AppHandler struct {
	AuthHandler
}

func NewAppHandler() *AppHandler {
	return &AppHandler{
		AuthHandler: NewAuthHandler(),
	}
}
