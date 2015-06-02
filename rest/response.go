package rest

import (
    "encoding/json"
    "net/http"
)

func Response(w http.ResponseWriter, statuscode int, payload interface{}) error {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(statuscode)
    if err := json.NewEncoder(w).Encode(payload); err != nil {
        return err
    }

    return nil
}
