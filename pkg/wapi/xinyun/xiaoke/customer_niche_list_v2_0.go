package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerNicheList
// @id 2670
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=客户下关联商机列表)
func (client *Client) CustomerNicheListV2_0(request *CustomerNicheListRequestV2_0) (response *CustomerNicheListResponseV2_0, err error) {
	rpcResponse := CreateCustomerNicheListResponseV2_0()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerNicheListRequestV2_0 struct {
	*api.BaseRequest

	Wid      int    `json:"wid,omitempty"`
	PageNum  int    `json:"pageNum,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
	Key      string `json:"key,omitempty"`
}

type CustomerNicheListResponseV2_0 struct {
}

func CreateCustomerNicheListRequestV2_0() (request *CustomerNicheListRequestV2_0) {
	request = &CustomerNicheListRequestV2_0{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "2_0", "customer/nicheList", "api")
	request.Method = api.POST
	return
}

func CreateCustomerNicheListResponseV2_0() (response *api.BaseResponse[CustomerNicheListResponseV2_0]) {
	response = api.CreateResponse[CustomerNicheListResponseV2_0](&CustomerNicheListResponseV2_0{})
	return
}
