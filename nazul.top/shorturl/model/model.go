package model

import "time"

type Mapper struct {
	Id         int64     `json:"id"`
	ShortUrl   string    `json:"shortUrl"`
	Url        string    `json:"url"`
	CreateTime time.Time `json:"createTime"`
}
