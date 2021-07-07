package service

import (
	"context"

	"github.com/diogofgomespinheiro/codebank/dto"
	"github.com/diogofgomespinheiro/codebank/infrastructure/grpc/pb"
	"github.com/diogofgomespinheiro/codebank/usecases"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TransactionService struct {
	ProcessTransactionUseCase usecases.UseCaseTransaction
	pb.UnimplementedPaymentServiceServer
}

func NewTransactionService() *TransactionService {
	return &TransactionService{}
}

func (t *TransactionService) Payment(ctx context.Context, in *pb.PaymentRequest) (*empty.Empty, error) {
	transactionDto := dto.Transaction{
		Name:            in.GetCreditCard().Name,
		Number:          in.GetCreditCard().Number,
		ExpirationMonth: in.GetCreditCard().ExpirationMonth,
		ExpirationYear:  in.GetCreditCard().ExpirationYear,
		CVV:             in.GetCreditCard().GetCvv(),
		Amount:          in.GetAmount(),
		Store:           in.GetStore(),
		Description:     in.GetDescription(),
	}

	transaction, err := t.ProcessTransactionUseCase.ProcessTransaction(transactionDto)

	if err != nil {
		return &empty.Empty{}, status.Error(codes.FailedPrecondition, err.Error())
	}

	if transaction.Status != "approved" {
		return &empty.Empty{}, status.Error(codes.FailedPrecondition, "transaction rejected by bank")
	}

	return &empty.Empty{}, nil
}
