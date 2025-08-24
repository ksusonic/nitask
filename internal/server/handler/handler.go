package handler

type Handler struct {
	*Deps
}

func New(deps *Deps) *Handler {
	return &Handler{
		Deps: deps,
	}
}
