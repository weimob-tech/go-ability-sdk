package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GuideBindGuideRelationship
// @id 1076
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=绑定导购关系链)
func (client *Client) GuideBindGuideRelationship(request *GuideBindGuideRelationshipRequest) (response *GuideBindGuideRelationshipResponse, err error) {
	rpcResponse := CreateGuideBindGuideRelationshipResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GuideBindGuideRelationshipRequest struct {
	*api.BaseRequest

	GuiderPhone string `json:"guiderPhone,omitempty"`
	GuiderWid   int64  `json:"guiderWid,omitempty"`
	JobNumber   string `json:"jobNumber,omitempty"`
	CustomerWid int64  `json:"customerWid,omitempty"`
}

type GuideBindGuideRelationshipResponse struct {
}

func CreateGuideBindGuideRelationshipRequest() (request *GuideBindGuideRelationshipRequest) {
	request = &GuideBindGuideRelationshipRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "guide/bindGuideRelationship", "api")
	request.Method = api.POST
	return
}

func CreateGuideBindGuideRelationshipResponse() (response *api.BaseResponse[GuideBindGuideRelationshipResponse]) {
	response = api.CreateResponse[GuideBindGuideRelationshipResponse](&GuideBindGuideRelationshipResponse{})
	return
}
