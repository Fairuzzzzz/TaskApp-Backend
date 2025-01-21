package configs

type Config struct {
	Service Service
	Database Database
}

type Database struct {
	Datasourcename string
}

type Service struct {
	Port string
}