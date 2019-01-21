# Playlist Zipper
[![Build Status](https://travis-ci.com/charliemaiors/playlist-zipper.svg?branch=master)](https://travis-ci.com/charliemaiors/playlist-zipper)
[![Maintainability](https://api.codeclimate.com/v1/badges/27b04bc52c80354a96f3/maintainability)](https://codeclimate.com/github/charliemaiors/playlist-zipper/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/27b04bc52c80354a96f3/test_coverage)](https://codeclimate.com/github/charliemaiors/playlist-zipper/test_coverage)

Playlist zipper will parse a playlist definition file, and add each file to a zip archive.

## Supported formats

The playlist supported formats are:

* m3u
* zpl (default, this tool was born to read Groove Music playlists)
* xspf

The archive formats are:

* bz2
* tar
* xz
* zip (default)

## Installation

```go get github.com/charliemaiors/playlist-zipper```

## Usage

Run ```playlist-zipper (---playlist-ext m3u --archive-type xz)``` and the executable will search into current directory for every file with the target extension, parse it and produce the archive with the same name of the playlist and the desidered type.