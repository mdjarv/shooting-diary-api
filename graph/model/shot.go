package model

type Shot struct {
	ID       string  `json:"id"`
	Score    int     `json:"score"`
	Inner    bool    `json:"inner" gorm:"not null;default:false"`
	SeriesID string  `json:"seriesId"`
	Series   *Series `json:"series"`
}
