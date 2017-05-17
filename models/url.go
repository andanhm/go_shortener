package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)
// UrlInfo holds options for mongodb collection tblShortURL
type UrlInfo struct {
	// MongoDB _id
	Id       bson.ObjectId `json:"id,omitempty" bson:"_id"`
	// Sort URL of the created for the provided long url by the client
	LongUrl  string        `json:"longUrl,omitempty"`
	// Sort URL of the created for the provided long url by the client
	ShortUrl string        `json:"shortUrl,omitempty"`
	// RequestTimeStamp is the time at which collection created
	RequestTimeStamp     time.Time     `json:"requestTimeStamp,omitempty"`
}
