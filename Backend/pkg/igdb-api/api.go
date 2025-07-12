package igdbapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type IGDB_Client struct{
	HttpClient *http.Client
	AuthToken string
	clientId string
}

//type 

// Creates a IGDB client to communicate with the IGDB API
func NewClient(client_id string, client_secret string) *IGDB_Client{
	client := http.Client{}
	url:= "https://id.twitch.tv/oauth2/token?client_id="+client_id+"&client_secret="+client_secret+"&grant_type=client_credentials"
	buf := strings.NewReader("")
	resp, err:= client.Post(url,"Application/json",buf)
	if err!=nil{
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body ,err:= ioutil.ReadAll(resp.Body)
	if err!=nil{
		log.Fatal(err)
	}
	var tok struct {
		Access_token string `json:"access_token"`
	}

	if err := json.Unmarshal(body, &tok); err != nil {
		log.Fatal(err)
	}

	return &IGDB_Client{ HttpClient: &client, AuthToken: tok.Access_token, clientId: client_id}
}

func (ig *IGDB_Client) Search(name string) []byte {
	query := fmt.Sprintf("fields *; search \"%s\"; limit 50;",name)
	body := strings.NewReader(query)
	req, err := http.NewRequest("POST", "https://api.igdb.com/v4/games", body)
	if err != nil {
		log.Fatal(err)
	}
	token := "Bearer "+ig.AuthToken
	req.Header.Set("Client-ID", ig.clientId)
	req.Header.Set("Authorization", token)
	resp, err := ig.HttpClient.Do(req)
	if err!=nil{
		log.Fatal(err)
	}
	b ,err:= ioutil.ReadAll(resp.Body)
	if err!=nil{
		log.Fatal(err)
	}
	// var tok struct {
	// 	Access_token string `json:"access_token"`
	// }

	// if err := json.Unmarshal(b, &tok); err != nil {
	// 	log.Fatal(err)
	// }
	return b
}