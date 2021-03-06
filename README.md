# webexec

![Test](https://github.com/tuzig/webexec/workflows/Test/badge.svg)

webexec is a terminal server running over WebRTC with http for signalling.
Webexec listens for connection requests, executes commands over pseudo ttys
and pipes their I/O over WebRTC data channels.

webexec is currently single session - each client that authenticates has
access to the same list of "Panes" and layout geomtry. 

webexec exposes TCP port 7777 by default, to support signalling.
There's a single endpoint `/connect`: The client exchanges tokens with the
server and then initiates a WebRTC connection.

## Install

The easiest way to install is to download the 
[latest release](https://github.com/tuzig/webexec/releases) tar ball for your
system and extract it to get the binary file. 
We recommended moving webexec to a system-wide tools folder such as 
`/usr/local/bin`.

If you want to include it in your go project:

``` console

% GO111MODULE=on go get github.com/tuzig/webexec/...

```

## Quickstart


```console
$ ./webexec start
```

This will launch a user's agent and report its process id 

## Contributing

We welcome bug reports, ideas for new features and pull requests.
If you are new to open source, DON'T PANIC. Just follow these simple
steps:

1. Fork it and clone it `git clone <your-fork-url>`
2. Create your feature branch `git checkout -b my-new-feature`
3. Write tests for your new feature and run them `go test -v`
4. Commit the failed tests `git commit -am 'Testing ... '`
4. Write the code that psses the tests and commit 
4. Push to the branch `git push --set-upstream origin my-new-feature`
5. Create a new Pull Request
