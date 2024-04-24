package utils

import (
	"daltondiaz/async-jobs/models"
	"encoding/json"
)


func MarshalJobArgs(args models.Args) (string, error) {
    json, err := json.Marshal(args)
    if err != nil {
        return "", err
    }
    return string(json), nil
}

func UnmarshalJobArgs(argStr string) (models.Args, error) {
    var args models.Args
    err := json.Unmarshal([]byte(argStr), &args)
    if err != nil {
        return args, err
    }
    return args, nil
}
