package global

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/militch/exswap/config"
	"github.com/spf13/viper"
)

var (
	GVA_VP     *viper.Viper
	GVA_DB     *gorm.DB
	GVA_CONFIG *config.Server
)
