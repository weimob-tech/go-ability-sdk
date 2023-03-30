package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomAttributeInsertList
// @id 2134
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2134?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=批量新增自定义字段自定义选项)
func (client *Client) CustomAttributeInsertList(request *CustomAttributeInsertListRequest) (response *CustomAttributeInsertListResponse, err error) {
	rpcResponse := CreateCustomAttributeInsertListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomAttributeInsertListRequest struct {
	*api.BaseRequest

	BasicInfo             CustomAttributeInsertListRequestBasicInfo               `json:"basicInfo,omitempty"`
	CustomerOptionDtoList []CustomAttributeInsertListRequestCustomerOptionDtoList `json:"customerOptionDtoList,omitempty"`
	AttributeId           int64                                                   `json:"attributeId,omitempty"`
}

type CustomAttributeInsertListResponse struct {
	IsSuccess bool `json:"isSuccess,omitempty"`
}

func CreateCustomAttributeInsertListRequest() (request *CustomAttributeInsertListRequest) {
	request = &CustomAttributeInsertListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "custom/attribute/insertList", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomAttributeInsertListResponse() (response *api.BaseResponse[CustomAttributeInsertListResponse]) {
	response = api.CreateResponse[CustomAttributeInsertListResponse](&CustomAttributeInsertListResponse{})
	return
}

type CustomAttributeInsertListRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type CustomAttributeInsertListRequestCustomerOptionDtoList struct {
	ChildList  []CustomAttributeInsertListRequestChildList `json:"childList,omitempty"`
	OptionName string                                      `json:"optionName,omitempty"`
	SortNum    int64                                       `json:"sortNum,omitempty"`
}

type CustomAttributeInsertListRequestChildList struct {
	ChildList  []CustomAttributeInsertListRequestChildList2 `json:"childList,omitempty"`
	OptionName string                                       `json:"optionName,omitempty"`
	SortNum    int64                                        `json:"sortNum,omitempty"`
}

type CustomAttributeInsertListRequestChildList2 struct {
	OptionName string `json:"optionName,omitempty"`
	SortNum    int64  `json:"sortNum,omitempty"`
}
