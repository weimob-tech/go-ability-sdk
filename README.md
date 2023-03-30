# go-ability-sdk
微盟云开放平台的GO项目的能力SDK，根据开放能力的元数据规范，提供一套GO语言的能力定义，帮助开发者减少开发成本，使其可以快速接入微盟云的开放生态。

## 介绍
### 功能列表
* API能力定义
	* Client
	* Request
	* Response
* SPI能力定义
	* Service
	* Request
	* Response
* MSG能力定义
	* Service
	* Message

### 项目结构

```
|-- .editorconfig
|-- .gitignore
|-- go.mod
|-- LICENSE.MD
|-- pkg/                         # 源码目录
|   |-- sdk/       
|       |-- api/    #API数据结构及抽象
|       |-- auth/   #授权数据结构及抽象    
|       |-- msg/    #MSG数据结构及抽象
|       |-- spi/    #SPI数据结构及抽象
|   |-- wapi/       # API能力定义
|       |-- wos/                 # Wos开放能力
|       |-- xinyun/              # Xinyun开放能力
|   |-- wspi/       # SPI能力定义
|       |-- wos/                 # Wos开放能力
|       |-- xinyun/              # Xinyun开放能力
|   |-- wmsg/       # MSG能力定义
|       |-- wos/                 # Wos开放能力
|       |-- xinyun/              # Xinyun开放能力
```

## 快速开始
1. 使用go module管理包，在go项目工程的go.module添加依赖

   ``` json
	  require (
		     github.com/weimob-tech/go-ability-sdk v0.0.0
	  )
   ```
2. 安装包，使用go命令
	* `GO111MODULE=on GOPROXY=https://goproxy.cn,direct go mod tidy`
3. Api调用
  ```go
  //API调用的逻辑
  func testApi(shopId, shopType string){
	  clientInfo := auth.GetClientInfo("product-a")
	  client, err := weimob_shop.NewClientWithClientKey(clientInfo.ClientId, clientInfo.ClientSecret)
	  if err != nil {
		  panic(err)
	  }

	  req := weimob_shop.CreateCommentGoodsGetRequest()
	  req.ShopId = shopId
	  req.ShopType = shopType
	  req.QueryParameter = weimob_shop.CommentGoodsGetRequestQueryParameter{CommentIds: []int64{1, 2}}
	  req.BasicInfo = weimob_shop.CommentGoodsGetRequestBasicInfo{Vid: 55}
	  res, err := client.CommentGoodsGet(req)

	  if err != nil {
		  panic(err)
	  }
	  wlog.I().Msgf("api response: %v", codec.ToJson(res))
  }
  ```
4. Spi实现

   ```go
   //实现SPI的逻辑
   func main() {
         shop, err := weimob_shop.NewService()
         if err != nil {
              panic(err)
         }
         shop.OnPaasWeimobShopCouponPaasBatchLockCouponServiceInvocation(
             func(ctx context.Context, request *spi.WosInvocationRequest[weimob_shop.PaasWeimobShopCouponPaasBatchLockCouponRequest]) (response spi.InvocationResponse[weimob_shop.PaasWeimobShopCouponPaasBatchLockCouponResponse], err error) {
                 wlog.I().Msgf("got request: %v", codec.ToJson(request))
                 ret := weimob_shop.PaasWeimobShopCouponPaasBatchLockCouponResponse{Success: true}
                 return spi.Ok(&ret), nil
         })
    }
   ``` 
5. Msg实现

  ```go
  //实现消息订阅逻辑
  func main() {
	  shop, err := weimob_shop.NewService()
	  if err != nil {
		  panic(err)
	  }
	  shop.OnMallOrderMallOrderMessageEvent(
		  func(ctx context.Context, message *msg.WosOpenMessage[weimob_shop.MallOrderMallOrderMessageEvent]) (response x.Result, err error) {
			  wlog.I().Msgf("got request: %v", codec.ToJson(message))
			  event := message.GetMsgBody()
			  wlog.I().Msgf("got request body: %v", codec.ToJson(event))
			  return x.Ok(), nil
		  })
  }
  ```

## 使用文档
* [能力文档](http://doc.weimobcloud.com/list?tag=2396&menuId=19&childMenuId=1&isold=2)
* [开发者入驻](http://doc.weimobcloud.com/word?menuId=46&childMenuId=47&tag=2970&isold=2)
* [应用开发](http://doc.weimobcloud.com/word?menuId=53&childMenuId=54&tag=2488&isold=2)
* [GO工程](http://doc.weimobcloud.com/word?menuId=53&childMenuId=54&tag=2488&isold=2)

## 贡献方法
* 申请加入weimob_tech

## 联系我们
* Weimob-tech@weimob.com
