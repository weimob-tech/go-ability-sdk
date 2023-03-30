package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MbpQueryCardTemplate
// @id 1779
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询会员卡模板信息)
func (client *Client) MbpQueryCardTemplate(request *MbpQueryCardTemplateRequest) (response *MbpQueryCardTemplateResponse, err error) {
	rpcResponse := CreateMbpQueryCardTemplateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MbpQueryCardTemplateRequest struct {
	*api.BaseRequest

	CardTemplateId int64 `json:"cardTemplateId,omitempty"`
	PID            int64 `json:"PID,omitempty"`
}

type MbpQueryCardTemplateResponse struct {
}

func CreateMbpQueryCardTemplateRequest() (request *MbpQueryCardTemplateRequest) {
	request = &MbpQueryCardTemplateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "mbp/queryCardTemplate", "api")
	request.Method = api.POST
	return
}

func CreateMbpQueryCardTemplateResponse() (response *api.BaseResponse[MbpQueryCardTemplateResponse]) {
	response = api.CreateResponse[MbpQueryCardTemplateResponse](&MbpQueryCardTemplateResponse{})
	return
}
