package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func handleChirp(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	body := chirp{}
	err := decoder.Decode(&body)
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		respondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	if len(body.Body) > 140 {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long")
		return
	}

	cleanedRes := profanityCheck(body.Body)

	respondWithJSON(w, http.StatusOK, validResponse{
		CleanedResponse: cleanedRes,
	})
}

func profanityCheck(msg string) string {
	split := strings.Split(msg, " ")
	for i, val := range split {
		lower := strings.ToLower(val)
		if lower == "kerfuffle" || lower == "sharbert" || lower == "fornax" {
			split[i] = "****"
		}
	}
	joined := strings.Join(split, " ")
	return joined
}
