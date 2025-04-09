package data

import (
	"context"

	"github.com/itsgn/sgn-layout/internal/biz"
	"github.com/itsgn/sgn-layout/internal/domain"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"

	exampleV1 "github.com/itsgn/sgn-layout/api/example/v1"
)

type ExampleEntity struct {
	BaseFields
	Name   string `gorm:"type:varchar(255);not null;comment:name"`
	Status bool   `gorm:"not null;comment:Status 0 Frozen 1 Normal"`
	Domain string `gorm:"type:varchar(255);not null;comment:domain"`
}

func (ExampleEntity) TableName() string {
	return "example"
}

type ExampleRepo struct {
	data *Data
	log  *log.Helper
}

func NewExampleRepo(data *Data, logger log.Logger) biz.ExampleRepo {
	r := &ExampleRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/example")),
	}
	return r
}

// searchParam 
func (r ExampleRepo) searchParam(ctx context.Context, params map[string]interface{}) *gorm.DB {
	conn := r.data.db.Model(&ExampleEntity{})
	if Id, ok := params["id"]; ok && Id.(int64) != 0 {
		conn = conn.Where("id = ?", Id)
	}
	if Id, ok := params["id_neq"]; ok && Id.(int64) != 0 {
		conn = conn.Where("id != ?", Id)
	}
	if v, ok := params["name"]; ok && v.(string) != "" {
		conn = conn.Where("name LIKE ?", "%"+v.(string)+"%")
	}
	if v, ok := params["name_eq"]; ok && v.(string) != "" {
		conn = conn.Where("name = ?", v)
	}
	if v, ok := params["status"]; ok && v != nil {
		conn = conn.Where("status = ?", v)
	}
	// start time
	if v, ok := params["created_at_start"]; ok && v.(string) != "" {
		conn = conn.Where("created_at >= ?", v.(string)+" 00:00:00")
	}
	// end time
	if v, ok := params["created_at_end"]; ok && v.(string) != "" {
		conn = conn.Where("created_at <= ?", v.(string)+" 23:59:59")
	}
	if v, ok := params["is_deleted"]; ok {
		if v.(bool) {
			conn = conn.Unscoped()
		}
	}
	conn = getDbWithDomain(ctx, conn)
	return conn
}

func (r ExampleRepo) ListExample(ctx context.Context, page, pageSize int64, params map[string]interface{}) ([]*domain.Example, error) {
	list := []*ExampleEntity{}
	conn := r.searchParam(ctx, params)
	err := conn.Scopes(Paginate(page, pageSize)).Find(&list).Error
	if err != nil {
		return nil, exampleV1.ErrorSystemError("Failed to get list").WithCause(err)
	}

	rv := make([]*domain.Example, 0, len(list))
	for _, record := range list {
		rv = append(rv, domain.ToDomainExample(record))
	}
	return rv, nil
}
func (r ExampleRepo) GetExampleCount(ctx context.Context, params map[string]interface{}) (count int64, err error) {
	if len(params) == 0 {
		return 0, exampleV1.ErrorBadRequest("Parameter cannot be empty")
	}
	r.searchParam(ctx, params).Count(&count)
	return count, nil
}

func (r ExampleRepo) GetExampleByParams(ctx context.Context, params map[string]interface{}) (record *ExampleEntity, err error) {
	if len(params) == 0 {
		return &ExampleEntity{}, exampleV1.ErrorBadRequest("Parameter cannot be empty")
	}
	conn := r.searchParam(ctx, params)
	if err = conn.First(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &ExampleEntity{}, exampleV1.ErrorRecordNotFound("Data does not exist")
		}
		return record, exampleV1.ErrorSystemError("Failed to query records").WithCause(err)
	}
	return record, nil
}

func (r ExampleRepo) CreateExample(ctx context.Context, example *domain.Example) (*domain.Example, error) {
	record := &ExampleEntity{}
	record.Id = example.Id
	record.Name = example.Name
	record.Status = example.Status
	record.Domain = getDomain(ctx)
	if err := r.data.db.Model(record).Create(record).Error; err != nil {
		return nil, exampleV1.ErrorSystemError("Creation failed").WithCause(err)
	}
	return domain.ToDomainExample(record), nil
}

func (r ExampleRepo) UpdateExample(ctx context.Context, example *domain.Example) error {
	// Find records by Id
	record, err := r.GetExampleByParams(ctx, map[string]interface{}{
		"id": example.Id,
	})
	if err != nil {
		return err
	}
	// Update fields
	record.Name = example.Name
	record.Status = example.Status
	record.Domain = getDomain(ctx)
	if err := r.data.db.Model(&record).Where("id = ?", record.Id).Save(&record).Error; err != nil {
		return exampleV1.ErrorSystemError("Update failed").WithCause(err)
	}

	return nil
}

func (r ExampleRepo) GetExample(ctx context.Context, params map[string]interface{}) (*domain.Example, error) {
	// Find records by Id
	record, err := r.GetExampleByParams(ctx, params)
	if err != nil {
		return nil, err
	}
	// Return data
	response := domain.ToDomainExample(record)
	return response, nil
}

func (r ExampleRepo) DeleteExample(ctx context.Context, domain *domain.Example) error {
	if err := r.data.db.Where("id = ?", domain.Id).Delete(&ExampleEntity{}).Error; err != nil {
		return exampleV1.ErrorSystemError("Failed to delete data").WithCause(err)
	}
	return nil
}

func (r ExampleRepo) RecoverExample(ctx context.Context, domain *domain.Example) error {
	if err := r.data.db.Model(ExampleEntity{}).Unscoped().Where("id = ?", domain.Id).UpdateColumn("deleted_at", nil).Error; err != nil {
		return exampleV1.ErrorSystemError("Failed to restore data").WithCause(err)
	}
	return nil
}
