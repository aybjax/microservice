package main

import (
	"go_microservice/encrypt_string"
	"go_microservice/helpers"
	"log"
	"net/http"
	"os"

	kitlog "github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	key := "111023043350789514532147"
	message := "I am A Message"

	log.Println("       Key:", key)
	log.Println("  Original:", message)

	encrypted := encrypt_string.EncryptString(key, message)

	log.Println(" Encrypted:", encrypted)

	log.Println("Descripted:", encrypt_string.DecryptString(key, encrypted))

	log.Println("*********************************************************")

	logger := kitlog.NewLogfmtLogger(os.Stderr)
	field_keys := []string{"method", "error"}

	request_count := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "encryption",
		Subsystem: "my_service",
		Name: "request_count",
		Help: "Number of requests received",
	}, field_keys)

	request_latency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "encryption",
		Subsystem: "my_service",
		Name: "request_latency_microseconds",
		Help: "Total duration of requests in microseconds",
	}, field_keys)

	var svc helpers.EncryptService

	svc = helpers.EncryptServiceInstance{}
	svc = helpers.LoggingMiddleware{
		Logger: logger,
		Next: svc,
	}
	svc = helpers.InstrumentingMiddleware{
		RequestCount: request_count,
		RequestLatency: request_latency,
		Next: svc,
	}

	encryptHandler := httptransport.NewServer(helpers.MakeEncryptEndpoint(svc),
			helpers.DecodeEncryptRequest,
			helpers.EncodeResponse)
	
	decryptHandler := httptransport.NewServer(helpers.MakeDecryptEndpoint(svc),
			helpers.DecodeDecryptRequest,
			helpers.EncodeResponse)
	
	http.Handle("/encrypt", encryptHandler)
	http.Handle("/decrypt", decryptHandler)
	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}