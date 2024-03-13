package handler

import (
	"app/internal"
	"app/platform/tools"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

// VehicleJSON is a struct that represents a vehicle in JSON format
type VehicleJSON struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

type MultipleVehicleJSON struct {
	Vehicles []VehicleJSON `json:"vehicles"`
}

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv internal.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv internal.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// Create is a method that returns a handler for the route POST /vehicles
func (h *VehicleDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - read body to bytes
		body, err := io.ReadAll(r.Body)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, nil)
			return
		}
		// - unmarshal body to array string any for validations
		bodyMap := map[string]any{}
		err = json.Unmarshal(body, &bodyMap)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, nil)
			return
		}

		// process
		// - validate body
		if err = tools.ValidateField(bodyMap, "brand", "model", "registration", "color", "year", "passengers", "max_speed", "fuel_type", "transmission", "weight", "height", "length", "width"); err != nil {
			var fieldError *tools.FieldError
			if errors.As(err, &fieldError) {
				response.JSON(w, http.StatusBadRequest, map[string]any{
					"message": errors.Join(internal.ErrFieldsMissing, errors.New(fieldError.Error())).Error(),
				})
				return
			}
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "internal error",
			})
			return
		}
		// - unmarshal body to vehicle
		var vehicle VehicleJSON
		err = json.Unmarshal(body, &vehicle)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": internal.ErrFieldsMissing.Error(),
			})
			return
		}
		// - create vehicle
		err = h.sv.Create(internal.Vehicle{
			Id: vehicle.ID,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           vehicle.Brand,
				Model:           vehicle.Model,
				Registration:    vehicle.Registration,
				Color:           vehicle.Color,
				FabricationYear: vehicle.FabricationYear,
				Capacity:        vehicle.Capacity,
				MaxSpeed:        vehicle.MaxSpeed,
				FuelType:        vehicle.FuelType,
				Transmission:    vehicle.Transmission,
				Weight:          vehicle.Weight,
				Dimensions: internal.Dimensions{
					Height: vehicle.Height,
					Length: vehicle.Length,
					Width:  vehicle.Width,
				},
			},
		})
		if err != nil {
			response.JSON(w, http.StatusConflict, map[string]any{
				"message": err.Error(),
			})
			return
		}

		// response
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": internal.MesgVehicleCreated,
		})
	}
}

// GetByColorAndYear is a method that returns a handler for the route GET /vehicles/color/{color}/year/{year}
func (h *VehicleDefault) GetByColorAndYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get color and year from url
		color := chi.URLParam(r, "color")
		year, err := strconv.Atoi(chi.URLParam(r, "year"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": err.Error(),
			})
			return
		}

		// process
		// - get vehicles by color and year
		v, err := h.sv.GetByColorAndYear(color, year)
		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]any{
				"message": err.Error(),
			})
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})

	}
}

// GetByBrandAndYearRange is a method that returns a map of vehicles for the route GET /vehicles/brand/{brand}/between/{start_year}/{end_year}
func (h *VehicleDefault) GetByBrandAndYearRange() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get brand and year range from url
		brand := chi.URLParam(r, "brand")
		yearStart, err := strconv.Atoi(chi.URLParam(r, "start_year"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": err.Error(),
			})
			return
		}
		yearEnd, err := strconv.Atoi(chi.URLParam(r, "end_year"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": err.Error(),
			})
			return
		}
		// validate year range
		if yearStart > yearEnd {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "400 Bad Request: year range is invalid",
			})
			return
		}

		// process
		// - get vehicles by brand and year range
		v, err := h.sv.GetByBrandAndYearRange(brand, yearStart, yearEnd)
		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]any{
				"message": err.Error(),
			})
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// GetAverageSpeedByBrand is a method that returns a map of vehicles for the route GET /vehicles/average-speed/brand/{brand}
func (h *VehicleDefault) GetAverageSpeedByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get brand from url
		brand := chi.URLParam(r, "brand")

		// process
		// - get average speed by brand
		v, err := h.sv.GetAverageSpeedByBrand(brand)
		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]any{
				"message": err.Error(),
			})
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    v,
		})
	}
}

