# Image Size Tool

This Go program calculates the size, in megapixels, of all JPG & PNG files in
folders descending from the current working directory, and groups the results
into buckets.

## Usage

```sh
# Defaults to all JPG & PNG files, including nested directories,
# descending from the current working directory.
img_size

# or

img_size '<glob_pattern>'
img_size '**/*.png' # Only PNGs
img_size '**/*.jpg' # Only JPGs
```

## Download

* [Windows](https://github.com/remove-bg/integration/raw/master/img_size/dist/windows/img_size.exe)
* [MacOS](https://github.com/remove-bg/integration/raw/master/img_size/dist/macos/img_size)
* [Linux](https://github.com/remove-bg/integration/raw/master/img_size/dist/linux/img_size)

## Development

- Recommended: Go 1.13+

This program uses [Go modules](https://github.com/golang/go/wiki/Modules) to
manage dependencies, and allow development outside the `$GOPATH`.

### Build

Run `./bin/build` to build for Linux, Mac and Windows.
