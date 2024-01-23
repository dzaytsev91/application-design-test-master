package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"application-design-test-master/internal/models"
	st "application-design-test-master/internal/storage"
)

// CreateOrder create new order for booking rooms
func CreateOrder(w http.ResponseWriter, r *http.Request, storage st.Storage) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()
	var newOrder models.Order
	err := json.NewDecoder(r.Body).Decode(&newOrder)
	if err != nil {
		http.Error(w, "Error while decoding data", http.StatusInternalServerError)
		return
	}

	err = storage.GetAvailability(ctx, newOrder.HotelID, newOrder.RoomTypeID, time.Time(newOrder.From), time.Time(newOrder.To))
	if err != nil {
		if errors.Is(err, st.ErrNotAvailableHotel) || errors.Is(err, st.ErrNotAvailableDatesToBook) || errors.Is(err, st.ErrNotAvailableFreeRooms) || errors.Is(err, st.ErrNotAvailableRoomType) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err != nil {
		http.Error(w, "Error while encoding data", http.StatusInternalServerError)
		return
	}
}
