package types

import "encoding/json"

type Operation int

const (
	ENRICH Operation = iota
	CONSULT
	RESULT
)

type Situation string

const (
	ATIVA   = "ATIVA"
	INATIVA = "INATIVA"
)

type Segmento int

const (
	TECH = iota
	AGROTECH
	FINANCAS
	SAUDE
	EDUCACAO
	VAREJO
	INDUSTRIA
	CONSTRUCAO
	LOGISTICA
	ENERGIA
)

var segmentoNome = map[Segmento]string{
	TECH:       "TECH",
	AGROTECH:   "AGROTECH",
	FINANCAS:   "FINANCAS",
	SAUDE:      "SAUDE",
	EDUCACAO:   "EDUCACAO",
	VAREJO:     "VAREJO",
	INDUSTRIA:  "INDUSTRIA",
	CONSTRUCAO: "CONSTRUCAO",
	LOGISTICA:  "LOGISTICA",
	ENERGIA:    "ENERGIA",
}

func (sn Segmento) String() string {
	return segmentoNome[sn]
}

func (s Segmento) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}
