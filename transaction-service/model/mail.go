package model

type Mail struct {
	Topic string      `json:"topic"`
	Key   string      `json:"key"`
	Data  interface{} `json:"data"`
}
