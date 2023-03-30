package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FieldsClue
// @id 1646
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询线索字段列表)
func (client *Client) FieldsClue(request *FieldsClueRequest) (response *FieldsClueResponse, err error) {
	rpcResponse := CreateFieldsClueResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FieldsClueRequest struct {
	*api.BaseRequest

	Wid int64 `json:"wid,omitempty"`
}

type FieldsClueResponse struct {
}

func CreateFieldsClueRequest() (request *FieldsClueRequest) {
	request = &FieldsClueRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "fields/clue", "api")
	request.Method = api.POST
	return
}

func CreateFieldsClueResponse() (response *api.BaseResponse[FieldsClueResponse]) {
	response = api.CreateResponse[FieldsClueResponse](&FieldsClueResponse{})
	return
}
