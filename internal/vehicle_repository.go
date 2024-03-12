package internal

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
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
}
