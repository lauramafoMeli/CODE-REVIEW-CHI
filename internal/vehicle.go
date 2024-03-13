package internal

import "errors"

// Dimensions is a struct that represents a dimension in 3d
type Dimensions struct {
	// Height is the height of the dimension
	Height float64
	// Length is the length of the dimension
	Length float64
	// Width is the width of the dimension
	Width float64
}

// VehicleAttributes is a struct that represents the attributes of a vehicle
type VehicleAttributes struct {
	// Brand is the brand of the vehicle
	Brand string
	// Model is the model of the vehicle
	Model string
	// Registration is the registration of the vehicle
	Registration string
	// Color is the color of the vehicle
	Color string
	// FabricationYear is the fabrication year of the vehicle
	FabricationYear int
	// Capacity is the capacity of people of the vehicle
	Capacity int
	// MaxSpeed is the maximum speed of the vehicle
	MaxSpeed float64
	// FuelType is the fuel type of the vehicle
	FuelType string
	// Transmission is the transmission of the vehicle
	Transmission string
	// Weight is the weight of the vehicle
	Weight float64
	// Dimensions is the dimensions of the vehicle
	Dimensions
}

// Vehicle is a struct that represents a vehicle
type Vehicle struct {
	// Id is the unique identifier of the vehicle
	Id int

	// VehicleAttribue is the attributes of a vehicle
	VehicleAttributes
}

// Var for different types of errors and messages
var (
	//messages
	MesgVehicleCreated      = "201 Created: Vehículo creado exitosamente."
	MesgVehicleUpdatedSpeed = "200 OK: Velocidad del vehículo actualizada exitosamente."
	MesgVehicleDeleted      = "204 No Content: Vehículo eliminado exitosamente."

	// errors
	ErrVehicleAlreadyExists          = errors.New("409 Conflict: Identificador del vehículo ya existente.")
	ErrFieldsMissing                 = errors.New("400 Bad Request: Datos del vehículo mal formados o incompletos.")
	ErrVehicleNotFound               = errors.New("404 Not Found: No se encontraron vehículos con esos criterios.")
	ErrVehicleNotFoundByBrand        = errors.New("404 Not Found: No se encontraron vehículos de esa marca.")
	ErrVehicleNotFoundByTransmission = errors.New("404 Not Found: No se encontraron vehiculos con ese tipo de transmisión.")
)
