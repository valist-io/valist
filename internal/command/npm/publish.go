package npm

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"

	"github.com/valist-io/registry/internal/config"
	"github.com/valist-io/registry/internal/core/client"
	"github.com/valist-io/registry/internal/registry/npm"
)

func NewPublishCommand() *cli.Command {
	return &cli.Command{
		Name:  "publish",
		Usage: "Publish an npm package",
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				cli.ShowSubcommandHelpAndExit(c, 1)
			}

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

			client, err := client.NewClientWithMetaTx(c.Context, cfg, account)
			if err != nil {
				return err
			}
			defer client.Close()

			registryAddr := "localhost:9000"
			registryPath := fmt.Sprintf("http://%s/@%s", registryAddr, c.Args().Get(0))

			// run registry in background
			go http.ListenAndServe(registryAddr, npm.NewHandler(client))

			cmd := exec.Command("npm", "publish", "--registry", registryPath)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			return cmd.Run()
		},
	}
}
