package dtos

type UnsignedTx struct {
	To    string `json:"to"`
	Data  string `json:"data"`
	Value string `json:"value"`
	Gas   string `json:"gas"`
}
