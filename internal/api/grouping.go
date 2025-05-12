package api

import "github.com/CaptainFallaway/XDH/internal/models"

type Grouping struct {
	ValidScans   []models.Scan `json:"validScans"`
	InvalidScans []models.Scan `json:"invalidScans"`
}
