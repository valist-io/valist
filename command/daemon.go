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
	"github.com/valist-io/valist/http"
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

	server, err := http.NewServer(client, config)
	if err != nil {
		return err
	}

	go server.ListenAndServe() //nolint:errcheck
	fmt.Println("server running on", server.Addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	fmt.Println("shutting down")

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	server.Shutdown(ctx) //nolint:errcheck
	return nil
}
