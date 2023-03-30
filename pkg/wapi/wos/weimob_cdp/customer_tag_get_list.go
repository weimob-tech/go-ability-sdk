package weimob_cdp

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerTagGetList
// @id 1560
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1560?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询客户的标签列表)
func (client *Client) CustomerTagGetList(request *CustomerTagGetListRequest) (response *CustomerTagGetListResponse, err error) {
	rpcResponse := CreateCustomerTagGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerTagGetListRequest struct {
	*api.BaseRequest

	WidList      []int64 `json:"widList,omitempty"`
	IsFilterCrow bool    `json:"isFilterCrow,omitempty"`
	VidType      int64   `json:"vidType,omitempty"`
	Vid          int64   `json:"vid,omitempty"`
}

type CustomerTagGetListResponse struct {
	CustomerTagListMap []CustomerTagGetListResponseCustomerTagListMap `json:"customerTagListMap,omitempty"`
}

func CreateCustomerTagGetListRequest() (request *CustomerTagGetListRequest) {
	request = &CustomerTagGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_cdp", "v2.0", "customer/tag/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomerTagGetListResponse() (response *api.BaseResponse[CustomerTagGetListResponse]) {
	response = api.CreateResponse[CustomerTagGetListResponse](&CustomerTagGetListResponse{})
	return
}

type CustomerTagGetListResponseCustomerTagListMap struct {
	Wid []CustomerTagGetListResponseWid `json:"wid,omitempty"`
}

type CustomerTagGetListResponseWid struct {
	UpdateTime            CustomerTagGetListResponseUpdateTime              `json:"updateTime,omitempty"`
	CustomerTagAttrVoList []CustomerTagGetListResponseCustomerTagAttrVoList `json:"customerTagAttrVoList,omitempty"`
	TagId                 int64                                             `json:"tagId,omitempty"`
	TagName               string                                            `json:"tagName,omitempty"`
	SortIndex             int64                                             `json:"sortIndex,omitempty"`
	TagType               int64                                             `json:"tagType,omitempty"`
}

type CustomerTagGetListResponseUpdateTime struct {
}

type CustomerTagGetListResponseCustomerTagAttrVoList struct {
	AttrId   int64  `json:"attrId,omitempty"`
	AttrName string `json:"attrName,omitempty"`
}
