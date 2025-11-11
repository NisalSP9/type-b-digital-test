package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"unicode"
)

func CheckName(w http.ResponseWriter, r *http.Request) {

	name := strings.TrimSpace(r.URL.Query().Get("name"))

	log.Printf("read query parameter name : %v", name)

	w.Header().Set("Content-Type", "application/json")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid Input",
		}); err != nil {
			log.Println("error while encoding response")
		}
		return
	}

	nameLowerCase := strings.ToLower(name)

	firstRune := rune(nameLowerCase[0])

	if unicode.IsLetter(firstRune) && firstRune >= 'a' && firstRune <= 'm' {

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(map[string]string{
			"message": "Hello " + name,
		}); err != nil {
			log.Println("error while encoding response")
		}
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid Input",
		}); err != nil {
			log.Println("error while encoding response")
		}
	}

}
