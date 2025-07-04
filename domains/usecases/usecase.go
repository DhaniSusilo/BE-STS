package usecases

import (
	"context"
	"errors"
	interfaces "learning-backend/domains"
	"learning-backend/domains/entities"
	"learning-backend/domains/models/requests"
	"learning-backend/domains/models/responses"
	"learning-backend/shared/constant"
	sharedresponse "learning-backend/shared/models/responses"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UseCase struct {
	repo interfaces.Repository
}



func NewUseCase(repo interfaces.Repository) interfaces.UseCase {
	return UseCase{repo: repo}
}

func (u UseCase) Login(ctx context.Context, request *requests.LoginRequest) (*responses.LoginResponse, error) {
	user, err := u.repo.FindByUsername(ctx, request.Username)

	if err != nil {
		return nil, err
	}

	if user.Enabled == 0 {
		return nil, errors.New("user is disabled")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	if err != nil {
		return nil, errors.New("Unauthorized")
	}

	accessToken, err := u.repo.CreateTokens(ctx, user.Id.String())

	if err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"token_id": accessToken,
			"exp":      time.Now().Add(24 * time.Hour).Unix(),
		})

	tokenString, err := token.SignedString(constant.JWT_SECRET)

	if err != nil {
		return nil, err
	}

	return &responses.LoginResponse{
		AccessToken: tokenString,
		For:         user.For,
		Level:       user.Level,
		Name:        user.FirstName + " " + user.LastName,
	}, nil
}

func (u UseCase) SignUp(ctx context.Context, request *requests.SignUpRequest) (*sharedresponse.BasicResponse, error) {
	if request.Password != request.ConfirmPassword {
		return nil, errors.New("password must be same with confirm password")
	}

	b, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)

	if err != nil {
		return nil, errors.New("error when hashing password")
	}

	request.Password = string(b)

	err = u.repo.Create(ctx, request)

	if err != nil {
		return nil, err
	}

	return &sharedresponse.BasicResponse{
		Data: struct {
			Message string
		}{
			Message: "Account created",
		},
	}, nil
}

func (u UseCase) AddMember(ctx context.Context, request *requests.MemberRegistration) (*sharedresponse.BasicResponse, error) {

	member := &entities.Member{
		NIK:       request.NIK,
		Nama:      request.Nama,
		NoHp:      request.NoHp,
		Provinsi:  request.Provinsi,
		Kabupaten: request.Kabupaten,
		Kecamatan: request.Kecamatan,
		Kelurahan: request.Kelurahan,
		CreatedAt: time.Now(),
	}

	if err := u.repo.AddMember(ctx, member); err != nil {
		return nil, err
	}

	return &sharedresponse.BasicResponse{
		Data: struct {
			Message string
		}{
			Message: "Member successfully registered",
		},
	}, nil
}

// GetDashboardData implements interfaces.UseCase.
func (u UseCase) GetDashboardData(ctx context.Context, request *requests.GetDashboardData) (*sharedresponse.BasicResponse, error) {
	// Call the repository to get the data
	data, err := u.repo.GetDashboardData(ctx, request)
	if err != nil {
		return nil, err
	}

	// Wrap the data in a basic response
	return &sharedresponse.BasicResponse{
		Data: struct {
			Message string
			Data    *responses.GetDashboardData
		}{

			Message: "Success get dashboard data",
			Data:    data,
		},
	}, nil
}

func (u UseCase) GetRekapitulasi(ctx context.Context, req *requests.RekapitulasiRequest) (*responses.RekapitulasiResponse, error) {
	// Check user level
	if req.Level == "Kelurahan" {
		// Fetch detailed members list for that kelurahan
		members, total, err := u.repo.GetMembersByKelurahan(ctx, req.Wilayah, req.Page, req.RowsPerPage)
		if err != nil {
			return nil, err
		}
		return &responses.RekapitulasiResponse{Members: *members, TotalCount: total}, nil
	} else {
		// Fetch aggregated counts for wilayahs up to kecamatan level
		aggregated, total, err := u.repo.GetAggregatedRekapitulasi(ctx, req.Level, req.Wilayah, req.Page, req.RowsPerPage)
		if err != nil {
			return nil, err
		}
		return &responses.RekapitulasiResponse{Aggregated: *aggregated, TotalCount: total}, nil
	}
}

// GetAllUser implements interfaces.UseCase.
func (u UseCase) GetAllUser(ctx context.Context) (*sharedresponse.BasicResponse, error) {
	users, err := u.repo.GetAllUser(ctx)
	if err != nil {
		return nil, err
	}

	return &sharedresponse.BasicResponse{
		Data: struct{
			Message string
			Data *[]entities.User

		}{

			Message: "Successfully retrieved all users",
			Data:    users,
		},
	}, nil
}
