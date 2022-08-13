package response

import (
	"encoding/json"
	jsoniter "github.com/json-iterator/go"
)

type BaseParam struct {
	C  string // 接口
	Cv string // 客户端版本
	Pv string // 协议版本
	T  string // 时间戳
}

type BaseResponse struct {
	Code string      `json:"code"`
	Desc string      `json:"desc"`
	Data interface{} `json:"data"`
}

func (bp BaseResponse) GetJsonBytes() []byte {
	bytes, _ := jsoniter.Marshal(bp)
	return bytes
}

type Response struct {
	data interface{}
}

func Resp() *Response {
	return &Response{}
}

func (r *Response) Json(data interface{}) *Response {
	b, err := json.Marshal(data)
	if err != nil {
		r.data = ""
	}
	r.data = string(b)
	return r
}

func (r *Response) String(data string) *Response {
	r.data = data
	return r
}

func (r *Response) Byte(data []byte) *Response {
	r.data = data
	return r
}

func (r *Response) GetData() interface{} {
	return r.data
}
