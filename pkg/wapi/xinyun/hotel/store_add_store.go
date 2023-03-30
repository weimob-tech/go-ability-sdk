package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StoreAddStore
// @id 517
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新增酒店)
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
}

type StoreAddStoreResponse struct {
}

func CreateStoreAddStoreRequest() (request *StoreAddStoreRequest) {
	request = &StoreAddStoreRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "store/addStore", "api")
	request.Method = api.POST
	return
}

func CreateStoreAddStoreResponse() (response *api.BaseResponse[StoreAddStoreResponse]) {
	response = api.CreateResponse[StoreAddStoreResponse](&StoreAddStoreResponse{})
	return
}
