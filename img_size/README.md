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

## Development

- Recommended: Go 1.13+

This program uses [Go modules](https://github.com/golang/go/wiki/Modules) to
manage dependencies, and allow development outside the `$GOPATH`.
