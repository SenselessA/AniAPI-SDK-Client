# Kitsu-SDK-Client

** Usage

** Init
```
  aniApiClient, err := AniAPI.NewClient()

	if err != nil {
		log.Fatal(err)
	}
```

** GET anime list
```
  animeList, err := aniApiClient.GetAnimeList()
	if err != nil {
		log.Fatal(err)
	}

	for _, anime := range animeList {
		fmt.Println(anime.Titles)
	}
```

** POST update user
```
  user, err := aniApiClient.UpdateUser()
	if err != nil {
		log.Fatal(err)
	}

	println(user)
```