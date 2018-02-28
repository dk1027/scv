package scv

import (
	"github.com/dk1027/scv/testtools"
	"testing"
)

func TestEngine(t *testing.T) {
	// doc := []byte(`[
	//   {
	//     "trigger": {"type": "watchdir", "param": "/Users/ltse/go/src/scv"},
	//     "action": {"type": "build", "param": "echo 'hello world'"}
	//   }
	// ]`)
	// taskConfig, _ := ReadConfig(doc)
	// NewEngine(taskConfig)
	testtools.AssertTrue(t, true)
}
