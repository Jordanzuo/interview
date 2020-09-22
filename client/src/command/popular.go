package command

import (
	"encoding/json"
	"fmt"

	"interview.com/cloudcade/chat/client/src/model"
)

type PopularParameter struct{}

type PopularResponseData struct {
	MostPopularWord string
}

func Popular(requestID int64, paramList []string) (requestObj *model.RequestObject, callback func(interface{}), err error) {
	parameter := &PopularParameter{}
	requestObj = model.NewRequestObject(requestID, Command_popular, parameter)
	callback = popularCallback
	return
}

func popularCallback(data interface{}) {
	var responseData *PopularResponseData

	bytes, _ := json.Marshal(data)
	err := json.Unmarshal(bytes, &responseData)
	if err != nil {
		fmt.Printf("Unmarshal popular result failed. Error:%s\n", err)
		return
	}

	fmt.Printf("%s\n", responseData.MostPopularWord)
}
