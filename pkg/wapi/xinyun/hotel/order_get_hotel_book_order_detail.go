package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderGetHotelBookOrderDetail
// @id 511
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=酒店订房订单详情)
func (client *Client) OrderGetHotelBookOrderDetail(request *OrderGetBookOrderDetailRequest) (response *OrderGetBookOrderDetailResponse, err error) {
	rpcResponse := CreateOrderGetBookOrderDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderGetBookOrderDetailRequest struct {
	*api.BaseRequest

	OrderNo string `json:"orderNo,omitempty"`
}

type OrderGetBookOrderDetailResponse struct {
	Details              []OrderGetBookOrderDetailResponseDetails        `json:"details,omitempty"`
	RoomTimePrices       []OrderGetBookOrderDetailResponseRoomTimePrices `json:"roomTimePrices,omitempty"`
	VoucherList          []OrderGetBookOrderDetailResponseVoucherList    `json:"voucherList,omitempty"`
	BlanceAmount         int64                                           `json:"blanceAmount,omitempty"`
	ConsumeMemberPoint   int64                                           `json:"consumeMemberPoint,omitempty"`
	Count                int64                                           `json:"count,omitempty"`
	CreateTime           int64                                           `json:"createTime,omitempty"`
	IsMember             bool                                            `json:"isMember,omitempty"`
	IsShowCheckInButton  int64                                           `json:"isShowCheckInButton,omitempty"`
	MemberDiscount       int64                                           `json:"memberDiscount,omitempty"`
	MemberLevel          int64                                           `json:"memberLevel,omitempty"`
	MemberLevelName      string                                          `json:"memberLevelName,omitempty"`
	MemberStatus         int64                                           `json:"memberStatus,omitempty"`
	Mobile               string                                          `json:"mobile,omitempty"`
	Name                 string                                          `json:"name,omitempty"`
	OrderNo              string                                          `json:"orderNo,omitempty"`
	OrderStatus          int64                                           `json:"orderStatus,omitempty"`
	PayTime              int64                                           `json:"payTime,omitempty"`
	RankName             string                                          `json:"rankName,omitempty"`
	RealAmount           int64                                           `json:"realAmount,omitempty"`
	StoreId              int64                                           `json:"storeId,omitempty"`
	SubOrderType         int64                                           `json:"subOrderType,omitempty"`
	TotalCharge          int64                                           `json:"totalCharge,omitempty"`
	TotalProfit          int64                                           `json:"totalProfit,omitempty"`
	Wid                  int64                                           `json:"wid,omitempty"`
	ContinuityType       int64                                           `json:"continuityType,omitempty"`
	ContinuityPrice      int64                                           `json:"continuityPrice,omitempty"`
	ContinuityRealAmount int64                                           `json:"continuityRealAmount,omitempty"`
	CouponRealAmount     int64                                           `json:"couponRealAmount,omitempty"`
	CouponCode           string                                          `json:"couponCode,omitempty"`
	MemberRealAmount     int64                                           `json:"memberRealAmount,omitempty"`
	SettleStatus         int64                                           `json:"settleStatus,omitempty"`
}

func CreateOrderGetBookOrderDetailRequest() (request *OrderGetBookOrderDetailRequest) {
	request = &OrderGetBookOrderDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "order/getHotelBookOrderDetail", "api")
	request.Method = api.POST
	return
}

func CreateOrderGetBookOrderDetailResponse() (response *api.BaseResponse[OrderGetBookOrderDetailResponse]) {
	response = api.CreateResponse[OrderGetBookOrderDetailResponse](&OrderGetBookOrderDetailResponse{})
	return
}

type OrderGetBookOrderDetailResponseDetails struct {
	LodgerInfos      []OrderGetBookOrderDetailResponseLodgerInfos `json:"lodgerInfos,omitempty"`
	CheckInTime      int64                                        `json:"checkInTime,omitempty"`
	CheckOutTime     int64                                        `json:"checkOutTime,omitempty"`
	OrgRoomCharge    int64                                        `json:"orgRoomCharge,omitempty"`
	RoomTotalCharge  int64                                        `json:"roomTotalCharge,omitempty"`
	RoomTypeId       int64                                        `json:"roomTypeId,omitempty"`
	RoomTypeName     string                                       `json:"roomTypeName,omitempty"`
	CashPledge       int64                                        `json:"cashPledge,omitempty"`
	StatusNo         int64                                        `json:"statusNo,omitempty"`
	RealCheckOutTime int64                                        `json:"realCheckOutTime,omitempty"`
	RealCheckInTime  int64                                        `json:"realCheckInTime,omitempty"`
	RoomNumber       string                                       `json:"roomNumber,omitempty"`
	RoomId           int64                                        `json:"roomId,omitempty"`
	CardType         int64                                        `json:"cardType,omitempty"`
	CardNo           string                                       `json:"cardNo,omitempty"`
}

type OrderGetBookOrderDetailResponseLodgerInfos struct {
	Mobile    string `json:"mobile,omitempty"`
	Name      string `json:"name,omitempty"`
	OrderType int64  `json:"orderType,omitempty"`
}

type OrderGetBookOrderDetailResponseRoomTimePrices struct {
	Date    string `json:"date,omitempty"`
	Price   int64  `json:"price,omitempty"`
	WeekDay string `json:"weekDay,omitempty"`
}

type OrderGetBookOrderDetailResponseVoucherList struct {
	CashCouponSnNo            string  `json:"cashCouponSnNo,omitempty"`
	Name                      string  `json:"name,omitempty"`
	SellPrice                 int64   `json:"sellPrice,omitempty"`
	CashCouponAmount          int64   `json:"cashCouponAmount,omitempty"`
	CashCouponDeductionAmount int64   `json:"cashCouponDeductionAmount,omitempty"`
	VoucherJson               []int64 `json:"voucherJson,omitempty"`
	VoucherRealAmount         int64   `json:"voucherRealAmount,omitempty"`
}
