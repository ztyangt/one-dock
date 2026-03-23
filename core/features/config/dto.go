package config

type GetReq struct {
	ConfigKey string `query:"config_key" validate:"required" alias:"配置项key"`
}

type CreateReq struct {
	ConfigKey string `json:"config_key" validate:"required" alias:"配置项key"`
	Name      string `json:"name" validate:"required" alias:"配置名称"`
	Public    bool   `json:"public" alias:"是否公开"`
	Value     any    `json:"value" alias:"配置值"`
}

type UpdateReq struct {
	ConfigKey string `json:"config_key" validate:"required" alias:"配置项key"`
	Name      string `json:"name"  alias:"配置名称"`
	Value     any    `json:"value" validate:"required" alias:"配置项值"`
}

type DeleteReq struct {
	ConfigKey string `json:"config_key" validate:"required" alias:"配置项key"`
}

type ConfigResp struct {
	ConfigKey string `json:"config_key"`
	Name      string `json:"name"`
	Value     any    `json:"value"`
}
