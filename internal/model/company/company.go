package company

//Lógica para a configuração e organização dos dados no DB

import (
	//"time"

	"gorm.io/gorm"
)

type SituacaoCadastral int

const (
	INATIVA SituacaoCadastral = iota
	ATIVA
)

var situacaoNome = map[SituacaoCadastral]string{
	INATIVA: "INATIVA",
	ATIVA:   "ATIVA",
}

func (sn SituacaoCadastral) String() string {
	return situacaoNome[sn]
}

type Company struct {
	gorm.Model
	Cnpj               string
	Razao_social       string
	Municipio          string //normalizado IBGE
	Segmento           int16  //código
	Situacao_cadastral string //enum (IMPLEMENTAR)
}

/*
type CompanyResponse struct {
	id                 uint
	cnpj               string
	razao_social       string
	municipio          string //normalizado IBGE
	segmento           int16  //código
	situacao_cadastral string //enum (IMPLEMENTAR)
	updated_at         time.Time
}
*/
