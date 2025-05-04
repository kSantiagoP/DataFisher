package mock

type MockProvider struct {
	emails EmailDatabase
	phones PhoneDatabase
}

func NewMockProvider() (*MockProvider, error) {
	emails, phones, err := LoadMocks()
	if err != nil {
		return nil, err
	}
	return &MockProvider{
		emails: emails,
		phones: phones,
	}, nil
}

func (p *MockProvider) GetEmailsByCnpj(cnpj string) []string {
	return p.emails[cnpj]
}

func (p *MockProvider) GetPhonesByCnpj(cnpj string) []string {
	return p.phones[cnpj]
}
