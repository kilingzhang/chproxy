package main

import "github.com/prometheus/client_golang/prometheus"

var (
	requestSum      *prometheus.CounterVec
	statusCodes     *prometheus.CounterVec
	requestErrors   *prometheus.CounterVec
	requestSuccess  *prometheus.CounterVec
	userTimeouts    *prometheus.CounterVec
	clusterTimeouts *prometheus.CounterVec
)

func init() {
	statusCodes = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "status_codes",
			Help: "Distribution by status codes counter",
		},
		[]string{"host", "code"},
	)

	userTimeouts = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "initial_timeouts",
			Help: "Number of timeouts for initial user",
		},
		[]string{"user", "host"},
	)

	clusterTimeouts = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "execution_timeouts",
			Help: "Number of timeouts for execution user",
		},
		[]string{"user", "host"},
	)

	requestErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "request_errors",
			Help: "Number of errors returned by target. Including amount of timeouts",
		},
		[]string{"user", "cluster_user", "host"},
	)

	requestSum = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "request_sum",
			Help: "Total number of sent requests",
		},
		[]string{"user", "cluster_user", "host"},
	)

	requestSuccess = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "request_success",
			Help: "Total number of sent success requests",
		},
		[]string{"user", "cluster_user", "host"},
	)

	prometheus.MustRegister(statusCodes, userTimeouts, clusterTimeouts, requestErrors,
		requestSum, requestSuccess)
}
