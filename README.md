# Docker Container Manager

## Description
Docker Container Manager is a powerful CLI tool built with Go for managing Docker containers efficiently. Designed with a modular architecture, it provides seamless container orchestration, extension points for plugins, and supports containerized deployment via Docker Compose. It leverages Hexagonal Architecture for clean separation of concerns, layered architecture for command and infrastructure separation, and an event-driven pattern for asynchronous operations.

## Tech Stack
- Go
- Docker
- Docker Compose
- Event-Driven Architecture Patterns
- Plugin Architecture for Extensibility

## Folder Structure
docker-container-manager/
│
├── cmd/                   # CLI commands and entrypoints
│   ├── root.go           # Root command setup
│   ├── manage.go         # Container management commands
│   └── plugin/           # Plugin command modules
│
├── internal/              # Core application logic (Hexagonal Architecture)
│   ├── container/        # Container domain logic
│   ├── event/            # Event handling for async operations
│   └── plugin/           # Plugin system interface and implementations
│
├── pkg/                   # Infrastructure layer: Docker client wrappers, config loaders
│
├── plugins/               # External plugin modules
│
├── docker-compose.yml     # Docker Compose setup for containerized deployment
│
├── README.md              # Project documentation
│
├── go.mod                 # Go module file
│
└── go.sum                 # Go checksum files

## How to Run Locally
1. Clone the repository:
git clone https://github.com/yourusername/docker-container-manager.git
cd docker-container-manager
2. Build the CLI tool:
go build -o dcm ./cmd
3. (Optional) Run using Docker Compose:
docker-compose up -d
4. Execute commands:
./dcm --help
./dcm manage list

## Environment Variables
Configure the following environment variables as needed:
- `DOCKER_HOST` : Docker daemon address (default: `unix:///var/run/docker.sock`)
- `DPM_LOG_LEVEL` : Logging level (`debug`, `info`, `warn`, `error`)
- `PLUGIN_PATH` : Directory for external plugins

Set environment variables:
export DOCKER_HOST=unix:///var/run/docker.sock
export DPM_LOG_LEVEL=info
export PLUGIN_PATH=./plugins

## Contributing
Contributions are welcome! Please follow these steps:
1. Fork the repo
2. Create a new branch `git checkout -b feature/your-feature`
3. Make your changes
4. Commit with descriptive messages
5. Push to your branch
6. Create a pull request

Please adhere to the code style and include tests for new features.

## License
This project is licensed under the MIT License. See the `LICENSE` file for details.