package repository

import (
	"app/internal"
	"fmt"
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
	for index, value := range r.db {
		if value.Color == color && value.FabricationYear == year {
			v[index] = value
		}
	}

	if len(v) == 0 {
		err = internal.ErrVehicleNotFound
	}

	return
}

// GetByBrandAndYearRange is a method that returns a map of vehicles by brand and year range
func (r *VehicleMap) GetByBrandAndYearRange(brand string, startYear int, finishYear int) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for index, value := range r.db {
		if value.Brand == brand && value.FabricationYear >= startYear && value.FabricationYear <= finishYear {
			v[index] = value
		}
	}

	if len(v) == 0 {
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

// UpdateSpeed is a method that updates the speed of a vehicle
func (r *VehicleMap) Update(id int, fields map[string]any) (err error) {
	vehicle, ok := r.db[id]
	if !ok {
		err = internal.ErrVehicleNotFound
		return
	}

	/*vehicle.MaxSpeed = speed // Update the speed of the copied vehicle
	r.db[id] = vehicle       // Assign the updated vehicle back to the map
	return*/

	for key, value := range fields {
		switch key {
		case "speed":
			vehicle.MaxSpeed, ok = value.(float64)
			if !ok {
				err = internal.ErrFieldsMissing
				return
			}
		case "fuel_type":
			vehicle.FuelType, ok = value.(string)
			if !ok {
				err = internal.ErrFieldsMissing
				return
			}
		default:
			err = internal.ErrFieldsMissing
			return
		}
	}
	r.db[id] = vehicle
	return
}

// GetByFuelType is a method that returns a map of vehicles by fuel type
func (r *VehicleMap) GetByFuelType(fuelType string) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for index, value := range r.db {
		if value.FuelType == fuelType {
			v[index] = value
		}
	}

	if len(v) == 0 {
		err = internal.ErrVehicleNotFound
	}

	return
}

// Delete is a method that deletes a vehicle
func (r *VehicleMap) Delete(id int) (err error) {
	if _, ok := r.db[id]; !ok {
		err = internal.ErrVehicleNotFound
		return
	}

	delete(r.db, id)
	return
}

// GetByTransmission is a method that returns a map of vehicles by transmission type
func (r *VehicleMap) GetByTransmission(transmission string) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for index, value := range r.db {
		if value.Transmission == transmission {
			v[index] = value
		}
	}

	if len(v) == 0 {
		err = internal.ErrVehicleNotFoundByTransmission
	}

	return
}

// GetAverageCapacityByBrand is a method that returns the average capacity of vehicles by brand
func (r *VehicleMap) GetAverageCapacityByBrand(brand string) (averageCapacity float64, err error) {
	averageCapacity = 0.0
	count := 0

	for _, value := range r.db {
		if value.Brand == brand {
			averageCapacity += float64(value.Capacity)
			count++
		}
	}

	if count == 0 {
		err = internal.ErrVehicleNotFoundByBrand
	}

	averageCapacity /= float64(count)
	return
}

// GetByDimensions is a method that returns a map of vehicles by dimension
func (r *VehicleMap) GetByDimensions(dimensions map[string]float64) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	_, ok_max_length := dimensions["max_length"]
	_, ok_max_width := dimensions["max_width"]

	fmt.Println(dimensions)

	if !ok_max_length && !ok_max_width {
		v, err = r.FindAll()
	} else if !ok_max_length {
		for index, value := range r.db {
			if value.Width <= dimensions["max_width"] && value.Width >= dimensions["min_width"] {
				v[index] = value
			}
		}
	} else if !ok_max_width {
		for index, value := range r.db {
			if value.Height <= dimensions["max_length"] && value.Height >= dimensions["min_length"] {
				v[index] = value
			}
		}
	} else {
		for index, value := range r.db {
			if value.Height <= dimensions["max_length"] && value.Height >= dimensions["min_length"] && value.Width <= dimensions["max_width"] && value.Width >= dimensions["min_width"] {
				v[index] = value
			}
		}
	}

	return
}

// GetByWeight is a method that returns a map of vehicles by weight
func (r *VehicleMap) GetByWeight(weight map[string]float64) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	fmt.Println(weight)

	// copy db
	_, ok_max_weight := weight["max"]
	_, ok_min_weight := weight["min"]

	if !ok_max_weight && !ok_min_weight {
		v, err = r.FindAll()
	} else if !ok_max_weight {
		for index, value := range r.db {
			if value.Weight >= weight["min"] {
				v[index] = value
			}
		}
	} else if !ok_min_weight {
		for index, value := range r.db {
			if value.Weight <= weight["max"] {
				v[index] = value
			}
		}
	} else {
		for index, value := range r.db {
			if value.Weight <= weight["max"] && value.Weight >= weight["min"] {
				v[index] = value
			}
		}
	}

	return
}
