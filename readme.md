# URL Parser

A simple command-line tool written in Go to extract specific components from a given URL. Especially useful in Bash scripts or automation workflows where URL processing is required.

## Features

- Extract various components of a URL, including:
  - Host
  - Domain
  - Subdomain
  - Top-level domain (TLD)
  - Port
  - Path
  - Fragment
  - Scheme (protocol)
  - Registered domain from a subdomain (`domain.tld`)
  - Query parameter values

## Installation

### [Download the latest binary](https://github.com/singh-inder/url-parser/releases/latest)

### wget

Use wget to download pre-compiled binaries:

For instance, VERSION=v1.0.0 and BINARY=url-parser-linux-amd64

#### Plain binary

```bash
wget https://github.com/singh-inder/url-parser/releases/download/${VERSION}/${BINARY} -O /usr/bin/url-parser &&\
    chmod +x /usr/bin/url-parser
```

#### Latest version

```bash
wget https://github.com/singh-inder/url-parser/releases/latest/download/url-parser-linux-amd64 -O /usr/bin/url-parser &&\
    chmod +x /usr/bin/url-parser
```

## Usage

```bash
url-parser --url <URL> --get <component>
```

### Required Flags

- `--url`: The URL to parse. Example: `https://example.com`.
- `--get`: The component to extract. Options:
  - `host`: Full host (e.g., `subdomain.example.com`)
  - `domain`: Domain name (e.g., `example`)
  - `subdomain`: Subdomain (e.g., `sub`)
  - `tld`: Top-level domain (e.g., `com`)
  - `port`: Port (if specified, e.g., `8080`)
  - `path`: URL path (e.g., `/path/to/resource`)
  - `fragment`: Fragment (e.g., `section1`)
  - `scheme`: URL scheme (e.g., `https`)
  - `registeredDomain`: Registered domain from a subdomain (e.g., `example.com`)
  - `query.<paramName>`: Value of a specific query parameter (e.g., `query.user` extracts the value of the `user` parameter).

### Examples

#### Extract the domain from a URL:

```bash
url-parser --url "https://sub.example.com/path?user=123" --get domain
# Output: example
```

#### Get the value of a query parameter:

```bash
url-parser --url "https://example.com/path?user=123" --get query.user
# Output: 123
```

#### Extract the TLD:

```bash
url-parser --url "https://example.com" --get tld
# Output: com
```

#### Extract the scheme:

```bash
url-parser --url "https://example.com" --get scheme
# Output: https
```

#### Extract the full host:

```bash
url-parser --url "https://sub.example.com" --get host
# Output: sub.example.com
```

#### Extract the registered domain from a subdomain:

```bash
url-parser --url "https://sub.example.com" --get registeredDomain
# Output: example.com
```

## Contributing

Feel free to submit issues or create pull requests to improve the project. Contributions are always welcome!

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgments

This cli uses [go-tld](https://github.com/jpillora/go-tld) package for extracting URL components.
