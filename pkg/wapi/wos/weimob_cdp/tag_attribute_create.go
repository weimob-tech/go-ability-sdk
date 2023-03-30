package weimob_cdp

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// TagAttributeCreate
// @id 1714
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1714?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=创建标签属性)
func (client *Client) TagAttributeCreate(request *TagAttributeCreateRequest) (response *TagAttributeCreateResponse, err error) {
	rpcResponse := CreateTagAttributeCreateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type TagAttributeCreateRequest struct {
	*api.BaseRequest

	TagId           int64  `json:"tagId,omitempty"`
	AttrName        string `json:"attrName,omitempty"`
	VidType         int64  `json:"vidType,omitempty"`
	Vid             int64  `json:"vid,omitempty"`
	OperationSource int64  `json:"operationSource,omitempty"`
}

type TagAttributeCreateResponse struct {
	AttrId int64 `json:"attrId,omitempty"`
}

func CreateTagAttributeCreateRequest() (request *TagAttributeCreateRequest) {
	request = &TagAttributeCreateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_cdp", "v2.0", "tag/attribute/create", "apigw")
	request.Method = api.POST
	return
}

func CreateTagAttributeCreateResponse() (response *api.BaseResponse[TagAttributeCreateResponse]) {
	response = api.CreateResponse[TagAttributeCreateResponse](&TagAttributeCreateResponse{})
	return
}
