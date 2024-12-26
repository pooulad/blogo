package config

const (
	Development environment = "development"
	Production  environment = "production"
	SslMode                 = "disable"
)

type Config struct {
	Environment environment `json:"environment"`
	AppUrl      string      `json:"app_url"`
	Port        string      `json:"port"`
	DB          DB          `json:"db"`
}

type DB struct {
	Postgresql Postgresql `json:"postgresql"`
}

type Postgresql struct {
	Port     string `json:"port"`
	Host     string `json:"host"`
	DbName   string `json:"dbname"`
	Username string `json:"username"`
	Password string `json:"password"`
	SslMode  string `json:"sslmode"`
}

type environment string

// func (e environment) String() string {
// 	return string(e)
// }

// func (e *environment) UnmarshalJSON(data []byte) error {
// 	var v interface{}
// 	if err := json.Unmarshal(data, &v); err != nil {
// 		return err
// 	}

// 	switch t := v.(type) {
// 	case string:
// 		env := environment(t)
// 		if env != Development && env != Production {
// 			return fmt.Errorf("unsupported environment %s. Options are: %s - %s", env, Production, Development)
// 		}
// 		*e = env
// 		return nil

// 	default:
// 		return fmt.Errorf("environment must string format")
// 	}
// }
