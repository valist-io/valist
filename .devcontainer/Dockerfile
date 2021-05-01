FROM ubuntu:20.04

# Install global dependencies
RUN apt update && apt install -y git curl zsh python3 python3-pip && \
  pip3 install slither-analyzer

# Create devcontainer user
ARG USERNAME=user
ARG USER_UID=1000
ARG USER_GID=$USER_UID
RUN groupadd --gid $USER_GID $USERNAME && \
    useradd --uid $USER_UID --gid $USER_GID --shell /usr/bin/zsh -m $USERNAME \
    #
    # [Optional] Add sudo support. Omit if you don't need to install software after connecting.
    && apt install -y sudo \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME
USER $USERNAME

# Install Oh My Zsh
RUN curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh | zsh
ENV DISABLE_AUTO_UPDATE=true

# Install Node.js
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.38.0/install.sh | zsh && \
    NVM_DIR="$HOME/.nvm" && . "$NVM_DIR/nvm.sh" && nvm install 14
    # Install npm-based global dependencies
    #npm install -g eslint
