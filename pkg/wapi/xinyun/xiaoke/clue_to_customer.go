package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ClueToCustomer
// @id 1722
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=线索转客户)
func (client *Client) ClueToCustomer(request *ClueToCustomerRequest) (response *ClueToCustomerResponse, err error) {
	rpcResponse := CreateClueToCustomerResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ClueToCustomerRequest struct {
	*api.BaseRequest

	CustomerInfo ClueToCustomerRequestCustomerInfo `json:"customerInfo,omitempty"`
	ContactsInfo ClueToCustomerRequestContactsInfo `json:"contactsInfo,omitempty"`
	NicheInfo    ClueToCustomerRequestNicheInfo    `json:"nicheInfo,omitempty"`
	ClueKey      string                            `json:"clueKey,omitempty"`
	Wid          int64                             `json:"wid,omitempty"`
}

type ClueToCustomerResponse struct {
}

func CreateClueToCustomerRequest() (request *ClueToCustomerRequest) {
	request = &ClueToCustomerRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "clue/toCustomer", "api")
	request.Method = api.POST
	return
}

func CreateClueToCustomerResponse() (response *api.BaseResponse[ClueToCustomerResponse]) {
	response = api.CreateResponse[ClueToCustomerResponse](&ClueToCustomerResponse{})
	return
}

type ClueToCustomerRequestCustomerInfo struct {
}

type ClueToCustomerRequestContactsInfo struct {
}

type ClueToCustomerRequestNicheInfo struct {
}
