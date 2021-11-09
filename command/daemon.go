package command

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/valist-io/valist/core/client"
	"github.com/valist-io/valist/core/config"
	"github.com/valist-io/valist/registry"
)

// Daemon runs the valist node daemon indefinitely.
func Daemon(ctx context.Context) error {
	config := ctx.Value(ConfigKey).(*config.Config)
	client := ctx.Value(ClientKey).(*client.Client)

	fmt.Printf(`

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

	apiServer, err := registry.NewServer(client, config)
	if err != nil {
		return err
	}
	go apiServer.ListenAndServe() //nolint:errcheck

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	fmt.Println("Shutting down")

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	apiServer.Shutdown(ctx) //nolint:errcheck
	return nil
}
