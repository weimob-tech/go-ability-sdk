package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomOptionDeleteList
// @id 2130
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2130?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=批量删除自定义字段选项信息)
func (client *Client) CustomOptionDeleteList(request *CustomOptionDeleteListRequest) (response *CustomOptionDeleteListResponse, err error) {
	rpcResponse := CreateCustomOptionDeleteListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomOptionDeleteListRequest struct {
	*api.BaseRequest

	BasicInfo        CustomOptionDeleteListRequestBasicInfo          `json:"basicInfo,omitempty"`
	CustomOptionList []CustomOptionDeleteListRequestCustomOptionList `json:"customOptionList,omitempty"`
	AttributeId      string                                          `json:"attributeId,omitempty"`
}

type CustomOptionDeleteListResponse struct {
	IsSuccess bool `json:"isSuccess,omitempty"`
}

func CreateCustomOptionDeleteListRequest() (request *CustomOptionDeleteListRequest) {
	request = &CustomOptionDeleteListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "custom/option/deleteList", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomOptionDeleteListResponse() (response *api.BaseResponse[CustomOptionDeleteListResponse]) {
	response = api.CreateResponse[CustomOptionDeleteListResponse](&CustomOptionDeleteListResponse{})
	return
}

type CustomOptionDeleteListRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type CustomOptionDeleteListRequestCustomOptionList struct {
	Id int64 `json:"id,omitempty"`
}
