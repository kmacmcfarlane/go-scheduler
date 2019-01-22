# go-scheduler
This is a distributed job scheduler for linux systems built in Go. It can run arbitrary docker images on a cluster of computers.

# Install

## Get the package
`go get github.com/kmacmcfarlane/go-scheduler`

## Build the project
`make`

## Generate Certificates with EasyRSA
The cluster uses x509 certificates to authenticate the client, master, and node applications.

# CLI
The go-scheduler CLI client allows you to control work run on the cluster.

## Install Client Certificate
The client certificate is used to authenticate with your cluster. The certificate must have file permission 600.

`mkdir ~/.go-scheduler`
`cp ~/Downloads/client.crt ~/.go-scheduler/client.crt`
`chmod 600 ~/.go-scheduler/client.crt`

## Start a Job

`go-scheduler-cli start --image redis --name redisCache --master 192.168.1.80`

## Stop a Job

`go-scheduler-cli stop --name redisCache --master 192.168.1.80`

## Query Job Status

`go-scheduler-cli query --name redisCache --master 192.168.1.80`

## Stream Log Output

`go-scheduler-cli log --name redisCache --master 192.168.1.80`

## Status Codes
go-scheduler-cli returns the following status codes:

0: success
1: error parsing command-line arguments
2: error parsing command-line arguments
3: application error