package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsSizetableDefaultparamGetList
// @id 2155
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2155?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=通过类目类型获取默认尺码表表头参数)
func (client *Client) GoodsSizetableDefaultparamGetList(request *GoodsSizetableDefaultparamGetListRequest) (response *GoodsSizetableDefaultparamGetListResponse, err error) {
	rpcResponse := CreateGoodsSizetableDefaultparamGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsSizetableDefaultparamGetListRequest struct {
	*api.BaseRequest

	BasicInfo    GoodsSizetableDefaultparamGetListRequestBasicInfo `json:"basicInfo,omitempty"`
	CategoryType int64                                             `json:"categoryType,omitempty"`
}

type GoodsSizetableDefaultparamGetListResponse struct {
	ParamName         string `json:"paramName,omitempty"`
	IsSystemParam     bool   `json:"isSystemParam,omitempty"`
	IsDefaultSelected bool   `json:"isDefaultSelected,omitempty"`
}

func CreateGoodsSizetableDefaultparamGetListRequest() (request *GoodsSizetableDefaultparamGetListRequest) {
	request = &GoodsSizetableDefaultparamGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/sizetable/defaultparam/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsSizetableDefaultparamGetListResponse() (response *api.BaseResponse[GoodsSizetableDefaultparamGetListResponse]) {
	response = api.CreateResponse[GoodsSizetableDefaultparamGetListResponse](&GoodsSizetableDefaultparamGetListResponse{})
	return
}

type GoodsSizetableDefaultparamGetListRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
