package user

type Service struct {
	dbRepo dbRepoProvider
}

func NewService(dbRepo dbRepoProvider) *Service {
	return &Service{
		dbRepo: dbRepo,
	}
}
