package model

import (
	//"time"
)

type Game struct{
	Name string `json:"name"`
	Description string `json:"description"`
	Score int `json:"score"`
	Studio string `json:"studio"`
	ReleaseDate string `json:"releasedate"`
	ImageCover ImageForGame	`json:image`

}

type ImageForGame struct{
	ImageLink string `json:"imagelink"`
}