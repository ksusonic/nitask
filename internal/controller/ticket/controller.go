package ticket

import "log/slog"

type Controller struct {
	repository Repository
	log        *slog.Logger
}

func New(
	repository Repository,
	log *slog.Logger,
) *Controller {
	return &Controller{
		repository: repository,
		log:        log,
	}
}
