package command

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/valist-io/valist/core"
	"github.com/valist-io/valist/http"
)

// Daemon runs the valist node daemon indefinitely.
func Daemon(ctx context.Context) error {
	config := ctx.Value(ConfigKey).(*core.Config)
	client := ctx.Value(ClientKey).(*core.Client)

	logger.Info(`

@@@  @@@   @@@@@@   @@@       @@@   @@@@@@   @@@@@@@
@@@  @@@  @@@@@@@@  @@@       @@@  @@@@@@@   @@@@@@@
@@!  @@@  @@!  @@@  @@!       @@!  !@@         @@!
!@!  @!@  !@!  @!@  !@!       !@!  !@!         !@!
@!@  !@!  @!@!@!@!  @!!       !!@  !!@@!!      @!!
!@!  !!!  !!!@!!!!  !!!       !!!   !!@!!!     !!!
:!:  !!:  !!:  !!!  !!:       !!:       !:!    !!:
 ::!!:!   :!:  !:!   :!:      :!:      !:!     :!:
  ::::    ::   :::   :: ::::   ::  :::: ::      ::
   :       :   : :  : :: : :  :    :: : :       :

`)

	server, err := http.NewServer(client, config.ApiAddress)
	if err != nil {
		return err
	}

	go server.ListenAndServe() //nolint:errcheck
	logger.Info("server running on %s", server.Addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	logger.Info("shutting down")

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	server.Shutdown(ctx) //nolint:errcheck
	return nil
}
