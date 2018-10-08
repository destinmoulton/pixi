### Pixi

Pixi is a simple CLI file browser and video player. Pixi was built to browse and play videos on a Raspberry Pi, but it should be compatible with any flavor of linux. The default video player is omxplayer, but is easily configurable.

### Requirements

-   go 1.10
-   omxplayer
-   xterm\*

Xterm is not required, but is recommended so that omxplayer gains keyboard focus when a video is played.

### Installation

```
go get github.com/destinmoulton/pixi
```

### Configuring the Video Player

The video player can be configured on the Settings screen. Launch pixi and press 's' to access the settings. Command line parameters can be altered to match your Pi configuration (ie audio output).

### Keyboard Commands

| Key            | Description                                       |
| -------------- | ------------------------------------------------- |
| up/down arrows | Select directories/files                          |
| left arrow     | Navigate to parent directory                      |
| right arrow    | Navigate into directory                           |
| enter/return   | Play selected file with omxplayer                 |
| h              | Show history of played files                      |
| c              | Clear history of played files (when history open) |
| s              | Open player settings                              |
| \>/.           | Toggle viewing hidden files                       |
| q/Ctrl+c       | Quit pixi                                         |


### Debugging

To enable logging, pass the -log flag. This will create a pixi.log file in the working directory.

```
$ go run pixi.go -log
```

### License

MIT
