# Govee Go Client
The `govee-go` client provides an easy way to interact with your [Govee](https://us.govee.com/) devices using Golang.

## Prerequisites

### Govee API Key
You will need a Govee API key to use this module. Instructions on how to get your Govee API key can be found [here](https://developer.govee.com/reference/apply-you-govee-api-key).

## Install the package

### In a Go module

Install the latest version of the `govee-go` client:

   ```sh
   go get github.com/bugfix4u/govee-go/govee
   ```

## Usage

In a Go file, import the `govee-go` package to use it in your code, for example:

```go
import (
	"fmt"
	"log"

	"github.com/bugfix4u/govee-go/govee"
)
```

## Example Code

Example code can be found in the [`example` folder](./govee/example/README.md)

## Govee Device Support

A list of supported Govee devices can be found [here](https://developer.govee.com/docs/support-product-model)

## Contribution

To contribute to this project, fork the repository on GitHub and send a pull request to the `main` branch.

## License

The Govee Go Client is released under the [MIT License](https://opensource.org/licenses/MIT)