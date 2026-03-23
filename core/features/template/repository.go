package template

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	IsExist(ctx context.Context, id int64) (bool, error)
	Create(ctx context.Context, cfg *TemplateModel) error
	Update(ctx context.Context, cfg *TemplateModel) error
	Delete(ctx context.Context, id int64) error
	Find(ctx context.Context, id int64) (*TemplateModel, error)
}

type repository struct {
	DB *gorm.DB
}

func newRepository(db *gorm.DB) Repository {
	return &repository{DB: db}
}

// IsExist 检查模板是否存在
func (r *repository) IsExist(ctx context.Context, id int64) (bool, error) {
	var cfg TemplateModel
	err := r.DB.WithContext(ctx).Where("id = ?", id).First(&cfg).Error

	if err == nil {
		return true, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return false, err
}

// Create 创建配置
func (r *repository) Create(ctx context.Context, cfg *TemplateModel) error {
	return r.DB.WithContext(ctx).Create(cfg).Error
}

// Update 更新配置
func (r *repository) Update(ctx context.Context, cfg *TemplateModel) error {
	// 使用 Updates 会更新非零值字段。如果 cfg 中包含了主键 ID，GORM 会自动以 ID 为条件更新。
	// 如果需要更新所有字段（包含零值），请使用 r.DB.WithContext(ctx).Save(cfg).Error
	return r.DB.WithContext(ctx).Model(cfg).Updates(cfg).Error
}

// Find 根据模板ID获取模板
func (r *repository) Find(ctx context.Context, id int64) (*TemplateModel, error) {
	var cfg TemplateModel
	err := r.DB.WithContext(ctx).Where("id = ?", id).First(&cfg).Error

	if err != nil {
		// 拦截 GORM 的未找到错误，转化为业务定义的错误
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("数据不存在！")
		}
		// 返回其他数据库异常（如连接失败等）
		return nil, err
	}

	return &cfg, nil
}

// Delete 删除模板
func (r *repository) Delete(ctx context.Context, id int64) error {
	err := r.DB.WithContext(ctx).Where("id = ?", id).Delete(&TemplateModel{}).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("数据不存在！")
	}
	return err
}
