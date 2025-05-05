package data_api

import "time"

func GetEmailsByCnpj(cnpj string) []string {
	time.Sleep(100 * time.Millisecond)
	return provider.GetEmailsByCnpj(cnpj)
}
