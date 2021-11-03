package AniAPI

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client struct {
	client *http.Client
}

func NewClient() (*Client, error) {

	return &Client{
		client: http.DefaultClient,
	}, nil
}

func (c Client) GetAnimeList() ([]Anime, error) {
	resp, err := c.client.Get("https://api.aniapi.com/v1/anime")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r animeListResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return r.AnimeList.Data, nil
}

func (c Client) UpdateUser() (User, error) {
	values := map[string]interface{}{"id": 637, "gender": 1, "localization": "en"}

	json_data, err := json.Marshal(values)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.aniapi.com/v1/user", bytes.NewBuffer(json_data))

	req.Header.Add("Authorization", `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjYzNyIsIm5iZiI6MTYzNTk1MjcyNywiZXhwIjoxNjM4NTQ0NzI3LCJpYXQiOjE2MzU5NTI3Mjd9.g1CGLqWYkQqhrF_bMkNWd2smkbEnB5PNZkOXms3F7jc`)
	req.Header.Add("Content-Type", `application/json`)
	req.Header.Add("Accept", `application/json`)

	resp, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r userResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return r.User, nil
}
