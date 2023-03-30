package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// TicketVerifyQueryRecord
// @id 1092
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询核销记录)
func (client *Client) TicketVerifyQueryRecord(request *TicketVerifyQueryRecordRequest) (response *TicketVerifyQueryRecordResponse, err error) {
	rpcResponse := CreateTicketVerifyQueryRecordResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type TicketVerifyQueryRecordRequest struct {
	*api.BaseRequest

	Code            string `json:"code,omitempty"`
	Phone           string `json:"phone,omitempty"`
	OrderNo         string `json:"orderNo,omitempty"`
	GoodsName       string `json:"goodsName,omitempty"`
	Receiver        string `json:"receiver,omitempty"`
	VerifyClient    int    `json:"verifyClient,omitempty"`
	VerifyTimeStart int    `json:"verifyTimeStart,omitempty"`
	VerifyTimeEnd   int    `json:"verifyTimeEnd,omitempty"`
	PageIndex       int    `json:"pageIndex,omitempty"`
	PageSize        int    `json:"pageSize,omitempty"`
}

type TicketVerifyQueryRecordResponse struct {
	Results    TicketVerifyQueryRecordResponseResults `json:"results,omitempty"`
	TotalCount int64                                  `json:"totalCount,omitempty"`
}

func CreateTicketVerifyQueryRecordRequest() (request *TicketVerifyQueryRecordRequest) {
	request = &TicketVerifyQueryRecordRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "ticketVerify/queryRecord", "api")
	request.Method = api.POST
	return
}

func CreateTicketVerifyQueryRecordResponse() (response *api.BaseResponse[TicketVerifyQueryRecordResponse]) {
	response = api.CreateResponse[TicketVerifyQueryRecordResponse](&TicketVerifyQueryRecordResponse{})
	return
}

type TicketVerifyQueryRecordResponseResults struct {
	Code            string `json:"code,omitempty"`
	OrderNo         string `json:"orderNo,omitempty"`
	OperationClient int64  `json:"operationClient,omitempty"`
	CodeSource      int64  `json:"codeSource,omitempty"`
	Status          int64  `json:"status,omitempty"`
	GoodsName       string `json:"goodsName,omitempty"`
	ScenicName      string `json:"scenicName,omitempty"`
	ScenicImg       string `json:"scenicImg,omitempty"`
	Price           int64  `json:"price,omitempty"`
	Receiver        string `json:"receiver,omitempty"`
	Phone           string `json:"phone,omitempty"`
	NumberOfPeople  int64  `json:"numberOfPeople,omitempty"`
	VisitDate       string `json:"visitDate,omitempty"`
	CreateTime      string `json:"createTime,omitempty"`
}
