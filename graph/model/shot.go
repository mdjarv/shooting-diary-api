package model

type Shot struct {
	ID       string  `json:"id"`
	Score    int     `json:"score"`
	SeriesID string  `json:"seriesId"`
	Series   *Series `json:"series"`
}
