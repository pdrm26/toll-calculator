package main

import (
	"flag"
	"net"
	"net/http"
	"os"

	"github.com/go-kit/log"
	"github.com/pdrm26/toll-calculator/go-kit-aggregator-service/aggsvc/aggendpoint"
	"github.com/pdrm26/toll-calculator/go-kit-aggregator-service/aggsvc/aggservice"
	"github.com/pdrm26/toll-calculator/go-kit-aggregator-service/aggsvc/aggtransport"
)

func main() {
	fs := flag.NewFlagSet("aggsvc", flag.ExitOnError)
	var (
		httpAddr = fs.String("http-addr", ":3000", "HTTP listen address")
	)

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var (
		service     = aggservice.New()
		endpoints   = aggendpoint.New(service)
		httpHandler = aggtransport.NewHTTPHandler(endpoints, logger)
	)

	httpListener, err := net.Listen("tcp", *httpAddr)
	if err != nil {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		os.Exit(1)
	}

	http.Serve(httpListener, httpHandler)

}
