# labe

a river wm

## dev
### regenerating proto package from availible protocols
```sh
go generate ./internal/proto/proto.go
```
### running a nested live reloading session
```sh
river &
go tool air
```
> [!TIP]
> you might want to run this in 2 terminal sessions, the combined output is a bit of a mess
### running live reloading session from a seperate TTY
```sh
WAYLAND_DISPLAY=<wayland-display> \ # sets the proper wayland display, defaults to nested"wayland-2"
LABE_LOGLEVEL=<log-level> \ # sets the log level, for dev you will most likely want debug
CLICOLOR_FORCE=1 \ # forces log to produce ansi escape codes despite being piped
go tool air 2>&1 | tee /tmp/labe.log
```
while just go tool air would work with the wayland display defined, this way you can `tail` the logs from within the session `tail -f /tmp/labe.log` and the formating stays colored with the ansi codes
