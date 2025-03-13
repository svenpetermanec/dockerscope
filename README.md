# Dockerscope

Dockerscope is an interactive, terminal-based tool for managing and inspecting Docker containers. Written in Go, it leverages a modular panel system to provide a user-friendly interface for executing Docker commands, listing resources, and displaying inspection details—all from your terminal.

## Features

- Interactive Terminal UI:
    Utilizes multiple panels for command selection, resource display, and detailed inspection output.

- Docker Command Execution:
    Executes Docker commands seamlessly through a custom executor, making it easy to manage containers and resources.

- Modular Architecture:
    Organized into dedicated packages (e.g., docker, panel, terminal) to promote clarity and extensibility.

- Robust Error Handling:
    Implements error recovery with logging to a file for easier troubleshooting.

## Installation
### Prerequisites

- [Go](https://go.dev/dl/) (version 1.16 or higher recommended)
- [Docker](https://www.docker.com/) installed and running on your machine

### Building from Source

Clone the repository:
```bash
  git clone https://github.com/svenpetermanec/dockerscope.git
  cd dockerscope
```

Build the application:

```bash
  go build -o dockerscope
```

## Usage

Run the executable:

```bash
  ./dockerscope
```

When launched, you'll see three panels:

    Command Panel:
    Lists Docker commands. Use ↑/↓ to navigate. Press Tab to move to the next panel (Resource Panel) and Shift+Tab to return to the previous panel.

    Resource Panel:
    Displays resources for the selected command. Use ↑/↓ to navigate. Press Tab to go to the next panel (Inspect Panel) and Shift+Tab to return to the Command Panel.

    Inspect Panel:
    Shows detailed information for the selected resource. Scroll with ↑/↓ and press Shift+Tab to return to the Resource Panel.

Any errors are logged to a file named `log`.
