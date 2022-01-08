package internal

import (
	"fmt"

	customer "github.com/dignelidxdx/HackthonGo/internal/customers"
	"github.com/dignelidxdx/HackthonGo/internal/models"
	"github.com/dignelidxdx/HackthonGo/pkg/util"
)

type BackUpService interface {
	IsSaved(nameFile string) (bool, error)
	ToSave(id int, nameFile string) (bool, error)
	SaveFile(nameFile string) ([][]string, error)
	SaveElementToDB(key string) error
}

type backUpService struct {
	repository BackUpRepository
	customerS  customer.CustomerService
}

func NewBackUpService(repository BackUpRepository, service customer.CustomerService) BackUpService {
	return &backUpService{repository: repository, customerS: service}
}

func (s *backUpService) IsSaved(nameFile string) (bool, error) {

	return s.repository.isLocked(nameFile)

}

func (s *backUpService) ToSave(id int, nameFile string) (bool, error) {

	return s.repository.SaveToLock(nameFile, id)

}

func (s *backUpService) SaveFile(nameFile string) ([][]string, error) {

	err := util.ConvertToCsv(nameFile)
	if err != nil {
		return nil, err
	}
	lines, err := util.ReadCsv(nameFile)
	if err != nil {
		panic(err)
	}
	return lines, nil

}

func (s *backUpService) SaveElementToDB(key string) error {

	fmt.Println(key)
	isSaved, err := s.IsSaved(key)

	if err != nil {
		return err
	}
	if isSaved {
		return fmt.Errorf("ya se guardo el txt de %v", key)
	}

	err = util.ConvertToCsv(key)
	if err != nil {
		return err
	}
	lines, err := util.ReadCsv(key)
	if err != nil {
		panic(err)
	}

	var indx int

	switch key {
	case "customers":
		customers := []models.Customer{}

		// Loop through lines & turn into object
		for _, line := range lines {
			data := models.Customer{
				LastName:  line[1],
				FirstName: line[2],
				Condition: line[3],
			}
			customers = append(customers, data)

		}
		fmt.Println("antes del service customer")
		err = s.customerS.SaveFile(customers)
		fmt.Println("despues del service")
		if err != nil {
			return err
		}
		indx = 1
	case "sales":
	case "products":
	case "invoices":
	default:
		return fmt.Errorf("la key es incorrecta")
	}

	fmt.Printf("sdfdsfs")
	_, err = s.ToSave(indx, key)
	if err != nil {
		return err
	}
	return nil
}
