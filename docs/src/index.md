# Getting Started

## Installation

### CLI

#### Package managers
!!! note
    Installing from package managers coming soon!

#### Go Install

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

```
valist daemon
```

This will start the universal package registry API and web server:

```

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

### SDK

The Valist JS-SDK can be installed by running the following command:

```bash
echo "@valist:registry=http://localhost:9000/api/npm" >> .npmrc

npm install @valist/sdk
```
