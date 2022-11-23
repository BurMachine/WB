package handlers

import (
	"burmachine/TestSolution/storage"
	"fmt"
	"io"
	"net/http"
)

type Storage struct {
	StorageH *storage.Storage
}

func (str *Storage) SetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Write([]byte("ERROR"))
		println("[SET]: It's not 'POST' method")
		return
	}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("ERROR"))
		fmt.Printf("[SET]: Body data reading error")
		return
	}
	err = str.StorageH.Set(data)
	if err != nil {
		w.Write([]byte("ERROR"))
		fmt.Printf("[SET]: Storage Set() error - %v", err)
		return
	}
	w.Write([]byte("OK"))
}
