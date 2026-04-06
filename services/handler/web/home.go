package web

type HomeService struct {
}

var HomeServiceInstance = new(HomeService)

func (s *HomeService) GetHomeArticle() {
}

func (s *HomeService) Info() {
}

func (s *HomeService) GetRecommendArticle() {
}

func (s *HomeService) GetTagCloud() {
}

func (s *HomeService) Search() {
}
