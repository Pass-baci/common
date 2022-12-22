package common

import (
	"github.com/asim/go-micro/v3/util/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"strconv"
)

func PrometheusBoot(port int) {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		if err := http.ListenAndServe("0.0.0.0:"+strconv.Itoa(port), nil); err != nil {
			log.Fatalf("启动Prometheus失败 err: %s", err.Error())
		}
		log.Info("启动Prometheus成功，端口为：" + strconv.Itoa(port))
	}()
}
