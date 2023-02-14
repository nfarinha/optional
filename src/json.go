package optional

import (
	"encoding/json"
)

func (o Optional[T]) MarshalJSON() ([]byte, error) {
	if o.defined {
		return json.Marshal(o.value)
	}
	return json.Marshal(nil)
}

func (o *Optional[T]) UnmarshalJSON(data []byte) error {
	var value T

	o.defined = false
	if string(data) == "null" {
		return nil
	}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	o.defined = true
	o.value = value
	return nil
}
