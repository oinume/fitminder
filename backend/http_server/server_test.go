package http_server

import (
	"os"
	"path/filepath"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/oinume/fitminder/backend/config"
)

func TestMain(m *testing.M) {
	config.MustProcessDefault()
	config.DefaultVars.SetTemplateDir(filepath.Join("..", "..", config.DefaultTemplateDir))
	//dbURL := model.ReplaceToTestDBURL(config.DefaultVars.XODBURL())
	//db, err := dburl.Open(dbURL)
	//if err != nil {
	//	panic(err)
	//}
	//realDB = db
	os.Exit(m.Run())
}
