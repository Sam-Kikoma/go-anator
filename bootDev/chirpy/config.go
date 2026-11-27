package main

import (
	"sync/atomic"

	"github.com/Sam-Kikoma/go-anator/bootDev/chirpy/internal/database"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	db             *database.Queries
	platform       string
}
