package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"github.com/things-go/log"
	"github.com/things-go/ormat/pkg/deploy"
	"github.com/things-go/ormat/view"
)

// Global global config
var Global = &Config{}

// Config custom config
type Config struct {
	Deploy     string            `yaml:"deploy" json:"deploy" binding:"oneof=local dev debug uat prod"` // 布署环境
	Database   *Database         `yaml:"database" json:"database"`                                      // 数据库连接信息
	OutDir     string            `yaml:"outDir" json:"outDir" binding:"required"`                       // 文件输出路径
	TypeDefine map[string]string `yaml:"typeDefine" json:"typeDefine"`                                  // 自定义数据类型
	TableNames []string          `yaml:"tableNames" json:"tableNames"`                                  // 指定输出表
	View       view.Config       `yaml:"view" json:"view"`
}

func init() {
	viper.SetDefault("deploy", "prod")
	viper.SetDefault("outDir", "./model")
	viper.SetDefault("view.dbTag", "gorm")
	viper.SetDefault("view.isNullToPoint", true)
	viper.SetDefault("view.isCommentTag", true)
	viper.SetDefault("view.webTags", []view.WebTag{
		{
			Kind:    view.WebTagSnakeCase,
			Tag:     "json",
			HasOmit: true,
		},
	})
	viper.SetDefault("view.protobuf", view.Protobuf{
		Enabled: false,
		Merge:   false,
		Dir:     "./model",
		Package: "typing",
		Options: map[string]string{"go_package": "typing"},
	})
}

func (cc *Config) Load() error {
	err := viper.Unmarshal(cc, func(c *mapstructure.DecoderConfig) { c.TagName = "yaml" })
	if err != nil {
		return err
	}
	if cc.Database != nil {
		err = cc.Database.Parse()
		if err != nil {
			return err
		}
	}
	validate := validator.New()
	validate.SetTagName("binding")
	err = validate.Struct(cc)
	if err != nil {
		return err
	}
	deploy.MustSetDeploy(cc.Deploy)
	log.ReplaceGlobals(log.NewLogger(log.WithConfig(log.Config{Level: "info", Format: "console"})))
	return nil
}
