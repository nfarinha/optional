package optional_test

import (
	"encoding/json"
	"optional"
	"testing"
)

func TestMarshall(t *testing.T) {
	type nullstruct struct {
		Text optional.Optional[string] `json:"text,omitempty"`
	}

	a := nullstruct{}
	text, _ := json.Marshal(a)
	t.Log(a)
	t.Logf(`%s`, text)
}
