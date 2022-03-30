package cashier

import (
	"gorm.io/gorm"
)

// Repository for inserting data to database
type Repository interface {
	Save(cashier Cashier) (Cashier, error)
	Update(cashier Cashier) (Cashier, error)
	FindByID(ID int) (Cashier, error)
	DeleteOne(ID int) (Cashier, error)
	FindAll(limit int, skip int) (cashier []Cashier, count int64, error error)
}

type repository struct {
	db *gorm.DB
}

// NewRepository function received main
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(cashier Cashier) (Cashier, error) {
	err := r.db.Create(&cashier).Error
	if err != nil {
		return cashier, err
	}

	return cashier, nil
}

func (r *repository) FindAll(limit int, skip int) (cashier []Cashier, ciunt int64, error error) {

	var cashiers []Cashier
	var count int64
	r.db.Find(&cashiers).Count(&count)
	err := r.db.Offset(skip).Limit(limit).Find(&cashiers).Limit(1).Error

	if err != nil {
		return cashiers,count , err
	}
	return cashiers, count, nil
}


func (r *repository) Update(cashier Cashier) (Cashier, error) {
	err := r.db.Save(&cashier).Error
	if err != nil {
		return cashier, err
	}

	return cashier, nil
}

func (r *repository) DeleteOne(ID int) (Cashier, error) {
	var cashier Cashier

	err := r.db.Where("id = ?", ID).Delete(&cashier).Error
	
	if err != nil {
		return cashier, err
	}

	return cashier, nil

}

func (r *repository) FindByID(ID int) (Cashier, error) {
	var cashier Cashier

	err := r.db.Where("id = ?", ID).Find(&cashier).Error
	
	if err != nil {
		return cashier, err
	}

	return cashier, nil
}