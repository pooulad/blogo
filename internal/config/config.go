package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func New() (*Config, error) {
	var config Config

	// get config from env
	var (
		env        = os.Getenv("ENVIRONMENT")
		app_url    = os.Getenv("APP_URL")
		port       = os.Getenv("PORT")
		dbPort     = os.Getenv("DB_PORT")
		dbName     = os.Getenv("DB_NAME")
		dbHost     = os.Getenv("DB_HOST")
		dbUsername = os.Getenv("DB_USERNAME")
		dbPassword = os.Getenv("DB_PASSWORD")
		dbSslmode  = os.Getenv("DB_SSLMODE")
		configFile = os.Getenv("CONFIG_FILE")
	)

	// check config from command-line
	flag.StringVar((*string)(&config.Environment), "env", env, "application environment: Production or Development mode")
	flag.StringVar(&config.AppUrl, "app_url", app_url, "application url")
	flag.StringVar(&config.Port, "port", port, "application port")
	flag.StringVar(&config.DB.Postgresql.Port, "db_port", dbPort, "database port")
	flag.StringVar(&config.DB.Postgresql.DbName, "db_name", dbName, "database name")
	flag.StringVar(&config.DB.Postgresql.Host, "db_host", dbHost, "database host")
	flag.StringVar(&config.DB.Postgresql.Username, "db_username", dbUsername, "database username")
	flag.StringVar(&config.DB.Postgresql.Password, "db_password", dbPassword, "database password")
	flag.StringVar(&config.DB.Postgresql.SslMode, "db_sslmode", dbSslmode, "database sslmode")
	flag.StringVar(&configFile, "cfg", configFile, "confige file")

	flag.Parse()

	fmt.Println(configFile)
	
	err := readAppConfig(&config, configFile)
	if err != nil {
		return nil, err
	}
	
	fmt.Println(config)

	err = verifyAppConfig(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func readAppConfig(cfg *Config, configFile string) error {
	if configFile == "" {
		return nil
	}

	file, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &cfg)
	if err != nil {
		return err
	}

	return nil
}

func verifyAppConfig(cfg *Config) error {
	// set default environment
	if cfg.Environment == "" {
		cfg.Environment = Production
	}

	if cfg.Port == "" {
		return fmt.Errorf("port must be specified")
	}
	if cfg.AppUrl == "" {
		return fmt.Errorf("app url must be specified")
	}
	port, err := strconv.Atoi(cfg.Port)
	if err != nil || port < 0 || port > 65535 {
		return fmt.Errorf("port must be a number between 0 and 65535")
	}

	if cfg.DB.Postgresql.Port == "" {
		return fmt.Errorf("sql port must be specified")
	}
	if cfg.DB.Postgresql.Host == "" {
		return fmt.Errorf("sql host must be specified")
	}
	if cfg.DB.Postgresql.Username == "" {
		return fmt.Errorf("sql username must be specified")
	}
	if cfg.DB.Postgresql.DbName == "" {
		return fmt.Errorf("sql username must be specified")
	}
	if cfg.DB.Postgresql.SslMode == "" {
		cfg.DB.Postgresql.SslMode = SslMode
	}

	return nil
}
