package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ScenicTicketSaveScenicTicket
// @id 1093
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新增/修改门票产品)
func (client *Client) ScenicTicketSaveScenicTicket(request *ScenicTicketSaveScenicTicketRequest) (response *ScenicTicketSaveScenicTicketResponse, err error) {
	rpcResponse := CreateScenicTicketSaveScenicTicketResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ScenicTicketSaveScenicTicketRequest struct {
	*api.BaseRequest

	GoodsCode     string `json:"goodsCode,omitempty"`
	GoodsName     string `json:"goodsName,omitempty"`
	ScenicCode    string `json:"scenicCode,omitempty"`
	AdvanceDay    int    `json:"advanceDay,omitempty"`
	AdvanceHour   int    `json:"advanceHour,omitempty"`
	AdvanceMinute int    `json:"advanceMinute,omitempty"`
	CheckinAddr   string `json:"checkinAddr,omitempty"`
	CheckinTime   string `json:"checkinTime,omitempty"`
	InparkMode    int    `json:"inparkMode,omitempty"`
	TravelerInfo  int    `json:"travelerInfo,omitempty"`
	BuyLimit      int    `json:"buyLimit,omitempty"`
	IsMemPrice    int    `json:"isMemPrice,omitempty"`
	IsPoint       int    `json:"isPoint,omitempty"`
	Intro         string `json:"intro,omitempty"`
}

type ScenicTicketSaveScenicTicketResponse struct {
}

func CreateScenicTicketSaveScenicTicketRequest() (request *ScenicTicketSaveScenicTicketRequest) {
	request = &ScenicTicketSaveScenicTicketRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "scenicTicket/saveScenicTicket", "api")
	request.Method = api.POST
	return
}

func CreateScenicTicketSaveScenicTicketResponse() (response *api.BaseResponse[ScenicTicketSaveScenicTicketResponse]) {
	response = api.CreateResponse[ScenicTicketSaveScenicTicketResponse](&ScenicTicketSaveScenicTicketResponse{})
	return
}
