package code_organization

import "fmt"

func GenericFunctionsDemo() {
	tags := []string{"go", "generic", "functions", "go", "generic"}
	fmt.Println("All Tags:", tags)
	fmt.Println("Unique Tags:", unique(tags))
	fmt.Println("Unique Tags:", unique2(tags))
}

func unique[T comparable](tags []T) []T {
	var result []T
	seen := make(map[T]bool)

	for _, tag := range tags {
		if !seen[tag] {
			seen[tag] = true
			result = append(result, tag)
		}
	}

	return result
}

func unique2[T comparable](items []T) []T {
	result := make([]T, 0, len(items)) // Pre-allocate for speed!
	seen := make(map[T]struct{})       // 0-byte values

	for _, item := range items {
		if _, exists := seen[item]; !exists {
			seen[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
