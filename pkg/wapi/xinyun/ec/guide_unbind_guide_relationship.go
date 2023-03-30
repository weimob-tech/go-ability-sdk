package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GuideUnbindGuideRelationship
// @id 1071
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=解绑导购关系链)
func (client *Client) GuideUnbindGuideRelationship(request *GuideUnbindGuideRelationshipRequest) (response *GuideUnbindGuideRelationshipResponse, err error) {
	rpcResponse := CreateGuideUnbindGuideRelationshipResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GuideUnbindGuideRelationshipRequest struct {
	*api.BaseRequest

	GuiderPhone   string `json:"guiderPhone,omitempty"`
	CustomerPhone string `json:"customerPhone,omitempty"`
	CustomerWid   string `json:"customerWid,omitempty"`
	GuiderWid     int64  `json:"guiderWid,omitempty"`
	JobNumber     string `json:"jobNumber,omitempty"`
}

type GuideUnbindGuideRelationshipResponse struct {
}

func CreateGuideUnbindGuideRelationshipRequest() (request *GuideUnbindGuideRelationshipRequest) {
	request = &GuideUnbindGuideRelationshipRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "guide/unbindGuideRelationship", "api")
	request.Method = api.POST
	return
}

func CreateGuideUnbindGuideRelationshipResponse() (response *api.BaseResponse[GuideUnbindGuideRelationshipResponse]) {
	response = api.CreateResponse[GuideUnbindGuideRelationshipResponse](&GuideUnbindGuideRelationshipResponse{})
	return
}
