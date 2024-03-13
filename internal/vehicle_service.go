package internal

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	// Create is a method that creates a vehicle
	Create(v Vehicle) (err error)
	// GetByColorAndYear is a method that returns a map of vehicles by color and year
	GetByColorAndYear(color string, year int) (v map[int]Vehicle, err error)
	// GetByBrandAndYearRange is a method that returns a map of vehicles by brand and year range
	GetByBrandAndYearRange(brand string, startYear int, finishYear int) (v map[int]Vehicle, err error)
	// GetAverageSpeedByBrand is a method that returns the average speed of vehicles by brand
	GetAverageSpeedByBrand(brand string) (averageSpeed float64, err error)
	// CreateMultiple is a method that creates multiple vehicles
	CreateMultiple(v []Vehicle) (err error)
	// UpdateSpeed is a method that updates the speed of a vehicle
	UpdateSpeed(id int, speed float64) (err error)
	// GetByFuelType is a method that returns a map of vehicles by fuel type
	GetByFuelType(fuelType string) (v map[int]Vehicle, err error)
	// Delete is a method that deletes a vehicle
	Delete(id int) (err error)
	// GetByTransmission is a method that returns a map of vehicles by transmission type
	GetByTransmission(transmission string) (v map[int]Vehicle, err error)
}
