package main

import (
	"flag"
	"os"

	"github.com/oinume/fitminder/backend/cli"
	"github.com/oinume/fitminder/backend/config"
)

func main() {
	m := &serverMain{
		streams: cli.NewStreams(),
	}
	if err := m.run(os.Args); err != nil {
		cli.WriteError(m.streams.Err, err)
		os.Exit(cli.ExitError)
	}
	os.Exit(cli.ExitOK)
}

type serverMain struct {
	streams *cli.Streams
}

func (m *serverMain) run(args []string) error {
	flagSet := flag.NewFlagSet("server", flag.ContinueOnError)
	flagSet.SetOutput(m.streams.Err)
	var (
	//logLevel        = flag.String("log-level", "info", "Log level")
	)
	if err := flagSet.Parse(args[1:]); err != nil {
		return err
	}

	config.MustProcessDefault()

	//db, err := dburl.Open(config.DefaultVars.XODBURL())
	//if err != nil {
	//	return err
	//}
	//svc := service.New(db)
	//server := http_server.New(db, svc)
	//
	//fmt.Printf("Listening on port %v\n", config.DefaultVars.HTTPPort)
	//return http.ListenAndServe(
	//	fmt.Sprintf(":%d", config.DefaultVars.HTTPPort),
	//	server.NewRouter(),
	//)
	return nil
}
