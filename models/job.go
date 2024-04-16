package models

type Job struct {
	Id          int64
	Description string `json:"description"`
	Name        string `json:"name"`
	Cron        string `json:"cron"`
	Enabled     bool   `json:"enabled"`
	Executed    int    `json:"executed"`
	Args        string `json:"args"`
}

type Arg struct {
	arguments string
	file      string
}
