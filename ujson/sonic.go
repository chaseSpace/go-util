package ujson

import (
	"github.com/bytedance/sonic"
)

// MustJSON 将对象序列化为JSON字符串，如果出错则panic
func MustJSON(v interface{}) string {
	data, err := sonic.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(data)
}

// MustUnmarshal 将JSON字符串反序列化为对象，如果出错则panic
func MustUnmarshal(data string, v interface{}) {
	err := sonic.Unmarshal([]byte(data), v)
	if err != nil {
		panic(err)
	}
}

// MustUnmarshalBytes 将JSON字节切片反序列化为对象，如果出错则panic
func MustUnmarshalBytes(data []byte, v interface{}) {
	err := sonic.Unmarshal(data, v)
	if err != nil {
		panic(err)
	}
}
