package services

import (
	"task-session-1/models"
	"task-session-1/repositories"
	"time"
)

type TransactionService struct {
	transactionRepo *repositories.TransactionRepository
}

func NewTransactionService(transactionRepo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{transactionRepo: transactionRepo}
}

func (s *TransactionService) Checkout(items []models.CheckoutItem, useLock bool) (*models.Transaction, error) {
	return s.transactionRepo.CreateTransaction(items)
}

func (s *TransactionService) GetReport(startDate, endDate time.Time) (*models.ReportResponse, error) {
	return s.transactionRepo.GetReport(startDate, endDate)
}
