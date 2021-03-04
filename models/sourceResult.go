package models

import "time"

type SourceResult struct {
	Type string
}

type TrelloResult struct {
	SourceResult
	Cards []Card
}

type Card struct {
	Name string `json:"name"`
	Id string `json:"id"`
	Labels []Label `json:"labels"`
	DateLastActivity time.Time `json:"dateLastActivity"`
}

type Label struct {
	Name string `json:"name"`
	Id string `json:"id"`
	IdBoard string `json:"IdBoard"`
}