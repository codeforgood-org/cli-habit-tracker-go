#!/bin/sh
set -e

# Habit Tracker Installation Script
# This script detects the platform and downloads the appropriate binary

REPO="codeforgood-org/cli-habit-tracker-go"
BINARY_NAME="habit"
INSTALL_DIR="/usr/local/bin"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

info() {
    printf "${GREEN}==>${NC} %s\n" "$1"
}

warn() {
    printf "${YELLOW}Warning:${NC} %s\n" "$1"
}

error() {
    printf "${RED}Error:${NC} %s\n" "$1"
    exit 1
}

# Detect OS and architecture
detect_platform() {
    OS="$(uname -s)"
    ARCH="$(uname -m)"

    case "$OS" in
        Linux*)
            PLATFORM="Linux"
            ;;
        Darwin*)
            PLATFORM="Darwin"
            ;;
        MINGW*|MSYS*|CYGWIN*)
            PLATFORM="Windows"
            BINARY_NAME="habit.exe"
            ;;
        *)
            error "Unsupported operating system: $OS"
            ;;
    esac

    case "$ARCH" in
        x86_64|amd64)
            ARCH="x86_64"
            ;;
        aarch64|arm64)
            ARCH="arm64"
            ;;
        armv7l|armv7)
            ARCH="armv7"
            ;;
        *)
            error "Unsupported architecture: $ARCH"
            ;;
    esac

    info "Detected platform: ${PLATFORM}_${ARCH}"
}

# Get latest release version
get_latest_version() {
    info "Fetching latest version..."

    if command -v curl > /dev/null 2>&1; then
        VERSION=$(curl -sfL "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
    elif command -v wget > /dev/null 2>&1; then
        VERSION=$(wget -qO- "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
    else
        error "Neither curl nor wget found. Please install one of them."
    fi

    if [ -z "$VERSION" ]; then
        error "Failed to fetch latest version"
    fi

    info "Latest version: $VERSION"
}

# Download binary
download_binary() {
    ARCHIVE_NAME="habit-tracker_${VERSION#v}_${PLATFORM}_${ARCH}.tar.gz"
    if [ "$PLATFORM" = "Windows" ]; then
        ARCHIVE_NAME="habit-tracker_${VERSION#v}_${PLATFORM}_${ARCH}.zip"
    fi

    URL="https://github.com/${REPO}/releases/download/${VERSION}/${ARCHIVE_NAME}"

    info "Downloading from: $URL"

    TEMP_DIR=$(mktemp -d)
    cd "$TEMP_DIR"

    if command -v curl > /dev/null 2>&1; then
        curl -sfLO "$URL" || error "Failed to download binary"
    elif command -v wget > /dev/null 2>&1; then
        wget -q "$URL" || error "Failed to download binary"
    fi

    info "Download complete"
}

# Extract and install
install_binary() {
    info "Extracting archive..."

    if [ "$PLATFORM" = "Windows" ]; then
        unzip -q "$ARCHIVE_NAME" || error "Failed to extract archive"
    else
        tar xzf "$ARCHIVE_NAME" || error "Failed to extract archive"
    fi

    info "Installing to ${INSTALL_DIR}..."

    # Check if we need sudo
    if [ -w "$INSTALL_DIR" ]; then
        mv "$BINARY_NAME" "${INSTALL_DIR}/${BINARY_NAME}"
    else
        if command -v sudo > /dev/null 2>&1; then
            sudo mv "$BINARY_NAME" "${INSTALL_DIR}/${BINARY_NAME}"
            sudo chmod +x "${INSTALL_DIR}/${BINARY_NAME}"
        else
            error "No write permission to ${INSTALL_DIR} and sudo not available"
        fi
    fi

    # Clean up
    cd -
    rm -rf "$TEMP_DIR"

    info "Installation complete!"
}

# Install shell completions
install_completions() {
    info "Installing shell completions..."

    COMPLETION_DIR=""

    # Detect shell
    if [ -n "$BASH_VERSION" ]; then
        COMPLETION_DIR="$HOME/.bash_completion.d"
    elif [ -n "$ZSH_VERSION" ]; then
        COMPLETION_DIR="$HOME/.zsh/completions"
    fi

    if [ -n "$COMPLETION_DIR" ]; then
        mkdir -p "$COMPLETION_DIR"
        # Note: Completions would need to be downloaded separately
        # For now, just inform the user
        warn "Shell completions available at: https://github.com/${REPO}/tree/main/completions"
    fi
}

# Verify installation
verify_installation() {
    if command -v "$BINARY_NAME" > /dev/null 2>&1; then
        VERSION_OUTPUT=$("$BINARY_NAME" version)
        info "Verification successful: $VERSION_OUTPUT"
        echo ""
        info "Run '${BINARY_NAME} help' to get started!"
    else
        warn "Binary installed but not found in PATH. You may need to restart your shell."
    fi
}

# Main installation process
main() {
    echo ""
    info "Habit Tracker Installation"
    echo ""

    detect_platform
    get_latest_version
    download_binary
    install_binary
    install_completions
    verify_installation

    echo ""
    info "Installation complete!"
    echo ""
}

main "$@"
