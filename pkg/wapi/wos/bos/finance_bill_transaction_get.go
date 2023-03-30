package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FinanceBillTransactionGet
// @id 2232
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2232?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询支付交易账单)
func (client *Client) FinanceBillTransactionGet(request *FinanceBillTransactionGetRequest) (response *FinanceBillTransactionGetResponse, err error) {
	rpcResponse := CreateFinanceBillTransactionGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FinanceBillTransactionGetRequest struct {
	*api.BaseRequest

	PayTime  string `json:"payTime,omitempty"`
	PageNum  int64  `json:"pageNum,omitempty"`
	PageSize int64  `json:"pageSize,omitempty"`
}

type FinanceBillTransactionGetResponse struct {
	List        []FinanceBillTransactionGetResponselist `json:"list,omitempty"`
	PageCount   int64                                   `json:"pageCount,omitempty"`
	RecordCount int64                                   `json:"recordCount,omitempty"`
	PageNum     int64                                   `json:"pageNum,omitempty"`
	PageSize    int64                                   `json:"pageSize,omitempty"`
}

func CreateFinanceBillTransactionGetRequest() (request *FinanceBillTransactionGetRequest) {
	request = &FinanceBillTransactionGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "finance/bill/transaction/get", "apigw")
	request.Method = api.POST
	return
}

func CreateFinanceBillTransactionGetResponse() (response *api.BaseResponse[FinanceBillTransactionGetResponse]) {
	response = api.CreateResponse[FinanceBillTransactionGetResponse](&FinanceBillTransactionGetResponse{})
	return
}

type FinanceBillTransactionGetResponselist struct {
	PayTime                       string `json:"payTime,omitempty"`
	OrderNos                      string `json:"orderNos,omitempty"`
	TransNo                       string `json:"transNo,omitempty"`
	ChannelTrxNo                  string `json:"channelTrxNo,omitempty"`
	InteractId                    string `json:"interactId,omitempty"`
	PayModeDec                    string `json:"payModeDec,omitempty"`
	Name                          string `json:"name,omitempty"`
	Fee                           string `json:"fee,omitempty"`
	OnlineSharingProcessStatusDec string `json:"onlineSharingProcessStatusDec,omitempty"`
	CheckStatusDec                string `json:"checkStatusDec,omitempty"`
	PayAmount                     string `json:"payAmount,omitempty"`
	ArriveAmount                  string `json:"arriveAmount,omitempty"`
	ArriveTime                    string `json:"arriveTime,omitempty"`
	IsSharingSplitDec             string `json:"isSharingSplitDec,omitempty"`
	RefundFee                     string `json:"refundFee,omitempty"`
	UnFreezeTime                  string `json:"unFreezeTime,omitempty"`
	RefundAmount                  string `json:"refundAmount,omitempty"`
	ChannelMerchantNo             string `json:"channelMerchantNo,omitempty"`
	SettlementAmount              string `json:"settlementAmount,omitempty"`
	CouponAmount                  string `json:"couponAmount,omitempty"`
	Currency                      string `json:"currency,omitempty"`
}
