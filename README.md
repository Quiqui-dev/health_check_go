# health_check_go

An app to check the health of a domain on a specific port, written in Go.

# Build

To build the app run the following from the root of the repo:

```
go build src/*.go
```

# Usage

To use the application, once built, run the following command:

```
./check -d "<domain>" -p "<port>" -t <timeout>
```

# Expected Results

## Success

Upon a successful check the application will respond into the terminal with the following:

```
[UP] <domain> is reachable,
From: <local IP address + port>
To: <remote IP of domain + port>
```

# Failure

Upon a failed check the application will respond into the terminal with the following:

```
[DOWN] <domain> is unreachable,
From: <local IP address + port>
To: <remote IP of domain + port>
```
