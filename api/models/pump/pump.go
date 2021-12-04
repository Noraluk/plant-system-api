package pumpmodel

type Pump struct {
	ID        int  `json:"id"`
	IsActive  bool `json:"is_active"`
	IsWorking bool `json:"is_working"`
}

type PumpIsWorkingReq struct {
	ID        int  `json:"id" validate:"required"`
	IsWorking bool `json:"is_working" validate:"required"`
}

type PumpIsWorkingResponse struct {
	ID        int  `json:"id"`
	IsWorking bool `json:"is_working"`
}

type PumpActiveReq struct {
	ID       int  `json:"id" validate:"required"`
	IsActive bool `json:"is_active" validate:"required"`
}

type PumpActiveResponse struct {
	ID       int  `json:"id"`
	IsActive bool `json:"is_active"`
}
