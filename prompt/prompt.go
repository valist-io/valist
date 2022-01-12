package prompt

import (
	"errors"
	"os"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
)

var ErrNonInteractive = errors.New("prompt in non-interactive environment")

type Prompt struct {
	inner promptui.Prompt
}

// Run returns the value of the prompt. If a CI environment is detected ErrNonInteractive is returned.
func (p Prompt) Run() (string, error) {
	if ci, _ := strconv.ParseBool(os.Getenv("CI")); ci {
		return "", ErrNonInteractive
	}

	return p.inner.Run()
}

// RunFlag returns the value of the flag if set, otherwise the prompt is run normally.
func (p Prompt) RunFlag(c *cli.Context, flag string) (string, error) {
	if c.IsSet(flag) {
		return c.String(flag), nil
	}

	return p.Run()
}

func NewAccountPassphrase() Prompt {
	return Prompt{promptui.Prompt{
		Label:       "New account passphrase",
		Mask:        '*',
		HideEntered: true,
		Validate:    ValidateMinLength(5),
	}}
}

func AccountPrivateKey() Prompt {
	return Prompt{promptui.Prompt{
		Label:       "Hex-encoded ECDSA private key",
		Mask:        '*',
		HideEntered: true,
		Validate:    ValidateMinLength(32),
	}}
}

func AccountPassphrase() Prompt {
	return Prompt{promptui.Prompt{
		Label:       "Account passphrase",
		Mask:        '*',
		HideEntered: true,
		Validate:    ValidateMinLength(5),
	}}
}

func OrganizationName(value string) Prompt {
	return Prompt{promptui.Prompt{
		Label:    "Organization full name",
		Default:  value,
		Validate: ValidateMinLength(1),
	}}
}

func OrganizationDescription(value string) Prompt {
	return Prompt{promptui.Prompt{
		Label:    "Organization description",
		Default:  value,
		Validate: ValidateMinLength(1),
	}}
}

func OrganizationHomepage(value string) Prompt {
	return Prompt{promptui.Prompt{
		Label:   "Organization homepage",
		Default: value,
	}}
}

func RepositoryName(value string) Prompt {
	return Prompt{promptui.Prompt{
		Label:    "Repository full name",
		Default:  value,
		Validate: ValidateMinLength(1),
	}}
}

func RepositoryDescription(value string) Prompt {
	return Prompt{promptui.Prompt{
		Label:    "Repository description",
		Default:  value,
		Validate: ValidateMinLength(1),
	}}
}

func RepositoryHomepage(value string) Prompt {
	return Prompt{promptui.Prompt{
		Label:   "Homepage",
		Default: value,
	}}
}

func RepositoryURL(value string) Prompt {
	return Prompt{promptui.Prompt{
		Label:   "Source code repository url",
		Default: value,
	}}
}

func ReleaseTag(value string) Prompt {
	return Prompt{promptui.Prompt{
		Label:   "Latest release tag",
		Default: value,
	}}
}

func ReleaseMultiArch() Prompt {
	return Prompt{promptui.Prompt{
		Label:     "Are you building for multiple architecures? (y,N)",
		IsConfirm: true,
		Validate:  ValidateYesNo(),
	}}
}

func StatsOptIn() Prompt {
	return Prompt{promptui.Prompt{
		Label:     "Would you like to participate in the Valist package popularity contest? (Y,n)",
		IsConfirm: true,
		Validate:  ValidateYesNo(),
	}}
}
