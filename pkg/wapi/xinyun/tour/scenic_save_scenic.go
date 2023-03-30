package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ScenicSaveScenic
// @id 1103
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新增/修改景区产品)
func (client *Client) ScenicSaveScenic(request *ScenicSaveScenicRequest) (response *ScenicSaveScenicResponse, err error) {
	rpcResponse := CreateScenicSaveScenicResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ScenicSaveScenicRequest struct {
	*api.BaseRequest

	ScenicCode    string   `json:"scenicCode,omitempty"`
	ScenicName    string   `json:"scenicName,omitempty"`
	Status        int      `json:"status,omitempty"`
	Level         string   `json:"level,omitempty"`
	CityCode      string   `json:"cityCode,omitempty"`
	Address       string   `json:"address,omitempty"`
	OpenTime      string   `json:"openTime,omitempty"`
	Contact       string   `json:"contact,omitempty"`
	Intro         string   `json:"intro,omitempty"`
	AroundTraffic string   `json:"aroundTraffic,omitempty"`
	AroundRepast  string   `json:"aroundRepast,omitempty"`
	GoodsImages   []string `json:"goodsImages,omitempty"`
}

type ScenicSaveScenicResponse struct {
}

func CreateScenicSaveScenicRequest() (request *ScenicSaveScenicRequest) {
	request = &ScenicSaveScenicRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "scenic/saveScenic", "api")
	request.Method = api.POST
	return
}

func CreateScenicSaveScenicResponse() (response *api.BaseResponse[ScenicSaveScenicResponse]) {
	response = api.CreateResponse[ScenicSaveScenicResponse](&ScenicSaveScenicResponse{})
	return
}
