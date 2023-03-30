package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// NewmembercardGetConsumingLog
// @id 121
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=单个消费明细)
func (client *Client) NewmembercardGetConsumingLog(request *NewmembercardGetConsumingLogRequest) (response *NewmembercardGetConsumingLogResponse, err error) {
	rpcResponse := CreateNewmembercardGetConsumingLogResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type NewmembercardGetConsumingLogRequest struct {
	*api.BaseRequest

	Id int64 `json:"id,omitempty"`
}

type NewmembercardGetConsumingLogResponse struct {
}

func CreateNewmembercardGetConsumingLogRequest() (request *NewmembercardGetConsumingLogRequest) {
	request = &NewmembercardGetConsumingLogRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "newmembercard/GetConsumingLog", "api")
	request.Method = api.POST
	return
}

func CreateNewmembercardGetConsumingLogResponse() (response *api.BaseResponse[NewmembercardGetConsumingLogResponse]) {
	response = api.CreateResponse[NewmembercardGetConsumingLogResponse](&NewmembercardGetConsumingLogResponse{})
	return
}
