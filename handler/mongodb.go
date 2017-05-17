// Package handler provides primitives for handling the mongo db queries collections.
package handler

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"go_shortener/config"
	"time"
	"fmt"
	"go_shortener/models"
)

var mongoDB *mgo.Database

// Open establishes a new session to the cluster identified by the given seed
// server(s). The session will enable communication with all of the servers in
// the cluster, so the seed servers are used only to find out about the cluster
// topology.
// It returns the mongodb session
func Open() (*mgo.Database, error) {
	if mongoDB == nil {
		session, err := mgo.Dial(config.MONGO_DB_URL)
		if err != nil {
			fmt.Printf("%s\n",err.Error())
			return nil, err
		}
		//defer mongo.Close()

		// Optional. Switch the session to a monotonic behavior.
		session.SetMode(mgo.Monotonic, true)
		session.DB(config.DB_NAME)
		return session.DB(config.DB_NAME), nil
	}
	return mongoDB, nil
}

// Allows to create a index for the date filed in mongodb
// db.getCollection('tblShortUrl').getIndices({}) List all indexes on all collections in a database
func CreateUrlIndex(connection *mgo.Collection) {
	index := mgo.Index{
		Key: []string{"shorturl", "longurl", "requesttimetstamp"},
	}
	err := connection.EnsureIndex(index)
	if err != nil {
		fmt.Printf("%s Unable to create the mongodb index\n",err.Error())
	}
	fmt.Println("tblShortUrl index's created")
}

func InsertUrlInfo(connection *mgo.Collection, longUrl string, shortUrl string) error {
	CreateUrlIndex(connection)
	url := models.UrlInfo{
		Id:               bson.NewObjectId(),
		LongUrl:          longUrl,
		ShortUrl:         shortUrl,
		RequestTimeStamp: time.Now(),
	}

	err := connection.Insert(url)
	if err != nil {
		return err
	}
	return nil
}
func CheckShortUrl(connection *mgo.Collection, shortUrl string) (error, models.UrlInfo) {
	var urlInfo models.UrlInfo
	err := connection.Find(bson.M{"shorturl": shortUrl}).One(&urlInfo)
	if err != nil && err.Error() != "not found" {
		return err, urlInfo
	}
	return nil, urlInfo
}

func CheckLongUrl(connection *mgo.Collection, longUrl string) (error, models.UrlInfo) {
	var urlInfo models.UrlInfo
	err := connection.Find(bson.M{"longurl": longUrl}).One(&urlInfo)
	if err != nil && err.Error() != "not found" {
		return err, urlInfo
	}
	return nil, urlInfo

}

func Fetch(connection *mgo.Collection) (error, []models.UrlInfo) {
	var urlInfoList []models.UrlInfo
	err := connection.Find(nil).All(&urlInfoList)
	if err != nil && err.Error() != "not found" {
		return err, urlInfoList
	}
	return nil, urlInfoList
}
