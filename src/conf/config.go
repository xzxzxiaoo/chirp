package conf

import (
	"chirp/src/databases"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

type Logger struct {
	LogsDir   string
	LogsName  string
	LogsLevel string
}

type Parameter struct {
	ListenPort int `yaml:"listen-port",required:"true"`
	//Logger
	MysqlConnector *databases.DBConfig
	LdapEnable     bool
}

type ApiServerConfiguration struct {
	C_dbConfig *databases.DBConfig  `yaml:"mysql",required:"true"`
	LdapEnable bool                 `yaml:"ldap-enable",required:"true"`
	DBManager  *databases.DBManager `yaml:"-"`
	ListenPort int                  `yaml:"listen-port",required:"true"`
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}
	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	c.watchConfig()

	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		// 如果指定了配置文件，则解析指定的配置文件
		viper.SetConfigFile(c.Name)
	} else {
		// 如果没有指定配置文件，则解析默认的配置文件
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	// 设置配置文件格式为YAML
	viper.SetConfigType("yaml")
	// viper解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file changed: %s\n", e.Name)
	})
}

func Configured() *Parameter {

	if err := Init("conf/config.yaml"); err != nil {
		panic(err)
	}

	mysql := &databases.DBConfig{viper.GetString("common.store.mysql.dbname"),
		viper.GetString("common.store.mysql.address"),
		viper.GetString("common.store.mysql.username"),
		viper.GetString("common.store.mysql.password"),
		true}

	ldap_enable := viper.GetBool("ldap-enable")

	para := &Parameter{viper.GetInt("listen-port"), mysql, ldap_enable}

	return para

}

func (self *ApiServerConfiguration) InitDatabaseConnecter() {
	var err error
	self.DBManager, err = databases.New(self.C_dbConfig)
	if err != nil {
		panic(err)
	}

}
