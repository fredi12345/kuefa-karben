package rest

import (
	"net/http"
	"time"

	"github.com/fredi12345/kuefa-karben/src/storage"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type (
	// CreateEventRequest is the request definition for creating events
	CreateEventRequest struct {
		// the topic of the event
		// @Required
		Theme string `json:"theme" validate:"required,lte=256"`

		// the topic of the event
		// @Required
		ImageID string `json:"imageID" validate:"required,uuid4"`

		// the date when the event begins
		// @Required
		StartingDate time.Time `json:"startingDate" validate:"required"`

		// the closing date for signing up to the event
		// @Required
		ClosingDate time.Time `json:"closingDate" validate:"required"`

		// the starter of the event
		// @Required
		Starter string `json:"starter" validate:"required,lte=512"`

		// the main dish of the event
		// @Required
		MainDish string `json:"mainDish" validate:"required,lte=512"`

		// the dessert of the event
		// @Required
		Dessert string `json:"dessert" validate:"required,lte=512"`

		// the description of the event
		// @Required
		Description string `json:"description" validate:"required,lte=2028"`
	}

	// CreateEventResponse is the response after an event was created successfully
	CreateEventResponse struct {
		// a UUIDv4 to identify the event
		// @Required
		ID string `json:"id"`
	}

	// GetEventsRequest is the request definition for fetching a list of events
	GetEventsRequest struct {
		Limit  int `query:"limit" validate:"gte=0"`
		Offset int `query:"offset" validate:"gte=0"`
	}

	// GetEventsResponse is the response containing the list of events
	GetEventsResponse struct {
		// the list of events
		// @Required
		Events []EventTeaser `json:"events"`
	}

	EventTeaser struct {
		// a UUIDv4 to identify the event
		// @Required
		ID string `json:"id"`

		// link to the event thumbnail
		// @Required
		ThumbnailURL string `json:"thumbnailURL"`

		// the topic of the event
		// @Required
		Theme string `json:"theme"`

		// the time when the event begins
		// @Required
		Date time.Time `json:"date"`
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
func (s *Server) CreateEvent(c echo.Context, request CreateEventRequest) error {
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
		s.l.Error("could not create event", zap.Error(err))
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, CreateEventResponse{
		ID: eventID,
	})
}

// GetEvents lists events according to the filter parameters.
//
// @OperationID GetEvents
// @Title get filtered event list
// @Param limit query int false "limit the number of events in the list, 0 means unlimited"
// @Param offset query int false "number of events to skip in the list"
// @Success 200 object GetEventsResponse "List of the events"
// @Failure 400 object ErrorResponse "Error while listing events"
// @Failure 500 object ErrorResponse "Error while listing events"
// @Route /events [get]
func (s *Server) GetEvents(c echo.Context, request GetEventsRequest) error {
	if request.Limit == 0 {
		request.Limit = 99999
	}

	dbTeasers, err := s.db.GetEventList(request.Offset, request.Limit)
	if err != nil {
		return echo.ErrInternalServerError
	}

	responseTeasers := make([]EventTeaser, 0, len(dbTeasers))
	for _, event := range dbTeasers {
		responseTeasers = append(responseTeasers, EventTeaser{
			ID:           event.ID,
			ThumbnailURL: s.formatThumbnailURL(event.ImageID),
			Theme:        event.Theme,
			Date:         event.EventDate,
		})
	}

	return c.JSON(http.StatusOK, GetEventsResponse{Events: responseTeasers})
}
