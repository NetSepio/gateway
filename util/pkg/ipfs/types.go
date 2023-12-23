package ipfs

import "encoding/json"

func UnmarshalNFTStorageRes(data []byte) (NFTStorageRes, error) {
	var r NFTStorageRes
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *NFTStorageRes) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type NFTStorageRes struct {
	Ok    bool  `json:"ok"`
	Value Value `json:"value"`
	Error Error `json:"error"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Value struct {
	Cid     string `json:"cid"`
	Size    int64  `json:"size"`
	Created string `json:"created"`
	Type    string `json:"type"`
	Scope   string `json:"scope"`
	Pin     Pin    `json:"pin"`
	Files   []File `json:"files"`
	Deals   []Deal `json:"deals"`
}

type Deal struct {
	BatchRootCid   string `json:"batchRootCid"`
	LastChange     string `json:"lastChange"`
	Miner          string `json:"miner"`
	Network        string `json:"network"`
	PieceCid       string `json:"pieceCid"`
	Status         string `json:"status"`
	StatusText     string `json:"statusText"`
	ChainDealID    int64  `json:"chainDealID"`
	DealActivation string `json:"dealActivation"`
	DealExpiration string `json:"dealExpiration"`
}

type File struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Pin struct {
	Cid     string `json:"cid"`
	Name    string `json:"name"`
	Meta    Meta   `json:"meta"`
	Status  string `json:"status"`
	Created string `json:"created"`
	Size    int64  `json:"size"`
}

type Meta struct {
}
