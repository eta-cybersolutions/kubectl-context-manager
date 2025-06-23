# kubectl-context-manager

**Interactive Kubernetes context manager** CLI tool.

This tool helps interactively select a Kubernetes context from your `~/.kube/config` using arrow keys — with zero memorization.

---

## Prerequisites

To build and run this tool, you need:

- **Go (>= 1.18)**  
  Install from: https://go.dev/dl/

- **kubectl** installed and configured  
  - Must be available in your system's `$PATH`
  - Should have access to a valid `~/.kube/config` file

- **Linux/macOS/WSL** (Windows is also possible, but not tested)

---

## Features

- Reads and parses your `~/.kube/config` file
- Shows a friendly CLI selection menu with current context marked
- Automatically switches to the selected context using `kubectl`
- Fast, native Go binary — no dependencies

---

## Installation

### Clone & Build

```bash
git clone https://github.com/eta-cybersolutions/kubectl-context-manager
cd kubectl-context-manager
make build
make install
```