package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env              string `yaml:"env" env-default:"local"`
	StoragePath      string `yaml:"storage_path" env-required:"true"`
	TelegramBotToken string `yaml:"tgbot_token" env-required:"true"`
	NotionToken      string `yaml:"notion_token"`
	OurDiary         `yaml:"ourdiary"`
}

type OurDiary struct {
	Token      string `yaml:"token"`
	TasksId    string `yaml:"tasks_id"`
	ScheduleId string `yaml:"schedule_id"`
	Password   string `yaml:"password"`
	First      First  `yaml:"first"`
	Second     Second `yaml:"second"`
}

type First Person
type Second Person

type Person struct {
	PageId string `yaml:"page_id"`
	UserId string `yaml:"user_id"`
	Email  string `yaml:"email"`
	Enter  string `yaml:"enter"`
}

func init() {
	if err := godotenv.Load("local.env"); err != nil {
		log.Print("No .env file found")
	}
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
