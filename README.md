### Pixi

Pixi is a simple command line video browser for the Raspberry PI. Video files are played via omxplayer.

### Requirements

-   go 1.10
-   omxplayer
-   xterm

Xterm is required so that omxplayer gains keyboard focus when a video is played.

### Installation

```
go get github.com/destinmoulton/pixi
```

### Debugging

To enable logging, pass the -log flag. This will create a pixi.log file in the working directory.

```
$ go run pixi.go -log
```

### License

MIT
