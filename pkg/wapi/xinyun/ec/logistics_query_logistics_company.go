package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// LogisticsQueryLogisticsCompany
// @id 340
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询可用物流公司)
func (client *Client) LogisticsQueryLogisticsCompany(request *LogisticsQueryLogisticsCompanyRequest) (response *LogisticsQueryLogisticsCompanyResponse, err error) {
	rpcResponse := CreateLogisticsQueryLogisticsCompanyResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type LogisticsQueryLogisticsCompanyRequest struct {
	*api.BaseRequest
}

type LogisticsQueryLogisticsCompanyResponse struct {
	CompanyName string `json:"companyName,omitempty"`
	CompanyCode string `json:"companyCode,omitempty"`
}

func CreateLogisticsQueryLogisticsCompanyRequest() (request *LogisticsQueryLogisticsCompanyRequest) {
	request = &LogisticsQueryLogisticsCompanyRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "logistics/queryLogisticsCompany", "api")
	request.Method = api.POST
	return
}

func CreateLogisticsQueryLogisticsCompanyResponse() (response *api.BaseResponse[LogisticsQueryLogisticsCompanyResponse]) {
	response = api.CreateResponse[LogisticsQueryLogisticsCompanyResponse](&LogisticsQueryLogisticsCompanyResponse{})
	return
}
