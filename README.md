# xdg

[![GoDoc](https://godoc.org/github.com/tehbilly/xdg?status.svg)](https://godoc.org/github.com/tehbilly/xdg)

A small utility library to easily get the XDG Base Directory specification's
directories in a cross-platform manner. Windows works!

It does this by making no attempt to use `%APPDATA%` or the like, simply
using the user's home directory.
