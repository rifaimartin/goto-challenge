package cashier

import (
	"errors"
)

// Service is interface
type Service interface {
	CreateCashier(input CasherInput) (Cashier, error)
	UpdateCashier(cashierId int, input CasherInput) (Cashier, error)
	Login(input LoginInput, cashierId int) (Cashier, error)
	Logout(input LoginInput, cashierId int) (Cashier, error)
	GetcashierByID(ID int) (Cashier, error)
	DeleteCashier(ID int) (Cashier, error)
	GetAllCashiers(limit int, skip int) (cashier []Cashier , count int64, error error)
}

type service struct {
	repository Repository
}

// NewService e
func NewService(repository Repository) *service {
	return &service{repository}
}


func (s *service) Login(input LoginInput, cashierId int) (Cashier, error) {
	passCode := input.PassCode

	cashier, err := s.repository.FindByID(cashierId)
	if err != nil {
		return cashier, err
	}

	if cashier.ID == 0 {
		return cashier, errors.New("No cashier found on that email")
	}

	//condition for match password
	if passCode != "123456" {
		return cashier,  errors.New("Passcode Not Match")
	}

	return cashier, nil
}

func (s *service) Logout(input LoginInput, cashierId int) (Cashier, error) {
	passCode := input.PassCode

	cashier, err := s.repository.FindByID(cashierId)
	if err != nil {
		return cashier, err
	}

	if cashier.ID == 0 {
		return cashier, errors.New("No cashier found on that email")
	}


	//condition for match password
	if passCode != "123456" {
		return cashier,  errors.New("Passcode Not Match")
	}

	return cashier, nil
}

func (s *service) GetcashierByID(ID int) (Cashier, error) {
	cashier, err := s.repository.FindByID(ID)

	if err != nil {
		return cashier, err
	}

	if cashier.ID == 0 {
		return cashier, errors.New("No cashier found with that ID")
	}

	return cashier, nil
}

func (s *service) GetAllCashiers(limit int, skip int) (cashier []Cashier, count int64, error error) {

	cashiers,count,err := s.repository.FindAll(limit,skip)

	if err == nil {
		return cashiers,count, err
	}

	return cashiers, count, nil
}

func (s *service) CreateCashier(input CasherInput) (Cashier, error) {
	cashier := Cashier{}
	cashier.Name = input.Name
	cashier.Passcode = "123456"

	newCashier, err := s.repository.Save(cashier)
	if err != nil {
		return newCashier, err
	}

	return newCashier, nil
}

func (s *service) UpdateCashier(cashierId int, inputData CasherInput) (Cashier, error) {
	cashier, err := s.repository.FindByID(cashierId)

	if err != nil {
		return cashier, err
	}

	if cashier.ID == 0 {
		return cashier, errors.New("No cashier found with that ID")
	}

	cashier.Name = inputData.Name

	updatedCampaign, err := s.repository.Update(cashier)
	if err != nil {
		return updatedCampaign, err
	}

	return updatedCampaign, nil

}

func (s *service) DeleteCashier(ID int) (Cashier, error) {

	dataCashier, err := s.repository.FindByID(ID)

	if err != nil {
		return dataCashier, err
	}

	if dataCashier.ID == 0 {
		return dataCashier, errors.New("No cashier found with that ID")
	}
	cashier, err := s.repository.DeleteOne(ID)
	
	if err != nil {
		return cashier, err
	}

	return cashier, nil
}
