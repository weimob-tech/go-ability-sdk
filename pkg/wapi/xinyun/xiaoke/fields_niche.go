package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FieldsNiche
// @id 1742
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询商机字段列表)
func (client *Client) FieldsNiche(request *FieldsNicheRequest) (response *FieldsNicheResponse, err error) {
	rpcResponse := CreateFieldsNicheResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FieldsNicheRequest struct {
	*api.BaseRequest

	Wid int64 `json:"wid,omitempty"`
}

type FieldsNicheResponse struct {
}

func CreateFieldsNicheRequest() (request *FieldsNicheRequest) {
	request = &FieldsNicheRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "fields/niche", "api")
	request.Method = api.POST
	return
}

func CreateFieldsNicheResponse() (response *api.BaseResponse[FieldsNicheResponse]) {
	response = api.CreateResponse[FieldsNicheResponse](&FieldsNicheResponse{})
	return
}
