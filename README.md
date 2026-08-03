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
