// Package ja  实现了一个简单的 json 的 DOM 树构造工具.
package ja

import (
	"encoding/json"
)

const (
	NULL   = 0
	STRING = 1
	BOOL   = 2
	NUMBER = 3
	MAP    = 4
	ARRAY  = 5
)

type Anchor struct {
	Ref interface{}
}

func New(data []byte) (Anchor, error) {
	var v interface{}
	err := json.Unmarshal(data, &v)
	if nil != err {
		return Anchor{nil}, err
	}

	return Anchor{v}, nil
}

func (c Anchor) Unmarshal(obj interface{}) error {
	bytes, err := json.Marshal(c.Ref)
	if nil != err {
		return err
	}

	return json.Unmarshal(bytes, obj)
}

func (c Anchor) Marshal() ([]byte, error) {
	return json.Marshal(c.Ref)
}

func (c Anchor) Type() int {
	switch c.Ref.(type) {
	case json.Number:
		return NUMBER
	case string:
		return STRING
	case bool:
		return BOOL
	case map[string]interface{}:
		return MAP
	case []interface{}:
		return ARRAY
	default:
		return NULL
	}
}

func (c Anchor) String(def string) string {
	if nil == c.Ref {
		return def
	}

	v, ok := c.Ref.(string)
	if !ok {
		return def
	}
	return v
}

func (c Anchor) Bool(def bool) bool {
	if nil == c.Ref {
		return def
	}

	v, ok := c.Ref.(bool)
	if !ok {
		return def
	}
	return v
}

func (c Anchor) Int(def int64) int64 {
	if nil == c.Ref {
		return def
	}

	v, ok := c.Ref.(json.Number)
	if !ok {
		return def
	}

	t, err := v.Int64()
	if nil != err {
		return def
	}

	return t
}

func (c Anchor) Float(def float64) float64 {
	if nil == c.Ref {
		return def
	}

	v, ok := c.Ref.(json.Number)
	if !ok {
		return 0
	}

	t, err := v.Float64()
	if nil != err {
		return 0
	}

	return t
}

func (c Anchor) Map() map[string]interface{} {
	v, ok := c.Ref.(map[string]interface{})
	if !ok {
		return nil
	}

	return v
}

func (c Anchor) Array() []interface{} {
	v, ok := c.Ref.([]interface{})
	if !ok {
		return nil
	}

	return v
}

func (c Anchor) Quote(key string) Anchor {
	m, ok := c.Ref.(map[string]interface{})
	if !ok {
		return Anchor{nil} //	如果无法转换成 map
	}

	v, exist := m[key]
	if !exist {
		return Anchor{nil}
	}

	return Anchor{v}
}

func (c Anchor) Index(index int) Anchor {
	m, ok := c.Ref.([]interface{})
	if !ok {
		return Anchor{nil} //	如果无法转换成 Array
	}

	if index < 0 || len(m) <= index {
		c.Ref = nil
		return Anchor{nil}
	}

	return Anchor{m[index]}
}
