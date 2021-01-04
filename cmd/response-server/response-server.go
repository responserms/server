package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/responserms/server/pkg/config"
	"github.com/responserms/server/server"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()

	app.Name = "response-server"
	app.Usage = "response-server"
	app.UsageText = "Start Response Server"

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "config",
			Aliases:     []string{"c"},
			Usage:       "The path to the HCL or JSON file used to configure Response Server.",
			DefaultText: "server.hcl",
			Value:       "server.hcl",
			EnvVars:     []string{"RESPONSE_SERVER_CONFIG"},
		},
		&cli.BoolFlag{
			Name:    "dev",
			Aliases: []string{"d"},
			Usage:   "Start Response Server in zero-configuration development mode.",
			Value:   false,
			EnvVars: []string{"RESPONSE_SERVER_DEV"},
		},
		&cli.BoolFlag{
			Name:  "generate-key",
			Usage: "Generate and print an encryption key for Response Server.",
			Value: false,
		},
		&cli.StringFlag{
			Name:    "level",
			Aliases: []string{"l"},
			Usage:   "Set the log level for Response Server: off, info, warn, error, debug, trace,",
			Value:   "info",
			EnvVars: []string{"RESPONSE_SERVER_LEVEL"},
		},
	}

	app.Action = handleCommand

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("Response Server ran into a problem: %s\n", err.Error())
	}
}

func handleCommand(ctx *cli.Context) error {
	if ctx.Bool("generate-key") {
		return handleGenerateKey()
	}

	cfg, err := createConfig(ctx)
	if err != nil {
		return err
	}

	// set the log level from the --level flag
	cfg.SetLogLevelFromStr(ctx.String("level"))

	svr, err := server.New(cfg)
	if err != nil {
		return err
	}

	svrContext := ctx.Context

	// Start() is not blocking
	errChan := make(chan error, 1)
	svr.Start(svrContext, errChan)

	sigs := make(chan os.Signal, 1)
	stop := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs

		fmt.Println()
		fmt.Println(sig)

		stop <- true
	}()

	go func() {
		err := <-errChan

		fmt.Println()
		fmt.Println(err)

		stop <- true
	}()

	<-stop

	if err := svr.Shutdown(svrContext); err != nil {
		fmt.Printf("failed to shutdown: %s\n", err)
	}

	return nil
}

func createDevelopmentConfig() (*config.Config, error) {
	return config.NewDevelopment()
}

func createStandardConfig(path string) (*config.Config, error) {
	cfg, diags := config.NewFromFile(path)

	if diags.HasErrors() {
		err := diags.WriteText(os.Stdout, 0, true)
		if err != nil {
			return nil, err
		}

		return nil, diags
	}

	return cfg, nil
}

func createConfig(ctx *cli.Context) (*config.Config, error) {
	switch {
	case ctx.Bool("dev"):
		if ctx.String("config") != "server.hcl" {
			return nil, errors.New("the --dev flag does not support using configuration files, remove --config")
		}

		return createDevelopmentConfig()
	default:
		return createStandardConfig(ctx.String("config"))
	}
}

func handleGenerateKey() error {
	genKey, err := config.GenerateEncryptionKey()
	if err != nil {
		return err
	}

	fmt.Println(genKey)

	return nil
}
