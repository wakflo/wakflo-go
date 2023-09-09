package main

import (
	"github.com/wakflo/wakflo-go/sdk"
	"testing"
)

func TestPluginFn(t *testing.T) {
	expectedResult := "Hello, World!"
	result := PluginFn(sdk.TaskContext{})

	if result.Output["message"] == "Hello, World!" {
		t.Errorf("got %s, want %s", result.Output["message"], expectedResult)
	}
}
