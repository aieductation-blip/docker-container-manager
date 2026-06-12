# ARCHITECTURE.md for Docker Container Manager

## Overview
`docker-container-manager` is a CLI tool developed in Go for managing Docker containers. It leverages Docker SDK for Go to interact with Docker Engine and provides an extensible command system via plugins. The system is designed with a hexagonal architecture to decouple core logic from CLI interfaces, and employs containerization for deployment.

---

## 1. Architecture Patterns and Rationale

- **Hexagonal Architecture for CLI Command Separation**  
  Isolates core container management logic from CLI command parsing and user interaction, allowing independent testing and future UI expansion.

- **Layered Architecture**  
  Separates command layer (CLI parsing), application layer (business logic), and infrastructure layer (Docker SDK integration).

- **Plugin Architecture**  
  Supports extensible CLI commands through dynamically loaded plugins, enabling custom container operations without recompiling the core.

- **Containerized Deployment with Docker Compose**  
  Deploys the CLI tool as a Docker container via Docker Compose, ensuring consistent environment setup.

- **Event-Driven Pattern**  
  Uses a message bus (Go channels) internally for handling asynchronous Docker events, such as container start or stop notifications.

---

## 2. Tech Stack and Tool Selection

| Layer | Components | Rationale |
|--------|--------------|------------|
| **Core Logic** | Go (Version 1.20+) | Chosen for its concurrency model and static binaries, enabling efficient CLI tool with minimal dependencies. |
| **CLI Framework** | `cobra` | Provides structured commands, nested subcommands, and plugin support, simplifying command management. |
| **Communication with Docker** | Docker SDK for Go (`github.com/docker/docker/client`) | Direct API access to Docker Engine, providing robust and version-aware Docker interactions over REST API. |
| **Plugin System** | Go plugins (`plugin` package) | Facilitates dynamic extension loading at runtime, enabling new commands without recompiling core. |
| **Container Deployment** | Docker Compose (`docker-compose.yml`) | Simplifies deployment in development and CI environments, bundling the CLI tool with its runtime environment. |
| **Async Operations** | Go channels and goroutines | Implements event-driven notifications for asynchronous container events, like container state changes. |

---

## 3. System Components and Data Flow

### Component Diagram

                +-----------------------+
                |   CLI Layer (cobra)   |
                +-----------------------+
                          |
                 Parse Commands / Subcommands
                          |
+------------------------------------------------------------+
|                        Application Layer                   |
|  +------------+    +----------------+   +--------------+  |
|  | Commands   | -> | Command Handlers| ->| Plugin Loader| -> [Load plugins dynamically]
|  +------------+    +----------------+   +--------------+  |
|                        |                                    |
|               Business Logic Layer                        |
|        +---------------------------+                      |
|        | Container Management Core |                      |
|        +---------------------------+                      |
|                /|\                                         |
+------------------------------------------------------------+
                          |
        Interact via Docker SDK (Docker Engine API HTTP)
                          |
    +------------------------------------------------------------+
    | Infrastructure Layer (Docker SDK Client)                    |
    +------------------------------------------------------------+
                          |
                   Docker Daemon (API @ Unix socket or TCP)

### Data Flow Example: Starting a Container

1. User executes: `docker-manager start container_id`
2. CLI (`cobra`) parses command, calls corresponding Handler.
3. Handler invokes the Core `StartContainer` function (application layer).
4. `StartContainer` uses Docker SDK (`client.ContainerStart(ctx, containerID, options)`) to send HTTP request to Docker API.
5. Docker API responds with success or error.
6. On success, an internal event (via Go channels) is emitted.
7. Event listener logs or triggers notifications.

### Real Endpoints and Data

- **Start container:** POST `/containers/{id}/start` via Docker SDK
- **Stop container:** POST `/containers/{id}/stop`
- **List containers:** GET `/containers/json`
  
**Sample Table Name (if DB is used for augmented metadata):** `containers`, `container_events`

---

## 4. Deployment Architecture

- The CLI is packaged as a Docker image (`docker/docker-container-manager:latest`).
- `docker-compose.yml` deploys an environment with:
  - The CLI container connected to Docker Engine.
  - Optional plugins mounted as volumes.
  
Example `docker-compose.yml` snippet:

version: '3.8'
services:
  container-manager:
    image: docker/docker-container-manager:latest
    volumes:
      - ./plugins:/plugins
      - /var/run/docker.sock:/var/run/docker.sock
    environment:
      - DOCKER_HOST=unix:///var/run/docker.sock
    restart: unless-stopped

---

## 5. Technology Choice Justifications

- **Go** over Python or Node.js due to its compile-time performance, simple static binaries, and native concurrency support for event-driven async handling.
- **Docker SDK for Go** over Docker CLI commands via subprocess, because direct SDK calls are more efficient, reliable, and less error-prone.
- **cobra** over other CLI frameworks like urfave/cli because it has built-in support for nested commands and plugin architecture, necessary for extensibility.
- **Docker Compose** as deployment method instead of manual Docker run commands to simplify multi-container setup, secrets management, and volume mounting.
- **Go's plugin package** over other plugin systems because it offers native, compiled plugin support, ideal for stable plugin architectures within the same language.

---

## 6. Conclusion

This architecture leverages the strengths of Go and Docker to produce a modular, performant, and extensible CLI tool for container management. The layered, hexagonal design ensures the core logic remains independent of clilayer and infrastructure changes, enabling robust testing and future expansion.