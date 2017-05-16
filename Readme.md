<p align="center"><a href="http://isomorphicgo.org" target="_blank"><img src="https://github.com/isomorphicgo/isogoapp/blob/master/static/images/isomorphic_go_logo.png"></a></p>

# Isomorphic Go App (isogoapp)

A basic, barebones web app, intented to be used as a starting point for developing an [Isomorphic Go](http://isomorphicgo.org) application.

## Installation

The instructions assume that you are using a Unix-like operating system (e.g. BSD, Linux, Mac OS, etc).

If you are using Windows, you may need to adapt the installation instructions, most notably for setting up the `ISOGO_APP_ROOT` environment variable.

### Prerequisites

It is assumed that the latest version of Go is already installed on your system and that you have properly configured your Go workspace.

### Get the isogoapp project
`go get -u github.com/isomorphicgo/isogoapp`


### Define the `ISOGO_APP_ROOT` environment variable.
`export ISOGO_APP_ROOT=${GOPATH}/src/github.com/isomorphicgo/isogoapp`


### Build the Client-Side Go App
`cd $ISOGO_APP_ROOT/client`

`gopherjs build`

### Run the Web Server Instance

`cd $ISOGO_APP_ROOT`

`go run isogoapp.go`


The web server runs locally on port 8080: [localhost:8080](http://localhost:8080).

If all goes well, you should see the Isomorphic Go logo, along with a message rendered on the web page.


## Where to now?

### Kick

Once you've installed, and confirmed, that the basic Isomorphic Go application is working, the next logical step is to install [Kick](https://github.com/isomorphicgo/kick).

Kick automatically recompiles Go code, and it can take both the `go` and `gopherjs` commands into consideration. 

Kick performs an *instant kickstart* of the web server instance, upon the modification of a Go source file. [Check out Kick](https://github.com/isomorphicgo/kick)

### The Isomorphic Go Project
More information on the benefits of Isomorphic Go applications can be found at the [Isomorphic Go Website](http://isomorphicgo.org).

## License
The `isogoapp` is licensed under the BSD License. Read the [LICENSE](https://github.com/isomorphicgo/isogoapp/blob/master/LICENSE) file for more information.

