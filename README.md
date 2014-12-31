## Outkast

This is a port of a node.js module that I create.  It is a fun project that randomly removes a player from a game.

#### Installation

go get github.com/Dahs81/go-outkast

#### Usage

```
$ go-outkast Bill Lisa Jen Mike
# returns three of the four names.  Randomly removes one of the names.
```

```
# Removing multiple items from list of args
$ go-outkast --rand Bill Lisa Jen Mike
OR
$ go-outkast -r 2 Bill Lisa Jen Mike
# Randomly removes 2 items
```
