package config

import (
	"encoding/json"
	"flag"
	"os"
)

type Config struct {
	Evn    string `json:"env"`
	Server Server `json:"server"`
}
type Server struct {
	Port string `json:"port"`
	Host string `json:"host"`
}

var Cng Config

var (
	env  = flag.String("ENV", " ", "env")
	port = flag.String("PORT", " ", "port")
	host = flag.String("HOST", " ", "host")
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
	switch {
	case *env != " ":
		cng.Evn = *env
	case *port != " ":
		cng.Server.Port = *port
	case *host != " ":
		cng.Server.Host = *host
	}
}

func init() {
	mostParseFile("./config/auth/config.json", &Cng)
	parseFlags(&Cng)
}
