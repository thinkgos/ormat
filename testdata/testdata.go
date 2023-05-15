// Code generated by ormat. DO NOT EDIT.
// version: v0.9.2

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

/* protobuf field helper
// TestData1 公告-面向所有人的消息
message TestData1 {
  int64 id = 1 [(grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field) = { type: [ INTEGER ] }, (things_go.seaql.field) = { type: "bigint NOT NULL AUTO_INCREMENT" }];
  // 标题
  string title = 2 [(things_go.seaql.field) = { type: "varchar(255) NOT NULL DEFAULT ''" }];
  // 内容
  string content = 3 [(things_go.seaql.field) = { type: "varchar(2048) DEFAULT NULL" }];
  // 值1,[0:空,1:键1,2:键2,3:键3]
  float value1 = 4 [(things_go.seaql.field) = { type: "float NOT NULL DEFAULT '1'" }];
  // 值2,0:空,1:键1,2:键2,3:键3
  string value2 = 5 [(things_go.seaql.field) = { type: "decimal(10,1) NOT NULL DEFAULT '2'" }];
  // 值3
  double value3 = 6 [(things_go.seaql.field) = { type: "double(16,2) NOT NULL DEFAULT '3'" }];
  // 值4
  string value4 = 7 [(things_go.seaql.field) = { type: "enum('00','SH') NOT NULL DEFAULT '00'" }];
  // 发布时间
  int64 created_at = 8 [(grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field) = { type: [ INTEGER ] }, (things_go.seaql.field) = { type: "datetime NOT NULL" }];
  int64 updated_at = 9 [(grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field) = { type: [ INTEGER ] }, (things_go.seaql.field) = { type: "datetime DEFAULT NULL" }];
}
*/
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

/* protobuf field helper
// TestData2 公告-面向所有人的消息
message TestData2 {
  int64 id = 1 [(grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field) = { type: [ INTEGER ] }, (things_go.seaql.field) = { type: "bigint NOT NULL AUTO_INCREMENT" }];
  // 标题
  string title = 2 [(things_go.seaql.field) = { type: "varchar(255) NOT NULL DEFAULT ''" }];
  // 内容
  string content = 3 [(things_go.seaql.field) = { type: "varchar(2048) NOT NULL DEFAULT ''" }];
  // 值1,0:空,1:键1,2:键2,3:键3
  float value1 = 4 [(things_go.seaql.field) = { type: "float NOT NULL DEFAULT '1'" }];
  // 值2,[0:空,1:键1,2:键2,3:键3]
  float value2 = 5 [(things_go.seaql.field) = { type: "float(10,1) NOT NULL DEFAULT '2'" }];
  // 值3
  double value3 = 6 [(things_go.seaql.field) = { type: "double(16,2) NOT NULL DEFAULT '3'" }];
  // 值4
  string value4 = 7 [(things_go.seaql.field) = { type: "enum('00','SH') NOT NULL DEFAULT '00'" }];
  // 发布时间
  int64 created_at = 8 [(grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field) = { type: [ INTEGER ] }, (things_go.seaql.field) = { type: "datetime NOT NULL" }];
  int64 updated_at = 9 [(grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field) = { type: [ INTEGER ] }, (things_go.seaql.field) = { type: "datetime NOT NULL" }];
}
*/
