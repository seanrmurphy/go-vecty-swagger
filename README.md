# go-vecty-swagger Frontend

This repository contains an experiment in using a Swagger client within a Go
frontend application to communicate with a backend hosted on AWS (also written
in Go).

Please check this repo for the backend:

https://github.com/seanrmurphy/go-fullstack

# Downloading the repo

The repo can be downloaded as follows:

`git clone https://github.com/seanrmurphy/go-vecty-swagger $GOPATH/github.com/seanrmurphy/go-vecty-swagger`

# Generating the client

The client code is provided in this repo, so it is not necessary to generate it per
se. It was generated with the stratoscale go-swagger generator and can be generated
using the following in the root directory of the repository - the directory containing
the `swagger.yaml` file:

`docker run --rm -e GOPATH=$GOPATH:/go -v $PWD:$PWD -w $PWD -u (id -u):(id -u) stratoscale/swagger:v1.0.27 generate client`

# Building the frontend

See the documentation inside the `frontend` directory for instructions on building
the frontend.





