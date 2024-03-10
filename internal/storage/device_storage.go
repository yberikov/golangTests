package storage

import (
	"errors"
	"homework/internal/models"
	"sync"
)

type DeviceStorage struct {
	sync.Mutex
	devices map[string]models.Device
}

var (
	ErrNoSuchDevice        = errors.New("there is no such device")
	ErrDeviceAlreadyExists = errors.New("such device is already in database")
)

func NewDeviceStorage() *DeviceStorage {
	return &DeviceStorage{
		devices: make(map[string]models.Device),
	}
}

func (s *DeviceStorage) GetDeviceBySerialNum(serialNum string) (models.Device, error) {
	defer s.Unlock()
	s.Lock()
	if val, ok := s.devices[serialNum]; ok {
		return val, nil
	}
	return models.Device{}, ErrNoSuchDevice
}

func (s *DeviceStorage) CreateDevice(device models.Device) error {
	defer s.Unlock()
	s.Lock()
	if _, ok := s.devices[device.SerialNum]; ok {
		return ErrDeviceAlreadyExists
	}
	s.devices[device.SerialNum] = device
	return nil
}

func (s *DeviceStorage) DeleteDeviceBySerialNum(serialNum string) error {
	defer s.Unlock()
	s.Lock()
	if _, ok := s.devices[serialNum]; ok {
		delete(s.devices, serialNum)
		return nil
	}
	return ErrNoSuchDevice
}

func (s *DeviceStorage) UpdateDevice(device models.Device) error {
	defer s.Unlock()
	s.Lock()
	if _, ok := s.devices[device.SerialNum]; ok {
		s.devices[device.SerialNum] = device
		return nil
	}
	return ErrNoSuchDevice
}
