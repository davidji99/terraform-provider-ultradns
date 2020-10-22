// +build ignore

// This a hack to populate the version in the custom binary file as this provider is not official.

package main

import (
	"fmt"
	"terraform-provider-ultradns/version"
)

var ver = version.ProviderVersion

func main() {
	fmt.Println(ver)
}
