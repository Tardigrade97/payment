package service

import (
	"github.com/Tardigrade97/payment/domain/model"
	"github.com/Tardigrade97/payment/domain/repository"
)

type IPaymentDataService interface {
	AddPayment(payment *model.Payment) (int64, error)
	DeletePayment(int64) error
	UpdatePayment(payment *model.Payment) error
	FindPaymentByID(int64) (*model.Payment, error)
	FindAllPayment() ([]model.Payment, error)
}

type PaymentDataService struct {
	PaymentRepository repository.IPaymentRepository
}

//创建
func NewPaymentDataService(paymentRepository repository.IPaymentRepository) IPaymentDataService {
	return &PaymentDataService{paymentRepository}
}
func (u *PaymentDataService) AddPayment(payment *model.Payment) (int64, error) {
	return u.PaymentRepository.CreatePayment(payment)
}
func (u *PaymentDataService) DeletePayment(paymentID int64) error {
	return u.PaymentRepository.DeletePaymentByID(paymentID)
}
func (u *PaymentDataService) UpdatePayment(payment *model.Payment) error {
	return u.PaymentRepository.UpdatePayment(payment)
}
func (u *PaymentDataService) FindPaymentByID(pamentID int64) (*model.Payment, error) {
	return u.PaymentRepository.FindPaymentByID(pamentID)
}
func (u *PaymentDataService) FindAllPayment() ([]model.Payment, error) {
	return u.PaymentRepository.FindAll()
}
