package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// AiImportNumberSync
// @id 3725
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=导入号码（同步）)
func (client *Client) AiImportNumberSync(request *AiImportNumberSyncRequest) (response *AiImportNumberSyncResponse, err error) {
	rpcResponse := CreateAiImportNumberSyncResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type AiImportNumberSyncRequest struct {
	*api.BaseRequest

	Phones AiImportNumberSyncRequestPhones `json:"phones,omitempty"`
	TaskId int64                           `json:"taskId,omitempty"`
}

type AiImportNumberSyncResponse struct {
	PlaceFailDetail AiImportNumberSyncResponsePlaceFailDetail `json:"placeFailDetail,omitempty"`
	Total           int64                                     `json:"total,omitempty"`
	SuccessNum      int64                                     `json:"successNum,omitempty"`
	PlaceFailNum    int64                                     `json:"placeFailNum,omitempty"`
	Phone           string                                    `json:"phone,omitempty"`
	CustomerName    string                                    `json:"customerName,omitempty"`
	ErrorMsg        string                                    `json:"errorMsg,omitempty"`
}

func CreateAiImportNumberSyncRequest() (request *AiImportNumberSyncRequest) {
	request = &AiImportNumberSyncRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "ai/importNumberSync", "api")
	request.Method = api.POST
	return
}

func CreateAiImportNumberSyncResponse() (response *api.BaseResponse[AiImportNumberSyncResponse]) {
	response = api.CreateResponse[AiImportNumberSyncResponse](&AiImportNumberSyncResponse{})
	return
}

type AiImportNumberSyncRequestPhones struct {
	CustomerName string `json:"customerName,omitempty"`
	CompanyName  string `json:"companyName,omitempty"`
	Phone        string `json:"phone,omitempty"`
}

type AiImportNumberSyncResponsePlaceFailDetail struct {
}
