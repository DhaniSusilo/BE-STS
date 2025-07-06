package interfaces

import (
	"context"
	"learning-backend/domains/entities"
	"learning-backend/domains/models/requests"
	"learning-backend/domains/models/responses"
	sharedresponses "learning-backend/shared/models/responses"
)

type UseCase interface {
	Login(ctx context.Context, request *requests.LoginRequest) (*responses.LoginResponse, error)
	SignUp(ctx context.Context, request *requests.SignUpRequest) (*sharedresponses.BasicResponse, error)
	AddMember(ctx context.Context, request *requests.MemberRegistration) (*sharedresponses.BasicResponse, error)
	GetDashboardData(ctx context.Context,request *requests.GetDashboardData) (*sharedresponses.BasicResponse, error)
	GetRekapitulasi(ctx context.Context, req *requests.RekapitulasiRequest) (*responses.RekapitulasiResponse, error)
	GetAllUser(ctx context.Context) (*sharedresponses.BasicResponse,error)
	UpdateUser(ctx context.Context, username string, level string, forValue string) (*sharedresponses.BasicResponse, error)
	DeleteUser(ctx context.Context, userID string) (*sharedresponses.BasicResponse, error)
}

type Repository interface {
	FindByUsername(ctx context.Context, username string) (*entities.User, error)
	Create(ctx context.Context, request *requests.SignUpRequest) error
	CreateTokens(ctx context.Context,userId string) (string,error)
	AddMember(ctx context.Context,request *entities.Member) error
	GetDashboardData(ctx context.Context, request *requests.GetDashboardData) (*responses.GetDashboardData,error)
	GetAggregatedRekapitulasi(ctx context.Context, level, wilayah string, page, rowsPerPage int) (*[]responses.RekapitulasiResult, int, error)
	GetMembersByKelurahan(ctx context.Context, kelurahan string, page, rowsPerPage int) (*[]responses.MemberDetail, int, error)
	GetAllUser(ctx context.Context) (*[]entities.User,error)
    UpdateUser(ctx context.Context, username string, level , forValue string) error
	DeleteUser(ctx context.Context, userID string) error
}
