package handlers

import (
	"fmt"
	"net/http"
)

func (str *Storage) BackupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Write([]byte("ERROR"))
		println("[BACKUP]: It's not 'GET' method")
		return
	}
	err := str.StorageH.BackupCreate()
	if err != nil {
		w.Write([]byte("ERROR"))
		fmt.Printf("[BACKUP]: Backup processing error - %v", err)
		return
	}
	w.Write([]byte("OK"))
}
