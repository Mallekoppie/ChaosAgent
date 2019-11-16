# Introduction

Used to distribute tests to agents. The agents then make http requests to the remote target.

The main purpose is to determine the maximum theoretical through of a service. The tests are repeated and not randomized at this point in time.

## Commands

### Build proto file

In root ChaosGenerator directory

> protoc --go_out=plugins=grpc:. ./Chaos/chaos.proto