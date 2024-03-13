package repository

import (
	"app/internal"
)

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]internal.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]internal.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

// Create is a method that creates a vehicle
func (r *VehicleMap) Create(v internal.Vehicle) (err error) {
	// validate vehicle ID
	for _, value := range r.db {
		if value.Id == v.Id {
			err = internal.ErrVehicleAlreadyExists
			return
		}
	}
	// add vehicle to db
	r.db[v.Id] = v
	return
}

// GetByColorAndYear is a method that returns a map of vehicles by color and year
func (r *VehicleMap) GetByColorAndYear(color string, year int) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	index := 0
	for _, value := range r.db {
		if value.Color == color && value.FabricationYear == year {
			v[index] = value
			index++
		}
	}

	if index == 0 {
		err = internal.ErrVehicleNotFound
	}

	return
}

// GetByBrandAndYearRange is a method that returns a map of vehicles by brand and year range
func (r *VehicleMap) GetByBrandAndYearRange(brand string, startYear int, finishYear int) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	index := 0
	for _, value := range r.db {
		if value.Brand == brand && value.FabricationYear >= startYear && value.FabricationYear <= finishYear {
			v[index] = value
			index++
		}
	}

	if index == 0 {
		err = internal.ErrVehicleNotFound
	}

	return
}

// GetAverageSpeedByBrand is a method that returns the average speed of vehicles by brand
func (r *VehicleMap) GetAverageSpeedByBrand(brand string) (averageSpeed float64, err error) {
	averageSpeed = 0.0
	count := 0

	for _, value := range r.db {
		if value.Brand == brand {
			averageSpeed += value.MaxSpeed
			count++
		}
	}

	if count == 0 {
		err = internal.ErrVehicleNotFoundByBrand
	}

	averageSpeed /= float64(count)
	return

}

// CreateMultiple is a method that creates multiple vehicles
func (r *VehicleMap) CreateMultiple(v []internal.Vehicle) (err error) {
	// Validate vehicles ID
	for _, vehicle := range v {
		for _, value := range r.db {
			if value.Id == vehicle.Id {
				err = internal.ErrVehicleAlreadyExists
				return
			}
		}

	}

	// Add vehicles to db
	for _, vehicle := range v {
		r.db[vehicle.Id] = vehicle
	}

	return
}
