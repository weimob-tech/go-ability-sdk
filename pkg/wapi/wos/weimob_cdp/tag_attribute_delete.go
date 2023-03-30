package weimob_cdp

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// TagAttributeDelete
// @id 1712
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1712?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=删除标签属性)
func (client *Client) TagAttributeDelete(request *TagAttributeDeleteRequest) (response *TagAttributeDeleteResponse, err error) {
	rpcResponse := CreateTagAttributeDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type TagAttributeDeleteRequest struct {
	*api.BaseRequest

	TagId           int64   `json:"tagId,omitempty"`
	AttrIdList      []int64 `json:"attrIdList,omitempty"`
	VidType         int64   `json:"vidType,omitempty"`
	Vid             int64   `json:"vid,omitempty"`
	OperationSource int64   `json:"operationSource,omitempty"`
}

type TagAttributeDeleteResponse struct {
	IsSuccess bool `json:"isSuccess,omitempty"`
}

func CreateTagAttributeDeleteRequest() (request *TagAttributeDeleteRequest) {
	request = &TagAttributeDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_cdp", "v2.0", "tag/attribute/delete", "apigw")
	request.Method = api.POST
	return
}

func CreateTagAttributeDeleteResponse() (response *api.BaseResponse[TagAttributeDeleteResponse]) {
	response = api.CreateResponse[TagAttributeDeleteResponse](&TagAttributeDeleteResponse{})
	return
}
