package main

type (
	IServicesInfs interface {
		GetTest() string
	}
	serviceDeps struct {
		repo IRepositoryInfs
	}
)

func NewService(repo IRepositoryInfs) IServicesInfs {
	return &serviceDeps{
		repo: repo,
	}
}

func (s *serviceDeps) GetTest() string {
	return s.repo.GetTest()
}
