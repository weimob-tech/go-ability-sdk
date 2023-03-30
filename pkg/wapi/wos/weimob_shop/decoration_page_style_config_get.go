package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// DecorationPageStyleConfigGet
// @id 2197
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2197?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询分类模板样式信息)
func (client *Client) DecorationPageStyleConfigGet(request *DecorationPageStyleConfigGetRequest) (response *DecorationPageStyleConfigGetResponse, err error) {
	rpcResponse := CreateDecorationPageStyleConfigGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type DecorationPageStyleConfigGetRequest struct {
	*api.BaseRequest

	BasicInfo DecorationPageStyleConfigGetRequestBasicInfo `json:"basicInfo,omitempty"`
	PageType  int64                                        `json:"pageType,omitempty"`
}

type DecorationPageStyleConfigGetResponse struct {
	StyleExData                      DecorationPageStyleConfigGetResponseStyleExData                        `json:"styleExData,omitempty"`
	PageStyleCategoryNavigationOutVO []DecorationPageStyleConfigGetResponsePageStyleCategoryNavigationOutVO `json:"PageStyleCategoryNavigationOutVO,omitempty"`
	BosId                            int64                                                                  `json:"bosId,omitempty"`
	CategoryTemplateADStatus         int64                                                                  `json:"categoryTemplateADStatus,omitempty"`
	Cid                              int64                                                                  `json:"cid,omitempty"`
	Description                      string                                                                 `json:"description,omitempty"`
	PageType                         int64                                                                  `json:"pageType,omitempty"`
	StyleType                        int64                                                                  `json:"styleType,omitempty"`
	Title                            string                                                                 `json:"title,omitempty"`
	Vid                              int64                                                                  `json:"vid,omitempty"`
}

func CreateDecorationPageStyleConfigGetRequest() (request *DecorationPageStyleConfigGetRequest) {
	request = &DecorationPageStyleConfigGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "decoration/page/style/config/get", "apigw")
	request.Method = api.POST
	return
}

func CreateDecorationPageStyleConfigGetResponse() (response *api.BaseResponse[DecorationPageStyleConfigGetResponse]) {
	response = api.CreateResponse[DecorationPageStyleConfigGetResponse](&DecorationPageStyleConfigGetResponse{})
	return
}

type DecorationPageStyleConfigGetRequestBasicInfo struct {
	BosId int64 `json:"bosId,omitempty"`
	Vid   int64 `json:"vid,omitempty"`
	Cid   int64 `json:"cid,omitempty"`
}

type DecorationPageStyleConfigGetResponseStyleExData struct {
	BackColor           string `json:"backColor,omitempty"`
	ContentColor        string `json:"contentColor,omitempty"`
	ImageSize           int64  `json:"imageSize,omitempty"`
	ShowCart            int64  `json:"showCart,omitempty"`
	ShowGoodsSaleNum    int64  `json:"showGoodsSaleNum,omitempty"`
	ShowGoodsTag        int64  `json:"showGoodsTag,omitempty"`
	ShowImage           int64  `json:"showImage,omitempty"`
	TopNavBarColorStyle int64  `json:"topNavBarColorStyle,omitempty"`
}

type DecorationPageStyleConfigGetResponsePageStyleCategoryNavigationOutVO struct {
	Title      string `json:"title,omitempty"`
	ModuleSign string `json:"moduleSign,omitempty"`
	ModuleJson string `json:"moduleJson,omitempty"`
}
