package repositories

import (
	"context"
	"fmt"
	interfaces "learning-backend/domains"
	"learning-backend/domains/entities"
	"learning-backend/domains/models/requests"
	"learning-backend/domains/models/responses"
	"learning-backend/infrastructures"
	"time"

	"github.com/google/uuid"
)

type databaseRepository struct {
	db infrastructures.Database
}


func NewDatabaseRepository(db infrastructures.Database) interfaces.Repository {
	return databaseRepository{
		db: db,
	}
}

func (repo databaseRepository) FindByUsername(ctx context.Context, username string) (*entities.User, error) {
	var user entities.User

	result := repo.db.GetInstance().First(&user, "username = ?", username)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repo databaseRepository) Create(ctx context.Context, request *requests.SignUpRequest) error {
	result := repo.db.GetInstance().Create(
		&entities.User{
			Id:        uuid.New(),
			Username:  request.UserName,
			FirstName: request.FirstName,
			LastName:  request.LastName,
			Password:  request.Password,
			Level:     request.Level,
			For:       request.For,
		},
	)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// CreateTokens implements users.UserRepository.
func (repo databaseRepository) CreateTokens(ctx context.Context, userId string) (string, error) {
	tokenID := uuid.New()

	return tokenID.String(), nil
}

// AddMember implements interfaces.Repository.
func (repo databaseRepository) AddMember(ctx context.Context, request *entities.Member) error {
	// Use the GORM DB instance to insert the member into the database
	if err := repo.db.GetInstance().WithContext(ctx).Create(request).Error; err != nil {
		return err
	}
	return nil
}

// GetDashboardData implements interfaces.Repository.
func (repo databaseRepository) GetDashboardData(ctx context.Context, request *requests.GetDashboardData) (*responses.GetDashboardData, error) {
	var total int64
	var hariIni int64

	// Filter base query depending on level
	query := repo.db.GetInstance().WithContext(ctx).Model(&entities.Member{})

	switch request.Level {
	case "Provinsi":
		query = query.Where("provinsi = ?", request.For)
	case "Kabupaten":
		query = query.Where("kabupaten = ?", request.For)
	case "Kecamatan":
		query = query.Where("kecamatan = ?", request.For)
	case "Kelurahan":
		query = query.Where("kelurahan = ?", request.For)
		// "pusat" sees all data, no filter needed
	}

	// Count total registrations
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	// Count today's registrations
	startOfDay := time.Now().Truncate(24 * time.Hour)
	if err := query.Where("created_at >= ?", startOfDay).Count(&hariIni).Error; err != nil {
		return nil, err
	}

	return &responses.GetDashboardData{
		HariIni: int(hariIni),
		Total:   int(total),
	}, nil
}

func (repo databaseRepository) GetAggregatedRekapitulasi(ctx context.Context, level string, wilayah string, page int, rowsPerPage int) (*[]responses.RekapitulasiResult, int, error) {
	var results []responses.RekapitulasiResult
	var total int64
	var groupByField string
	var whereClause map[string]interface{} = make(map[string]interface{})

	offset := (page - 1) * rowsPerPage

	switch level {
	case "Pusat":
		groupByField = "provinsi"
	case "Provinsi":
		groupByField = "kabupaten"
		whereClause["provinsi"] = wilayah
	case "Kabupaten":
		groupByField = "kecamatan"
		whereClause["kabupaten"] = wilayah
	case "Kecamatan":
		groupByField = "kelurahan"
		whereClause["kecamatan"] = wilayah
	default:
		return nil, 0, fmt.Errorf("unsupported level: %s", level)
	}

	// Count total groups
	err := repo.db.GetInstance().
		Model(&entities.Member{}).
		Where(whereClause).
		Select(groupByField).
		Group(groupByField).
		Count(&total).Error

	if err != nil {
		return nil, 0, err
	}

	// Fetch paginated aggregated data
	err = repo.db.GetInstance().
		Model(&entities.Member{}).
		Select(fmt.Sprintf("%s as wilayah, COUNT(*) as total_anggota", groupByField)).
		Where(whereClause).
		Group(groupByField).
		Order(groupByField).
		Scan(&results).Error

	if err != nil {
		return nil, 0, err
	}

	// Add numbering
	for i := range results {
		results[i].No = offset + i + 1
	}

	return &results, int(total), nil
}

func (repo databaseRepository) GetMembersByKelurahan(ctx context.Context, kelurahan string, page int, rowsPerPage int) (*[]responses.MemberDetail, int, error) {
	var members []responses.MemberDetail
	var total int64

	offset := (page - 1) * rowsPerPage

	// Count total matching members
	err := repo.db.GetInstance().
		Model(&entities.Member{}).
		Where("kelurahan = ?", kelurahan).
		Count(&total).Error

	if err != nil {
		return nil, 0, err
	}

	// Fetch member details
	err = repo.db.GetInstance().
		Model(&entities.Member{}).
		Where("kelurahan = ?", kelurahan).
		Select("nama, nik, no_hp, provinsi, kabupaten, kecamatan, kelurahan, created_at as tanggal_daftar").
		Limit(rowsPerPage).
		Offset(offset).
		Order("created_at DESC").
		Scan(&members).Error

	if err != nil {
		return nil, 0, err
	}

	return &members, int(total), nil
}

// GetAllUser implements interfaces.Repository.
func (repo databaseRepository) GetAllUser(ctx context.Context) (*[]entities.User, error) {
	var users []entities.User

	// Query all users with the given context
	result := repo.db.GetInstance().WithContext(ctx).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return &users, nil
}

func (repo databaseRepository) UpdateUser(ctx context.Context, username string, level string, forValue string) error {
	err := repo.db.GetInstance().WithContext(ctx).Model(&entities.User{}).
		Where("username = ?", username).
		Updates(map[string]interface{}{
			"level": level,
			"for":   forValue,
		}).Error
	// fmt.Println(err)
	return err
}

func (repo databaseRepository) DeleteUser(ctx context.Context, userID string) error {
    return repo.db.GetInstance().WithContext(ctx).Where("id = ?", userID).Delete(&entities.User{}).Error
}

