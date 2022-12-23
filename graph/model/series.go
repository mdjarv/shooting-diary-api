package model

type Series struct {
	ID        string   `json:"id"`
	SessionID string   `json:"sessionId"`
	Order     int      `json:"order" gorm:"not null;default:0"`
	Session   *Session `json:"session"`
	Shots     []Shot   `json:"shots"`
}
