package handler

import (
	"encoding/json"
	"net/http"
)

func RemoveSubscriptionHandler(w http.ResponseWriter, r *http.Request) {

	// Decode JSON

	var e Event
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Remove subscription from database

	RemoveSubscription(&e)
	w.WriteHeader(http.StatusOK)
	printDB()

}
