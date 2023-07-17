// Package main rssbot main package
package main

import (
	"context"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/forest33/rssbot/business/entity"
	"github.com/forest33/rssbot/business/usecase"
	"github.com/forest33/rssbot/pkg/database"
	"github.com/forest33/rssbot/pkg/logger"

	db "github.com/forest33/rssbot/adapter/database"
)

var (
	zlog    *logger.Zerolog
	homeDir string
)

var (
	// AppName application name
	AppName string
	// AppVersion application version
	AppVersion string
	// AppURL application homepage
	AppURL = "https://github.com/forest33/rssbot"
	// BuiltAt build date
	BuiltAt string
)

var (
	cfg = &entity.Config{}
	dbi *database.Database

	feedsRepo         *db.FeedsRepository
	usersRepo         *db.UsersRepository
	subscriptionsRepo *db.SubscriptionsRepository

	parserUseCase *usecase.ParserUseCase
	botUseCase    *usecase.BotUseCase

	ctx    context.Context
	cancel context.CancelFunc
)

const (
	applicationName = "rssbot"
)

func init() {
	if cfg.Runtime.GoMaxProcs == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	} else {
		runtime.GOMAXPROCS(cfg.Runtime.GoMaxProcs)
	}

	ctx, cancel = context.WithCancel(context.Background())
}

func main() {
	defer shutdown()

	zlog = logger.NewZerolog(logger.ZeroConfig{
		Level:             cfg.Logger.Level,
		TimeFieldFormat:   cfg.Logger.TimeFieldFormat,
		PrettyPrint:       cfg.Logger.PrettyPrint,
		DisableSampling:   cfg.Logger.DisableSampling,
		RedirectStdLogger: cfg.Logger.RedirectStdLogger,
		ErrorStack:        cfg.Logger.ErrorStack,
		ShowCaller:        cfg.Logger.ShowCaller,
	})

	initDatabase()
	initAdapters()
	initClients()
	initUseCases()

	botUseCase.Start()
	parserUseCase.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func initAdapters() {
	feedsRepo = db.NewFeedsRepository(dbi, zlog)
	usersRepo = db.NewUsersRepository(dbi, zlog)
	subscriptionsRepo = db.NewSubscriptionsRepository(dbi, zlog)
}

func initClients() {
}

func initUseCases() {
	var err error

	parserUseCase, err = usecase.NewParserUseCase(ctx, cfg.Parser, zlog, feedsRepo)
	if err != nil {
		zlog.Fatal(err)
	}
	botUseCase, err = usecase.NewBotUseCase(ctx, cfg.Bot, zlog, dbi, feedsRepo, usersRepo, subscriptionsRepo)
	if err != nil {
		zlog.Fatal(err)
	}

	usecase.SetParserUseCase(parserUseCase)
	usecase.SetBotUseCase(botUseCase)
}

func shutdown() {
	cancel()
	dbi.Close()
}
