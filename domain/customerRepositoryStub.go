package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{
            Id:          "1",
            Name:        "John",
            City:        "London",
            Zipcode:     "12345",
            DateOfBirth: "2000-01-01",
			Status:      "1",
		},
		{
            Id:          "2",
            Name:        "Paul",
            City:        "Liverpool",
            Zipcode:     "1589745",
            DateOfBirth: "1956-03-08",
			Status:      "1",
		},
	}
	return CustomerRepositoryStub{customers: customers}
}