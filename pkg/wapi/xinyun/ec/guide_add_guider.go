package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GuideAddGuider
// @id 1177
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新增导购)
func (client *Client) GuideAddGuider(request *GuideAddGuiderRequest) (response *GuideAddGuiderResponse, err error) {
	rpcResponse := CreateGuideAddGuiderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GuideAddGuiderRequest struct {
	*api.BaseRequest

	Data              GuideAddGuiderRequestData `json:"data,omitempty"`
	StoreNumber       string                    `json:"storeNumber,omitempty"`
	GuiderPhone       string                    `json:"guiderPhone,omitempty"`
	GuiderName        string                    `json:"guiderName,omitempty"`
	JobNumber         string                    `json:"jobNumber,omitempty"`
	IsUsed            int                       `json:"isUsed,omitempty"`
	IsExclusiveCus    int                       `json:"isExclusiveCus,omitempty"`
	Zone              string                    `json:"zone,omitempty"`
	PhysicalStoreCode string                    `json:"physicalStoreCode,omitempty"`
}

type GuideAddGuiderResponse struct {
}

func CreateGuideAddGuiderRequest() (request *GuideAddGuiderRequest) {
	request = &GuideAddGuiderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "guide/addGuider", "api")
	request.Method = api.POST
	return
}

func CreateGuideAddGuiderResponse() (response *api.BaseResponse[GuideAddGuiderResponse]) {
	response = api.CreateResponse[GuideAddGuiderResponse](&GuideAddGuiderResponse{})
	return
}

type GuideAddGuiderRequestData struct {
}
