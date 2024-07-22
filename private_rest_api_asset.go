package mybybitapi

// bybit AssetTransferQueryInterTransferList PrivateRest接口 GET 查詢劃轉紀錄 (單帳號內)
func (client *PrivateRestClient) NewAssetTransferQueryInterTransferList() *AssetTransferQueryInterTransferListAPI {
	return &AssetTransferQueryInterTransferListAPI{
		client: client,
		req:    &AssetTransferQueryInterTransferListReq{},
	}
}
func (api *AssetTransferQueryInterTransferListAPI) Do() (*BybitRestRes[AssetTransferQueryInterTransferListRes], error) {
	url := bybitHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[AssetTransferQueryInterTransferList])
	return bybitCallAPIWithSecret[AssetTransferQueryInterTransferListRes](api.client.c, url, NIL_REQBODY, GET)
}

// bybit AssetTransferQueryTransferCoinList:  PrivateRest接口 //GET 帳戶類型間可劃轉的幣種
func (client *PrivateRestClient) NewAssetTransferQueryTransferCoinList() *AssetTransferQueryTransferCoinListAPI {
	return &AssetTransferQueryTransferCoinListAPI{
		client: client,
		req:    &AssetTransferQueryTransferCoinListReq{},
	}
}
func (api *AssetTransferQueryTransferCoinListAPI) Do() (*BybitRestRes[AssetTransferQueryTransferCoinListRes], error) {
	url := bybitHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[AssetTransferQueryTransferCoinList])
	return bybitCallAPIWithSecret[AssetTransferQueryTransferCoinListRes](api.client.c, url, NIL_REQBODY, GET)
}

// bybit AssetTransferInterTransfer:  PrivateRest接口 //POST 劃轉 (單帳號內)
func (client *PrivateRestClient) NewAssetTransferInterTransfer() *AssetTransferInterTransferAPI {
	return &AssetTransferInterTransferAPI{
		client: client,
		req:    &AssetTransferInterTransferReq{},
	}
}
func (api *AssetTransferInterTransferAPI) Do() (*BybitRestRes[AssetTransferInterTransferRes], error) {
	url := bybitHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[AssetTransferInterTransfer])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return bybitCallAPIWithSecret[AssetTransferInterTransferRes](api.client.c, url, reqBody, POST)

}
