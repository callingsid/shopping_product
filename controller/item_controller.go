package controller

import (
	"encoding/json"
	"fmt"
	"github.com/callingsid/shopping_product/domain"
	"github.com/callingsid/shopping_product/service"
	domain2 "github.com/callingsid/shopping_utils/domain"
	"github.com/callingsid/shopping_utils/logger"
	"github.com/callingsid/shopping_utils/queue"
	"github.com/callingsid/shopping_utils/rest_errors"
	"gopkg.in/mgo.v2/bson"
)

const (
	topic = "items"
	topic_ops_res = "shop2.operation.res"
)


type Result struct {
	Value  map[string]interface{}
	Err rest_errors.RestErr
}

func HandleProductRequest(request domain2.Request)  {
	logger.Info(fmt.Sprintf("received request On Items channel, ready for processing : %s", request	))
	go process(request)
}

func process(request domain2.Request) {
	switch request.Method {
	case "POST":
		{
			var item domain.Item
			err := json.Unmarshal([]byte(request.Data), &item)
			if err != nil {
				logger.Error("error when trying to unmarshal item POST request", err)
				//return nil, rest_errors.NewInternalServerError("error when trying to unmarshal item POST request", errors.New("json error"))
			}
			//Save the data to mongodb
			item.Id = bson.NewObjectId()
			result, err3 := service.ItemsService.Create(item)
			if err3 != nil {
				//return nil, err3
			}
			logger.Info(fmt.Sprintf("The value of result returned in POST of items is %s", result))

			//return result data. Attach the UID to the data.
			kdata := make(map[string]interface{})
			kdata["data"] = result
			kdata["uid"] = request.UID
			kdata["topic"] = "shop2.operation.res"
			go postResponse(kdata)
		}
	case "GET":
		{
			logger.Info(fmt.Sprintf("Processing get request"))
			var itemID bson.ObjectId
			err := json.Unmarshal([]byte(request.Data), &itemID)
			if err != nil {
				logger.Error("error when trying to unmarshal item GET request", err)
				panic(err)
			}
			//Get the data from mongodb
			logger.Info(fmt.Sprintf("The value of itemID trying to fect is %s", itemID))
			kdata := make(map[string]interface{})
			result, err := service.ItemsService.Get(itemID)
			if err != nil {
				kdata["data"] = err
			} else {
				kdata["data"] = result
			}
			kdata["uid"] = request.UID
			kdata["topic"] = "shop2.operation.res"
			logger.Info(fmt.Sprintf("The value of result returned in GET of items is %s", result))
			go postResponse(kdata)
		}
	}
}

func postResponse(data interface{}) {
	logger.Info(fmt.Sprintf("value of data in reponse is %s", data))
	if err :=  queue.PClient.Publish(topic_ops_res, data); err != nil {
		logger.Error("Kafka error: %s\n", err)
		logger.Error("Panic cant post response to response queue", err)
		panic(err)
	}
}

