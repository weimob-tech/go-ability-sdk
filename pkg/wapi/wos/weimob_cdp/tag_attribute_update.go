package weimob_cdp

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// TagAttributeUpdate
// @id 1713
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1713?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=修改标签属性)
func (client *Client) TagAttributeUpdate(request *TagAttributeUpdateRequest) (response *TagAttributeUpdateResponse, err error) {
	rpcResponse := CreateTagAttributeUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type TagAttributeUpdateRequest struct {
	*api.BaseRequest

	AttValueMap     TagAttributeUpdateRequestAttValueMap `json:"attValueMap,omitempty"`
	TagId           int64                                `json:"tagId,omitempty"`
	VidType         int64                                `json:"vidType,omitempty"`
	Vid             int64                                `json:"vid,omitempty"`
	OperationSource int64                                `json:"operationSource,omitempty"`
}

type TagAttributeUpdateResponse struct {
	IsSuccess bool `json:"isSuccess,omitempty"`
}

func CreateTagAttributeUpdateRequest() (request *TagAttributeUpdateRequest) {
	request = &TagAttributeUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_cdp", "v2.0", "tag/attribute/update", "apigw")
	request.Method = api.POST
	return
}

func CreateTagAttributeUpdateResponse() (response *api.BaseResponse[TagAttributeUpdateResponse]) {
	response = api.CreateResponse[TagAttributeUpdateResponse](&TagAttributeUpdateResponse{})
	return
}

type TagAttributeUpdateRequestAttValueMap struct {
}
