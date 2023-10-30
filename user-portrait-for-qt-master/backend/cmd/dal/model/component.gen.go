// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameComponent = "component"

// Component mapped from table <component>
type Component struct {
	ComponentID   int64   `gorm:"column:component_id;type:bigint;primaryKey;autoIncrement:true" json:"component_id"` // 组件ID
	ComponentName string  `gorm:"column:component_name;type:text;not null" json:"component_name"`                    // 组件名
	ComponentType int64   `gorm:"column:component_type;type:int;not null;default:-1" json:"component_type"`          // 组件类型
	AppID         int64   `gorm:"column:app_id;type:bigint;not null" json:"app_id"`                                  // 应用ID
	ComponentDesc *string `gorm:"column:component_desc;type:text" json:"component_desc"`                             // 组件描述
}

// TableName Component's table name
func (*Component) TableName() string {
	return TableNameComponent
}
