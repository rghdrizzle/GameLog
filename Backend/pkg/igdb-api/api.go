package igdbapi

import (
	"io"
	"log"
	"net/http"
)

type IGDB_Client struct{
	HttpClient *http.Client
	AuthToken string
}


func NewClient(client_id string, client_secret string) *IGDB_Client{
	client := http.Client{}
	url:= "https://id.twitch.tv/oauth2/token?client_id="+client_id+"&client_secret="+client_secret+"&grant_type=client_credentials"
	buf := io.Reader
	resp, err:= client.Post(url,"Application/json",buf)
	if err!=nil{
		log.Fatal(err)
	}
	return &IGDB_Client{ HttpClient: &client}
}