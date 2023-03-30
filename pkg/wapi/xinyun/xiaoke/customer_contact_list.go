package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerContactList
// @id 1712
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=客户下关联联系人列表)
func (client *Client) CustomerContactList(request *CustomerContactListRequest) (response *CustomerContactListResponse, err error) {
	rpcResponse := CreateCustomerContactListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerContactListRequest struct {
	*api.BaseRequest

	Wid      int    `json:"wid,omitempty"`
	PageNum  int    `json:"pageNum,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
	Key      string `json:"key,omitempty"`
}

type CustomerContactListResponse struct {
}

func CreateCustomerContactListRequest() (request *CustomerContactListRequest) {
	request = &CustomerContactListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "customer/contactList", "api")
	request.Method = api.POST
	return
}

func CreateCustomerContactListResponse() (response *api.BaseResponse[CustomerContactListResponse]) {
	response = api.CreateResponse[CustomerContactListResponse](&CustomerContactListResponse{})
	return
}
