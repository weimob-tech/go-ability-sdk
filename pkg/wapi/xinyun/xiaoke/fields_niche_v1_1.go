package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FieldsNiche
// @id 2672
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询商机字段列表)
func (client *Client) FieldsNicheV1_1(request *FieldsNicheRequestV1_1) (response *FieldsNicheResponseV1_1, err error) {
	rpcResponse := CreateFieldsNicheResponseV1_1()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FieldsNicheRequestV1_1 struct {
	*api.BaseRequest

	Wid int64 `json:"wid,omitempty"`
}

type FieldsNicheResponseV1_1 struct {
}

func CreateFieldsNicheRequestV1_1() (request *FieldsNicheRequestV1_1) {
	request = &FieldsNicheRequestV1_1{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_1", "fields/niche", "api")
	request.Method = api.POST
	return
}

func CreateFieldsNicheResponseV1_1() (response *api.BaseResponse[FieldsNicheResponseV1_1]) {
	response = api.CreateResponse[FieldsNicheResponseV1_1](&FieldsNicheResponseV1_1{})
	return
}
