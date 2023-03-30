package kld

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// KLDMemberCardGetConsumingLogPageListAndTotal
// @id 216
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取消费明细)
func (client *Client) KLDMemberCardGetConsumingLogPageListAndTotal(request *KLDMemberCardGetConsumingLogPageListAndTotalRequest) (response *KLDMemberCardGetConsumingLogPageListAndTotalResponse, err error) {
	rpcResponse := CreateKLDMemberCardGetConsumingLogPageListAndTotalResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type KLDMemberCardGetConsumingLogPageListAndTotalRequest struct {
	*api.BaseRequest

	MemberCardNo    string  `json:"memberCardNo,omitempty"`
	Name            string  `json:"name,omitempty"`
	StoreId         int64   `json:"storeId,omitempty"`
	CouponType      int     `json:"couponType,omitempty"`
	Phone           string  `json:"phone,omitempty"`
	AmountMin       float64 `json:"amountMin,omitempty"`
	AmountMax       float64 `json:"amountMax,omitempty"`
	Sort            int     `json:"sort,omitempty"`
	Begintime       int64   `json:"begintime,omitempty"`
	Endtime         int64   `json:"endtime,omitempty"`
	PageIndex       int     `json:"pageIndex,omitempty"`
	PageSize        int     `json:"pageSize,omitempty"`
	IsOnlyEffective bool    `json:"isOnlyEffective,omitempty"`
}

type KLDMemberCardGetConsumingLogPageListAndTotalResponse struct {
}

func CreateKLDMemberCardGetConsumingLogPageListAndTotalRequest() (request *KLDMemberCardGetConsumingLogPageListAndTotalRequest) {
	request = &KLDMemberCardGetConsumingLogPageListAndTotalRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("kld", "1_0", "KLDMemberCard/GetConsumingLogPageListAndTotal", "api")
	request.Method = api.POST
	return
}

func CreateKLDMemberCardGetConsumingLogPageListAndTotalResponse() (response *api.BaseResponse[KLDMemberCardGetConsumingLogPageListAndTotalResponse]) {
	response = api.CreateResponse[KLDMemberCardGetConsumingLogPageListAndTotalResponse](&KLDMemberCardGetConsumingLogPageListAndTotalResponse{})
	return
}
