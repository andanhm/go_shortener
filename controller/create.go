package controller

import (
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/andanhm/go_shortener/handler"
	"github.com/andanhm/go_shortener/models"
	"github.com/andanhm/go_shortener/utilities"
)

func getShortUrl() string {
	return "http://bms.co/" + utilities.Hash()
}
func CreateURL(connection *mgo.Collection, longUrl string) (models.Response, int) {
	if longUrl == "" {
		response := models.Response{
			Status: false,
			Error:  models.ErrorObj{Code: 1001, Message: "Required parameter not provided"},
			Data:   models.DataObj{},
		}
		return response, http.StatusBadRequest
	}
	var shortUrl = getShortUrl()
	err, urlInfo := handler.CheckLongUrl(connection, longUrl)
	if err != nil {
		response := models.Response{
			Status: false,
			Error:  models.ErrorObj{Code: 1002, Message: "Unable to create short url created"},
			Data:   models.DataObj{},
		}
		return response, http.StatusInternalServerError
	}
	if urlInfo.LongUrl != "" {
		response := models.Response{
			Status: true,
			Error:  models.ErrorObj{},
			Data:   models.DataObj{ShortUrl: urlInfo.ShortUrl, Message: "Successfully short url created"},
		}
		return response, http.StatusCreated
	}
	err, urlInfo = handler.CheckShortUrl(connection, shortUrl)
	if err != nil {
		response := models.Response{
			Status: false,
			Error:  models.ErrorObj{Code: 1003, Message: "Unable to create short url created"},
			Data:   models.DataObj{},
		}
		return response, http.StatusInternalServerError
	}
	if urlInfo.ShortUrl != "" {
		shortUrl = getShortUrl()
	}

	err = handler.InsertUrlInfo(connection, longUrl, shortUrl)
	if err != nil {
		response := models.Response{
			Status: false,
			Error:  models.ErrorObj{Code: 1004, Message: "Unable to create short url created"},
			Data:   models.DataObj{},
		}
		return response, http.StatusInternalServerError
	}
	response := models.Response{
		Status: true,
		Error:  models.ErrorObj{},
		Data:   models.DataObj{ShortUrl: shortUrl, Message: "Successfully short url created"},
	}
	return response, http.StatusCreated
}
