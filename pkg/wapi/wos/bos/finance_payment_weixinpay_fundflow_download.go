package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FinancePaymentWeixinpayFundflowDownload
// @id 2260
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2260?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=微信二级商户资金账单下载)
func (client *Client) FinancePaymentWeixinpayFundflowDownload(request *FinancePaymentWeixinpayFundflowDownloadRequest) (response *FinancePaymentWeixinpayFundflowDownloadResponse, err error) {
	rpcResponse := CreateFinancePaymentWeixinpayFundflowDownloadResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FinancePaymentWeixinpayFundflowDownloadRequest struct {
	*api.BaseRequest

	EntryMerchantNo string `json:"entryMerchantNo,omitempty"`
	BillDate        string `json:"billDate,omitempty"`
}

type FinancePaymentWeixinpayFundflowDownloadResponse struct {
	DownloadBillList  []FinancePaymentWeixinpayFundflowDownloadResponseDownloadBillList `json:"downloadBillList,omitempty"`
	DownloadBillCount int64                                                             `json:"downloadBillCount,omitempty"`
}

func CreateFinancePaymentWeixinpayFundflowDownloadRequest() (request *FinancePaymentWeixinpayFundflowDownloadRequest) {
	request = &FinancePaymentWeixinpayFundflowDownloadRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "finance/payment/weixinpay/fundflow/download", "apigw")
	request.Method = api.POST
	return
}

func CreateFinancePaymentWeixinpayFundflowDownloadResponse() (response *api.BaseResponse[FinancePaymentWeixinpayFundflowDownloadResponse]) {
	response = api.CreateResponse[FinancePaymentWeixinpayFundflowDownloadResponse](&FinancePaymentWeixinpayFundflowDownloadResponse{})
	return
}

type FinancePaymentWeixinpayFundflowDownloadResponseDownloadBillList struct {
	BillSequence int64  `json:"billSequence,omitempty"`
	DownloadUrl  string `json:"downloadUrl,omitempty"`
}
