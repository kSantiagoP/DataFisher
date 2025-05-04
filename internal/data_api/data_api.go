package data_api

import (
	"fmt"

	"github.com/kSantiagoP/DataFisher/internal/data_api/mock"
)

//All the logic to data api extern calls

var (
	provider *mock.MockProvider
)

func Init() error {
	var err error
	provider, err = mock.NewMockProvider()
	if err != nil {
		return fmt.Errorf("error initializing providers: %v", err)
	}
	return nil
}
