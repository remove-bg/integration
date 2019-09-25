package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math"
	"os"

	"github.com/bmatcuk/doublestar"
)

func main() {
	paths := findImagePaths()
	pathsMegapixels := calculateImageSizes(paths)
	buckets := countIntoBuckets(pathsMegapixels)

	fmt.Println("Resolution: Count\n-----------------")
	for _, bucket := range buckets {
		fmt.Println(bucket.Description())
	}
}

func findImagePaths() []string {
	paths, err := doublestar.Glob("./**/*.{jpg,JPG,jpeg,JPEG,png,PNG}")

	if err != nil {
		log.Fatal(err)
	}

	return paths
}

func calculateImageSizes(paths []string) map[string]float64 {
	result := make(map[string]float64, len(paths))

	for _, path := range paths {
		result[path] = calculateImageSize(path)
	}

	return result
}

func calculateImageSize(path string) float64 {
	file, _ := os.Open(path)
	defer file.Close()

	config, _, _ := image.DecodeConfig(file)
	return float64(config.Width*config.Height) / oneMillion
}

const oneMillion = 1000000.0

func countIntoBuckets(pathsMegapixels map[string]float64) []*bucket {
	buckets := createBuckets()

	for _, megapixels := range pathsMegapixels {
		for _, bucket := range buckets {
			if bucket.Match(megapixels) {
				bucket.Increment()
			}
		}
	}

	return buckets
}

func createBuckets() []*bucket {
	return []*bucket{
		&bucket{0.00, 0.25, 0},
		&bucket{0.25, 1, 0},
		&bucket{1, 2, 0},
		&bucket{2, 4, 0},
		&bucket{4, 8, 0},
		&bucket{8, 12, 0},
		&bucket{12, 16, 0},
		&bucket{16, 20, 0},
		&bucket{20, 25, 0},
		&bucket{25, 30, 0},
		&bucket{30, 35, 0},
		&bucket{35, 40, 0},
		&bucket{40, 45, 0},
		&bucket{45, 50, 0},
		&bucket{50, 75, 0},
		&bucket{75, 100, 0},
		&bucket{100, 200, 0},
		&bucket{200, math.MaxInt64, 0},
	}
}

type bucket struct {
	start float64
	end   float64
	count int
}

func (b bucket) Match(value float64) bool {
	return value > b.start && value <= b.end
}

func (b *bucket) Increment() {
	b.count++
}

func (b bucket) Description() string {
	return fmt.Sprintf("%.2f-%.2f:\t%d", b.start, b.end, b.count)
}
