# Go wrapper for Trezor's Shamir Secret Sharing C implementation

## Usage

### Example

See the example under `./cmd/shamir-example/main.go`.

### Build

```
go build -v ./cmd/shamir-example/...
```

### Run

```
./shamir-example
```

Output:

```
[91 188 226 91 254 197 225]
```

## Development

### Prerequisite

Get `trezor-firmware` source code:
```
git submodule update --init --recursive
```

### Build

```
go build -v ./...
```

### Test

```
go test -v ./...
```
