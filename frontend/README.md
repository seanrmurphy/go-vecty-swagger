# The go-vecty-swagger Frontend

The frontend provided here is the TodoMVC example provided by the vecty project
with some modest modifications to communicate with the backend using Swagger.

The frontend can be run directly on your local machine or it can be run from an
S3 bucket with CORS enabled. Basic instructions for each of these options are
provided below.

## Building the application

The application has been developed and tested with go 1.14.1. An earlier version
was tested with tinygo but as of this writing tinygo is still based on a go 1.13
variant and hence it does not support the newer 'syscall/js' library functionality
of go 1.14. For this reason go 1.14 was chosen, even though it results in wasm
files which are significantly larger than those of tinygo.

The application can be built using the go compiler as follows

`env GOOS=js GOARCH=wasm go build -o build/out.wasm ./src`

(The `build_fe.sh` script contains this command).

## Setting the backend

The backend is specified as a variable (`RestEndpoint`) in the main entry point
for the application - in the `todo-client.go` file. Modify this based on the
configuration of the backend.

## wasm_exec.js

`wasm_exec.js` is a Javascript file which provides the glue between the wasm executable
and the browser based Javascript and is specific to the version of the go compiler
that you are using. You will need to copy the correct version of this file from the
go toolchain you are using to the local directory; typically it is stored in `/usr/local/go/bin/misc/wasm`.

`cp /usr/local/go/bin/misc/wasm/wasm_exec.js .`

## Running the application locally

The application can be run using the simple python server provided as follows -
note that it is dependent on python3:

`./server.py`

This will run the server on all host interfaces on port 4443.

You can open a browser and point it a `https://localhost:4443` and see the FE in
operation. Note a self-signed certificate is provided in this repo; the browser
will ask for authorization to trust this certificate before proceeding.

If the backend is not operational, you should still see the application, but there
will be errors logged in the console.

## Running the application on S3

The application can be run directly off S3 as follows:
- Create a bucket for the application on S3
- Modify the `s3_upload.sh` script, inserting the bucket name as appropriate
- Create a `policy.json` file from the file `policy.json.template` in the directoy `json_config`
- Upload to S3 using the script:

`./s3_upload.sh`

- Point your browser at the endpoint specified to check that it is running on S3.

## Blocking access to the bucket

Modify the script `./s3_block_access.sh` to include the bucket name and run this
script to disable access to the bucket in S3.

