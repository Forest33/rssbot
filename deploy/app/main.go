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
	log *logger.Zerolog
	cfg = &entity.Config{}
	dbi *database.Database

	feedsRepo         *db.FeedsRepository
	feedItemsRepo     *db.FeedItemsRepository
	usersRepo         *db.UsersRepository
	subscriptionsRepo *db.SubscriptionsRepository

	ctx    context.Context
	cancel context.CancelFunc
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

	log = logger.NewZerolog(logger.ZeroConfig{
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
	initUseCases()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func initAdapters() {
	feedsRepo = db.NewFeedsRepository(dbi, log)
	feedItemsRepo = db.NewFeedItemsRepository(dbi, log)
	usersRepo = db.NewUsersRepository(dbi, log)
	subscriptionsRepo = db.NewSubscriptionsRepository(dbi, log)
}

func initUseCases() {
	parserUseCase, err := usecase.NewParserUseCase(ctx, cfg.Parser, log, feedsRepo)
	if err != nil {
		log.Fatal(err)
	}

	botUseCase, err := usecase.NewBotUseCase(ctx, cfg.Bot, log, dbi, feedsRepo, feedItemsRepo, usersRepo, subscriptionsRepo)
	if err != nil {
		log.Fatal(err)
	}

	cleanerUseCase, err := usecase.NewCleanerUseCase(ctx, cfg.Cleaner, log, feedsRepo, feedItemsRepo)
	if err != nil {
		log.Fatal(err)
	}

	botUseCase.Start()
	parserUseCase.Start()
	cleanerUseCase.Start()

	usecase.SetParserUseCase(parserUseCase)
	usecase.SetBotUseCase(botUseCase)
}

func shutdown() {
	cancel()
	dbi.Close()
}
