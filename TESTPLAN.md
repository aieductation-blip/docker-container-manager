# TESTPLAN.md

## 1. Test Strategy
This project utilizes a layered approach within the Hexagonal Architecture to ensure clear separation of concerns and facilitate focused testing. Unit tests will cover individual Go functions within core logic, particularly Docker API wrappers and CLI command handlers. Integration tests will validate interactions with Docker Engine API through mocked HTTP clients for commands like create, start, stop, and remove containers. E2E tests will simulate user workflows via CLI commands, verifying command parsing, execution, and expected Docker operations, using a dedicated test Docker environment spun up via Docker Compose. The event-driven pattern will be tested by asserting message passing correctness for asynchronous container management tasks.

## 2. Test Levels

**Unit Tests**
- Test `CreateContainer(id, image)` function ensuring correct Docker API request payload.
- Test `StartContainer(id)` with mocked Docker responses.
- Test `parseCLIArgs()` to confirm correct command parsing and flag handling.

**Integration Tests**
- Send HTTP POST to `/containers/create` with body `{"Image":"nginx"}` and verify container creation.
- Test `GET /containers/{id}/json` API call returns valid container state.
- Run commands like `docker-container-manager start --id <container_id>` and verify container starts in Docker.

**E2E Tests**
- Run `docker-container-manager create --image nginx` followed by `start`, `stop`, and `remove` commands, verifying each step in the Docker engine.
- Validate that logs are emitted appropriately during container lifecycle operations.
- Test concurrent command executions to verify thread safety and message passing.

## 3. Test Cases Table

| ID  | Test Case                                                     | Type            | Priority | Expected Result                                              |
|-----|---------------------------------------------------------------|-----------------|----------|--------------------------------------------------------------|
| TC01| Create a container with valid image `nginx` via CLI `create` | Unit            | High     | Docker API receives `POST /containers/create` with correct spec |
| TC02| List containers using CLI `list` command                       | E2E             | High     | Properly lists all running/stopped containers in Docker    |
| TC03| Start container with ID `abc123` via CLI                        | Integration     | High     | Docker API responds with container `abc123` started successfully|
| TC04| Stop container with ID `abc123` via CLI                         | Integration     | High     | Docker API responds with container `abc123` stopped successfully|
| TC05| Remove container with ID `abc123` via CLI                        | Integration     | Medium   | Docker API responses with successful removal status        |
| TC06| CLI command with invalid flag produces error                   | Unit            | Medium   | Proper error message and exit code                          |
| TC07| Concurrent container creation requests via CLI                 | E2E             | Low      | All containers are created and managed correctly without race conditions |
| TC08| API call to `/containers/xyz/json` returns container info     | Integration     | High     | JSON payload with container `xyz` status and config info     |
| TC09| Event message is correctly dispatched after container stop     | E2E             | Medium   | Message is received on message bus indicating container stopped |
| TC10| CLI command for `remove` on non-existing container             | Unit            | High     | Graceful error with message: "Container not found"        |

## 4. Edge Cases
- Creating a container with an invalid or empty image name.
- Attempting to start a container that is already running.
- Removing a container that does not exist.
- Executing CLI commands without necessary permissions or Docker daemon access.
- Handling Docker API disconnections or timeout errors gracefully.

## 5. Test Data Requirements
- Fixtures representing Docker container JSON responses (e.g., container status, network info).
- Sample Docker images (`nginx`, `redis`, `alpine`) preloaded in test Docker environment.
- Mocked message bus messages for event verification.
- Seed containers with known IDs such as `abc123`, `xyz789` for testing start/stop/remove actions.

## 6. Tools & Setup
- **Go testing framework** (`go test ./...`) with mock HTTP client (`net/http/httptest`) for API mocking.
- **Docker CLI** installed locally with access to a test Docker environment.
- **Docker Compose** file to deploy isolated Docker daemon for integration and E2E tests.
- **Testcontainers-Go** for spinning up lightweight Docker containers during testing:
  ```
  go get github.com/testcontainers/testcontainers-go
  ```
- **Message Bus Mock** setup for event-driven pattern testing, e.g., in-memory channel or mock message broker.
- **Run setup commands:**
  ```
  docker-compose -f docker-compose.test.yml up -d
  go test -v ./...
  ```