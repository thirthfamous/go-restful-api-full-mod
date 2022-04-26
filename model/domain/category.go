package domain

type Category struct {
	Id   int `gorm:"primaryKey"`
	Name string
}

// TableName overrides the table name used by User to `profiles`
func (Category) TableName() string {
	return "category"
}
