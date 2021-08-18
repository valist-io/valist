package command

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/registry/internal/config"
	"github.com/valist-io/registry/internal/core/client"
	"github.com/valist-io/registry/internal/http"
	"github.com/valist-io/registry/web"
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
		Name:  "daemon",
		Usage: "Runs a relay node",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "account",
				Value: "default",
				Usage: "Account to authenticate with",
			},
		},
		Action: func(c *cli.Context) error {
			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}

			cfg, err := config.Load(home)
			if err != nil {
				return err
			}

			var account accounts.Account
			if c.IsSet("account") {
				account.Address = common.HexToAddress(c.String("account"))
			} else {
				account.Address = cfg.Accounts.Default
			}

			client, err := client.NewClient(c.Context, cfg, account)
			if err != nil {
				return err
			}
			defer client.Close()
			
			fmt.Println(banner)
			fmt.Println("Api server running on", cfg.HTTP.ApiAddr)
			fmt.Println("Web server running on", cfg.HTTP.WebAddr)

			apiServer := http.NewServer(client, cfg.HTTP.ApiAddr)
			webServer := web.NewServer(cfg.HTTP.WebAddr)

			go webServer.ListenAndServe() //nolint:errcheck
			go apiServer.ListenAndServe() //nolint:errcheck

			quit := make(chan os.Signal, 1)
			signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

			<-quit
			fmt.Println("Shutting down")

			ctx, cancel := context.WithTimeout(c.Context, 30*time.Second)
			defer cancel()

			apiServer.Shutdown(ctx)
			webServer.Shutdown(ctx)

			return nil
		},
	}
}
