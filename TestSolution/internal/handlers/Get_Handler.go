package handlers

import (
	"fmt"
	"net/http"
)

func (str *Storage) GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		println("[GET]: It's not 'GET' method")
		return
	}
	userId := r.URL.Query().Get("user_id")
	if userId == "" {
		println("[GET]: user_id does not exist in query")
		return
	}
	data, err := str.StorageH.Get(userId)
	if err != nil {
		fmt.Printf("[GET]: storage Get() error - %v", err)
		return
	}
	w.Write(data)
}
