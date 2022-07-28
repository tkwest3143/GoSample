package model

import (
	"encoding/xml"
	"io"
	"net/http"
	"strike/go/db"
	"time"
)

type XmlData struct {
	Channel struct {
		Item []struct {
			Title       string `xml:"title"`
			Link        string `xml:"link"`
			Date        string `xml:"pubDate"`
			Image       string `xml:"image"`
			Description string `xml:"description"`
		} `xml:"item"`
	} `xml:"channel"`
}
type NewsData struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	PubDate     string `json:"pubDate"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

type NewsCategory struct {
	ID       string `json:"category_id"`
	Title    string `json:"title"`
	FileName string `json:"file_name"`
}

func GetAllNewsCategoryData() ([]NewsCategory, error) {
	newsData := db.NewsMasterSelectAll()
	NewsCategoryList := []NewsCategory{}
	for _, item := range newsData {
		NewsCategoryList = append(NewsCategoryList, NewsCategory{Title: item.Title, ID: item.ID, FileName: item.FileName})
	}
	return NewsCategoryList, nil
}

func GetNewsListByUser(userId string) ([]NewsData, error) {
	category := db.NewsMasterSelectById(userId)
	return readXmlFromHttp("https://news.yahoo.co.jp/rss/categories/" + category.FileName)
}

func GetNewsListByCategory(categoryId string) ([]NewsData, error) {
	category := db.NewsMasterSelectById(categoryId)
	return readXmlFromHttp("https://news.yahoo.co.jp/rss/categories/" + category.FileName)
}

func readXmlFromHttp(url string) ([]NewsData, error) {
	data, err := getXmlForHttp(url)
	if err != nil {
		return []NewsData{}, err
	}
	xmlData := XmlData{}
	xml.Unmarshal([]byte(data), &xmlData)

	newsDataList := []NewsData{}
	for _, item := range xmlData.Channel.Item {

		datetime, _ := time.Parse(time.RFC1123, item.Date)

		newsDataList = append(newsDataList, NewsData{Title: item.Title, Link: item.Link, PubDate: datetime.Format("2006/01/02 15:04:05"), Image: item.Image, Description: item.Description})
	}
	return newsDataList, nil

}
func getXmlForHttp(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
