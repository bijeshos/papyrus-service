package main

import (
	"context"
	"flag"
	"github.com/bijeshos/papyrus-service/user"
	"net/http"
	"os"

	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	var (
		listen = flag.String("listen", ":8080", "HTTP listen address")
		proxy  = flag.String("proxy", "", "Optional comma-separated list of URLs to proxy uppercase requests")
	)
	flag.Parse()

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "listen", *listen, "caller", log.DefaultCaller)

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "user_group",
		Subsystem: "add_user_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "user_group",
		Subsystem: "add_user_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "user_group",
		Subsystem: "add_user_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{})

	var svc user.Service
	svc = user.User{}
	svc = user.ProxyingMiddleware(context.Background(), *proxy, logger)(svc)
	svc = user.LoggingMiddleware(logger)(svc)
	svc = user.InstrumentingMiddleware(requestCount, requestLatency, countResult)(svc)

	addUserHandler := httptransport.NewServer(
		user.MakeAdduserEndpoint(svc),
		user.DecodeAdduserRequest,
		user.EncodeResponse,
	)

	http.Handle("/add-user", addUserHandler)

	http.Handle("/metrics", promhttp.Handler())
	logger.Log("msg", "HTTP", "addr", *listen)
	logger.Log("err", http.ListenAndServe(*listen, nil))
}
