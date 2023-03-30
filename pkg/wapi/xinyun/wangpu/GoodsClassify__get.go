package wangpu

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsClassifyGet
// @id 4
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取商品分组)
func (client *Client) GoodsClassifyGet(request *GoodsClassifyGetRequest) (response *GoodsClassifyGetResponse, err error) {
	rpcResponse := CreateGoodsClassifyGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsClassifyGetRequest struct {
	*api.BaseRequest

	ClassifyPid int64 `json:"classify_pid,omitempty"`
	PageNo      int   `json:"page_no,omitempty"`
	PageSize    int   `json:"page_size,omitempty"`
}

type GoodsClassifyGetResponse struct {
}

func CreateGoodsClassifyGetRequest() (request *GoodsClassifyGetRequest) {
	request = &GoodsClassifyGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wangpu", "1_0", "GoodsClassify/Get", "api")
	request.Method = api.POST
	return
}

func CreateGoodsClassifyGetResponse() (response *api.BaseResponse[GoodsClassifyGetResponse]) {
	response = api.CreateResponse[GoodsClassifyGetResponse](&GoodsClassifyGetResponse{})
	return
}
