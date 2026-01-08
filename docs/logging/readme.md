Logging Features

The Library provides structured logging capabilities built on top of [Logrus](https://github.com/sirupsen/logrus).

## Features

- **Structured Logging** - JSON and text formatters supported
- **Context-Aware Logging** - All log methods accept context.Context for request tracing
- **Automatic Source Tracking** - Automatically includes file, function name, and line number in log entries
- **Container ID Tracking** - Automatically tracks container/pod IDs from environment or generates unique IDs
- **Goroutine ID Support** - Optional goroutine ID tracking via context
- **Multiple Output Hooks** - Support for file, Graylog, and Logstash hooks
- **Custom Hooks** - Ability to register custom Logrus hooks at initialization

## Logger Interface

The Logger interface provides the following methods:

```go
type Logger interface {
    Debug(ctx context.Context, args ...interface{})
    Info(ctx context.Context, args ...interface{})
    Warn(ctx context.Context, args ...interface{})
    Error(ctx context.Context, args ...interface{})
    
    InfoF(ctx context.Context, format string, args ...interface{})
    WarnF(ctx context.Context, format string, args ...interface{})
    DebugF(ctx context.Context, format string, args ...interface{})
    ErrorF(ctx context.Context, format string, args ...interface{})
}
```

## Usage

### Basic Usage

The logger is automatically initialized when you call `bootstrap.Init()`. You can pass custom hooks during initialization using the `WithLogHook` option. You can access the logger from the returned `AppServer`:

```go
import (
    "context"
    "github.com/avinash92c/bootstrap-go/bootstrap"
    "github.com/sirupsen/logrus"
)

// Initialize without custom hooks
appsvr, router := bootstrap.Init()
logger := appsvr.Logger

// Or initialize with custom hooks
customHook := &MyCustomHook{} // Your custom logrus.Hook implementation
appsvr, router := bootstrap.Init(
    bootstrap.WithLogHook(customHook),
)
logger := appsvr.Logger

// Simple logging
logger.Info(context.Background(), "Application started")
logger.Error(context.Background(), "An error occurred", err)

// Formatted logging
logger.InfoF(context.Background(), "Server starting on port %d", port)
logger.ErrorF(context.Background(), "Failed to connect: %v", err)
```

You can pass multiple hooks by calling `WithLogHook` multiple times:

```go
appsvr, router := bootstrap.Init(
    bootstrap.WithLogHook(hook1),
    bootstrap.WithLogHook(hook2),
    bootstrap.WithLogHook(hook3),
)
```

### Getting Logger Instance

You can also retrieve the logger instance directly from the foundation package:

```go
import "github.com/avinash92c/bootstrap-go/foundation"

logger := foundation.GetLogger()
logger.Info(context.Background(), "Log message")
```

## Configuration

Logging is configured through your application's configuration file. See [Configuration Documentation](../config/config.md) for details.

### Logging Configuration Properties

| CONFIG NAME               | Description                                                               | Accepted Values                            |
| ------------------------- | ------------------------------------------------------------------------- | ------------------------------------------ |
| logging.level             | Logging Level                                                             | info/error/debug/warn/fatal                |
| logging.formatter.name    | LogData Formatting                                                        | json/text                                  |
| logging.formatter.options | [Formatting Options](https://github.com/heralight/logrus_mate#formatters) |                                            |
| logging.hooks             | [Hooks for Emitting Data](https://github.com/heralight/logrus_mate#hooks) | Currently Supported: graylog/logstash/file |

### Example Configuration

```yaml
logging:
  level: debug
  formatter:
    name: json
  hooks:
    - name: file
      options:
        file_path: "/var/log/app.log"
        maxsize: 5000
        maxbackups: 5
        maxage: 30
    - name: graylog
      options:
        address: "udp://graylog.example.com:12201"
    - name: logstash
      options:
        app_name: "myapp"
        protocol: "tcp"
        address: "logstash.example.com:5000"
```

## Supported Hooks

### File Hook

Writes logs to a file with automatic rotation using Lumberjack.

**Configuration Options:**
- `file_path` - Path to the log file
- `maxsize` - Maximum size in megabytes before rotation
- `maxbackups` - Maximum number of old log files to retain
- `maxage` - Maximum number of days to retain old log files

### Graylog Hook

Sends logs to a Graylog server via UDP.

**Configuration Options:**
- `address` - Graylog server address (e.g., "udp://graylog.example.com:12201")

### Logstash Hook

Sends logs to a Logstash server via TCP (asynchronous).

**Configuration Options:**
- `app_name` - Application name identifier
- `protocol` - Network protocol (typically "tcp")
- `address` - Logstash server address (e.g., "logstash.example.com:5000")
- `always_sent_fields` - Additional fields to include in every log entry (JSON object)

## Custom Hooks

You can register custom Logrus hooks during initialization using the `WithLogHook` option:

```go
import (
    "github.com/avinash92c/bootstrap-go/bootstrap"
    "github.com/sirupsen/logrus"
)

// Create a custom hook
customHook := &MyCustomHook{}

// Initialize with custom hook
appsvr, router := bootstrap.Init(
    bootstrap.WithLogHook(customHook),
)
```

### Example: Logstash Hook

Here's a concrete example using a Logstash hook with custom formatter:

```go
import (
    "os"
    "github.com/avinash92c/bootstrap-go/bootstrap"
    "github.com/sirupsen/logrus"
    logstash "github.com/avinash92c/logrus-logstash-async"
)

formatter := logstash.DefaultFormatter(logrus.Fields{
    "service": "orders",
    "env":     "prod",
})

// Create hook (stdout just for testing)
hook := logstash.New(os.Stdout, formatter)

appsvr, router := bootstrap.Init(
    bootstrap.WithLogHook(hook),
)
```

Note: In production, you would typically use a network connection (TCP/UDP) instead of `os.Stdout` for the Logstash hook.

## Automatic Fields

Every log entry automatically includes:

- **container** - Container/pod ID (from `CONTAINER_ID` environment variable or auto-generated)
- **src** - Source location in format `filename:functionname:linenumber`
- **routineid** - Goroutine ID (if set in context using the `goroutineIDKey`)

## Log Levels

Supported log levels (in order of severity):

1. **Debug** - Detailed information for debugging
2. **Info** - General informational messages
3. **Warn** - Warning messages
4. **Error** - Error messages
5. **Fatal** - Critical errors (not exposed via Logger interface, handled internally)

## Context Usage

All logging methods require a `context.Context` parameter. This enables:

- Request tracing across service boundaries
- Goroutine ID tracking (if set in context)
- Future support for distributed tracing

Example with request context:

```go
func handleRequest(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    logger := appsvr.Logger
    
    logger.Info(ctx, "Processing request")
    // ... handle request
    logger.InfoF(ctx, "Request completed in %v", duration)
}
```
