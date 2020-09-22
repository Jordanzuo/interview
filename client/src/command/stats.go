package command

import (
	"encoding/json"
	"fmt"

	"interview.com/cloudcade/chat/client/src/model"
)

type StatsParameter struct {
	Name string
}

type StatsResponseData struct {
	ActiveTime string
}

func Stats(requestID int64, paramList []string) (requestObj *model.RequestObject, callback func(interface{}), err error) {
	if len(paramList) != 2 {
		err = fmt.Errorf("Format: %s PlayerName", Command_stats)
		return
	}

	parameter := &StatsParameter{
		Name: paramList[1],
	}
	requestObj = model.NewRequestObject(requestID, Command_stats, parameter)
	callback = statsCallback
	return
}

func statsCallback(data interface{}) {
	var responseData *StatsResponseData

	bytes, _ := json.Marshal(data)
	err := json.Unmarshal(bytes, &responseData)
	if err != nil {
		fmt.Printf("Unmarshal stats result failed. Error:%s\n", err)
		return
	}

	fmt.Printf("%s\n", responseData.ActiveTime)
}
