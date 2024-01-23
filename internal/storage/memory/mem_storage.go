package memory

import (
	"context"
	"sync"
	"time"

	"application-design-test-master/internal/models"
	utils "application-design-test-master/internal/utils"
)

// Cluster in memory cluster storage
type Cluster interface {
	GetAvailability(ctx context.Context, hotelID uint64, RoomTypeID uint8, from time.Time, to time.Time) error
}

// Storage in memory storage
type Storage struct {
	availabilityList models.HotelAvailability
}

// New creates new cluster storage
func New() Cluster {
	return &Storage{
		availabilityList: models.HotelAvailability{
			Mux: &sync.RWMutex{},
			Availability: map[uint64]map[uint8]map[time.Time]*models.RoomAvailability{
				1: {
					2: {
						utils.NewDate(2024, 1, 1): &models.RoomAvailability{
							RoomTypeID: 1,
							Quota:      1,
						},
						utils.NewDate(2024, 1, 2): &models.RoomAvailability{
							RoomTypeID: 1,
							Quota:      1,
						},
						utils.NewDate(2024, 1, 3): &models.RoomAvailability{
							RoomTypeID: 1,
							Quota:      1,
						},
						utils.NewDate(2024, 1, 4): &models.RoomAvailability{
							RoomTypeID: 1,
							Quota:      1,
						},
						utils.NewDate(2024, 1, 5): &models.RoomAvailability{
							RoomTypeID: 1,
							Quota:      0,
						},
					},
				},
			},
		},
	}
}

// GetAvailability get availiable rooms for given dates
func (s *Storage) GetAvailability(_ context.Context, hotelID uint64, roomTypeID uint8, from time.Time, to time.Time) error {
	daysToBook := utils.DaysBetween(from, to)
	unavailableDays := make(map[time.Time]struct{})
	for _, day := range daysToBook {
		unavailableDays[day] = struct{}{}
	}

	err := s.availabilityList.FindFreeRooms(hotelID, roomTypeID, daysToBook)
	if err != nil {
		return err
	}

	return nil
}
