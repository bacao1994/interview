package global

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const prefix = "/interview"

var Config *VecConfig
var Uptime = time.Now().String()

type VecConfig struct {
	Prefix     string
	ApiVersion string
	ServerPort string
	ServerMode string

	JwtValidMinute          int64
	AESJWTKey               string
	HMACCombinePasswordKey  string
	RedisConnectionHost     string
	RedisConnectionPassword string

	MaxLimit   int // max record on a query db.
	DbHost     string
	DbPort     string
	DbUsername string
	DbPassword string
	DbName     string
	DbSSLMode  string
	DbTimeZone string

	DefaultMaxTodo uint
}

func FetchEnvironmentVariables() {
	Config = NewVecConfig()
	Config.GetConfig()
}

func NewVecConfig() *VecConfig {
	cf := VecConfig{}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	return &cf
}

func (cf *VecConfig) GetConfig() {
	port := os.Getenv("PORT")
	cf.ServerPort = port
	ginMode := strings.TrimSpace(strings.ToLower(os.Getenv("GIN_ENV")))
	switch ginMode {
	case gin.DebugMode, gin.ReleaseMode, gin.TestMode:
		gin.SetMode(ginMode)
		cf.ServerMode = ginMode
	default:
		gin.SetMode(gin.DebugMode)
		cf.ServerMode = gin.DebugMode
	}

	cf.Prefix = prefix
	cf.DbHost = os.Getenv("DATABASE_HOST")
	cf.DbPort = os.Getenv("DATABASE_PORT")
	cf.DbUsername = os.Getenv("DATABASE_USERNAME")
	cf.DbPassword = os.Getenv("DATABASE_PASSWORD")
	cf.DbName = os.Getenv("DATABASE_NAME")
	cf.DbSSLMode = os.Getenv("DATABASE_SSL_MODE")
	cf.DbTimeZone = os.Getenv("DATABASE_TIME_ZONE")
	cf.MaxLimit, _ = strconv.Atoi(os.Getenv("DATABASE_QUERY_MAX_LIMIT"))
	cf.AESJWTKey = os.Getenv("AES_JWT_KEY")
	cf.HMACCombinePasswordKey = os.Getenv("HMAC_COMBINE_PASSWORD_KEY")
	cf.ApiVersion = os.Getenv("API_VERSION")

	cf.JwtValidMinute , _ = strconv.ParseInt(os.Getenv("JWT_VALID_MINUTE"), 10, 64)

	cf.RedisConnectionHost = fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	cf.RedisConnectionPassword =  os.Getenv("REDIS_PASSWORD")
}
