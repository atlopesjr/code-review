package handler

import (
	"app/internal"
	"app/internal/repository"
	"app/internal/service"
	"encoding/json"
	"errors"
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

type VehicleSpeedJSON struct {
	MaxSpeed float64 `json:"max_speed"`
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
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

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

func (h *VehicleDefault) UpdateVehicleSpeedById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...
		id := chi.URLParam(r, "id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "invalid year param",
			})
		}

		var reqBody VehicleSpeedJSON
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "invalid request body",
			})
			return
		}

		// process
		// - get all vehicles
		err = h.sv.UpdateSpeedById(idInt, reqBody.MaxSpeed)
		if err != nil {
			response.JSON(w, handleError(err), nil)
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
		})
	}
}

func handleError(err error) int {
	if errors.Is(err, repository.NotFoundError) {
		return http.StatusNotFound
	}
	if errors.Is(err, service.ErrBadRequest) {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}
