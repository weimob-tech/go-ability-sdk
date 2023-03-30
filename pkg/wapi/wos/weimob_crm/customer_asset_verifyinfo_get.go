package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerAssetVerifyinfoGet
// @id 1648
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1648?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取资产核销客户授权校验二维码)
func (client *Client) CustomerAssetVerifyinfoGet(request *CustomerAssetVerifyinfoGetRequest) (response *CustomerAssetVerifyinfoGetResponse, err error) {
	rpcResponse := CreateCustomerAssetVerifyinfoGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerAssetVerifyinfoGetRequest struct {
	*api.BaseRequest

	OperatorWid int64  `json:"operatorWid,omitempty"`
	Wid         int64  `json:"wid,omitempty"`
	Vid         int64  `json:"vid,omitempty"`
	Point       int64  `json:"point,omitempty"`
	Balance     int64  `json:"balance,omitempty"`
	TradeToken  string `json:"tradeToken,omitempty"`
}

type CustomerAssetVerifyinfoGetResponse struct {
	QrCodeUrlList []CustomerAssetVerifyinfoGetResponseQrCodeUrlList `json:"qrCodeUrlList,omitempty"`
	CheckToken    string                                            `json:"checkToken,omitempty"`
}

func CreateCustomerAssetVerifyinfoGetRequest() (request *CustomerAssetVerifyinfoGetRequest) {
	request = &CustomerAssetVerifyinfoGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "customer/asset/verifyinfo/get", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomerAssetVerifyinfoGetResponse() (response *api.BaseResponse[CustomerAssetVerifyinfoGetResponse]) {
	response = api.CreateResponse[CustomerAssetVerifyinfoGetResponse](&CustomerAssetVerifyinfoGetResponse{})
	return
}

type CustomerAssetVerifyinfoGetResponseQrCodeUrlList struct {
	QrcodeUrl   string `json:"qrcodeUrl,omitempty"`
	Channel     int64  `json:"channel,omitempty"`
	ChannelName string `json:"channelName,omitempty"`
}
