package linn

import (
	"quaysports.com/server/pkg/core"
	"testing"
)

func TestAuth(t *testing.T) {
	core.File = "../../config.json"
	config, err := core.GetConfig()
	if err != nil {
		t.Errorf("Core config fail: %s", err.Error())
	}
	err = Auth(config)
	if err != nil {
		t.Errorf("Linn Auth error: %s", err.Error())
	}
}
