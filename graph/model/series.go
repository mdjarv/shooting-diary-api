package model

type Series struct {
	ID        string   `json:"id"`
	SessionID string   `json:"sessionId"`
	Session   *Session `json:"session"`
	Shots     []Shot   `json:"shots"`
}
