package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MerchantUpdateStoreSortNum
// @id 1284
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=批量修改门店排序值)
func (client *Client) MerchantUpdateStoreSortNum(request *MerchantUpdateStoreSortNumRequest) (response *MerchantUpdateStoreSortNumResponse, err error) {
	rpcResponse := CreateMerchantUpdateStoreSortNumResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MerchantUpdateStoreSortNumRequest struct {
	*api.BaseRequest

	Storelist []map[string]any `json:"storelist,omitempty"`
	SortNum   int              `json:"sortNum,omitempty"`
}

type MerchantUpdateStoreSortNumResponse struct {
}

func CreateMerchantUpdateStoreSortNumRequest() (request *MerchantUpdateStoreSortNumRequest) {
	request = &MerchantUpdateStoreSortNumRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "merchant/updateStoreSortNum", "api")
	request.Method = api.POST
	return
}

func CreateMerchantUpdateStoreSortNumResponse() (response *api.BaseResponse[MerchantUpdateStoreSortNumResponse]) {
	response = api.CreateResponse[MerchantUpdateStoreSortNumResponse](&MerchantUpdateStoreSortNumResponse{})
	return
}
