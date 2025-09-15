package notion

import (
	"encoding/json"

	"github.com/spf13/cast"
)

type Properties map[string]Property

func (m Properties) Exists(name string) bool {
	_, ok := m[name]
	return ok
}

func (m Properties) Set(name string, value any) {
	p, ok := m[name]
	if !ok {
		return
	}
	p.SetValue(value)
}

func (m Properties) Get(name string) any {
	p, ok := m[name]
	if !ok {
		return nil
	}
	return p.GetValue()
}

func (m Properties) GetString(name string) string {
	return cast.ToString(m.Get(name))
}

func (m Properties) Raw() map[string]any {
	result := map[string]any{}
	for k, v := range m {
		result[k] = v.GetValue()
	}
	return result
}

func (m *Properties) UnmarshalJSON(data []byte) error {
	temp := map[string]Property{}

	type dummyProp struct {
		Type string `json:"type"`
	}

	rawMsgs := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &rawMsgs); err != nil {
		return err
	}

	for name, raw := range rawMsgs {
		var dummy dummyProp
		if err := json.Unmarshal(raw, &dummy); err != nil {
			return err
		}
		prop := NewProperty(dummy.Type)
		if prop == nil {
			continue
		}
		if err := json.Unmarshal(raw, prop); err != nil {
			return err
		}
		temp[name] = prop
	}
	*m = temp
	return nil
}

func (m Properties) String() string {
	jsonBytes, _ := json.Marshal(m)
	return string(jsonBytes)
}

type PropertyConfigs map[string]PropertyConfig

func (m PropertyConfigs) Exists(name string) bool {
	_, ok := m[name]
	return ok
}

func (m *PropertyConfigs) UnmarshalJSON(data []byte) error {
	temp := map[string]PropertyConfig{}

	type dummyPropConfig struct {
		Type string `json:"type"`
	}

	rawMsgs := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &rawMsgs); err != nil {
		return err
	}

	for name, raw := range rawMsgs {
		var dummy dummyPropConfig
		if err := json.Unmarshal(raw, &dummy); err != nil {
			return err
		}
		prop := NewPropertyConfig(dummy.Type)
		if prop == nil {
			continue
		}
		if err := json.Unmarshal(raw, prop); err != nil {
			return err
		}
		temp[name] = prop
	}
	*m = temp
	return nil
}

func (m PropertyConfigs) String() string {
	jsonBytes, _ := json.Marshal(m)
	return string(jsonBytes)
}
