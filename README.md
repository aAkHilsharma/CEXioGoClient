# CryptoRateFetcher

CryptoRateFetcher is a simple Go application that retrieves cryptocurrency rates from the CEX.io API. It demonstrates the use of HTTP requests, JSON parsing, and concurrency in Go.

## Project Structure

- **main.go**: Initiates concurrent requests to fetch rates for a list of cryptocurrencies.
- **api/cex.go**: Contains logic for making API calls to CEX.io and parsing the responses.
- **api/responses.go**: Defines the `Response` struct for modeling API responses.
- **datatypes/data.go**: Defines the `Rate` struct for storing currency rates.
- **go.mod**: Manages project dependencies.

## Usage

1. Clone the repository.
2. Ensure you have Go installed on your system.
3. Run the application using `go run main.go`.

## Dependencies

This project uses Go's standard library for HTTP requests and JSON parsing. Ensure your Go environment is set up correctly to manage dependencies via `go.mod`.

## License

This project is open-source and available under the MIT License. 