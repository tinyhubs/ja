package ja

import (
	"testing"
)

func expect(t *testing.T, msg string, result bool) {
	if result {
		return
	}

	t.Error(msg)
}

func Test_1(t *testing.T) {
	s := `{
		"key1": [1, "string", null, true, false, {"a":"b","c":"d"}],
		"key2": {"m1":"sss", "m2":22, "m3":null, "m4":true, "m5":false}
	}
	`

	c, err := New([]byte(s))
	expect(t, "转换应该成功", nil == err)

	v := c.Quote("key1").Index(5).Quote("c").String("")
	expect(t, "通过链式引用成功", v == "d")
}

func Test_2(t *testing.T) {
	s := `{"key1":"value1","key2":22}`

	c, err := New([]byte(s))
	expect(t, "转换应该成功", nil == err)

	v := c.Map()
	expect(t, "通过链式引用成功", v["key1"].(string) == "d")
	//expect(t, "通过链式引用成功", v["key1"].json.N == )
}
