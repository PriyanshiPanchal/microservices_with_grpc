package impl

import (
	"grpc-microservices/internal/proto-files/domain"
	"grpc-microservices/internal/proto-files/service"
	"context"
	"log"
	"strconv"
)

type RepositoryServiceGrpcImpl struct {
}

func NewRepositoryServiceGrpcImpl() *RepositoryServiceGrpcImpl {
	return &RepositoryServiceGrpcImpl{}
}

func (serviceImpl *RepositoryServiceGrpcImpl) Add(ctx context.Context, in *domain.Repository) (*service.AddRepositoryResponse, error) {
	log.Println("Received request for adding repository with id " + strconv.FormatInt(in.Id, 10))

	log.Println("Repository persisted to the storage")

	return &service.AddRepositoryResponse{
		AddedRepository: in,
		Error:           nil,
	}, nil
}
