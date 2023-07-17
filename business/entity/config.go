// Package entity provides entities for business logic.
package entity

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	debugEnv = "rssbot_DEBUG"
)

// Config application configuration
type Config struct {
	Database *DatabaseConfig `json:"database"`
	Logger   *LoggerConfig   `json:"logger"`
	Runtime  *RuntimeConfig  `json:"runtime"`
	Bot      *BotConfig      `json:"bot"`
	Parser   *ParserConfig   `json:"parser"`
}

// DatabaseConfig database settings
type DatabaseConfig struct {
	Driver   string `json:"driver" default:"postgres"`
	Host     string `json:"host" default:"127.0.0.1"`
	Port     int    `json:"port" default:"5432"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	SSLMode  string `json:"sslmode" default:"disable"`
}

// LoggerConfig logger settings
type LoggerConfig struct {
	Level             string `json:"level" default:"debug"`
	TimeFieldFormat   string `json:"time_field_format" default:"2006-01-02T15:04:05Z07:00"`
	PrettyPrint       bool   `json:"pretty_print" default:"true"`
	DisableSampling   bool   `json:"disable_sampling" default:"false"`
	RedirectStdLogger bool   `json:"redirect_std_logger" default:"false"`
	ErrorStack        bool   `json:"error_stack" default:"false"`
	ShowCaller        bool   `json:"show_caller" default:"false"`
}

// RuntimeConfig runtime settings
type RuntimeConfig struct {
	GoMaxProcs int `json:"go_max_procs" default:"0"`
}

// BotConfig telegram settings
type BotConfig struct {
	Token                  string `json:"token"`
	UpdateTimeout          int    `json:"update_timeout" default:"60"`
	CommandWorkersPoolSize int    `json:"command_workers_pool_size" default:"5"`
	SenderWorkersPoolSize  int    `json:"sender_workers_pool_size" default:"20"`
	Debug                  bool   `json:"debug" default:"false"`
}

// ParserConfig feed parser settings
type ParserConfig struct {
	WorkersPoolSize     int   `json:"workers_pool_size" default:"5"`
	FeedUpdateFrequency int64 `json:"feed_update_frequency" default:"600"`
}

// Validate validation BotConfig
func (c BotConfig) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Token, validation.Required),
	)
}

// Validate validation ParserConfig
func (c ParserConfig) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.WorkersPoolSize, validation.Required, validation.Min(1)),
		validation.Field(&c.FeedUpdateFrequency, validation.Required, validation.Min(10)),
	)
}
