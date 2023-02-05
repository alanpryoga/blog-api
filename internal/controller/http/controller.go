package http

type Controller struct {
	userSvc userServiceProvider
}

func NewController(userSvc userServiceProvider) *Controller {
	return &Controller{
		userSvc: userSvc,
	}
}
