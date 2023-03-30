package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StoreEditStore
// @id 520
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新酒店)
func (client *Client) StoreEditStore(request *StoreEditStoreRequest) (response *StoreEditStoreResponse, err error) {
	rpcResponse := CreateStoreEditStoreResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StoreEditStoreRequest struct {
	*api.BaseRequest

	StoreName         string   `json:"storeName,omitempty"`
	BranchName        string   `json:"branchName,omitempty"`
	CountryId         string   `json:"countryId,omitempty"`
	Country           string   `json:"country,omitempty"`
	ProvinceId        string   `json:"provinceId,omitempty"`
	Province          string   `json:"province,omitempty"`
	CityId            string   `json:"cityId,omitempty"`
	City              string   `json:"city,omitempty"`
	DistrictId        string   `json:"districtId,omitempty"`
	District          string   `json:"district,omitempty"`
	Address           string   `json:"address,omitempty"`
	TencentLongitude  float32  `json:"tencentLongitude,omitempty"`
	TencentLatitude   float32  `json:"tencentLatitude,omitempty"`
	LogoList          []string `json:"logoList,omitempty"`
	Phone             string   `json:"phone,omitempty"`
	Introduction      string   `json:"introduction,omitempty"`
	Notice            string   `json:"notice,omitempty"`
	TrafficInfo       string   `json:"trafficInfo,omitempty"`
	MatingServiceList []int    `json:"matingServiceList,omitempty"`
	StoreId           int64    `json:"storeId,omitempty"`
}

type StoreEditStoreResponse struct {
}

func CreateStoreEditStoreRequest() (request *StoreEditStoreRequest) {
	request = &StoreEditStoreRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "store/editStore", "api")
	request.Method = api.POST
	return
}

func CreateStoreEditStoreResponse() (response *api.BaseResponse[StoreEditStoreResponse]) {
	response = api.CreateResponse[StoreEditStoreResponse](&StoreEditStoreResponse{})
	return
}
