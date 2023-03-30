package weimob_guide

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RewardCommissionGoodsGetList
// @id 1752
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1752?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询特殊商品提成明细)
func (client *Client) RewardCommissionGoodsGetList(request *RewardCommissionGoodsGetListRequest) (response *RewardCommissionGoodsGetListResponse, err error) {
	rpcResponse := CreateRewardCommissionGoodsGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RewardCommissionGoodsGetListRequest struct {
	*api.BaseRequest

	PageNum     int64  `json:"pageNum,omitempty"`
	PageSize    int64  `json:"pageSize,omitempty"`
	GuiderPhone string `json:"guiderPhone,omitempty"`
	GuiderWid   int64  `json:"guiderWid,omitempty"`
	JobNumber   string `json:"jobNumber,omitempty"`
	CreateTime  string `json:"createTime,omitempty"`
	EndTime     string `json:"endTime,omitempty"`
	Vid         int64  `json:"vid,omitempty"`
	VidType     int64  `json:"vidType,omitempty"`
}

type RewardCommissionGoodsGetListResponse struct {
	PageList   []RewardCommissionGoodsGetListResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                                          `json:"pageNum,omitempty"`
	PageSize   int64                                          `json:"pageSize,omitempty"`
	TotalCount int64                                          `json:"totalCount,omitempty"`
}

func CreateRewardCommissionGoodsGetListRequest() (request *RewardCommissionGoodsGetListRequest) {
	request = &RewardCommissionGoodsGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_guide", "v2.0", "reward/commission/goods/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateRewardCommissionGoodsGetListResponse() (response *api.BaseResponse[RewardCommissionGoodsGetListResponse]) {
	response = api.CreateResponse[RewardCommissionGoodsGetListResponse](&RewardCommissionGoodsGetListResponse{})
	return
}

type RewardCommissionGoodsGetListResponsePageList struct {
	TradeId                     int64  `json:"tradeId,omitempty"`
	TradeType                   int64  `json:"tradeType,omitempty"`
	GoodsId                     int64  `json:"goodsId,omitempty"`
	GoodsTitle                  string `json:"goodsTitle,omitempty"`
	SkuId                       string `json:"skuId,omitempty"`
	SkuNum                      int64  `json:"skuNum,omitempty"`
	SkuName                     string `json:"skuName,omitempty"`
	GuiderWid                   int64  `json:"guiderWid,omitempty"`
	GuiderName                  string `json:"guiderName,omitempty"`
	GuiderStoreName             string `json:"guiderStoreName,omitempty"`
	GuiderPhone                 string `json:"guiderPhone,omitempty"`
	PaymentAmount               int64  `json:"paymentAmount,omitempty"`
	BalanceDiscountAmount       int64  `json:"balanceDiscountAmount,omitempty"`
	PointsAmount                int64  `json:"pointsAmount,omitempty"`
	SharedPayAmount             int64  `json:"sharedPayAmount,omitempty"`
	SharedPointsAmount          int64  `json:"sharedPointsAmount,omitempty"`
	SharedBalanceDiscountAmount int64  `json:"sharedBalanceDiscountAmount,omitempty"`
	CommissionAmount            int64  `json:"commissionAmount,omitempty"`
	CommissionRewardAmount      int64  `json:"commissionRewardAmount,omitempty"`
	CommissionType              int64  `json:"commissionType,omitempty"`
	CommissionRatio             int64  `json:"commissionRatio,omitempty"`
	CreateTime                  string `json:"createTime,omitempty"`
	ItemCommissionAmount        int64  `json:"itemCommissionAmount,omitempty"`
}
