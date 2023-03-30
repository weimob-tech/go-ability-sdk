package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomAttributeDelete
// @id 2135
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2135?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=删除自定义字段)
func (client *Client) CustomAttributeDelete(request *CustomAttributeDeleteRequest) (response *CustomAttributeDeleteResponse, err error) {
	rpcResponse := CreateCustomAttributeDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomAttributeDeleteRequest struct {
	*api.BaseRequest

	BasicInfo CustomAttributeDeleteRequestBasicInfo `json:"basicInfo,omitempty"`
	Id        int64                                 `json:"id,omitempty"`
}

type CustomAttributeDeleteResponse struct {
}

func CreateCustomAttributeDeleteRequest() (request *CustomAttributeDeleteRequest) {
	request = &CustomAttributeDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "custom/attribute/delete", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomAttributeDeleteResponse() (response *api.BaseResponse[CustomAttributeDeleteResponse]) {
	response = api.CreateResponse[CustomAttributeDeleteResponse](&CustomAttributeDeleteResponse{})
	return
}

type CustomAttributeDeleteRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
