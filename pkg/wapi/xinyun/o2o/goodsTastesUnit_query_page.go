package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsTastesUnitQueryPage
// @id 154
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询口味)
func (client *Client) GoodsTastesUnitQueryPage(request *GoodsTastesUnitQueryPageRequest) (response *GoodsTastesUnitQueryPageResponse, err error) {
	rpcResponse := CreateGoodsTastesUnitQueryPageResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsTastesUnitQueryPageRequest struct {
	*api.BaseRequest

	PageNum  int `json:"pageNum,omitempty"`
	PageSize int `json:"pageSize,omitempty"`
}

type GoodsTastesUnitQueryPageResponse struct {
}

func CreateGoodsTastesUnitQueryPageRequest() (request *GoodsTastesUnitQueryPageRequest) {
	request = &GoodsTastesUnitQueryPageRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "goodsTastesUnit/queryPage", "api")
	request.Method = api.POST
	return
}

func CreateGoodsTastesUnitQueryPageResponse() (response *api.BaseResponse[GoodsTastesUnitQueryPageResponse]) {
	response = api.CreateResponse[GoodsTastesUnitQueryPageResponse](&GoodsTastesUnitQueryPageResponse{})
	return
}
