package template

import (
	"context"
)

type Service interface {
	CreateTemplate(ctx context.Context, req *CreateReq) error
	UpdateTemplate(ctx context.Context, req *UpdateReq) error
	GetTemplateById(ctx context.Context, id int64) (*TemplateResp, error)
	DeleteTemplate(ctx context.Context, req *DeleteReq) error
}

type service struct {
	repo Repository
}

// newService 私有构造函数
func newService(repo Repository) Service {
	return &service{repo: repo}
}

// CreateTemplate 创建模板
func (s *service) CreateTemplate(ctx context.Context, req *CreateReq) error {

	model := &TemplateModel{
		Name:  req.Name,
		Value: req.Value,
	}
	return s.repo.Create(ctx, model)
}

// UpdateTemplate 更新模板
func (s *service) UpdateTemplate(ctx context.Context, req *UpdateReq) error {

	existingConfig, err := s.repo.Find(ctx, req.ID)
	if err != nil {
		return err
	}

	// 1. 修改需要更新的字段
	existingConfig.Name = req.Name
	existingConfig.Value = req.Value

	// 4. 执行更新
	return s.repo.Update(ctx, existingConfig)
}

// GetTemplate 获取模板
func (s *service) GetTemplateById(ctx context.Context, id int64) (*TemplateResp, error) {
	// 1. 获取底层数据
	model, err := s.repo.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	// 2. 转换为 DTO 返回给表现层(Controller/Handler)
	return &TemplateResp{
		ID:    model.ID,
		Name:  model.Name,
		Value: model.Value,
	}, nil
}

// DeleteTemplate 删除模板
func (s *service) DeleteTemplate(ctx context.Context, req *DeleteReq) error {
	return s.repo.Delete(ctx, req.ID)
}
