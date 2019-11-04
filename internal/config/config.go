package config

import (
	"gopkg.in/gomail.v2"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"software-theory-final/internal/log"
	"software-theory-final/pkg/modules/mysql"
	"software-theory-final/pkg/modules/redis"
)

type HttpConf struct {
	Port        int      `yaml:"port"`
	Host        string   `yaml:"host"`
	AllowOrigin []string `yaml:"allow-origin"`
}

type MySQLConf struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"dbName"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type RedisConf struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   int    `yaml:"dbName"`
}

type MailConf struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type ServerConf struct {
	MySQLDB *MySQLConf `yaml:"mysql"`
	RedisDB *RedisConf `yaml:"redis"`
	Http    *HttpConf  `yaml:"http"`
	Mail    *MailConf  `yaml:"mail"`
}

func (c *ServerConf) ErrHandler(op string, err error) {
	if err != nil {
		log.Fatalf("%s found error: %v", op, err)
	}
}

func (c *ServerConf) GetConfig() *ServerConf {
	log.Info("Loading configs")
	ymlFile, err := ioutil.ReadFile("config/config.yaml")
	c.ErrHandler("openning file", err)

	err = yaml.Unmarshal(ymlFile, c)
	c.ErrHandler("unmarshal yaml", err)

	return c
}

var (
	mysqlClient *mysql.Client
	redisClient *redis.Client
	mailDialer  *gomail.Dialer
)

func GetMySQLClient() *mysql.Client {
	return mysqlClient
}

func GetRedisClient() *redis.Client {
	return redisClient
}

func GetMailDialer() *gomail.Dialer {
	return mailDialer
}

func InitMySQLClient(host string, port int, dbname, username, password string) error {
	log.Infof("trying connecting MySQL:%s:%d %s %s %s", host, port, dbname, username, password)
	mysqlClient = &mysql.Client{
		Host:     host,
		Port:     port,
		DbName:   dbname,
		Username: username,
		Password: password,
	}
	return mysqlClient.NewMySQLClient()
}

func InitRedisClient(host string, port int, username, password string, dbName int) error {
	log.Infof("trying connecting Redis:%s:%d %s %s", host, port, username, password)
	redisClient = &redis.Client{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		DBName:   dbName,
	}
	return redisClient.NewRedisClient()
}

func InitDB(config ServerConf) {
	host := config.MySQLDB.Host
	port := config.MySQLDB.Port
	dbName := config.MySQLDB.DBName
	username := config.MySQLDB.Username
	password := config.MySQLDB.Password

	err := InitMySQLClient(host, port, dbName, username, password)
	if err != nil {
		log.Errorf("Err connecting MySQL:%v", config.MySQLDB)
	}

	host = config.RedisDB.Host
	port = config.RedisDB.Port
	dbId := config.RedisDB.DBName
	username = config.RedisDB.Username
	password = config.RedisDB.Password
	err = InitRedisClient(host, port, username, password, dbId)
	if err != nil {
		log.Errorf("Error connecting Redis:%v", config.RedisDB)
	}
}

func InitScheme() {
	mysql.MigrateTables(mysqlClient)
}

func InitMailClient(conf *MailConf) {
	log.Infof("Trying connecting mail server:%s:%d %s", conf.Host, conf.Port, conf.Username)
	mailDialer = &gomail.Dialer{
		Host:     conf.Host,
		Port:     conf.Port,
		Username: conf.Username,
		Password: conf.Password,
		SSL:      true,
	}
}
