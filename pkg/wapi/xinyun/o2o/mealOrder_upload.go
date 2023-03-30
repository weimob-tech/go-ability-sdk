package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MealOrderUpload
// @id 262
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=上传订单（不建议再接入）)
func (client *Client) MealOrderUpload(request *MealOrderUploadRequest) (response *MealOrderUploadResponse, err error) {
	rpcResponse := CreateMealOrderUploadResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MealOrderUploadRequest struct {
	*api.BaseRequest

	MenuInfos          []map[string]any `json:"menuInfos,omitempty"`
	PayInfos           []map[string]any `json:"payInfos,omitempty"`
	StoreId            int64            `json:"storeId,omitempty"`
	PosName            string           `json:"posName,omitempty"`
	Operator           string           `json:"operator,omitempty"`
	OpenId             string           `json:"openId,omitempty"`
	ThirdOrderNo       string           `json:"thirdOrderNo,omitempty"`
	TotalAmount        int64            `json:"totalAmount,omitempty"`
	PayAmount          int64            `json:"payAmount,omitempty"`
	PayStatus          int              `json:"payStatus,omitempty"`
	GoodId             int64            `json:"goodId,omitempty"`
	GoodName           string           `json:"goodName,omitempty"`
	GoodNum            int              `json:"goodNum,omitempty"`
	GoodPrice          int64            `json:"goodPrice,omitempty"`
	TotalPrice         int64            `json:"totalPrice,omitempty"`
	GoodThirdId        string           `json:"goodThirdId,omitempty"`
	PayType            int              `json:"payType,omitempty"`
	Amount             int64            `json:"amount,omitempty"`
	CardCode           string           `json:"cardCode,omitempty"`
	Point              int64            `json:"point,omitempty"`
	PointRule          float64          `json:"pointRule,omitempty"`
	MemberCardCode     string           `json:"memberCardCode,omitempty"`
	DiscountAmountType int              `json:"discountAmountType,omitempty"`
	DiscountAmountRule string           `json:"discountAmountRule,omitempty"`
	PayWay             int              `json:"payWay,omitempty"`
}

type MealOrderUploadResponse struct {
}

func CreateMealOrderUploadRequest() (request *MealOrderUploadRequest) {
	request = &MealOrderUploadRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "mealOrder/upload", "api")
	request.Method = api.POST
	return
}

func CreateMealOrderUploadResponse() (response *api.BaseResponse[MealOrderUploadResponse]) {
	response = api.CreateResponse[MealOrderUploadResponse](&MealOrderUploadResponse{})
	return
}
