package app

import (
	"context"
	"github.com/Jumaniyozov/go-rest-template/internal/config"
	loggerpkg "github.com/Jumaniyozov/go-rest-template/internal/logger"
	"net"
	"net/http"
	"time"
)

const (
	shutdownTimeout = 10 * time.Second
)

type App struct {
	serviceProvider *serviceProvider
	restServer      *http.Handler
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	logger := loggerpkg.New(cfg).Logger

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

// Start starts up the server
func RunRest(ctx context.Context) error {

}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(a.grpcServer)

	desc.RegisterNoteV1Server(a.grpcServer, a.serviceProvider.NoteImpl(ctx))

	return nil
}

func (a *App) runRestServer() error {
	log.Printf("GRPC server is running on %s", a.serviceProvider.GRPCConfig().Address())

	list, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}
