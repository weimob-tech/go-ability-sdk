package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FinanceBillRefundGet
// @id 2231
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2231?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询支付退款账单)
func (client *Client) FinanceBillRefundGet(request *FinanceBillRefundGetRequest) (response *FinanceBillRefundGetResponse, err error) {
	rpcResponse := CreateFinanceBillRefundGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FinanceBillRefundGetRequest struct {
	*api.BaseRequest

	RefundTime string `json:"refundTime,omitempty"`
	PageNum    int64  `json:"pageNum,omitempty"`
	PageSize   int64  `json:"pageSize,omitempty"`
}

type FinanceBillRefundGetResponse struct {
	List        []FinanceBillRefundGetResponselist `json:"list,omitempty"`
	PageCount   int64                              `json:"pageCount,omitempty"`
	RecordCount int64                              `json:"recordCount,omitempty"`
	PageSize    int64                              `json:"pageSize,omitempty"`
	PageNum     int64                              `json:"pageNum,omitempty"`
}

func CreateFinanceBillRefundGetRequest() (request *FinanceBillRefundGetRequest) {
	request = &FinanceBillRefundGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "finance/bill/refund/get", "apigw")
	request.Method = api.POST
	return
}

func CreateFinanceBillRefundGetResponse() (response *api.BaseResponse[FinanceBillRefundGetResponse]) {
	response = api.CreateResponse[FinanceBillRefundGetResponse](&FinanceBillRefundGetResponse{})
	return
}

type FinanceBillRefundGetResponselist struct {
	RefundTime             string `json:"refundTime,omitempty"`
	RefundAmount           string `json:"refundAmount,omitempty"`
	SettlementRefundAmount string `json:"settlementRefundAmount,omitempty"`
	RefundFee              string `json:"refundFee,omitempty"`
	CouponRefundAmount     string `json:"couponRefundAmount,omitempty"`
	RefundStatus           int64  `json:"refundStatus,omitempty"`
	RefundStatusStr        string `json:"refundStatusStr,omitempty"`
	Remark                 string `json:"remark,omitempty"`
	CheckStatus            int64  `json:"checkStatus,omitempty"`
	CheckStatusStr         string `json:"checkStatusStr,omitempty"`
	CheckTime              string `json:"checkTime,omitempty"`
	StoreName              string `json:"storeName,omitempty"`
	BizRefundNo            string `json:"bizRefundNo,omitempty"`
	RefundChannelTrxNo     string `json:"refundChannelTrxNo,omitempty"`
	InteractId             string `json:"interactId,omitempty"`
	PayTime                string `json:"payTime,omitempty"`
	PaymentAmount          string `json:"paymentAmount,omitempty"`
	UnFreezeTime           string `json:"unFreezeTime,omitempty"`
	GoodsRefundNo          string `json:"goodsRefundNo,omitempty"`
	OrderNo                string `json:"orderNo,omitempty"`
	PaymentChannelTrxNo    string `json:"paymentChannelTrxNo,omitempty"`
	PayModeDec             string `json:"payModeDec,omitempty"`
	ChannelMerchantNo      string `json:"channelMerchantNo,omitempty"`
	Currency               string `json:"currency,omitempty"`
	ChannelCode            string `json:"channelCode,omitempty"`
	ChannelCodeDec         string `json:"channelCodeDec,omitempty"`
}
