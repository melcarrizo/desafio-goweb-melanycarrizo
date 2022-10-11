package tickets

import (
	"context"
	"desafio-goweb-melanycarrizo/internal/domain"
)

type Service interface {
	GetTicketsByCountry(ctx context.Context, destination string) ([]domain.Ticket, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetTicketsByCountry(ctx context.Context, destination string) ([]domain.Ticket, error) {

	return s.repository.GetTicketsByCountry(ctx, destination)
}

func (s *service) AverageDestination(ctx context.Context, destination string) (float64, error) {
	total, error := s.repository.GetAll(ctx)

	if error != nil {
		return 0, error
	}

	totalDestino, error := s.repository.GetTicketsByCountry(ctx, destination)

	cantidad := float64(len(totalDestino)) / float64(len(total))

	return cantidad, nil
}
