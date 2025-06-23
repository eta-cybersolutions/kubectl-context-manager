# kubectx-manager v1.0.0

 First public release of **kubectx-manager** — a simple, interactive CLI tool to switch between Kubernetes contexts from `~/.kube/config`.

## Features

- Interactive CLI menu with arrow key navigation
- Highlights the current context in green
- Fast native binary (written in Go)
- Works cross-platform (Linux, macOS, Windows)
- No runtime dependencies

## Assets

Precompiled binaries are available below:
- `kcswitch-linux` (Linux x86_64)
- `kcswitch-macos` (macOS x86_64)
- `kcswitch.exe` (Windows x86_64)

To use:
```bash
chmod +x kubectx-manager-linux
./kubectx-manager-linux
```
