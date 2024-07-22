package mybybitapi

type AssetTransferQueryInterTransferListRes struct {
	List           []AssetTransferQueryInterTransferListResRow `json:"list"`
	NextPageCursor string                                      `json:"nextPageCursor"`
}

type AssetTransferQueryInterTransferListResRow struct {
	TransferId      string `json:"transferId"`      //劃轉Id
	Coin            string `json:"coin"`            //劃轉幣種
	Amount          string `json:"amount"`          //劃轉金額
	FromAccountType string `json:"fromAccountType"` //劃出賬戶類型
	ToAccountType   string `json:"toAccountType"`   //劃入賬戶類型
	Timestamp       string `json:"timestamp"`       //劃轉創建時間戳 (毫秒)
	Status          string `json:"status"`          //劃轉狀態
}

type AssetTransferQueryTransferCoinListRes struct {
	List []string `json:"list"`
}

// transferId	string	UUID
// status	string	劃轉狀態
// STATUS_UNKNOWN
// SUCCESS
// PENDING
// FAILED

type AssetTransferInterTransferRes struct {
	TransferId string `json:"transferId"`
	Status     string `json:"status"`
}
