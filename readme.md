# Simple API

## Installation

After cloning the repository, run `go mod tidy` to install the dependencies.

Then run `go run main.go` to start the server. The server url will be `http://localhost:4001`.

## Endpoints

### POST `/convert`

#### Request Body

```json
{
    "command": "magick test.png test.pdf"
}
```
