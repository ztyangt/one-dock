package config

import (
	"context"
	"errors"
	e "one-dock/app/error"
	"one-dock/core/comm"
	"one-dock/pkgs/utils"
	"strings"

	"github.com/spf13/cast"
)

type Service interface {
	CreateConfig(ctx context.Context, req *CreateReq) error
	UpdateConfig(ctx context.Context, req *UpdateReq) error
	DeleteConfig(ctx context.Context, req *DeleteReq) error
	GetConfig(ctx context.Context, req *GetReq) (*ConfigResp, error)
}

type service struct {
	repo Repository
}

// newService 私有构造函数
func newService(repo Repository) Service {
	return &service{repo: repo}
}

// CreateConfig 创建配置项
func (s *service) CreateConfig(ctx context.Context, req *CreateReq) error {

	if strings.HasPrefix(req.ConfigKey, "SYSTEM-") {
		return errors.New("系统配置项不能自定义！")
	}

	// 检查配置项是否已存在
	exist, err := s.repo.IsExist(ctx, req.ConfigKey)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("配置项已存在！")
	}

	var value = req.Value
	// 如果值是结构体或映射，转换为 JSON 字符串
	if utils.Json.IsJson(value) {
		value = utils.Json.Encode(value)
	}

	return s.repo.Create(ctx, &ConfigModel{
		Name:      req.Name,
		ConfigKey: req.ConfigKey,
		Public:    req.Public,
		Value:     cast.ToString(value),
	})
}

// UpdateConfig 更新配置项
func (s *service) UpdateConfig(ctx context.Context, req *UpdateReq) error {
	// 检查配置项是否存在
	_, err := s.repo.FindConfigByKey(ctx, req.ConfigKey)
	if err != nil {
		return err
	}

	var value = req.Value
	// 如果值是结构体或映射，转换为 JSON 字符串
	if utils.Json.IsJson(value) {
		value = utils.Json.Encode(value)
	}

	return s.repo.Update(ctx, &ConfigModel{
		ConfigKey: req.ConfigKey,
		Name:      req.Name,
		Value:     cast.ToString(value),
	})
}

// GetConfig 获取配置项详情
func (s *service) GetConfig(ctx context.Context, req *GetReq) (*ConfigResp, error) {

	// 1. 获取数据
	model, err := s.repo.FindConfigByKey(ctx, req.ConfigKey)
	if err != nil {
		return nil, err
	}

	if model.Public != true {
		user := ctx.Value("user")
		if user == nil {
			return nil, e.New(403, "配置项未公开！")
		}
		// 检查用户是否为管理员
		isAdmin := comm.UserRole(cast.ToInt64(user.(map[string]any)["role"])) == comm.UserRoleAdmin
		if !isAdmin {
			return nil, e.New(403, "您无权查看该配置项！")
		}
	}

	// 2. 转换为 DTO 返回给表现层(Controller/Handler)
	var configValue any = model.Value
	isJson := utils.Json.IsJson(configValue)
	if isJson {
		configValue, err = utils.Json.Decode(model.Value)
		if err != nil {
			return nil, err
		}
	}

	return &ConfigResp{
		ConfigKey: req.ConfigKey,
		Name:      model.Name,
		Value:     configValue,
	}, nil
}

// DeleteConfig 删除配置项
func (s *service) DeleteConfig(ctx context.Context, req *DeleteReq) error {

	if strings.HasPrefix(req.ConfigKey, "SYSTEM-") {
		return errors.New("系统配置项不能删除！")
	}

	return s.repo.Delete(ctx, req.ConfigKey)
}
