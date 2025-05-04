package data_api

func GetEmailsByCnpj(cnpj string) []string {
	//latency
	return provider.GetEmailsByCnpj(cnpj)
}
