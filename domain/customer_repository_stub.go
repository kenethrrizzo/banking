package domain

// Adapter: Mock

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Keneth", City: "Guayaquil", ZipCode: "12345", DateOfBirth: "2000-02-12", Status: "1"},
		{Id: "1002", Name: "Camila", City: "Guayaquil", ZipCode: "12344", DateOfBirth: "2000-23-02", Status: "1"},
		{Id: "1003", Name: "Diana", City: "Quito", ZipCode: "12343", DateOfBirth: "2000-04-08", Status: "1"},
	}
	return CustomerRepositoryStub{customers}
}
