package config

var configInstance *Config = &Config{
	Log:      Log{},
	Database: Database{},
}

type Config struct {
	Log      Log `yaml:"log"`
	Database `yaml:"database"`
}

type Log struct {
	LogLevel string  `yaml:"loglevel"`
	FileLog  FileLog `yaml:"fileLog"`
}

type Database struct {
	Host     string `yaml:"host"`
	Username string `yaml:"user"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
	DBname   string `yaml:"dbName"`
}

type FileLog struct {
	Folder   string `yaml:"path"`
	FileName string `yaml:"filename"`
	MaxSize  int    `yaml:"maxSize"`
	MaxAge   int    `yaml:"maxAge"`
}
