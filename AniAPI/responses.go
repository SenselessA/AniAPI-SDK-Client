package AniAPI

import "time"

type Anime struct {
	AnilistId       int               `json:"anilist_id"`
	MalId           int               `json:"mal_id"`
	Format          int               `json:"format"`
	Status          int               `json:"status"`
	Titles          map[string]string `json:"titles"`
	Descriptions    map[string]string `json:"descriptions"`
	StartDate       time.Time         `json:"start_date"`
	EndTime         time.Time         `json:"end_date"`
	SeasonPeriod    int               `json:"season_period"`
	SeasonYear      int               `json:"season_year"`
	EpisodesCount   int               `json:"episodes_count"`
	EpisodeDuration int               `json:"episode_duration"`
	CoverImage      string            `json:"cover_image"`
	CoverColor      string            `json:"cover_color"`
	BannerImage     string            `json:"banner_image"`
	Genres          []string          `json:"genres"`
	Score           int               `json:"score"`
	Id              int               `json:"id"`
}

type animeListData struct {
	CurrentPage int     `json:"current_page"`
	Count       int     `json:"count"`
	Data        []Anime `json:"documents"`
}

type animeListResponse struct {
	AnimeList animeListData `json:"data"`
	Timestamp int           `json:"timestamp"`
	// скорее всего TimeStamp там нет
}

type User struct {
	Username string `json:"username"`
	Role     int    `json:"role"`
	Gender   int    `json:"gender"`
	Avatar   string `json:"avatar"`
	Id       int    `json:"id"`
}

type userResponse struct {
	User User `json:"data"`
}
