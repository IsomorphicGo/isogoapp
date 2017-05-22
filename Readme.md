<p align="center"><a href="http://isomorphicgo.org" target="_blank"><img src="https://github.com/isomorphicgo/isogoapp/blob/master/static/images/isomorphic_go_logo.png"></a></p>

# Isomorphic Go App (isogoapp)
-
[![Go Report Card](https://goreportcard.com/badge/github.com/isomorphicgo/isogoapp)](https://goreportcard.com/report/github.com/isomorphicgo/isogoapp)

A basic, barebones web app, intented to be used as a starting point for developing an [Isomorphic Go](http://isomorphicgo.org) application.

## Installation on Unix-like OS

The instructions assume that you are using a Unix-like operating system (e.g. BSD, Linux, Mac OS, etc).

### Prerequisites

It is assumed that the latest version of Go is already installed on your system and that you have properly configured your Go workspace.

### Get the isogoapp project
`go get -u github.com/isomorphicgo/isogoapp`


### Define the `ISOGO_APP_ROOT` environment variable.
`export ISOGO_APP_ROOT=${GOPATH}/src/github.com/isomorphicgo/isogoapp`


### Build the Client-Side Go App
`cd $ISOGO_APP_ROOT/client`

`go get ./...`

`gopherjs build`

### Run the Web Server Instance

`cd $ISOGO_APP_ROOT`

`go run isogoapp.go`


The web server runs locally on port 8080: [localhost:8080](http://localhost:8080).

If all goes well, you should see the Isomorphic Go logo, along with a message rendered on the web page.


## Installation on Windows OS

### Prerequisites

It is assumed that the latest version of Go is already installed on your system and that you have properly configured your Go workspace. Make sure that GOPATH, GOROOT and GOBIN are set up. GOROOT and GOPATH were set up during the installation and configuration of Golang. To set up GOBIN, you have to edit environment variables.

Open the envrionment variables window and,

1. Find a variable called 'Path' under User Variables section. Focus on it and click 'Edit'. A new pop-up would open and there you would find a button called New in the top-right corner. Click on it, and it will bring blinking cursor in the focus. The cursor will be placed at the end of the list of paths. So if your Go workplace is C:\GoWorkspace, then you would type `C:\GoWorkspace\bin` there.

2. In the 'System Variables' section, click on New. Enter `GOBIN` in the Variable Name field and, `C:\GoWorkspace\bin\` in the Variable value field.

We have assumed that your go workspace is called GoWorkspace and is located at 'C:\GoWorkspace'. It is highly possible that your go workspace's name is not same as the example taken here, please modify the name of your go workspace according to your configuration.

Save the changes made and run 'go env' from cmd. You should now see GOBIN to be pointing to bin directory inside your Go workspace folder.

### Get the isogoapp project
Enter the command `go get -u github.com/isomorphicgo/isogoapp` in cmd.

If Go is properly configured in your system and command is successfully executed, then you will see the following directory structure in src folder of your Go workspace.

C:\GoWorkspace\src\github.com\isomorphicgo\isogoapp

### Define the `ISOGO_APP_ROOT` environment variable.
Now, we again need to set up an environment variable. Open the environment variable screen and under System Variables section, create a New variable. On Windows 10, this is done by clicking on New.

In the Variable name field, enter `ISOGO_APP_ROOT`.

And, in the Variable value field, enter `%GOPATH%src\github.com\isomorphicgo\isogoapp`

Did you notice back-slash missing between %GOPATH% and src? It is not a typo. If your GOPATH has a trailing back-slash in it, then type the command as given above. And if it does not have back-slash, then enter `%GOPATH%\src\github.com\isomorphicgo\isogoapp`. Check this detail by typing 'go env' in cmd.

### Build the Client-Side Go App

Now close all the instances of cmd relevant to our work as they don't know about the environment variable changes that have been made. 

Open a new cmd window. And run this command `cd %ISOGO_APP_ROOT%/client`.

You would now be pointed to the client folder inside isogoapp folder.

Now run `go get ./...` command. And after its successful execution, run `gopherjs build` command. 

Make sure that you don't move away from client directory during the execution of all these three commands.

### Run the Web Server Instance

Now, open a new cmd windows and run `cd %ISOGO_APP_ROOT%`.

After this change of directory, run `go run isogoapp.go`.

The web server runs locally on port 8080: [localhost:8080](http://localhost:8080).

If all goes well, you should see the Isomorphic Go logo, along with a message rendered on the web page, when you visit the link above.


## Where to now?

### Kick

Once you've installed, and confirmed, that the basic Isomorphic Go application is working, the next logical step is to install [Kick](https://github.com/isomorphicgo/kick).

Kick automatically recompiles Go code, and it has the ability to take both the `go` and `gopherjs` commands into consideration. 

Kick performs an *instant kickstart* of the web server instance, upon the modification of a Go source file. [Check out Kick](https://github.com/isomorphicgo/kick)

Please note, that Kick will only work on Unix-like operating systems (e.g., BSD, Linux, Mac OS).

### The Isomorphic Go Project
More information on the benefits of Isomorphic Go applications can be found at the [Isomorphic Go Website](http://isomorphicgo.org).

## License
The `isogoapp` is licensed under the BSD License. Read the [LICENSE](https://github.com/isomorphicgo/isogoapp/blob/master/LICENSE) file for more information.

