package core

import (
	"reflect"
	"testing"
)

func TestGetConfig(t *testing.T) {
	File = "../../config.json"
	config, err := GetConfig()
	v := reflect.ValueOf(config)
	var fail bool
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).String() == "" {
			fail = false
		}
	}
	if fail {
		t.Error("empty config field!")
	}
	if err != nil {
		t.Errorf("GetConfig error: %s", err.Error())
	}
}
