package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GuideGetLastRelationLog
// @id 1876
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询最近一条有效临时关系链)
func (client *Client) GuideGetLastRelationLog(request *GuideGetLastRelationLogRequest) (response *GuideGetLastRelationLogResponse, err error) {
	rpcResponse := CreateGuideGetLastRelationLogResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GuideGetLastRelationLogRequest struct {
	*api.BaseRequest

	CustomerWid   int64  `json:"customerWid,omitempty"`
	CustomerPhone string `json:"customerPhone,omitempty"`
}

type GuideGetLastRelationLogResponse struct {
	GuiderWid       int64  `json:"guiderWid,omitempty"`
	GuiderPhone     string `json:"guiderPhone,omitempty"`
	GuiderName      string `json:"guiderName,omitempty"`
	CustomerWid     int64  `json:"customerWid,omitempty"`
	JobNumber       string `json:"jobNumber,omitempty"`
	IsUsed          int    `json:"isUsed,omitempty"`
	GzhQrCodePicUrl string `json:"gzhQrCodePicUrl,omitempty"`
	XcxQrCodePicUrl string `json:"xcxQrCodePicUrl,omitempty"`
	BindSceneType   int    `json:"bindSceneType,omitempty"`
	BindSceneDec    string `json:"bindSceneDec,omitempty"`
	StoreId         int64  `json:"storeId,omitempty"`
}

func CreateGuideGetLastRelationLogRequest() (request *GuideGetLastRelationLogRequest) {
	request = &GuideGetLastRelationLogRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "guide/getLastRelationLog", "api")
	request.Method = api.POST
	return
}

func CreateGuideGetLastRelationLogResponse() (response *api.BaseResponse[GuideGetLastRelationLogResponse]) {
	response = api.CreateResponse[GuideGetLastRelationLogResponse](&GuideGetLastRelationLogResponse{})
	return
}
