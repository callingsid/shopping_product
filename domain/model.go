package domain

import "gopkg.in/mgo.v2/bson"

type Item struct {
	_id              bson.ObjectId     `json:"_id" bson:"_id,omitempty"`
	Id                bson.ObjectId     `json:"id" bson:"id,omitempty"`
	Seller            float64       	`json:"seller" bson:"seller"`
	Title             string      		`json:"title" bson:"title"`
	Description       Description 		`json:"description" bson:"description"`
	Pictures          []Picture   `json:"pictures" bson:"pictures"`
	Video             string      `json:"video" bson:"video"`
	Price             float32     `json:"price" bson:"price"`
	AvailableQuantity int         `json:"available_quantity" bson:"availablequantity"`
	SoldQuantity      int         `json:"sold_quantity" bson:"soldquantity"`
	Status            string      `json:"status" bson:"status"`
}

type Description struct {
	PlainText string `json:"plain_text"`
	Html      string `json:"html"`
}

type Picture struct {
	Id  int64  `json:"id"`
	Url string `json:"url"`
}
