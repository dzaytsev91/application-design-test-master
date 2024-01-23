package storage

import (
	"context"
	"errors"
	"time"
)

// Storage interface
type Storage interface {
	GetAvailability(ctx context.Context, hotelID uint64, RoomTypeID uint8, from time.Time, to time.Time) error
}

var (
	// ErrNotAvailableHotel if no hotel in system
	ErrNotAvailableHotel = errors.New("not available hotel")
	// ErrNotAvailableRoomType if no room type in system
	ErrNotAvailableRoomType = errors.New("not available room type")
	// ErrNotAvailableDatesToBook not available dates to book
	ErrNotAvailableDatesToBook = errors.New("not available dates to book")
	// ErrNotAvailableFreeRooms no free rooms for given dates
	ErrNotAvailableFreeRooms = errors.New("no free rooms for given dates")
)
