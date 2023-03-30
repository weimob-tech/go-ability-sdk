package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GuideTransferGuider
// @id 1595
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=导购切换门店)
func (client *Client) GuideTransferGuider(request *GuideTransferGuiderRequest) (response *GuideTransferGuiderResponse, err error) {
	rpcResponse := CreateGuideTransferGuiderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GuideTransferGuiderRequest struct {
	*api.BaseRequest

	TargetStoreId     int64  `json:"targetStoreId,omitempty"`
	TargetStoreNumber string `json:"targetStoreNumber,omitempty"`
	GuiderPhone       string `json:"guiderPhone,omitempty"`
	GuiderWid         int64  `json:"guiderWid,omitempty"`
	KeepCustomers     int64  `json:"keepCustomers,omitempty"`
	PhysicalStoreCode string `json:"physicalStoreCode,omitempty"`
}

type GuideTransferGuiderResponse struct {
}

func CreateGuideTransferGuiderRequest() (request *GuideTransferGuiderRequest) {
	request = &GuideTransferGuiderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "guide/transferGuider", "api")
	request.Method = api.POST
	return
}

func CreateGuideTransferGuiderResponse() (response *api.BaseResponse[GuideTransferGuiderResponse]) {
	response = api.CreateResponse[GuideTransferGuiderResponse](&GuideTransferGuiderResponse{})
	return
}
