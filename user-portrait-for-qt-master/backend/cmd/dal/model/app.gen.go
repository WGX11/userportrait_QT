// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameApp = "app"

// App mapped from table <app>
type App struct {
	AppID                  int64   `gorm:"column:app_id;type:bigint;primaryKey;autoIncrement:true" json:"app_id"`       // 应用ID
	AppName                *string `gorm:"column:app_name;type:varchar(256)" json:"app_name"`                           // 应用名
	AveBehaviorDurationMap *string `gorm:"column:ave_behavior_duration_map;type:text" json:"ave_behavior_duration_map"` // 平均使用时长map
	MaxBehaviorDurationMap *string `gorm:"column:max_behavior_duration_map;type:text" json:"max_behavior_duration_map"` // 最大使用时长map
}

// TableName App's table name
func (*App) TableName() string {
	return TableNameApp
}
