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
	Args        Args   `json:"args"`
	ArgsStr     string 
	CronId      int `json:"cronId"`
}

type Args struct {
	Args []string `json:"args"` // argument
	Path string `json:"path"` // path of file
	Cmd  string `json:"cmd"`  // command to be executed
}
