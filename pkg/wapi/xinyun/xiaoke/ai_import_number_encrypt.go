package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// AiImportNumberEncrypt
// @id 3926
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=号码导入加密接口)
func (client *Client) AiImportNumberEncrypt(request *AiImportNumberEncryptRequest) (response *AiImportNumberEncryptResponse, err error) {
	rpcResponse := CreateAiImportNumberEncryptResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type AiImportNumberEncryptRequest struct {
	*api.BaseRequest

	Phones AiImportNumberEncryptRequestPhones `json:"phones,omitempty"`
	TaskId int64                              `json:"taskId,omitempty"`
}

type AiImportNumberEncryptResponse struct {
	Total           int64   `json:"total,omitempty"`
	SuccessNum      int64   `json:"successNum,omitempty"`
	PlaceFailDetail []int64 `json:"placeFailDetail,omitempty"`
	PlaceFailNum    int64   `json:"placeFailNum,omitempty"`
}

func CreateAiImportNumberEncryptRequest() (request *AiImportNumberEncryptRequest) {
	request = &AiImportNumberEncryptRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "ai/importNumberEncrypt", "api")
	request.Method = api.POST
	return
}

func CreateAiImportNumberEncryptResponse() (response *api.BaseResponse[AiImportNumberEncryptResponse]) {
	response = api.CreateResponse[AiImportNumberEncryptResponse](&AiImportNumberEncryptResponse{})
	return
}

type AiImportNumberEncryptRequestPhones struct {
	Phone        string `json:"phone,omitempty"`
	CompanyName  string `json:"companyName,omitempty"`
	CustomerName string `json:"customerName,omitempty"`
}
