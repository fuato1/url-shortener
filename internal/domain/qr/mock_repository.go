package qr

import "github.com/stretchr/testify/mock"

type QRMockRepository struct {
	mock.Mock
}

func (m QRMockRepository) GetAll() ([]QR, error) {
	args := m.Called()
	return args.Get(0).([]QR), args.Error(1)
}

func (m QRMockRepository) Add(qr QR) error {
	args := m.Called(qr)
	return args.Error(0)
}
