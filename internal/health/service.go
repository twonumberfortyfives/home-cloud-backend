package health

type HealthService struct{}

func (s *HealthService) CheckStatus() string {
	return "OK"
}
