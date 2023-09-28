## Algorand Tiny Watcher

Tiny package to watch for wallet state changes on the Algorand blockchain.

### Endpoints

##### [GET] /watch

Returns the state for each address wallet being watched

##### [POST] /watch/:address

Adds the address passed an a URL parameter into the watchlist

#### Configurations

.env file

```
PORT                     : port the server will run in                            (defaults to :8080)
ALGORAND_PROVIDER_URL    : the API for the algorand go SDK                        (defaults to algonode.io)
UPDATE_TIMEOUT_IN_SECONDS: the number of seconds the watcher will run its updates (defaults to 10 seconds)
```

### Running

```
docker compose  up -d --build
```

or

```
go get ./...
go run .
```

### Tests

```
go test ./...
```

### Future Tasks

- Better error handling. Improve and customize to the scopes needs
- Add more test coverage
- Notify user through email/sms/whatevs for account state changes
- Transform this into Tiny Watcher and become blockchain agnostic
