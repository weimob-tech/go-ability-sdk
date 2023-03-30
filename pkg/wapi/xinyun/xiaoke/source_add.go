package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// SourceAdd
// @id 2135
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新增客资录入方式)
func (client *Client) SourceAdd(request *SourceAddRequest) (response *SourceAddResponse, err error) {
	rpcResponse := CreateSourceAddResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type SourceAddRequest struct {
	*api.BaseRequest

	Wid   int64  `json:"wid,omitempty"`
	Code  int    `json:"code,omitempty"`
	Stage string `json:"stage,omitempty"`
}

type SourceAddResponse struct {
	IsOpen bool   `json:"isOpen,omitempty"`
	Name   string `json:"name,omitempty"`
	Sort   string `json:"sort,omitempty"`
}

func CreateSourceAddRequest() (request *SourceAddRequest) {
	request = &SourceAddRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "source/add", "api")
	request.Method = api.POST
	return
}

func CreateSourceAddResponse() (response *api.BaseResponse[SourceAddResponse]) {
	response = api.CreateResponse[SourceAddResponse](&SourceAddResponse{})
	return
}
