package main

import (
	"github.com/forest33/rssbot/deploy/app/migrations"
	"github.com/forest33/rssbot/pkg/database"
)

//go:generate go-bindata -o ./migrations/migrations.bindata.go -pkg migrations -ignore=\\*.go ./migrations/...

func initDatabase() {
	binDataCfg := &database.BinDataConfig{
		Dir:          database.DefaultMigrationsDir,
		AssetDirFunc: migrations.AssetDir,
		AssetFunc:    migrations.Asset,
	}

	var err error
	dbi, err = database.NewConnector(cfg.Database, binDataCfg, log)
	if err != nil {
		log.Fatal(err.Error())
	}
}
