package rpc

//todo implement request sign method
func (r *ApiRequest) Sign(apiKey string) {
	//todo
	r.Signature = r.ApiKey
}

func (r *ApiRequest) GetSign(apiKey string) {
	//todo
}
