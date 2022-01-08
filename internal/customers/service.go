package internal

import (
	"fmt"

	"github.com/dignelidxdx/HackthonGo/internal/models"
)

type CustomerService interface {
	SaveCustomer(models.Customer) (models.Customer, error)
	SaveFile(customers []models.Customer) error
}

type customerService struct {
	repository CustomerRepository
}

func NewCustomerService(repository CustomerRepository) CustomerService {
	return &customerService{repository: repository}
}

/*
func (s *customerService) SaveCustomerFile(nameFile string) error {

	isSaved, err := s.backupService.IsSaved(nameFile)

	if err != nil {
		return err
	}
	if isSaved {
		return fmt.Errorf("ya se guardo el txt")
	}

	err = util.ConvertToCsv(nameFile)
	if err != nil {
		return err
	}
	lines, err := util.ReadCsv(nameFile)
	if err != nil {
		panic(err)
	}
	clients := []models.Customer{}
	// Loop through lines & turn into object
	for _, line := range lines {
		data := models.Customer{
			LastName:  line[1],
			FirstName: line[2],
			Condition: line[3],
		}
		clients = append(clients, data)
		fmt.Println(data.LastName + " " + data.FirstName)
	}

	err = s.repository.SaveFile(clients)
	if err != nil {
		return err
	}
	_, err = s.backupService.ToSave(1, nameFile)
	if err != nil {
		return err
	}
	return nil

}*/

func (s *customerService) SaveCustomer(customer models.Customer) (models.Customer, error) {

	customer, err := s.repository.Save(customer)
	if err != nil {
		return models.Customer{}, err
	}
	return customer, nil
}

func (s *customerService) SaveFile(customers []models.Customer) error {

	fmt.Println("service...")
	err := s.repository.SaveFile(customers)
	if err != nil {
		return err
	}
	return nil
}
