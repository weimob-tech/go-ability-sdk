package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// SettingAddRoomType
// @id 528
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新增房型)
func (client *Client) SettingAddRoomType(request *SettingAddRoomTypeRequest) (response *SettingAddRoomTypeResponse, err error) {
	rpcResponse := CreateSettingAddRoomTypeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type SettingAddRoomTypeRequest struct {
	*api.BaseRequest

	ImgUrlVoList        []SettingAddRoomTypeRequestImgUrlVoList `json:"imgUrlVoList,omitempty"`
	VideoUrlList        []SettingAddRoomTypeRequestVideoUrlList `json:"videoUrlList,omitempty"`
	StoreId             int64                                   `json:"storeId,omitempty"`
	Name                string                                  `json:"name,omitempty"`
	Price               float64                                 `json:"price,omitempty"`
	NeedDeposit         bool                                    `json:"needDeposit,omitempty"`
	Deposit             float64                                 `json:"deposit,omitempty"`
	Payment             bool                                    `json:"payment,omitempty"`
	Area                int                                     `json:"area,omitempty"`
	Guests              int                                     `json:"guests,omitempty"`
	Detail              string                                  `json:"detail,omitempty"`
	FacilityList        []int                                   `json:"facilityList,omitempty"`
	IsOnlineDeposit     bool                                    `json:"isOnlineDeposit,omitempty"`
	CanCoupon           bool                                    `json:"canCoupon,omitempty"`
	CanCashCoupon       bool                                    `json:"canCashCoupon,omitempty"`
	DiscountsType       int                                     `json:"discountsType,omitempty"`
	ExtraBedType        int                                     `json:"extraBedType,omitempty"`
	ExtraBedText        string                                  `json:"extraBedText,omitempty"`
	ExtraBedPrice       float64                                 `json:"extraBedPrice,omitempty"`
	PanoramaUrl         string                                  `json:"panoramaUrl,omitempty"`
	PanoramaCoverPicUrl string                                  `json:"panoramaCoverPicUrl,omitempty"`
	NeedHousemate       bool                                    `json:"needHousemate,omitempty"`
	HousemateNum        int                                     `json:"housemateNum,omitempty"`
	Labels              []string                                `json:"labels,omitempty"`
}

type SettingAddRoomTypeResponse struct {
}

func CreateSettingAddRoomTypeRequest() (request *SettingAddRoomTypeRequest) {
	request = &SettingAddRoomTypeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "setting/addRoomType", "api")
	request.Method = api.POST
	return
}

func CreateSettingAddRoomTypeResponse() (response *api.BaseResponse[SettingAddRoomTypeResponse]) {
	response = api.CreateResponse[SettingAddRoomTypeResponse](&SettingAddRoomTypeResponse{})
	return
}

type SettingAddRoomTypeRequestImgUrlVoList struct {
	Url string `json:"url,omitempty"`
}

type SettingAddRoomTypeRequestVideoUrlList struct {
	Url         string `json:"url,omitempty"`
	Duration    int64  `json:"duration,omitempty"`
	CoverPicUrl string `json:"coverPicUrl,omitempty"`
}
