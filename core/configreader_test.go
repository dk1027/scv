package scv

import (
	"github.com/dk1027/scv/testtools"
	"testing"
)

var testConfig = []byte(`[
  {
    "trigger": {"type": "watchdir", "param": "trigger param"},
    "action": {"type": "build", "param": "action param"}
  }
]`)

func TestReadConfig(t *testing.T) {
	config, err := ReadConfig(testConfig)
	testtools.AssertEqual(t, 1, len(config))
	testtools.AssertTrue(t, err == nil)
	testtools.AssertEqual(t, config[0].Trigger.Type, "watchdir")
	testtools.AssertEqual(t, config[0].Trigger.Param, "trigger param")
	testtools.AssertEqual(t, config[0].Action.Type, "build")
	testtools.AssertEqual(t, config[0].Action.Param, "action param")
}

func TestUnmarshalError(t *testing.T) {
	_, err := ReadConfig([]byte(`[{
    xxx "task": "watch",
    "arg1": "/Users/ltse/go/src/hello",
    "arg2": "go build"
  }]`))
	testtools.AssertTrue(t, err != nil)
}
