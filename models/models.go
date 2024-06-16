package models

var EXECUTING = 1
var EXECUTED = 0

type Job struct {
	// Identifier of Job
	Id int64
	// Description of Job
	Description string `json:"description"`
	// Name of Job
	Name string `json:"name"`
	// Cron of Job
	Cron string `json:"cron"`
	// If true the Job is enabled, if false disabled
	Enabled bool `json:"enabled"`
	// Flag of control of Job, 0 - Not Executed 1 - Executing
	Executed int  `json:"executed"`
	Args     Args `json:"args"`
	ArgsStr  string
	// Id of cron in memory
	CronId int `json:"cronId"`
}

// Argument struct to build the full command to be executed
type Args struct {
	// Arguments of command and path to be executed ["arg1", "arg2", "argN"]
	Args []string `json:"args"`
	// Path of file with permission to be executed
	Path string `json:"path"`
	// Command to be executed
	Cmd string `json:"cmd"`
}
