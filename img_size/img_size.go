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
	globPattern := fetchGlobPattern()
	paths := findImagePaths(globPattern)
	imageMetadatas := collectMetadata(paths)
	buckets := countIntoBuckets(imageMetadatas)

	fmt.Println("\nResolution: Count\n-----------------")
	for _, bucket := range buckets {
		fmt.Println(bucket.Description())
	}
}

func fetchGlobPattern() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}

	return "./**/*.{jpg,JPG,jpeg,JPEG,png,PNG}"
}

func findImagePaths(globPattern string) []string {
	paths, err := doublestar.Glob(globPattern)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Pattern: '%s' matched %d files\n\n", globPattern, len(paths))

	return paths
}

func collectMetadata(paths []string) []imageMetadata {
	results := make([]imageMetadata, len(paths))

	for _, path := range paths {
		megapixels := calculateImageMegapixels(path)
		results = append(results, imageMetadata{path, megapixels})
	}

	return results
}

func calculateImageMegapixels(path string) float64 {
	file, _ := os.Open(path)
	defer file.Close()

	config, _, err := image.DecodeConfig(file)

	if err != nil {
		fmt.Printf("%s: %s\n", path, err)
		return -1
	}

	return float64(config.Width*config.Height) / oneMillion
}

const oneMillion = 1000000.0

func countIntoBuckets(imageMetadatas []imageMetadata) []*bucket {
	buckets := createBuckets()

	for _, metadata := range imageMetadatas {
		for _, bucket := range buckets {
			if bucket.Match(metadata) {
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

type imageMetadata struct {
	path       string
	megapixels float64
}

type bucket struct {
	start float64
	end   float64
	count int
}

func (b bucket) Match(metadata imageMetadata) bool {
	return metadata.megapixels > b.start && metadata.megapixels <= b.end
}

func (b *bucket) Increment() {
	b.count++
}

func (b bucket) Description() string {
	return fmt.Sprintf("%.2f-%.2f:\t%d", b.start, b.end, b.count)
}
