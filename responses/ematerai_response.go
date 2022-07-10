package responses

type LoginResponse struct {
	StatusCode int                    `json:"statusCode"`
	Message    string                 `json:"message"`
	Result     map[string]interface{} `json:"result"`
}

type SaldoResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}
