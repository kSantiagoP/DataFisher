package main

import (
	"github.com/kSantiagoP/DataFisher/internal/config"
	"github.com/kSantiagoP/DataFisher/internal/logger"
	"github.com/kSantiagoP/DataFisher/internal/model/company"
)

func main() {
	logg := logger.NewLogger("db_init")

	err := config.Init()
	if err != nil {
		logg.Errorf("Error initializing configs: %v", err)
		return
	}

	err = config.MigrateSchemas()
	if err != nil {
		logg.Errorf("Error initializing database: %v", err)
		return
	}

	err = config.InitDatabase()
	if err != nil {
		logg.Errorf("Error initializing database: %v", err)
		return
	}

	populaDatabase(logg)
}

func populaDatabase(logg *logger.Logger) {
	db := config.GetPostgresDB()

	companies := []company.Company{
		{
			Cnpj:               "35339277000128",
			Razao_social:       "RAZAO SOCIAL EMPRESA 1",
			Municipio:          "SÃO PAULO",
			Segmento:           1,
			Situacao_cadastral: "ATIVA",
		},
		{
			Cnpj:               "62367487000100",
			Razao_social:       "RAZAO SOCIAL EMPRESA 2",
			Municipio:          "RIO DE JANEIRO",
			Segmento:           2,
			Situacao_cadastral: "ATIVA",
		},
		{
			Cnpj:               "74666731000107",
			Razao_social:       "RAZAO SOCIAL EMPRESA 3",
			Municipio:          "BELO HORIZONTE",
			Segmento:           3,
			Situacao_cadastral: "ATIVA",
		},
		{
			Cnpj:               "04481464000118",
			Razao_social:       "TECH SOLUTIONS BRASIL",
			Municipio:          "SÃO PAULO",
			Segmento:           0,
			Situacao_cadastral: "ATIVA",
		},
		{
			Cnpj:               "31646212000174",
			Razao_social:       "AGROTECH INOVAÇÃO",
			Municipio:          "CAMPINAS",
			Segmento:           1,
			Situacao_cadastral: "ATIVA",
		},
		{
			Cnpj:               "96288576000175",
			Razao_social:       "FINANCIA FÁCIL",
			Municipio:          "RIO DE JANEIRO",
			Segmento:           2,
			Situacao_cadastral: "ATIVA",
		},
		{
			Cnpj:               "11428793000160",
			Razao_social:       "HEALTHCARE SYSTEMS",
			Municipio:          "BELO HORIZONTE",
			Segmento:           3,
			Situacao_cadastral: "ATIVA",
		},
		{
			Cnpj:               "37697954000105",
			Razao_social:       "ESCOLA DIGITAL",
			Municipio:          "CURITIBA",
			Segmento:           4,
			Situacao_cadastral: "ATIVA",
		},
		{
			Cnpj:               "45316650000189",
			Razao_social:       "MEGA VAREJO",
			Municipio:          "PORTO ALEGRE",
			Segmento:           5,
			Situacao_cadastral: "ATIVA",
		},
		{
			Cnpj:               "50185876000128",
			Razao_social:       "INDÚSTRIA FORTE",
			Municipio:          "SALVADOR",
			Segmento:           6,
			Situacao_cadastral: "ATIVA",
		},
		{
			Cnpj:               "21848274000105",
			Razao_social:       "CONSTRÓI BEM",
			Municipio:          "BRASÍLIA",
			Segmento:           7,
			Situacao_cadastral: "ATIVA",
		},
		{
			Cnpj:               "26277114000177",
			Razao_social:       "LOGÍSTICA RÁPIDA",
			Municipio:          "FORTALEZA",
			Segmento:           8,
			Situacao_cadastral: "ATIVA",
		},
		{
			Cnpj:               "34982038000129",
			Razao_social:       "ENERGIA VERDE",
			Municipio:          "RECIFE",
			Segmento:           9,
			Situacao_cadastral: "ATIVA",
		},
		{
			Cnpj:               "81049353000188",
			Razao_social:       "CODE MASTERS",
			Municipio:          "SÃO PAULO",
			Segmento:           0,
			Situacao_cadastral: "ATIVA",
		},
	}

	for _, comp := range companies {
		result := db.Where(company.Company{Cnpj: comp.Cnpj}).FirstOrCreate(&comp)
		if result.Error != nil {
			logg.Errorf("error populating database: %v", result.Error)
			return
		}
	}
}
