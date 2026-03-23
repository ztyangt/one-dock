package comm

import (
	e "one-dock/app/error"

	"github.com/jinzhu/copier"
)

type BaseService struct {
}

// ModelToDTO 将模型转换为DTO
func (s *BaseService) ModelToDTO(dto any, source any) error {
	if err := copier.Copy(dto, source); err != nil {
		return e.New(500, err.Error())
	}
	return nil
}
