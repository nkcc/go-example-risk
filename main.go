package main

import (
	_ "github.com/crewjam/saml"
	_ "github.com/p4gefau1t/trojan-go"
	_ "github.com/prometheus/client_golang/api"
	_ "github.com/rehacktive/caffeine/database"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("hello")
	logrus.Info(aws_secret)
}
