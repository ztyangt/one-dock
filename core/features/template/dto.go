package template

type GetByIdReq struct {
	ID int64 `query:"id" validate:"required" alias:"模板id"`
}

type DeleteReq struct {
	ID int64 `json:"id" validate:"required" alias:"模板id"`
}

type CreateReq struct {
	Name  string `json:"name" validate:"required" alias:"模板数据名称"`
	Value string `json:"value" validate:"required" alias:"模板数据值"`
}

type UpdateReq struct {
	ID    int64  `json:"id" validate:"required" alias:"模板id"`
	Name  string `json:"name" validate:"required" alias:"模板数据名称"`
	Value string `json:"value" validate:"required" alias:"模板数据值"`
}

type TemplateResp struct {
	ID    int64
	Name  string
	Value string
}
