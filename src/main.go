package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gopkg.in/unrolled/secure.v1"

	"github.com/devopsfaith/krakend/config"
	"github.com/devopsfaith/krakend/logging"
	"github.com/devopsfaith/krakend/proxy"
	"github.com/devopsfaith/krakend/router/gorilla"
	"github.com/devopsfaith/krakend/router/mux"
)

var (
	LOGGING_TYPE  string
	LOGGING_NAME  string
	CONFIGURATION string
	DEBUG         string
)

func init() {
	LOGGING_TYPE = os.Getenv("LOGGING_TYPE")
	LOGGING_NAME = os.Getenv("LOGGING_NAME")
	CONFIGURATION = os.Getenv("CONFIGURATION")
	DEBUG = os.Getenv("DEBUG")
}

func main() {
	// Loand config json
	parser := config.NewParser()
	config.RoutingPattern = config.BracketsRouterPatternBuilder

	dir, err := os.Getwd()
	if err != nil {
		log.Panicln(err.Error())
	}

	file := fmt.Sprintf("%v/src/%v", dir, CONFIGURATION)

	serviceConfig, err := parser.Parse(file)
	if err != nil {
		log.Fatalf("Config error. Detail: %v\n", err.Error())
	}

	// Config logs level
	debug, _ := strconv.ParseBool(DEBUG)
	serviceConfig.Debug = serviceConfig.Debug || debug
	logger, err := logging.NewLogger(LOGGING_TYPE, os.Stdout, LOGGING_NAME)
	if err != nil {
		log.Fatalf("Log error. Detail: %v\n", err.Error())
	}

	// Middlewares
	secureMiddleware := secure.New(secure.Options{
		STSSeconds:           315360000,
		STSIncludeSubdomains: true,
		STSPreload:           true,
		FrameDeny:            true,
		ContentTypeNosniff:   true,
		BrowserXssFilter:     true,
	})

	cfg := gorilla.DefaultConfig(
		customProxyFactory{
			logger,
			proxy.DefaultFactory(logger),
		}, logger)

	cfg.Middlewares = append(cfg.Middlewares, secureMiddleware)
	routerFactory := mux.NewFactory(cfg)
	routerFactory.New().Run(serviceConfig)
}
