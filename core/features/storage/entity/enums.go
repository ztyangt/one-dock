package entity

// Category 文件分类
type Category int64

const (
	OtherCategory    Category = iota // 其他分类
	ImageCategory                    // 图片分类
	VideoCategory                    // 视频分类
	AudioCategory                    // 音频分类
	DocumentCategory                 // 文档分类
)

// StrategyType 存储策略类型
type StrategyType string

const (
	StorageLocal StrategyType = "local" // 本地存储
	StorageOSS   StrategyType = "oss"   // 阿里云OSS
	StorageKODO  StrategyType = "kodo"  // 七牛云Kodo
	StorageCOS   StrategyType = "cos"   // 腾讯云COS
	StorageUPYUN StrategyType = "upyun" // 又拍云
	StorageS3    StrategyType = "s3"    // AWS S3
	StorageOBS   StrategyType = "obs"   // 华为云OBS
)
