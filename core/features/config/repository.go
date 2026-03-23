package config

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	IsExist(ctx context.Context, key string) (bool, error)
	Create(ctx context.Context, cfg *ConfigModel) error
	Update(ctx context.Context, cfg *ConfigModel) error
	FindConfigByKey(ctx context.Context, key string) (*ConfigModel, error)
	Delete(ctx context.Context, key string) error
}

type repository struct {
	DB *gorm.DB
}

func newRepository(db *gorm.DB) Repository {
	return &repository{DB: db}
}

// IsExist 检查配置项是否存在
func (r *repository) IsExist(ctx context.Context, key string) (bool, error) {
	var cfg ConfigModel
	err := r.DB.WithContext(ctx).Where("config_key = ?", key).First(&cfg).Error
	if err == nil {
		return true, nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return false, err
}

// Create 创建配置
func (r *repository) Create(ctx context.Context, cfg *ConfigModel) error {
	return r.DB.WithContext(ctx).Create(cfg).Error
}

// Update 更新配置
func (r *repository) Update(ctx context.Context, cfg *ConfigModel) error {
	// 使用 Updates 会更新非零值字段。如果 cfg 中包含了主键 ID，GORM 会自动以 ID 为条件更新。
	// 如果需要更新所有字段（包含零值），请使用 r.DB.WithContext(ctx).Save(cfg).Error
	return r.DB.WithContext(ctx).Model(cfg).Updates(cfg).Error
}

// FindConfigByKey 根据配置键获取配置
func (r *repository) FindConfigByKey(ctx context.Context, key string) (*ConfigModel, error) {
	var cfg ConfigModel
	err := r.DB.WithContext(ctx).Where("config_key = ?", key).First(&cfg).Error

	if err != nil {
		// 拦截 GORM 的未找到错误，转化为业务定义的错误
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("配置不存在！")
		}
		// 返回其他数据库异常（如连接失败等）
		return nil, err
	}

	return &cfg, nil
}

// Delete 删除配置
func (r *repository) Delete(ctx context.Context, key string) error {
	err := r.DB.WithContext(ctx).Where("config_key = ?", key).Delete(&ConfigModel{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("配置不存在！")
	}
	return err
}
