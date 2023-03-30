package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomOptionGetList
// @id 2133
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2133?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询自定义字段选项列表)
func (client *Client) CustomOptionGetList(request *CustomOptionGetListRequest) (response *CustomOptionGetListResponse, err error) {
	rpcResponse := CreateCustomOptionGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomOptionGetListRequest struct {
	*api.BaseRequest

	BasicInfo   CustomOptionGetListRequestBasicInfo `json:"basicInfo,omitempty"`
	AttributeId int64                               `json:"attributeId,omitempty"`
	ParentId    int64                               `json:"parentId,omitempty"`
	PageNum     int64                               `json:"pageNum,omitempty"`
	PageSize    int64                               `json:"pageSize,omitempty"`
	OnlyTitle   int64                               `json:"onlyTitle,omitempty"`
}

type CustomOptionGetListResponse struct {
	PageList   []CustomOptionGetListResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                                 `json:"pageNum,omitempty"`
	PageSize   int64                                 `json:"pageSize,omitempty"`
	TotalCount int64                                 `json:"totalCount,omitempty"`
}

func CreateCustomOptionGetListRequest() (request *CustomOptionGetListRequest) {
	request = &CustomOptionGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "custom/option/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomOptionGetListResponse() (response *api.BaseResponse[CustomOptionGetListResponse]) {
	response = api.CreateResponse[CustomOptionGetListResponse](&CustomOptionGetListResponse{})
	return
}

type CustomOptionGetListRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type CustomOptionGetListResponsePageList struct {
	AttributeId int64  `json:"attributeId,omitempty"`
	Id          int64  `json:"id,omitempty"`
	IsTitle     int64  `json:"isTitle,omitempty"`
	Level       int64  `json:"level,omitempty"`
	OptionName  string `json:"optionName,omitempty"`
	ParentId    int64  `json:"parentId,omitempty"`
	SortNum     int64  `json:"sortNum,omitempty"`
}
