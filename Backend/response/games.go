package response

import (
	"time"
)

type Game struct{
	Name string
	Description string
	Image string // Place holder for now
	Score int
	Studio string
	ReleaseDate time.Time
	

}