// CreateMultiple is a method that returns a handler for the route POST /vehicles/batch
func (h *VehicleDefault) CreateMultiple() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": err.Error(),
			})
			return
		}

		// - unmarshal body to vehicles map
		var vehicles map[string](map[string]any)
		err = json.Unmarshal(body, &vehicles)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": err.Error(),
			})
			return
		}

		// process
		// - VehicleJSON to map[string]any
		for _, value := range vehicles {
			// - validate vehicles map
			if err = tools.ValidateField(value, "brand", "model", "registration", "color", "year", "passengers", "max_speed", "fuel_type", "transmission", "weight", "height", "length", "width"); err != nil {
				var fieldError *tools.FieldError
				if errors.As(err, &fieldError) {
					response.JSON(w, http.StatusBadRequest, map[string]any{
						"message": errors.Join(internal.ErrFieldsMissing, errors.New(fieldError.Error())).Error(),
					})
					return
				}
				response.JSON(w, http.StatusBadRequest, map[string]any{
					"message": "internal error",
				})
				return
			}
		}

		// - create vehicles slice
		var vehiclesSend []internal.Vehicle
		var vehicle internal.Vehicle
		for _, value := range vehicles {
			jsonData, err := json.Marshal(value)
			if err != nil {
				response.JSON(w, http.StatusBadRequest, map[string]any{
					"message": err.Error(),
				})
				return
			}
			err = json.Unmarshal(jsonData, &vehicle)
			if err != nil {
				response.JSON(w, http.StatusBadRequest, map[string]any{
					"message": err.Error(),
				})
				return
			}
			vehiclesSend = append(vehiclesSend, internal.Vehicle{
				Id: vehicle.Id,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           vehicle.Brand,
					Model:           vehicle.Model,
					Registration:    vehicle.Registration,
					Color:           vehicle.Color,
					FabricationYear: vehicle.FabricationYear,
					Capacity:        vehicle.Capacity,
					MaxSpeed:        vehicle.MaxSpeed,
					FuelType:        vehicle.FuelType,
					Transmission:    vehicle.Transmission,
					Weight:          vehicle.Weight,
					Dimensions: internal.Dimensions{
						Height: vehicle.Height,
						Length: vehicle.Length,
						Width:  vehicle.Width,
					},
				},
			})
		}

		err = h.sv.CreateMultiple(vehiclesSend)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": err.Error(),
			})
			return
		}

		// response
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": internal.MesgVehicleCreated,
		})
	}

}

// UpdateSpeed is a method that returns a handler for the route PATCH /vehicles/{id}/update_speed
func (h *VehicleDefault) UpdateSpeed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get id from url
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": err.Error(),
			})
			return
		}

		// - get body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": err.Error(),
			})
			return
		}

		// - unmarshal body to speed map
		var speed map[string]any
		err = json.Unmarshal(body, &speed)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": err.Error(),
			})
			return
		}

		// process
		// - validate speed map
		if err = tools.ValidateField(speed, "speed"); err != nil {
			var fieldError *tools.FieldError
			if errors.As(err, &fieldError) {
				response.JSON(w, http.StatusBadRequest, map[string]any{
					"message": errors.Join(internal.ErrFieldsMissing, errors.New(fieldError.Error())).Error(),
				})
				return
			}
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "internal error",
			})
			return
		}
		// - update speed
		speedValue := speed["speed"].(float64)
		if speedValue < 0 {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "400 Bad Request: Velocidad mal formada o fuera de rango.",
			})
			return
		}
		err = h.sv.UpdateSpeed(id, speedValue)
		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]any{
				"message": err.Error(),
			})
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": internal.MesgVehicleUpdatedSpeed,
		})
	}
}

// GetByFuelType is a method that returns a handler for the route GET /vehicles/fuel_type/{type}
func (h *VehicleDefault) GetByFuelType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get fuel type from url
		fuelType := chi.URLParam(r, "type")

		// process
		// - get vehicles by fuel type
		vehicles, err := h.sv.GetByFuelType(fuelType)
		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]any{
				"message": err.Error(),
			})
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    vehicles,
		})
	}
}

// Delete is a method that returns a handler for the route DELETE /vehicles/{id}
func (h *VehicleDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get id from url
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": err.Error(),
			})
			return
		}

		// process
		// - delete vehicle
		err = h.sv.Delete(id)
		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]any{
				"message": err.Error(),
			})
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": internal.MesgVehicleDeleted,
		})
	}
}

// GetByTransmission is a method that returns a handler for the route GET /vehicles/transmission/{type}
func (h *VehicleDefault) GetByTransmission() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get transmission type from url
		transmission := chi.URLParam(r, "type")

		// process
		// - get vehicles by transmission
		vehicles, err := h.sv.GetByTransmission(transmission)
		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]any{
				"message": err.Error(),
			})
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    vehicles,
		})
	}
}
