package domain

import (
	"errors"
	"fmt"
	"github.com/callingsid/shopping_utils/db"
	"github.com/callingsid/shopping_utils/logger"
	"github.com/callingsid/shopping_utils/rest_errors"
	"github.com/mitchellh/mapstructure"
	"strings"
)

const (
	collection = "items"
)

func (i *Item) Save()  rest_errors.RestErr {
	err := db.Client.Create(collection, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}
	return nil
}

func (i *Item) Get()  rest_errors.RestErr {
	itemId := i.Id
	data, err := db.Client.Get(collection, itemId)
	logger.Info(fmt.Sprintf("the data in dao is %s", data))
	if err != nil {
		logger.Info(fmt.Sprintf("the err in dao is %s", err))
		if strings.Contains(err.Error(), "not found") {
			return rest_errors.NewNotFoundError(fmt.Sprintf("no item found with id %s", i.Id))
		}
		return rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", i.Id), errors.New("database error"))
	}
	err2 := mapstructure.Decode(data, i)
	if err2 != nil {
		logger.Error("error in converting map to item object ", err2)
		panic(err2)
	}
	return nil
}