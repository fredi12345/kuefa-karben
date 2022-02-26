package rest

import (
	"log"
	"net/http"
	"time"

	"github.com/fredi12345/kuefa-karben/src/storage"
	"github.com/labstack/echo/v4"
)

type (
	// CreateEventRequest is the request definition for creating events
	CreateEventRequest struct {
		// the topic of the event
		// @Required
		Theme string `json:"theme"`

		// the topic of the event
		// @Required
		ImageID string `json:"imageID"`

		// the date when the event begins
		// @Required
		StartingDate time.Time `json:"startingDate"`

		// the closing date for signing up to the event
		// @Required
		ClosingDate time.Time `json:"closingDate"`

		// the starter of the event
		// @Required
		Starter string `json:"starter"`

		// the main dish of the event
		// @Required
		MainDish string `json:"mainDish"`

		// the dessert of the event
		// @Required
		Dessert string `json:"dessert"`

		// the description of the event
		// @Required
		Description string `json:"description"`
	}

	// CreateEventResponse is the response after an event was created successfully
	CreateEventResponse struct {
		// a UUIDv4 to identify the event
		// @Required
		ID string `json:"id"`
	}
)

// CreateEvent allows to create an event.
//
// @OperationID CreateEvent
// @Title create an event
// @Param event body CreateEventRequest true "CreateEventRequest"
// @Success 200 object CreateEventResponse "Successfully created the event"
// @Failure 400 object ErrorResponse "Error while creating the event"
// @Failure 500 object ErrorResponse "Error while creating the event"
// @Route /events [post]
func (s *Server) CreateEvent(c echo.Context) error {
	var request CreateEventRequest
	err := c.Bind(&request)
	if err != nil {
		log.Printf("could not bind request: %v", err)
		return echo.ErrBadRequest
	}

	eventID, err := s.db.CreateEvent(storage.Event{
		Theme:       request.Theme,
		EventDate:   request.StartingDate,
		ClosingDate: request.ClosingDate,
		Starter:     request.Starter,
		MainDish:    request.MainDish,
		Dessert:     request.Dessert,
		InfoText:    request.Description,
		ImageID:     request.ImageID,
	})
	if err != nil {
		log.Printf("could not create event: %v", err)
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, CreateEventResponse{
		ID: eventID,
	})
}
