package data_api

func GetPhonesByCnpj(cnpj string) []string {
	//latency
	return provider.GetPhonesByCnpj(cnpj)
}
