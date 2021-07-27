package command

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/valist-io/registry/internal/http"
	"github.com/valist-io/registry/internal/impl"
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

var bindAddr = ":8080"

func NewDaemonCommand() *cli.Command {
	return &cli.Command{
		Name:  "daemon",
		Usage: "Runs a persistent Valist relay",
		Action: func(c *cli.Context) error {
			// TODO do we want a default transactor?
			client, err := impl.NewClient(c.Context, nil)
			if err != nil {
				return err
			}
			fmt.Println(banner)

			server := http.NewServer(client, bindAddr)
			fmt.Println("Server running on", bindAddr)
			go server.ListenAndServe()

			quit := make(chan os.Signal, 1)
			signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

			<-quit
			fmt.Println("Shutting down")

			ctx, cancel := context.WithTimeout(c.Context, 30*time.Second)
			defer cancel()

			return server.Shutdown(ctx)
		},
	}
}
