package main

import (
	"fmt"
	"log"

	"github.com/SenselessA/AniAPI-SDK-Client/AniAPI"
)

func main() {
	aniApiClient, err := AniAPI.NewClient()

	if err != nil {
		log.Fatal(err)
	}

	animeList, err := aniApiClient.GetAnimeList()
	if err != nil {
		log.Fatal(err)
	}

	for _, anime := range animeList {
		fmt.Println(anime.Titles)
	}

	fmt.Println("==========================================")

	user, err := aniApiClient.UpdateUser()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)
}
