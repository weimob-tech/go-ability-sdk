package weimob_cdp

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// TagUpdateList
// @id 1577
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1577?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=批量修改标签)
func (client *Client) TagUpdateList(request *TagUpdateListRequest) (response *TagUpdateListResponse, err error) {
	rpcResponse := CreateTagUpdateListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type TagUpdateListRequest struct {
	*api.BaseRequest

	OpenUpdateTagRequestList []TagUpdateListRequestOpenUpdateTagRequestList `json:"openUpdateTagRequestList,omitempty"`
	VidType                  int64                                          `json:"vidType,omitempty"`
	Vid                      int64                                          `json:"vid,omitempty"`
	OperationSource          int64                                          `json:"operationSource,omitempty"`
}

type TagUpdateListResponse struct {
	ListVo []TagUpdateListResponseListVo `json:"listVo,omitempty"`
}

func CreateTagUpdateListRequest() (request *TagUpdateListRequest) {
	request = &TagUpdateListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_cdp", "v2.0", "tag/updateList", "apigw")
	request.Method = api.POST
	return
}

func CreateTagUpdateListResponse() (response *api.BaseResponse[TagUpdateListResponse]) {
	response = api.CreateResponse[TagUpdateListResponse](&TagUpdateListResponse{})
	return
}

type TagUpdateListRequestOpenUpdateTagRequestList struct {
	AttrInfoList []TagUpdateListRequestAttrInfoList `json:"attrInfoList,omitempty"`
	TagId        int64                              `json:"tagId,omitempty"`
	TagName      string                             `json:"tagName,omitempty"`
	Status       int64                              `json:"status,omitempty"`
	TagType      int64                              `json:"tagType,omitempty"`
	Source       int64                              `json:"source,omitempty"`
}

type TagUpdateListRequestAttrInfoList struct {
	AttrId   int64  `json:"attrId,omitempty"`
	AttrName string `json:"attrName,omitempty"`
}

type TagUpdateListResponseListVo struct {
	IsSuccess bool  `json:"isSuccess,omitempty"`
	TagId     int64 `json:"tagId,omitempty"`
}
