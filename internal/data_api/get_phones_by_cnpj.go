package data_api

import "time"

func GetPhonesByCnpj(cnpj string) []string {
	time.Sleep(100 * time.Millisecond)
	return provider.GetPhonesByCnpj(cnpj)
}
