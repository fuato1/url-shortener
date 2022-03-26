package url

import "github.com/stretchr/testify/mock"

type URLMockRepository struct {
	mock.Mock
}

func (m URLMockRepository) GetAll() ([]ShortUrl, error) {
	args := m.Called()
	return args.Get(0).([]ShortUrl), args.Error(1)
}

func (m URLMockRepository) Add(url ShortUrl) error {
	args := m.Called(url)
	return args.Error(0)
}
