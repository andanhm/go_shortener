package models

import (
	"fmt"
	"time"
)

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("Mon Jan _2"))
	return []byte(stamp), nil
}

type Response struct {
	Status bool     `json:"status"`
	Error  ErrorObj `json:"error,omitempty"`
	Data   DataObj  `json:"data,omitempty"`
}
type DataObj struct {
	ShortUrl string `json:"shortUrl,omitempty"`
	Message  string `json:"message,omitempty"`
}
type ErrorObj struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
