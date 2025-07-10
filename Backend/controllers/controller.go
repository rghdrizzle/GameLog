package controllers

import (
	//"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rghdrizzle/gameLog/model"
	igdbapi "rghdrizzle/gameLog/pkg/igdb-api"

	"github.com/Henry-Sarabia/igdb"
	"github.com/gofiber/fiber/v2"
)


func GetRecommendedGames(c *fiber.Ctx) error {
	//Todo recommendation algorithm
	// Todo fetch games from a database
	games := []model.Game{}
	testGame := model.Game{
		Name: "Clair Obscur Expedition 33",
		Description: "Lead the members of Expedition 33 on their quest to destroy the Paintress so that she can never paint death again. Explore a world of wonders inspired by Belle Époque France and battle unique enemies in this turn-based RPG with real-time mechanics.",
		Score: 10,
		ImageCover: model.ImageForGame{ImageLink: "https://upload.wikimedia.org/wikipedia/en/thumb/5/5a/Clair_Obscur%2C_Expedition_33_Cover_1.webp/250px-Clair_Obscur%2C_Expedition_33_Cover_1.webp.png"},
		Studio: "Kepler Studios",
		ReleaseDate: "24-04-2025",

	}
	// This is only for testing so frontend developer can work with creating the frontend
	for i:=0; i<6; i++{
		games = append(games,testGame)
	}
	GetIGDBClient()
	return c.JSON(games)
	
}
func SearchGameDatabaseBasedOnText(searchText string) {

}

func GetIGDBClient()*igdb.Client{
	c := &http.Client{}
	client:= igdb.NewClient("",c)
	igdbapi
	games, err := client.Games.Search("zelda")
	if err!=nil{
		log.Fatal(err)
	}

	for _, game := range games{
		name := game.Name
		fmt.Printf(name)
	}

	return client
}


func Recommender(){

}