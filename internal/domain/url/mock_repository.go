package url

import "github.com/stretchr/testify/mock"

type URLMockRepository struct {
	mock.Mock
}

func (m URLMockRepository) GetAll() ([]URL, error) {
	args := m.Called()
	return args.Get(0).([]URL), args.Error(1)
}

func (m URLMockRepository) Add(url URL) error {
	args := m.Called(url)
	return args.Error(0)
}
