package command

import (
	"context"
	"fmt"
	"net/http"
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

const banner = `

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

`

func NewDaemonCommand() *cli.Command {
	return &cli.Command{
		Name:   "daemon",
		Usage:  "Runs a relay node",
		Before: lifecycle.SetupClient,
		Action: func(c *cli.Context) error {
			config := c.Context.Value(core.ConfigKey).(*config.Config)
			client := c.Context.Value(core.ClientKey).(*client.Client)

			fmt.Println(banner)

			addr := os.Getenv("VALIST_HTTP_ADDR")
			if addr == "" {
				addr = config.HTTP.BindAddr
			}

			handler := http.NewServeMux()
			handler.Handle("/api/", registry.NewHandler(client))
			handler.Handle("/", web.NewHandler())

			server := &http.Server{
				Addr:    addr,
				Handler: handler,
			}

			go server.ListenAndServe() //nolint:errcheck

			fmt.Println("HTTP server running on", addr)

			quit := make(chan os.Signal, 1)
			signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

			<-quit
			fmt.Println("Shutting down")

			ctx, cancel := context.WithTimeout(c.Context, 30*time.Second)
			defer cancel()

			server.Shutdown(ctx) //nolint:errcheck
			return nil
		},
	}
}
