package models

import (
	"encoding/json"
	"strings"
	"sync"
	"time"

	"application-design-test-master/internal/storage"
)

// BookingDate type alias
type BookingDate time.Time

// Order order struct
type Order struct {
	HotelID    uint64      `json:"hotel_id"`
	RoomTypeID uint8       `json:"room_type_id"`
	UserEmail  string      `json:"email"`
	From       BookingDate `json:"from"`
	To         BookingDate `json:"to"`
}

// UnmarshalJSON implement Unmarshaler interface
func (j *BookingDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = BookingDate(t)
	return nil
}

// MarshalJSON marshal json
func (j BookingDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}

// RoomAvailability struct
type RoomAvailability struct {
	RoomTypeID uint8  `json:"room_type_id"`
	Quota      uint32 `json:"quota"`
}

// HotelAvailability struct
type HotelAvailability struct {
	Availability map[uint64]map[uint8]map[time.Time]*RoomAvailability
	Mux          *sync.RWMutex
}

// FindFreeRooms find free rooms for given dates or return error
func (a *HotelAvailability) FindFreeRooms(hotelID uint64, roomTypeID uint8, daysToBook []time.Time) error {
	a.Mux.Lock()
	defer a.Mux.Unlock()
	hotelAvails, ok := a.Availability[hotelID]
	if !ok {
		return storage.ErrNotAvailableHotel
	}
	roomsAvail, ok := hotelAvails[roomTypeID]
	if !ok {
		return storage.ErrNotAvailableRoomType
	}
	for _, dayToBook := range daysToBook {
		avail, ok := roomsAvail[dayToBook]
		if !ok { // nolint
			return storage.ErrNotAvailableDatesToBook
		} else if avail.Quota < 1 {
			return storage.ErrNotAvailableFreeRooms
		} else {
			avail.Quota--
		}
	}
	return nil
}
