package http

//go:generate mockgen -source=service.go -destination=service_mock.go -package=http

type userServiceProvider interface {
}
