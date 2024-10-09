# netsquirt

Effortlessly serve a single file on the local network via http.

Written in Go.

    Usage:

        -file, --file [PATH]        specify path to the file to be served
        -port, --port [NUMBER]      specify a port number to run the server on
        -version, --version         display verison
        -h, -help, --help           display Usage

If you run the program without the `-file` parameter, it will serve a simple html page.

## Installing

Install via go:
 
    go install github.com/maciakl/netsquirt@latest

On Windows, this tool is distributed via `scoop` (see [scoop.sh](https://scoop.sh)).

First, you need to add my bucket:

    scoop bucket add maciak https://github.com/maciakl/bucket
    scoop update

 Next simply run:
 
    scoop install netsquirt

If you don't want to use `scoop` you can simply download the executable from the release page and extract it somewhere in your path.
