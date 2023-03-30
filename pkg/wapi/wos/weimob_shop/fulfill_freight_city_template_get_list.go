package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FulfillFreightCityTemplateGetList
// @id 1430
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1430?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询同城配送模板列表)
func (client *Client) FulfillFreightCityTemplateGetList(request *FulfillFreightCityTemplateGetListRequest) (response *FulfillFreightCityTemplateGetListResponse, err error) {
	rpcResponse := CreateFulfillFreightCityTemplateGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FulfillFreightCityTemplateGetListRequest struct {
	*api.BaseRequest

	BasicInfo FulfillFreightCityTemplateGetListRequestBasicInfo `json:"basicInfo,omitempty"`
}

type FulfillFreightCityTemplateGetListResponse struct {
	DefaultTemplate     FulfillFreightCityTemplateGetListResponseDefaultTemplate       `json:"defaultTemplate,omitempty"`
	FreightTemplateList []FulfillFreightCityTemplateGetListResponseFreightTemplateList `json:"freightTemplateList,omitempty"`
	Success             bool                                                           `json:"success,omitempty"`
}

func CreateFulfillFreightCityTemplateGetListRequest() (request *FulfillFreightCityTemplateGetListRequest) {
	request = &FulfillFreightCityTemplateGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "fulfill/freight/city/template/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateFulfillFreightCityTemplateGetListResponse() (response *api.BaseResponse[FulfillFreightCityTemplateGetListResponse]) {
	response = api.CreateResponse[FulfillFreightCityTemplateGetListResponse](&FulfillFreightCityTemplateGetListResponse{})
	return
}

type FulfillFreightCityTemplateGetListRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type FulfillFreightCityTemplateGetListResponseDefaultTemplate struct {
	Id           int64  `json:"id,omitempty"`
	TemplateName string `json:"templateName,omitempty"`
}

type FulfillFreightCityTemplateGetListResponseFreightTemplateList struct {
	Id           int64  `json:"id,omitempty"`
	TemplateName string `json:"templateName,omitempty"`
}
