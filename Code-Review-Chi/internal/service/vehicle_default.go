package service

import (
	"app/internal"
	"errors"
)

var ErrBadRequest = errors.New("invalid request")

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp internal.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

func (s *VehicleDefault) UpdateSpeedById(id int, newSpeed float64) (err error) {
	if newSpeed < 0 {
		err = ErrBadRequest
		return
	}
	err = s.rp.UpdateSpeedById(id, newSpeed)
	return
}
