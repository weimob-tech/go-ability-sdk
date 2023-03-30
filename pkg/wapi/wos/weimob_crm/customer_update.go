package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerUpdate
// @id 1745
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1745?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=更新客户资料信息)
func (client *Client) CustomerUpdate(request *CustomerUpdateRequest) (response *CustomerUpdateResponse, err error) {
	rpcResponse := CreateCustomerUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerUpdateRequest struct {
	*api.BaseRequest

	ExtMap          []CustomerUpdateRequestExtMap `json:"extMap,omitempty"`
	Vid             string                        `json:"vid,omitempty"`
	Wid             int64                         `json:"wid,omitempty"`
	Name            string                        `json:"name,omitempty"`
	Gender          int64                         `json:"gender,omitempty"`
	Birthday        int64                         `json:"birthday,omitempty"`
	IdentityCardNum string                        `json:"identityCardNum,omitempty"`
	Email           string                        `json:"email,omitempty"`
	Education       string                        `json:"education,omitempty"`
	Income          string                        `json:"income,omitempty"`
	Industry        string                        `json:"industry,omitempty"`
	Hobby           string                        `json:"hobby,omitempty"`
	Province        string                        `json:"province,omitempty"`
	ProvinceCode    string                        `json:"provinceCode,omitempty"`
	City            string                        `json:"city,omitempty"`
	CityCode        string                        `json:"cityCode,omitempty"`
	Area            string                        `json:"area,omitempty"`
	AreaCode        string                        `json:"areaCode,omitempty"`
	Address         string                        `json:"address,omitempty"`
}

type CustomerUpdateResponse struct {
	Result bool `json:"result,omitempty"`
}

func CreateCustomerUpdateRequest() (request *CustomerUpdateRequest) {
	request = &CustomerUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "customer/update", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomerUpdateResponse() (response *api.BaseResponse[CustomerUpdateResponse]) {
	response = api.CreateResponse[CustomerUpdateResponse](&CustomerUpdateResponse{})
	return
}

type CustomerUpdateRequestExtMap struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}
