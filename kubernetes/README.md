# Kubernetes Service

This is the Kubernetes service

Generated with

```
micro new micros/kubernetes --namespace=go.micro --alias=kubernetes --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.kubernetes
- Type: srv
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
./kubernetes-srv
```

Build a docker image
```
make docker
```