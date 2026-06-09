// Package pkge2 demonstrates how to use an external package in Go.
// go get <package-name> to install the package before using it
// example: go get github.com/google/uuid
// go get github.com/google/uuid@1.2.0 to install a specific version of the package
package utils

import "github.com/google/uuid"
import "fmt"
func ExternalPackageExample() {
	id1 := uuid.New()
	fmt.Println("Generated UUID:", id1)
}