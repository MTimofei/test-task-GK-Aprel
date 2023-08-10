package config

import (
	"encoding/json"
	"flag"
	"os"
)

type Config struct {
	Evn      string   `json:"env"`
	Server   Addr     `json:"server"`
	Postgres Postgres `json:"postgres"`
}
type Addr struct {
	Port string `json:"port"`
	Host string `json:"host"`
}

type Postgres struct {
	Addr     Addr   `json:"addr"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

var Cng Config

var (
	env              = flag.String("ENV", " ", "env")
	serverPort       = flag.String("SERVER-PORT", " ", "port")
	serverHost       = flag.String("SERVER-HOST", " ", "host")
	postgresUser     = flag.String("POSTGRES-USER", " ", "user")
	postgresPassword = flag.String("POSTGRES-PASSWORD", " ", "password")
	postgresDatabase = flag.String("POSTGRES-DATABASE", " ", "database")
	postgresHost     = flag.String("POSTGRES-HOST", " ", "host")
	postgresPort     = flag.String("POSTGRES-PORT", " ", "port")
)

// читаем содержимое конфиг файла
// если возникла ошибка паникуем
func mostParseFile(pathFile string, cng *Config) {
	f, err := os.Open(pathFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(cng)
	if err != nil {
		panic(err)
	}
}

// читаем флаги и заменяем значения, есле указаны
func parseFlags(cng *Config) {
	flag.Parse()

	if *env != " " {
		cng.Evn = *env
	}

	if *postgresPort != " " {
		cng.Server.Port = *serverPort
	}

	if *postgresHost != " " {
		cng.Server.Host = *serverHost
	}

	if *postgresDatabase != " " {
		cng.Postgres.Database = *postgresDatabase
	}

	if *postgresUser != " " {
		cng.Postgres.User = *postgresUser
	}

	if *postgresPassword != " " {
		cng.Postgres.Password = *postgresPassword
	}

	if *postgresPort != " " {
		cng.Postgres.Addr.Port = *postgresPort
	}

	if *postgresHost != " " {
		cng.Postgres.Addr.Host = *postgresHost
	}

}

func init() {
	mostParseFile("./config/auth/config.json", &Cng)
	parseFlags(&Cng)
}
