package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RoomPricePlanGetRoomPricePlanList
// @id 1516
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取房价计划列表)
func (client *Client) RoomPricePlanGetRoomPricePlanList(request *RoomPricePlanGetRoomPricePlanListRequest) (response *RoomPricePlanGetRoomPricePlanListResponse, err error) {
	rpcResponse := CreateRoomPricePlanGetRoomPricePlanListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RoomPricePlanGetRoomPricePlanListRequest struct {
	*api.BaseRequest

	StoreId    int64 `json:"storeId,omitempty"`
	RoomTypeId int64 `json:"roomTypeId,omitempty"`
}

type RoomPricePlanGetRoomPricePlanListResponse struct {
	PricePlanRepItems []RoomPricePlanGetRoomPricePlanListResponsePricePlanRepItems `json:"pricePlanRepItems,omitempty"`
	RoomTypeId        int64                                                        `json:"roomTypeId,omitempty"`
	Order             int                                                          `json:"order,omitempty"`
	RoomTypeName      string                                                       `json:"roomTypeName,omitempty"`
}

func CreateRoomPricePlanGetRoomPricePlanListRequest() (request *RoomPricePlanGetRoomPricePlanListRequest) {
	request = &RoomPricePlanGetRoomPricePlanListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "roomPricePlan/getRoomPricePlanList", "api")
	request.Method = api.POST
	return
}

func CreateRoomPricePlanGetRoomPricePlanListResponse() (response *api.BaseResponse[RoomPricePlanGetRoomPricePlanListResponse]) {
	response = api.CreateResponse[RoomPricePlanGetRoomPricePlanListResponse](&RoomPricePlanGetRoomPricePlanListResponse{})
	return
}

type RoomPricePlanGetRoomPricePlanListResponsePricePlanRepItems struct {
	PricePlanId       int64   `json:"pricePlanId,omitempty"`
	PlanName          string  `json:"planName,omitempty"`
	BreakfastCount    int     `json:"breakfastCount,omitempty"`
	PayType           int     `json:"payType,omitempty"`
	CanSalesCount     int     `json:"canSalesCount,omitempty"`
	Price             float64 `json:"price,omitempty"`
	CanCouponCode     int     `json:"canCouponCode,omitempty"`
	CanCashCode       int     `json:"canCashCode,omitempty"`
	PriceType         int     `json:"priceType,omitempty"`
	SalesStrategyType int     `json:"salesStrategyType,omitempty"`
	SalesContent      string  `json:"salesContent,omitempty"`
	SortValue         int     `json:"sortValue,omitempty"`
	IsOpen            int     `json:"isOpen,omitempty"`
	GiftContent       string  `json:"giftContent,omitempty"`
	MinPrice          float64 `json:"minPrice,omitempty"`
	MaxPrice          float64 `json:"maxPrice,omitempty"`
}
