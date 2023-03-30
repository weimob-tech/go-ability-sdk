package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// TicketVerifyQueryVerifyCode
// @id 1091
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询核销码)
func (client *Client) TicketVerifyQueryVerifyCode(request *TicketVerifyQueryVerifyCodeRequest) (response *TicketVerifyQueryVerifyCodeResponse, err error) {
	rpcResponse := CreateTicketVerifyQueryVerifyCodeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type TicketVerifyQueryVerifyCodeRequest struct {
	*api.BaseRequest

	Code      string `json:"code,omitempty"`
	Phone     string `json:"phone,omitempty"`
	OrderNo   string `json:"orderNo,omitempty"`
	PageIndex int    `json:"pageIndex,omitempty"`
	PageSize  int    `json:"pageSize,omitempty"`
}

type TicketVerifyQueryVerifyCodeResponse struct {
	Results    TicketVerifyQueryVerifyCodeResponseResults `json:"results,omitempty"`
	TotalCount int64                                      `json:"totalCount,omitempty"`
}

func CreateTicketVerifyQueryVerifyCodeRequest() (request *TicketVerifyQueryVerifyCodeRequest) {
	request = &TicketVerifyQueryVerifyCodeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "ticketVerify/queryVerifyCode", "api")
	request.Method = api.POST
	return
}

func CreateTicketVerifyQueryVerifyCodeResponse() (response *api.BaseResponse[TicketVerifyQueryVerifyCodeResponse]) {
	response = api.CreateResponse[TicketVerifyQueryVerifyCodeResponse](&TicketVerifyQueryVerifyCodeResponse{})
	return
}

type TicketVerifyQueryVerifyCodeResponseResults struct {
	Code           string `json:"code,omitempty"`
	OrderNo        string `json:"orderNo,omitempty"`
	GoodsName      string `json:"goodsName,omitempty"`
	ScenicName     string `json:"scenicName,omitempty"`
	ScenicImg      string `json:"scenicImg,omitempty"`
	Price          int64  `json:"price,omitempty"`
	NumberOfPeople int64  `json:"numberOfPeople,omitempty"`
	Status         int64  `json:"status,omitempty"`
	CodeSource     int64  `json:"codeSource,omitempty"`
	VisitDate      string `json:"visitDate,omitempty"`
	Phone          string `json:"phone,omitempty"`
	Receiver       string `json:"receiver,omitempty"`
}
