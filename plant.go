package models

import "time"

type Plant struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Light     string    `json:"light"`
	Watering  string    `json:"watering"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
