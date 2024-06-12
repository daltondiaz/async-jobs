package utils

import (
	"daltondiaz/async-jobs/models"
	"reflect"
	"testing"
)

func TestMarshalJobArgs(t *testing.T) {
	expected := "{\"args\":[\"10\"],\"path\":\"/home/dalton/Dev/personal/async-jobs/test.php\",\"cmd\":\"php\"}"
	var arg models.Args
	arg.Args = []string{"10"}
	arg.Path = "/home/dalton/Dev/personal/async-jobs/test.php"
	arg.Cmd = "php"
	result, _ := MarshalJobArgs(arg)
	if result != expected {
		t.Fatalf("expected:\n%s\ngot:\n%s", expected, result)
	}
}

func TestUnmarshalJobArgs(t *testing.T) {
	json := "{\"args\":[\"10\"],\"path\":\"/home/dalton/Dev/personal/async-jobs/test.php\",\"cmd\":\"php\"}"
	var exp models.Args
	exp.Args = []string{"10"}
	exp.Path = "/home/dalton/Dev/personal/async-jobs/test.php"
	exp.Cmd = "php"
	result, _ := UnmarshalJobArgs(json)
	if !reflect.DeepEqual(result, exp) {
		t.Fatalf("expected:\n%v\ngot:\n%v", exp, result)
	}
}
