📦 Standard Library (Comes with Go)
1. fmt
Purpose: Formatting and printing to console (like printf, println).

Example: fmt.Println("Hello, Go!")

2. net/http
Purpose: Building HTTP clients and servers.

Example: Creating a web server or making HTTP requests.

3. encoding/json
Purpose: Parsing and generating JSON.

Example: Converting a Go struct to JSON or parsing JSON into a Go struct.

4. os
Purpose: Interacting with the operating system (e.g. files, environment variables).

Example: os.Open("file.txt")

5. io / io/ioutil
Purpose: Input/output operations, reading and writing data streams.

Example: Reading a file or HTTP response body.

6. time
Purpose: Working with dates, durations, and timers.

Example: time.Now(), time.Sleep(2 * time.Second)

7. log
Purpose: Simple logging functionality.

Example: log.Println("Something happened")

🌐 Web & Networking
8. gorilla/mux
Purpose: A powerful URL router and dispatcher for building web APIs.

Example: Easier route matching with parameters like /users/{id}.

9. gin-gonic/gin
Purpose: A high-performance HTTP web framework.

Why: Very popular for REST APIs due to speed and simplicity.

🧪 Testing & Validation
10. stretchr/testify
Purpose: Assertions and mocking for unit testing.

Why: Makes writing tests easier and more readable.

11. go-playground/validator
Purpose: Struct and field validation.

Example: Ensuring fields like email or password match certain formats.

🔄 Utilities & Helpers
12. spf13/viper
Purpose: Managing configuration from files (YAML, JSON, etc.) and environment variables.

Why: Flexible and widely used in real-world projects.

13. spf13/cobra
Purpose: Creating command-line (CLI) applications.

Example: kubectl is built with Cobra.

14. sirupsen/logrus
Purpose: Structured logging (better than standard log).

Why: JSON logs, hooks, log levels—great for production systems.

🧵 Concurrency
15. golang.org/x/sync/errgroup
Purpose: Running multiple goroutines and handling their errors together.

Why: Simplifies concurrent tasks with shared error handling.

🧰 Miscellaneous
16. google/uuid
Purpose: Generating UUIDs (universally unique identifiers).

Use case: Unique IDs for database records, sessions, etc.

17. pkg/errors (by Go team)
Purpose: Better error wrapping and stack traces.

If you're building an API, gin, viper, validator, and testify are excellent tools to start with
