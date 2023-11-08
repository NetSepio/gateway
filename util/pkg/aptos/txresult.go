package aptos

import "encoding/json"

func UnmarshalTxResult(data []byte) (TxResult, error) {
	var r TxResult
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TxResult) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TxResult struct {
	Result Result `json:"Result"`
}

type Result struct {
	TransactionHash string `json:"transaction_hash"`
	GasUsed         int64  `json:"gas_used"`
	GasUnitPrice    int64  `json:"gas_unit_price"`
	Sender          string `json:"sender"`
	SequenceNumber  int64  `json:"sequence_number"`
	Success         bool   `json:"success"`
	TimestampUs     int64  `json:"timestamp_us"`
	Version         int64  `json:"version"`
	VMStatus        string `json:"vm_status"`
}
