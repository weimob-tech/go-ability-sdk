package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StoreAddStore
// @id 49
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=添加门店)
func (client *Client) StoreAddStore(request *StoreAddStoreRequest) (response *StoreAddStoreResponse, err error) {
	rpcResponse := CreateStoreAddStoreResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StoreAddStoreRequest struct {
	*api.BaseRequest

	StoreName          string  `json:"storeName,omitempty"`
	BusinessTypeParent string  `json:"businessTypeParent,omitempty"`
	BusinessType       string  `json:"businessType,omitempty"`
	BusinessTypeSon    string  `json:"businessTypeSon,omitempty"`
	Phone              string  `json:"phone,omitempty"`
	Province           string  `json:"province,omitempty"`
	City               string  `json:"city,omitempty"`
	District           string  `json:"district,omitempty"`
	Address            string  `json:"address,omitempty"`
	BranchName         string  `json:"branchName,omitempty"`
	BusinessHourStart  string  `json:"businessHourStart,omitempty"`
	BusinessHourEnd    string  `json:"businessHourEnd,omitempty"`
	Brief              string  `json:"brief,omitempty"`
	Recommend          string  `json:"recommend,omitempty"`
	SpecialService     string  `json:"specialService,omitempty"`
	PerPersonSpending  int     `json:"perPersonSpending,omitempty"`
	Pictures           string  `json:"pictures,omitempty"`
	MapLocationX       float32 `json:"mapLocationX,omitempty"`
	MapLocationY       float32 `json:"mapLocationY,omitempty"`
}

type StoreAddStoreResponse struct {
}

func CreateStoreAddStoreRequest() (request *StoreAddStoreRequest) {
	request = &StoreAddStoreRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "store/addStore", "api")
	request.Method = api.POST
	return
}

func CreateStoreAddStoreResponse() (response *api.BaseResponse[StoreAddStoreResponse]) {
	response = api.CreateResponse[StoreAddStoreResponse](&StoreAddStoreResponse{})
	return
}
