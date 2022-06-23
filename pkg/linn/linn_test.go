package linn

import (
	"quaysports.com/server/pkg/core"
	"testing"
)

func TestInit(t *testing.T) {
	core.File = "../../config.json"
	config, err := core.GetConfig()
	if err != nil {
		t.Errorf("Core config fail: %s", err.Error())
	}

	var done = make(chan InitResult)
	go Init(config, done)

	result := <-done
	if result.Err != nil {
		t.Errorf("Linn init result error: %s", result.Err.Error())
	}

	t.Log(result.authData.Server + " - " + result.authData.Token)
	if result.authData.Server == "" || result.authData.Token == "" {
		t.Error("Missing Linn.authProperties")
	}
}
