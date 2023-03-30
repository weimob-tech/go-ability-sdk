package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MicrobookGetBookList
// @id 1494
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取预约订单)
func (client *Client) MicrobookGetBookList(request *MicrobookGetBookListRequest) (response *MicrobookGetBookListResponse, err error) {
	rpcResponse := CreateMicrobookGetBookListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MicrobookGetBookListRequest struct {
	*api.BaseRequest

	SearchType int    `json:"searchType,omitempty"`
	StartTime  string `json:"startTime,omitempty"`
	EndTime    string `json:"endTime,omitempty"`
	PageIndex  int    `json:"pageIndex,omitempty"`
	PageSize   int    `json:"pageSize,omitempty"`
	Id         int64  `json:"id,omitempty"`
}

type MicrobookGetBookListResponse struct {
}

func CreateMicrobookGetBookListRequest() (request *MicrobookGetBookListRequest) {
	request = &MicrobookGetBookListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "microbook/getBookList", "api")
	request.Method = api.POST
	return
}

func CreateMicrobookGetBookListResponse() (response *api.BaseResponse[MicrobookGetBookListResponse]) {
	response = api.CreateResponse[MicrobookGetBookListResponse](&MicrobookGetBookListResponse{})
	return
}
