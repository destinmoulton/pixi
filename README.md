### Pixi

Pixi is a simple CLI file browser and video player. Pixi was built to browse and play videos on a Raspberry Pi, but it should be compatible with any flavor of linux. The default video player is omxplayer, but is easily configurable to whatever your platform supports.

### Features

- Navigate folders in an intuitive manner (inspired by [nnn](https://github.com/jarun/nnn))
- Play video files
- Video player history 
- Edit video player command
- Open other files (uses xdg-open)

### Requirements

-   go 1.10
-   omxplayer
-   xterm\*

\*Xterm is not required, but is recommended so that omxplayer gains keyboard focus when a video is played.

### Installation

```
git clone https://github.com/destinmoulton/pixi.git
cd pixi
go install pixi.go
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
| F5/Ctrl+r      | Refresh the list of files                         |
| q/Ctrl+c       | Quit pixi                                         |

### Debugging

To enable logging, pass the -log flag. This will create a pixi.log file in the working directory.

```
$ go run pixi.go -log
```

### License

MIT
