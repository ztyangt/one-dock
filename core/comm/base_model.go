package comm

import (
	"github.com/bwmarrin/snowflake"
	"gorm.io/plugin/soft_delete"
)

var SnowNode *snowflake.Node

type BaseModel struct {
	Remark    string                `gorm:"size:512;comment:备注" json:"remark"`
	CreatedAt int64                 `gorm:"autoCreateTime:milli; comment:创建时间;" json:"created_at"`
	UpdatedAt int64                 `gorm:"autoUpdateTime:milli; comment:更新时间;" json:"updated_at"`
	DeletedAt soft_delete.DeletedAt `gorm:"softDelete:milli;comment:删除时间; default:null;" json:"deleted_at"`
}

func init() {
	// 初始化雪花算法节点
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
	SnowNode = node
}
