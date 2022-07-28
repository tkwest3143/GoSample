package main

import (
	"strike/go/data"
	"strike/go/db"
)

func main() {
	NEWS_CATEGORY_ITEMS := [...]data.NewsMaster{
		{
			Title:    "domestic",
			FileName: "domestic.xml",
		},
		{
			Title:    "world",
			FileName: "world.xml",
		}, {
			Title:    "business",
			FileName: "business.xml",
		},
		{
			Title:    "entertainment",
			FileName: "entertainment.xml",
		},
		{
			Title:    "sports",
			FileName: "sports.xml",
		},
		{
			Title:    "it",
			FileName: "it.xml",
		},
		{
			Title:    "science",
			FileName: "science.xml",
		},
		{
			Title:    "life",
			FileName: "life.xml",
		},
		{
			Title:    "local",
			FileName: "local.xml",
		},
	}
	for _, item := range NEWS_CATEGORY_ITEMS {
		db.NewsMasterInsert((item))
	}
}
