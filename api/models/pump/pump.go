package pumpmodel

type Pump struct {
	ID        int  `json:"id"`
	IsActive  bool `json:"is_active"`
	IsWorking bool `json:"is_working"`
	IsAsk     bool `json:"is_ask"`
}

type PumpActiveReq struct {
	ID       int  `json:"id" `
	IsActive bool `json:"is_active" `
}

type PumpActiveResponse struct {
	ID       int  `json:"id"`
	IsActive bool `json:"is_active"`
}

type PumpAskingReq struct {
	ID    int  `json:"id" `
	IsAsk bool `json:"is_ask" `
}

type PumpAskingResponse struct {
	ID    int  `json:"id"`
	IsAsk bool `json:"is_ask"`
}
