package models

var EXECUTING = 1
var EXECUTED = 0

type Job struct {
	Id          int64
	Description string `json:"description"`
	Name        string `json:"name"`
	Cron        string `json:"cron"`
	Enabled     bool   `json:"enabled"`
	Executed    int    `json:"executed"` // 0 - Not Executed 1 - Executing
	Args        string `json:"args"`
}

type Arg struct {
	arguments string
	file      string
}
