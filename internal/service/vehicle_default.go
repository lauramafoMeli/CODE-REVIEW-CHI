package service

import "app/internal"

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

// Create is a method that creates a vehicle
func (s *VehicleDefault) Create(v internal.Vehicle) (err error) {
	err = s.rp.Create(v)
	return
}

// GetByColorAndYear is a method that returns a map of vehicles by color and year
func (s *VehicleDefault) GetByColorAndYear(color string, year int) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.GetByColorAndYear(color, year)
	return
}

// GetByBrandAndYearRange is a method that returns a map of vehicles by brand and year range
func (s *VehicleDefault) GetByBrandAndYearRange(brand string, startYear int, finishYear int) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.GetByBrandAndYearRange(brand, startYear, finishYear)
	return
}

// GetAverageSpeedByBrand is a method that returns the average speed of vehicles by brand
func (s *VehicleDefault) GetAverageSpeedByBrand(brand string) (averageSpeed float64, err error) {
	averageSpeed, err = s.rp.GetAverageSpeedByBrand(brand)
	return
}

// CreateMultiple is a method that creates multiple vehicles
func (s *VehicleDefault) CreateMultiple(v []internal.Vehicle) (err error) {
	err = s.rp.CreateMultiple(v)
	return
}

// UpdateSpeed is a method that updates the speed of a vehicle
func (s *VehicleDefault) Update(id int, fields map[string]any) (err error) {
	err = s.rp.Update(id, fields)
	return
}

// GetByFuelType is a method that returns a map of vehicles by fuel type
func (s *VehicleDefault) GetByFuelType(fuelType string) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.GetByFuelType(fuelType)
	return
}

// Delete is a method that deletes a vehicle
func (s *VehicleDefault) Delete(id int) (err error) {
	err = s.rp.Delete(id)
	return
}

// GetByTransmission is a method that returns a map of vehicles by transmission type
func (s *VehicleDefault) GetByTransmission(transmission string) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.GetByTransmission(transmission)
	return
}

// GetAverageCapacityByBrand is a method that returns the average capacity of vehicles by brand
func (s *VehicleDefault) GetAverageCapacityByBrand(brand string) (averageCapacity float64, err error) {
	averageCapacity, err = s.rp.GetAverageCapacityByBrand(brand)
	return
}

// GetByDimensions is a method that returns a map of vehicles by dimensions
func (s *VehicleDefault) GetByDimensions(dimensions map[string]float64) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.GetByDimensions(dimensions)
	return
}

// GetByWeight is a method that returns a map of vehicles by weight
func (s *VehicleDefault) GetByWeight(weight map[string]float64) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.GetByWeight(weight)
	return
}
