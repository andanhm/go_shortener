package controller

import (
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/andanhm/go_shortener/handler"
	"github.com/andanhm/go_shortener/models"
)

type ResponseFetch struct {
	Status bool             `json:"status"`
	Error  ErrorObjFetch    `json:"error,omitempty"`
	Data   []models.UrlInfo `json:"data,omitempty"`
}
type ErrorObjFetch struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func Fetch(connection *mgo.Collection, urlInfo models.UrlInfo) (ResponseFetch, int) {
	var urlInfoData []models.UrlInfo
	if urlInfo.LongUrl != "" {
		err, data := handler.CheckLongUrl(connection, urlInfo.LongUrl)
		if err != nil {
			return ResponseFetch{false, ErrorObjFetch{Code: 2001, Message: "Required parameter not provided "}, urlInfoData}, http.StatusBadRequest
		}
		if data.Id == "" {
			urlInfoData = nil
		} else {
			urlInfoData = []models.UrlInfo{data}
		}
		return ResponseFetch{true, ErrorObjFetch{}, urlInfoData}, http.StatusOK
	} else if urlInfo.ShortUrl != "" {
		err, data := handler.CheckShortUrl(connection, urlInfo.ShortUrl)
		if err != nil {
			return ResponseFetch{false, ErrorObjFetch{Code: 2002, Message: "Required parameter not provided "}, urlInfoData}, http.StatusBadRequest
		}
		if data.Id == "" {
			urlInfoData = nil
		} else {
			urlInfoData = []models.UrlInfo{data}
		}
		return ResponseFetch{true, ErrorObjFetch{}, urlInfoData}, http.StatusOK
	}
	err, data := handler.Fetch(connection)
	if err != nil {
		return ResponseFetch{false, ErrorObjFetch{Code: 2003, Message: "Required parameter not provided "}, data}, http.StatusBadRequest
	}
	return ResponseFetch{true, ErrorObjFetch{}, data}, http.StatusOK
}
