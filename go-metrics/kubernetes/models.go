package kubernetes

type Liveness struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

type Readyness struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}
