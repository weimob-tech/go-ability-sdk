package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomOptionUpdateList
// @id 2131
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2131?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=批量修改自定义字段选项信息)
func (client *Client) CustomOptionUpdateList(request *CustomOptionUpdateListRequest) (response *CustomOptionUpdateListResponse, err error) {
	rpcResponse := CreateCustomOptionUpdateListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomOptionUpdateListRequest struct {
	*api.BaseRequest

	BasicInfo        CustomOptionUpdateListRequestBasicInfo          `json:"basicInfo,omitempty"`
	CustomOptionList []CustomOptionUpdateListRequestCustomOptionList `json:"customOptionList,omitempty"`
	ParentId         int64                                           `json:"parentId,omitempty"`
	AttributeId      string                                          `json:"attributeId,omitempty"`
}

type CustomOptionUpdateListResponse struct {
	IsSuccess bool `json:"isSuccess,omitempty"`
}

func CreateCustomOptionUpdateListRequest() (request *CustomOptionUpdateListRequest) {
	request = &CustomOptionUpdateListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "custom/option/updateList", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomOptionUpdateListResponse() (response *api.BaseResponse[CustomOptionUpdateListResponse]) {
	response = api.CreateResponse[CustomOptionUpdateListResponse](&CustomOptionUpdateListResponse{})
	return
}

type CustomOptionUpdateListRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type CustomOptionUpdateListRequestCustomOptionList struct {
	Id         int64  `json:"id,omitempty"`
	OptionName string `json:"optionName,omitempty"`
	SortNum    int64  `json:"sortNum,omitempty"`
}
