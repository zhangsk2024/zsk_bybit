package mybybitapi

type AssetTransferQueryInterTransferListReq struct {
	TransferId *string `json:"transferId,omitempty"` //String	false	string	UUID. 使用創建劃轉時用的UUID
	Coin       *string `json:"coin,omitempty"`       //String	false	string	幣種
	Status     *string `json:"status,omitempty"`     //String	false	string	劃轉狀態
	StartTime  *int    `json:"startTime,omitempty"`  //String	false	integer	開始時間戳 (毫秒) 注意: 實際查詢時是秒級維度生效	當startTime & endTime都不傳入時, API默認返回30天的數據
	EndTime    *int    `json:"endTime,omitempty"`    //String	false	integer	結束時間戳 (毫秒) 注意: 實際查詢時是秒級維度生效
	Limit      *int    `json:"limit,omitempty"`      //String	false	integer	每頁數量限制. [1, 50]. 默認: 20
	Cursor     *string `json:"cursor,omitempty"`     //String	false	string	游標，用於翻頁
}

type AssetTransferQueryInterTransferListAPI struct {
	client *PrivateRestClient
	req    *AssetTransferQueryInterTransferListReq
}

func (api *AssetTransferQueryInterTransferListAPI) TransferId(transferId string) *AssetTransferQueryInterTransferListAPI {
	api.req.TransferId = GetPointer(transferId)
	return api
}

func (api *AssetTransferQueryInterTransferListAPI) Coin(coin string) *AssetTransferQueryInterTransferListAPI {
	api.req.Coin = GetPointer(coin)
	return api
}

func (api *AssetTransferQueryInterTransferListAPI) Status(status string) *AssetTransferQueryInterTransferListAPI {
	api.req.Status = GetPointer(status)
	return api
}

func (api *AssetTransferQueryInterTransferListAPI) StartTime(startTime int) *AssetTransferQueryInterTransferListAPI {
	api.req.StartTime = GetPointer(startTime)
	return api
}

func (api *AssetTransferQueryInterTransferListAPI) EndTime(endTime int) *AssetTransferQueryInterTransferListAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}

func (api *AssetTransferQueryInterTransferListAPI) Limit(limit int) *AssetTransferQueryInterTransferListAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

func (api *AssetTransferQueryInterTransferListAPI) Cursor(cursor string) *AssetTransferQueryInterTransferListAPI {
	api.req.Cursor = GetPointer(cursor)
	return api
}

type AssetTransferQueryTransferCoinListReq struct {
	FromAccountType *string `json:"fromAccountType"` //String	true	string	劃出帳戶類型
	ToAccountType   *string `json:"toAccountType"`   //String	true	string	劃入帳戶類型
}

type AssetTransferQueryTransferCoinListAPI struct {
	client *PrivateRestClient
	req    *AssetTransferQueryTransferCoinListReq
}

func (api *AssetTransferQueryTransferCoinListAPI) FromAccountType(fromAccountType string) *AssetTransferQueryTransferCoinListAPI {
	api.req.FromAccountType = GetPointer(fromAccountType)
	return api
}

func (api *AssetTransferQueryTransferCoinListAPI) ToAccountType(toAccountType string) *AssetTransferQueryTransferCoinListAPI {
	api.req.ToAccountType = GetPointer(toAccountType)
	return api
}

// AssetTransferInterTransfer:          "/v5/asset/transfer/inter-transfer",            //POST 劃轉 (單帳號內)
type AssetTransferInterTransferReq struct {
	TransferId      *string `json:"transferId"`      //String	true	string	UUID. 請自行手動生成UUID
	Coin            *string `json:"coin"`            //String	true	string	幣種
	Amount          *string `json:"amount"`          //String	true	string	劃入數量
	FromAccountType *string `json:"fromAccountType"` //String	true	string	轉出賬戶類型
	ToAccountType   *string `json:"toAccountType"`   //String	true	string	轉入賬戶類型
}

type AssetTransferInterTransferAPI struct {
	client *PrivateRestClient
	req    *AssetTransferInterTransferReq
}

func (api *AssetTransferInterTransferAPI) TransferId(transferId string) *AssetTransferInterTransferAPI {
	api.req.TransferId = GetPointer(transferId)
	return api
}

func (api *AssetTransferInterTransferAPI) Coin(coin string) *AssetTransferInterTransferAPI {
	api.req.Coin = GetPointer(coin)
	return api
}

func (api *AssetTransferInterTransferAPI) Amount(amount string) *AssetTransferInterTransferAPI {
	api.req.Amount = GetPointer(amount)
	return api
}

func (api *AssetTransferInterTransferAPI) FromAccountType(fromAccountType string) *AssetTransferInterTransferAPI {
	api.req.FromAccountType = GetPointer(fromAccountType)
	return api
}

func (api *AssetTransferInterTransferAPI) ToAccountType(toAccountType string) *AssetTransferInterTransferAPI {
	api.req.ToAccountType = GetPointer(toAccountType)
	return api
}
