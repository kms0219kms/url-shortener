package shortenService

type Service interface {
	RedirectToOriginal(shortened string) string
}

type service struct {
}

func NewService() *service {
	return &service{}
}

// -------------------------------------------------

func (s *service) RedirectToOriginal(shortened string) string {
	return shortened
}

func (s *service) CreateShorten(shortened string, original string) string {
	return shortened
}

func (s *service) GetShorten(shortened string) string {
	return shortened
}

func (s *service) UpdateShorten(shortened string, original string) string {
	return shortened
}

func (s *service) DeleteShorten(shortened string) string {
	return shortened
}
