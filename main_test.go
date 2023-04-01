package main

import (
	"testing"

	"github.com/Kong/go-pdk/bridge"
	"github.com/Kong/go-pdk/bridge/bridgetest"
	"github.com/Kong/go-pdk/service"
	"github.com/stretchr/testify/assert"
)

func TestPlugin(t *testing.T) {

	service := service.Service{
		PdkBridge: bridge.New(
			bridgetest.Mock(
				t,
				[]bridgetest.MockStep{
					{
						Method: "kong.service.set_upstream",
						Args:   bridge.WrapString("farm_4"),
						Ret:    nil,
					},
				},
			),
		),
	}

	assert.NoError(t, service.SetUpstream("farm_4"))
}
