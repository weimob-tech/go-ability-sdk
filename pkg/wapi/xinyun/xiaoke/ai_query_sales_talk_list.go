package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// AiQuerySalesTalkList
// @id 3739
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=AI外呼查询话术列表)
func (client *Client) AiQuerySalesTalkList(request *AiQuerySalesTalkListRequest) (response *AiQuerySalesTalkListResponse, err error) {
	rpcResponse := CreateAiQuerySalesTalkListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type AiQuerySalesTalkListRequest struct {
	*api.BaseRequest
}

type AiQuerySalesTalkListResponse struct {
	SalesTalkDtoList AiQuerySalesTalkListResponseSalesTalkDtoList `json:"salesTalkDtoList,omitempty"`
}

func CreateAiQuerySalesTalkListRequest() (request *AiQuerySalesTalkListRequest) {
	request = &AiQuerySalesTalkListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "ai/querySalesTalkList", "api")
	request.Method = api.POST
	return
}

func CreateAiQuerySalesTalkListResponse() (response *api.BaseResponse[AiQuerySalesTalkListResponse]) {
	response = api.CreateResponse[AiQuerySalesTalkListResponse](&AiQuerySalesTalkListResponse{})
	return
}

type AiQuerySalesTalkListResponseSalesTalkDtoList struct {
	ExtId   string `json:"extId,omitempty"`
	Source  int64  `json:"source,omitempty"`
	Id      int64  `json:"id,omitempty"`
	ExtName string `json:"extName,omitempty"`
	Status  int64  `json:"status,omitempty"`
}
