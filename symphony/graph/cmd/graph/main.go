// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	stdlog "log"
	"net"
	"net/http"
	"os"
	"syscall"

	"github.com/facebookincubator/symphony/cloud/ctxgroup"
	"github.com/facebookincubator/symphony/cloud/ctxutil"
	"github.com/facebookincubator/symphony/cloud/log"
	"github.com/facebookincubator/symphony/cloud/oc"
	"github.com/facebookincubator/symphony/cloud/server"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jessevdk/go-flags"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type cliFlags struct {
	HTTPAddress string     `env:"HTTP_ADDRESS" long:"http-address" default:":http" description:"the http address to listen on"`
	GRPCAddress string     `env:"GRPC_ADDRESS" long:"grpc-address" default:":https" description:"the grpc address to listen on"`
	MySQL       string     `env:"MYSQL_DSN" long:"mysql-dsn" description:"connection string to mysql"`
	Log         log.Config `group:"log" namespace:"log" env-namespace:"LOG"`
	Census      oc.Options `group:"oc" namespace:"oc" env-namespace:"OC"`
}

func main() {
	var cf cliFlags
	if _, err := flags.Parse(&cf); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		}
		os.Exit(1)
	}
	app, _, err := NewApplication(&cf)
	if err != nil {
		stdlog.Fatal(err)
	}

	ctx := ctxutil.WithSignal(context.Background(), os.Interrupt, syscall.SIGTERM)
	app.Info("starting application",
		zap.String("http", cf.HTTPAddress),
		zap.String("grpc", cf.GRPCAddress),
	)
	err = app.run(ctx)
	app.Info("terminating application", zap.Error(err))
}

type application struct {
	*zap.Logger
	http struct {
		*server.Server
		addr string
	}
	grpc struct {
		*grpc.Server
		addr string
	}
}

func (app *application) run(ctx context.Context) error {
	g := ctxgroup.WithContext(ctx)
	g.Go(func(context.Context) error {
		err := app.http.ListenAndServe(app.http.addr)
		if err != nil && err != http.ErrServerClosed {
			return errors.Wrap(err, "starting http server")
		}
		return nil
	})
	g.Go(func(context.Context) error {
		lis, err := net.Listen("tcp", app.grpc.addr)
		if err != nil {
			return errors.Wrap(err, "creating grpc listener")
		}
		if err = app.grpc.Serve(lis); err != nil && err != grpc.ErrServerStopped {
			return errors.Wrap(err, "starting grpc server")
		}
		return nil
	})
	<-ctx.Done()

	g.Go(func(context.Context) error {
		app.grpc.GracefulStop()
		return nil
	})
	g.Go(func(context.Context) error {
		_ = app.http.Shutdown(context.Background())
		return nil
	})
	return g.Wait()
}
