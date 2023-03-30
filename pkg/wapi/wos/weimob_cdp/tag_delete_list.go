package weimob_cdp

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// TagDeleteList
// @id 1576
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1576?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=批量删除标签)
func (client *Client) TagDeleteList(request *TagDeleteListRequest) (response *TagDeleteListResponse, err error) {
	rpcResponse := CreateTagDeleteListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type TagDeleteListRequest struct {
	*api.BaseRequest

	TagIdList       []int64 `json:"tagIdList,omitempty"`
	VidType         int64   `json:"vidType,omitempty"`
	Vid             int64   `json:"vid,omitempty"`
	OperationSource int64   `json:"operationSource,omitempty"`
}

type TagDeleteListResponse struct {
	ListVo []TagDeleteListResponseListVo `json:"listVo,omitempty"`
}

func CreateTagDeleteListRequest() (request *TagDeleteListRequest) {
	request = &TagDeleteListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_cdp", "v2.0", "tag/deleteList", "apigw")
	request.Method = api.POST
	return
}

func CreateTagDeleteListResponse() (response *api.BaseResponse[TagDeleteListResponse]) {
	response = api.CreateResponse[TagDeleteListResponse](&TagDeleteListResponse{})
	return
}

type TagDeleteListResponseListVo struct {
	IsSuccess bool  `json:"isSuccess,omitempty"`
	TagId     int64 `json:"tagId,omitempty"`
}
