package lib

import (
	"log"

	"github.com/spf13/viper"
)

// Env has environment stored
type Env struct {
	ServerPort           string `mapstructure:"PORT"`
	Environment          string `mapstructure:"ENV"`
	LogOutput            string `mapstructure:"LOG_OUTPUT"`
	LogLevel             string `mapstructure:"LOG_LEVEL"`
	DBUsername           string `mapstructure:"DB_USER"`
	DBPassword           string `mapstructure:"DB_PASS"`
	DBHost               string `mapstructure:"DB_HOST"`
	DBPort               string `mapstructure:"DB_PORT"`
	DBName               string `mapstructure:"DB_NAME"`
	JWTSecret            string `mapstructure:"JWT_SECRET"`
	JWTRefreshSecret     string `mapstructure:"JWT_REFRESH_SECRET"`
	NodeNumber           int64  `mapstructure:"NODE_NUMBER"`
	SupabaseUrl          string `mapstructure:"SUPABASE_URL"`
	SupabaseKey          string `mapstructure:"SUPABASE_KEY"`
	SaltSize             int    `mapstructure:"PASS_SALT_SIZE"`
	HashAlgorithm        string `mapstructure:"PASS_HASH_ALGORITHM"`
	TokenTTL             int    `mapstructure:"TOKEN_TIMEOUT"`
	RefreshTokenTokenTTL int    `mapstructure:"REFRESH_TOKEN_TIMEOUT"` // tính bằng giây
}

// NewEnv creates a new environment
func NewEnv() Env {

	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("☠️ cannot read configuration")
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("☠️ environment can't be loaded: ", err)
	}

	return env
}
