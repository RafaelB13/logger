# Logger

The `logger` package provides a simple and colorful solution for logging in Go applications. It is designed to be easy to use and offers formatted messages for different log levels.

## Installation

Use the following command to install the `logger` package:

```
go get -u github.com/rafaelb13/logger
```

## Usage

```
package main

import (
	"github.com/rafaelb13/logger"
)

func main() {
	// Example of how to use the logger
	log := logger.NewLogger("parameter")
	log.Debug("Debug message")
	log.Info("Information message")
	log.Warn("Warning message")
	log.Error("Error message")
}

```


## Contribution

If you want to contribute, please follow these steps:

1. Fork the repository
2. Create a branch for your feature (`git checkout -b feature/YourFeature`)
3. Commit your changes (`git commit -m 'Adding awesome feature'`)
4. Push to the branch (`git push origin feature/YourFeature`)
5. Open a pull request


## Contribution

If you want to contribute, please follow these steps:

1. Fork the repository
2. Create a branch for your feature (`git checkout -b feature/YourFeature`)
3. Commit your changes (`git commit -m 'Adding awesome feature'`)
4. Push to the branch (`git push origin feature/YourFeature`)
5. Open a pull request


## Version

1.21.3
