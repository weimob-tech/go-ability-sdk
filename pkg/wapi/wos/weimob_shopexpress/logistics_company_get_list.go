package weimob_shopexpress

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// LogisticsCompanyGetList
// @id 1955
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1955?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=物流公司列表)
func (client *Client) LogisticsCompanyGetList(request *LogisticsCompanyGetListRequest) (response *LogisticsCompanyGetListResponse, err error) {
	rpcResponse := CreateLogisticsCompanyGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type LogisticsCompanyGetListRequest struct {
	*api.BaseRequest

	CompanyNameKeyword string `json:"companyNameKeyword,omitempty"`
	PageIndex          int64  `json:"pageIndex,omitempty"`
	PageSize           int64  `json:"pageSize,omitempty"`
}

type LogisticsCompanyGetListResponse struct {
	Items      []LogisticsCompanyGetListResponseItems `json:"items,omitempty"`
	TotalCount int64                                  `json:"totalCount,omitempty"`
}

func CreateLogisticsCompanyGetListRequest() (request *LogisticsCompanyGetListRequest) {
	request = &LogisticsCompanyGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shopexpress", "v2.0", "logistics/company/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateLogisticsCompanyGetListResponse() (response *api.BaseResponse[LogisticsCompanyGetListResponse]) {
	response = api.CreateResponse[LogisticsCompanyGetListResponse](&LogisticsCompanyGetListResponse{})
	return
}

type LogisticsCompanyGetListResponseItems struct {
	CompanyName string `json:"companyName,omitempty"`
	CompanyCode string `json:"companyCode,omitempty"`
}
