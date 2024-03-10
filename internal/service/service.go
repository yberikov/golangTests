package service

import (
	"homework/internal/models"
)

//go:generate go run github.com/vektra/mockery/v2@v2.36.0 --name=DeviceStorage
type DeviceStorage interface {
	GetDeviceBySerialNum(serialNum string) (models.Device, error)
	CreateDevice(device models.Device) error
	DeleteDeviceBySerialNum(serialNum string) error
	UpdateDevice(device models.Device) error
}

type DeviceService struct {
	storage DeviceStorage
}

func NewService(storage DeviceStorage) *DeviceService {
	return &DeviceService{
		storage: storage,
	}
}

func (s *DeviceService) GetDevice(serialNum string) (models.Device, error) {
	device, err := s.storage.GetDeviceBySerialNum(serialNum)
	if err != nil {
		return models.Device{}, err
	}
	return device, nil
}

func (s *DeviceService) CreateDevice(device models.Device) error {
	if err := s.storage.CreateDevice(device); err != nil {
		return err
	}
	return nil
}

func (s *DeviceService) DeleteDevice(serialNum string) error {
	if err := s.storage.DeleteDeviceBySerialNum(serialNum); err != nil {
		return err
	}
	return nil
}

func (s *DeviceService) UpdateDevice(device models.Device) error {

	err := s.storage.UpdateDevice(device)
	if err != nil {
		return err
	}
	return nil
}
