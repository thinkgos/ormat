// Code generated by ormat. DO NOT EDIT.
// version: v0.12.0

package testdata

import (
	"time"
)

// TestData1 公告-面向所有人的消息
type TestData1 struct {
	Id        int64      `gorm:"column:id;autoIncrement:true;not null;primaryKey" json:"id,omitempty"`
	Title     string     `gorm:"column:title;type:varchar(255);not null;index:uk_title_created_at,priority:1;comment:标题" json:"title,omitempty"`         // 标题
	Content   *string    `gorm:"column:content;type:varchar(2048);default:null;comment:内容" json:"content,omitempty"`                                     // 内容
	Value1    float32    `gorm:"column:value1;type:float;not null;default:1;comment:值1,[0:空,1:键1,2:键2,3:键3]" json:"value1,omitempty"`                    // 值1,[0:空,1:键1,2:键2,3:键3]
	Value2    string     `gorm:"column:value2;type:decimal(10,1);not null;default:2;comment:值2,0:空,1:键1,2:键2,3:键3" json:"value2,omitempty"`              // 值2,0:空,1:键1,2:键2,3:键3
	Value3    float64    `gorm:"column:value3;type:double(16,2);not null;default:3;comment:值3" json:"value3,omitempty"`                                  // 值3
	Value4    string     `gorm:"column:value4;type:enum('00','SH');not null;default:00;comment:值4" json:"value4,omitempty"`                              // 值4
	CreatedAt time.Time  `gorm:"column:created_at;type:datetime;not null;index:uk_title_created_at,priority:2;comment:发布时间" json:"created_at,omitempty"` // 发布时间
	UpdatedAt *time.Time `gorm:"column:updated_at;type:datetime;default:null" json:"updated_at,omitempty"`
}

// TableName implement schema.Tabler interface
func (*TestData1) TableName() string {
	return "test_data1"
}

// SelectTestData1 database column name.
var SelectTestData1 = []string{
	"`test_data1`.`id`",
	"`test_data1`.`title`",
	"`test_data1`.`content`",
	"`test_data1`.`value1`",
	"`test_data1`.`value2`",
	"`test_data1`.`value3`",
	"`test_data1`.`value4`",
	"UNIX_TIMESTAMP(`test_data1`.`created_at`) AS `created_at`",
	"IFNULL(UNIX_TIMESTAMP(`test_data1`.`updated_at`), 0) AS `updated_at`",
}

// TestData2 公告-面向所有人的消息
type TestData2 struct {
	Id        int64     `gorm:"column:id;autoIncrement:true;not null;primaryKey" json:"id,omitempty"`
	Title     string    `gorm:"column:title;type:varchar(255);not null;index:uk_title_created_at,priority:1;comment:标题" json:"title,omitempty"`         // 标题
	Content   string    `gorm:"column:content;type:varchar(2048);not null;comment:内容" json:"content,omitempty"`                                         // 内容
	Value1    float32   `gorm:"column:value1;type:float;not null;default:1;comment:值1,0:空,1:键1,2:键2,3:键3" json:"value1,omitempty"`                      // 值1,0:空,1:键1,2:键2,3:键3
	Value2    float32   `gorm:"column:value2;type:float(10,1);not null;default:2;comment:值2,[0:空,1:键1,2:键2,3:键3]" json:"value2,omitempty"`              // 值2,[0:空,1:键1,2:键2,3:键3]
	Value3    float64   `gorm:"column:value3;type:double(16,2);not null;default:3;comment:值3" json:"value3,omitempty"`                                  // 值3
	Value4    string    `gorm:"column:value4;type:enum('00','SH');not null;default:00;comment:值4" json:"value4,omitempty"`                              // 值4
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null;index:uk_title_created_at,priority:2;comment:发布时间" json:"created_at,omitempty"` // 发布时间
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;not null" json:"updated_at,omitempty"`
}

// TableName implement schema.Tabler interface
func (*TestData2) TableName() string {
	return "test_data2"
}

// SelectTestData2 database column name.
var SelectTestData2 = []string{
	"`test_data2`.`id`",
	"`test_data2`.`title`",
	"`test_data2`.`content`",
	"`test_data2`.`value1`",
	"`test_data2`.`value2`",
	"`test_data2`.`value3`",
	"`test_data2`.`value4`",
	"UNIX_TIMESTAMP(`test_data2`.`created_at`) AS `created_at`",
	"UNIX_TIMESTAMP(`test_data2`.`updated_at`) AS `updated_at`",
}
