# Getting Started

Valist can be used as a command line interface (CLI) or imported as a TypeScript library.

## Installation

### CLI

#### Package managers

##### NPM

To install the CLI on any system, run the following:

```sh
npm install -g @valist/cli
```

This will download the appropriate Go binary from IPFS. This package is simply a cross-platform wrapper for the Go CLI.

!!! note
    Installing from more package managers coming soon!

##### Homebrew

```sh
brew install valist-io/valist/valist
```

This will add the [homebrew-valist](github.com/valist-io/homebrew-valist) tap and install the CLI automatically.

##### Go Install

You can install the CLI by running the following:

```sh
go install github.com/valist-io/valist
```

!!! note
    You will need Go installed and your PATH configured to include `$GOPATH/bin`. Visit [here](https://www.digitalocean.com/community/tutorials/how-to-build-and-install-go-programs) for more information.

#### Download from valist.io

You can download the latest version of the Valist CLI from app.valist.io/valist/cli -- You'll be able to select your OS and architecture from the `Install` tab or from the sidebar. Place this binary in your PATH or use it directly.

#### Manual installation

Or, you can use the following script to install the binary globally:

```sh
mkdir -p ~/.local/bin
curl https://gateway.valist.io/ipfs/QmZ9T6H7WTb6VrNaqFEwo7Mqj6jGxMe4vpR6srxsjy3otz/linux-amd64/valist -o ~/.local/bin/valist
```

Please note, you'll need to replace `linux-amd64` with your os and architecture. The following options are available:

* linux-amd64
* linux-arm64
* darwin-amd64
* darwin-arm64
* windows-amd64

Finally, ensure your PATH includes the ~/.local/bin folder:

```sh
export PATH="$PATH:$HOME/.local/bin"
echo PATH=\"\$PATH:$HOME/.local/bin\" >> ~/.zshrc # or .bashrc
```

#### Running the daemon

After you have valist installed, simply run the following:

```sh
valist daemon
```

This will start the universal package registry API and web server:

```sh

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


API server running on localhost:9000
```

### TypeScript/JavaScript SDK

You can install the JS-SDK by running the following:

```sh
npm install @valist/sdk
```
