package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GuideTransferGuideRelationship
// @id 1077
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=换绑导购关系链)
func (client *Client) GuideTransferGuideRelationship(request *GuideTransferGuideRelationshipRequest) (response *GuideTransferGuideRelationshipResponse, err error) {
	rpcResponse := CreateGuideTransferGuideRelationshipResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GuideTransferGuideRelationshipRequest struct {
	*api.BaseRequest

	GuiderPhone string `json:"guiderPhone,omitempty"`
	GuiderWid   string `json:"guiderWid,omitempty"`
	JobNumber   string `json:"jobNumber,omitempty"`
	CustomerWid int64  `json:"customerWid,omitempty"`
}

type GuideTransferGuideRelationshipResponse struct {
}

func CreateGuideTransferGuideRelationshipRequest() (request *GuideTransferGuideRelationshipRequest) {
	request = &GuideTransferGuideRelationshipRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "guide/transferGuideRelationship", "api")
	request.Method = api.POST
	return
}

func CreateGuideTransferGuideRelationshipResponse() (response *api.BaseResponse[GuideTransferGuideRelationshipResponse]) {
	response = api.CreateResponse[GuideTransferGuideRelationshipResponse](&GuideTransferGuideRelationshipResponse{})
	return
}
