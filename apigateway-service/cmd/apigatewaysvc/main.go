package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env"
	"github.com/gorilla/mux"
	"github.com/oklog/oklog/pkg/group"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	newsfeederTransports "github.com/michaljirman/newsapp/newsfeeder-service/pkg/transports"
)

// Config - main config struct
type Config struct {
	ServiceName   string `env:"SVC_NAME"`
	DebugAddr     string `env:"DEBUG_ADDRESS"`
	HTTPAddr      string `env:"HTTP_ADDRESS"`
	NewsFeederSvc string `env:"NEWSFEEDERSVC"`
}

// Get - Loads config from env
func Get() (*Config, error) {
	c := &Config{}
	if err := env.Parse(c); err != nil {
		return nil, errors.Wrap(err, "failed to load config")
	}

	return c, nil
}

func main() {
	logger := logrus.New()
	lvl, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		lvl = "debug"
	}
	ll, err := logrus.ParseLevel(lvl)
	if err != nil {
		ll = logrus.DebugLevel
	}
	logger.SetReportCaller(true)
	logger.SetLevel(ll)

	cfg, err := Get()
	if err != nil {
		logger.Errorf("Failed to load config: %+v", err)
	}

	// ctx := context.Background()
	r := mux.NewRouter()

	logger.Debugf("grpc dial .... %s", cfg.NewsFeederSvc)

	conn, err := grpc.Dial(cfg.NewsFeederSvc, grpc.WithInsecure())
	if err != nil {
		logger.Errorf("%+v", errors.Wrap(err, "failed to connect"))
		os.Exit(1)
	}
	defer conn.Close()

	endpoints := newsfeederTransports.NewGRPCClient(conn, logger)

	r.PathPrefix("/newsfeeder").Handler(http.StripPrefix("/newsfeeder", newsfeederTransports.NewHTTPHandler(endpoints, logger)))

	var g group.Group
	{
		// The debug listener mounts the http.DefaultServeMux, and serves up
		// stuff like the Prometheus metrics route, the Go debug and profiling
		// routes, and so on.
		debugListener, err := net.Listen("tcp", cfg.DebugAddr)
		if err != nil {
			logger.Debugf("transport=%s during=%s err=%+v", "debug/HTTP", "Listen", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Debugf("transport=%s addr=%s", "debug/HTTP", cfg.DebugAddr)
			return http.Serve(debugListener, http.DefaultServeMux)
		}, func(error) {
			debugListener.Close()
		})
	}
	{
		// The HTTP listener mounts the Go kit HTTP handler we created.
		httpListener, err := net.Listen("tcp", cfg.HTTPAddr)
		if err != nil {
			logger.Debugf("transport=%s during=%s err=%+v", "HTTP", "Listen", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Debugf("transport=%s addr=%s", "HTTP", cfg.HTTPAddr)
			return http.Serve(httpListener, r)
		}, func(error) {
			httpListener.Close()
		})
	}
	{
		// This function just sits and waits for ctrl-C.
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}
	logger.Debugf("group exited %+v", g.Run())
}
