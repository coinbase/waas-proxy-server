// Package main runs a proxy server that proxies requests to Coinbase WaaS servers.
package main

import (
	"context"

	"github.com/coinbase/waas-proxy-server/internal/proxyserver"
	"go.uber.org/zap"
)

// main starts the proxy server.
func main() {
	log := zap.NewExample()

	proxyServer, err := proxyserver.NewProxyServer(context.Background(), log)
	if err != nil {
		log.Error("Error creating proxy server", zap.Error(err))
	}

	if err := proxyServer.Serve(context.Background()); err != nil {
		log.Error("Error serving proxy server", zap.Error(err))
	}
}
