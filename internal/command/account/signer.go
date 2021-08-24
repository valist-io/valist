package account

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/urfave/cli/v2"

	"github.com/valist-io/registry/internal/core/config"
	"github.com/valist-io/registry/internal/signer"
)

func NewSignerCommand() *cli.Command {
	return &cli.Command{
		Name:  "signer",
		Usage: "Runs a stand-alone signer",
		Action: func(c *cli.Context) error {
			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}

			cfg, err := config.Load(home)
			if err != nil {
				return err
			}

			listener, _, err := signer.StartIPCEndpoint(cfg)
			if err != nil {
				return err
			}
			defer listener.Close()

			// ui.OnSignerStartup(core.StartupInfo{
			// 	Info: map[string]interface{}{
			// 		"intapi_version": core.InternalAPIVersion,
			// 		"extapi_version": core.ExternalAPIVersion,
			// 		"extapi_ipc":     cfg.Signer.IPCAddress,
			// 		"extapi_http":    "",
			// 	},
			// })

			quit := make(chan os.Signal, 1)
			signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

			<-quit
			fmt.Println("Shutting down")

			return nil
		},
	}
}
