package command

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/valist-io/valist/internal/core/client"
	"github.com/valist-io/valist/internal/core/config"
	"github.com/valist-io/valist/internal/registry"
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

	webAddr := os.Getenv("VALIST_WEB_ADDR")
	if webAddr == "" {
		webAddr = config.HTTP.WebAddr
	}

	webServer := web.NewServer(webAddr)

	go apiServer.ListenAndServe() //nolint:errcheck
	go webServer.ListenAndServe() //nolint:errcheck

	fmt.Println("Web server running on", webAddr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	fmt.Println("Shutting down")

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	apiServer.Shutdown(ctx) //nolint:errcheck
	webServer.Shutdown(ctx) //nolint:errcheck
	return nil
}
