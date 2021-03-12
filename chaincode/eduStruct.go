package main

type WalletCC struct {
	Id     string `json:"id"`
	Amount []byte `json:"amount"`
}

type TransactionCC struct {
	Id         string `json:"id"`
	CreateTime int64  `json:"create_time"`
	FromWallet string `json:"from_wallet"`
	ToWallet   string `json:"to_wallet"`
	Amount     []byte `json:"amount"`
}
