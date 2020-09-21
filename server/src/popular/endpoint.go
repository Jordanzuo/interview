package popular

import (
	"interview.com/cloudcade/chat/server/src/clientmgr"
	"interview.com/cloudcade/chat/server/src/model"
	playerModel "interview.com/cloudcade/chat/server/src/player/model"
)

func init() {
	clientmgr.RegisterFunction("/popular", popularFunc)
}

type PopularParameter struct{}

func (this *PopularParameter) verify() model.ResponseStatus {
	return model.Success
}

type PopularResponseData struct {
	MostPopularWord string
}

func newPopularResponseData(mostPopularWord string) *PopularResponseData {
	return &PopularResponseData{
		MostPopularWord: mostPopularWord,
	}
}

func popularFunc(requestObj *model.RequestObject, clientObj clientmgr.IClient, playerObj *playerModel.Player) *model.ResponseObject {
	var responseObj = model.NewResponseObject()
	var paramObj = new(PopularParameter)
	var rs model.ResponseStatus

	rs = requestObj.ParseParameter(&paramObj)
	if rs != model.Success {
		return responseObj.SetResponseStatus(rs)
	}

	rs = paramObj.verify()
	if rs != model.Success {
		return responseObj.SetResponseStatus(rs)
	}

	popularResponseDataObj := newPopularResponseData(mostPopularWord)
	responseObj.SetData(popularResponseDataObj)

	return responseObj
}
