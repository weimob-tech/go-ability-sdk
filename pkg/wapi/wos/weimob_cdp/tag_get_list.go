package weimob_cdp

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// TagGetList
// @id 1397
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1397?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询标签列表)
func (client *Client) TagGetList(request *TagGetListRequest) (response *TagGetListResponse, err error) {
	rpcResponse := CreateTagGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type TagGetListRequest struct {
	*api.BaseRequest

	TagList              []TagGetListRequestTagList `json:"tagList,omitempty"`
	HasCoverNumber       bool                       `json:"hasCoverNumber,omitempty"`
	HideEnabled          bool                       `json:"hideEnabled,omitempty"`
	HideConditionTag     bool                       `json:"hideConditionTag,omitempty"`
	NeedOtherPersonalTag bool                       `json:"needOtherPersonalTag,omitempty"`
	PageNum              int64                      `json:"pageNum,omitempty"`
	PageSize             int64                      `json:"pageSize,omitempty"`
	VidType              int64                      `json:"vidType,omitempty"`
	Vid                  int64                      `json:"vid,omitempty"`
	OperationSource      int64                      `json:"operationSource,omitempty"`
}

type TagGetListResponse struct {
	PageList   []TagGetListResponsePageList `json:"pageList,omitempty"`
	TotalCount int64                        `json:"totalCount,omitempty"`
	PageNum    int64                        `json:"pageNum,omitempty"`
	PageSize   int64                        `json:"pageSize,omitempty"`
}

func CreateTagGetListRequest() (request *TagGetListRequest) {
	request = &TagGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_cdp", "v2.0", "tag/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateTagGetListResponse() (response *api.BaseResponse[TagGetListResponse]) {
	response = api.CreateResponse[TagGetListResponse](&TagGetListResponse{})
	return
}

type TagGetListRequestTagList struct {
	TagId      int64   `json:"tagId,omitempty"`
	AttrIdList []int64 `json:"attrIdList,omitempty"`
}

type TagGetListResponsePageList struct {
	OpenTagAttrVoList []TagGetListResponseOpenTagAttrVoList `json:"openTagAttrVoList,omitempty"`
	TagId             int64                                 `json:"tagId,omitempty"`
	IsPersonal        int64                                 `json:"isPersonal,omitempty"`
	TagName           string                                `json:"tagName,omitempty"`
	TagType           int64                                 `json:"tagType,omitempty"`
	CoverNumber       int64                                 `json:"coverNumber,omitempty"`
	Source            int64                                 `json:"source,omitempty"`
	CreateDate        string                                `json:"createDate,omitempty"`
	UpdateDate        string                                `json:"updateDate,omitempty"`
}

type TagGetListResponseOpenTagAttrVoList struct {
	AttrId   int64  `json:"attrId,omitempty"`
	AttrName string `json:"attrName,omitempty"`
}
