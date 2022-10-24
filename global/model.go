package global

type GVA_MODEL struct {
	ID         uint64 `gorm:"primarykey" json:"id"`                  // 主键ID
	CreateTime int64  `gorm:"column:create_time" json:"create_time"` // 创建时间
	UpdateTime int64  `gorm:"column:update_time" json:"update_time"` // 更新时间
}
