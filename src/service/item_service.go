package service

import (
	"github.com/callingsid/shopping_product/src/domain"
	"github.com/callingsid/shopping_utils/rest_errors"
	"gopkg.in/mgo.v2/bson"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(domain.Item) (*domain.Item, rest_errors.RestErr)
	Get(bson.ObjectId) (*domain.Item, rest_errors.RestErr)

}

type itemsService struct{}

func (s *itemsService) Create(item domain.Item) (*domain.Item, rest_errors.RestErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Get(id bson.ObjectId) (*domain.Item, rest_errors.RestErr) {
	item := domain.Item{Id: id}

	if err := item.Get(); err != nil {
		return nil, err
	}
	return &item, nil
}

