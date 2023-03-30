package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CityFreightTemplateDeleteCityTemplate
// @id 2437
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除同城模板)
func (client *Client) CityFreightTemplateDeleteCityTemplate(request *CityFreightTemplateDeleteCityTemplateRequest) (response *CityFreightTemplateDeleteCityTemplateResponse, err error) {
	rpcResponse := CreateCityFreightTemplateDeleteCityTemplateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CityFreightTemplateDeleteCityTemplateRequest struct {
	*api.BaseRequest

	Pid           int64 `json:"pid,omitempty"`
	StoreId       int64 `json:"storeId,omitempty"`
	TemplateId    int64 `json:"templateId,omitempty"`
	NewTemplateId int64 `json:"newTemplateId,omitempty"`
}

type CityFreightTemplateDeleteCityTemplateResponse struct {
	Result string `json:"result,omitempty"`
}

func CreateCityFreightTemplateDeleteCityTemplateRequest() (request *CityFreightTemplateDeleteCityTemplateRequest) {
	request = &CityFreightTemplateDeleteCityTemplateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "cityFreightTemplate/deleteCityTemplate", "api")
	request.Method = api.POST
	return
}

func CreateCityFreightTemplateDeleteCityTemplateResponse() (response *api.BaseResponse[CityFreightTemplateDeleteCityTemplateResponse]) {
	response = api.CreateResponse[CityFreightTemplateDeleteCityTemplateResponse](&CityFreightTemplateDeleteCityTemplateResponse{})
	return
}
