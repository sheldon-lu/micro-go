# Kubernetes Service

This is the Kubernetes service

Generated with

```
micro new micro-go/kubernetes-web --namespace=lu.micro --alias=kubernetes --type=web
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: lu.micro.web.kubernetes
- Type: web
- Alias: kubernetes

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./kubernetes-web
```

Build a docker image
```
make docker
```