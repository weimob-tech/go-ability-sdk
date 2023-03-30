package weimob_cdp

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// TagInsertList
// @id 1578
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1578?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=批量插入标签)
func (client *Client) TagInsertList(request *TagInsertListRequest) (response *TagInsertListResponse, err error) {
	rpcResponse := CreateTagInsertListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type TagInsertListRequest struct {
	*api.BaseRequest

	OpenCreateTagRequestList []TagInsertListRequestOpenCreateTagRequestList `json:"openCreateTagRequestList,omitempty"`
	VidType                  int64                                          `json:"vidType,omitempty"`
	Vid                      int64                                          `json:"vid,omitempty"`
	OperationSource          int64                                          `json:"operationSource,omitempty"`
}

type TagInsertListResponse struct {
	ListVo []TagInsertListResponseListVo `json:"listVo,omitempty"`
}

func CreateTagInsertListRequest() (request *TagInsertListRequest) {
	request = &TagInsertListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_cdp", "v2.0", "tag/insertList", "apigw")
	request.Method = api.POST
	return
}

func CreateTagInsertListResponse() (response *api.BaseResponse[TagInsertListResponse]) {
	response = api.CreateResponse[TagInsertListResponse](&TagInsertListResponse{})
	return
}

type TagInsertListRequestOpenCreateTagRequestList struct {
	TagAttrRequestList []TagInsertListRequestTagAttrRequestList `json:"tagAttrRequestList,omitempty"`
	Source             int64                                    `json:"source,omitempty"`
	TagName            string                                   `json:"tagName,omitempty"`
	TagType            int64                                    `json:"tagType,omitempty"`
	IsPersonal         int64                                    `json:"isPersonal,omitempty"`
}

type TagInsertListRequestTagAttrRequestList struct {
	AttrName string `json:"attrName,omitempty"`
}

type TagInsertListResponseListVo struct {
	TagId     int64  `json:"tagId,omitempty"`
	TagName   string `json:"tagName,omitempty"`
	IsSuccess bool   `json:"isSuccess,omitempty"`
	Source    int64  `json:"source,omitempty"`
}
