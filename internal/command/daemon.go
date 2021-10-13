package command

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/valist-io/valist/internal/command/utils/lifecycle"
	"github.com/valist-io/valist/internal/core"
	"github.com/valist-io/valist/internal/core/client"
	"github.com/valist-io/valist/internal/core/config"
	"github.com/valist-io/valist/internal/registry"
	"github.com/valist-io/valist/web"
)

func NewDaemonCommand() *cli.Command {
	return &cli.Command{
		Name:   "daemon",
		Usage:  "Runs a relay node",
		Before: lifecycle.SetupClient,
		Action: func(c *cli.Context) error {
			config := c.Context.Value(core.ConfigKey).(*config.Config)
			client := c.Context.Value(core.ClientKey).(*client.Client)

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

			ctx, cancel := context.WithTimeout(c.Context, 30*time.Second)
			defer cancel()

			apiServer.Shutdown(ctx) //nolint:errcheck
			webServer.Shutdown(ctx) //nolint:errcheck
			return nil
		},
	}
}
