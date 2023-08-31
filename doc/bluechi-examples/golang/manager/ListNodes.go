package main
// SPDX-License-Identifier: LGPL-2.1-or-later

import (
	"github.com/dougsland/bluechi/tree/go/src/bindings"
)

func main() {

	for _, node := range Manager.ListNodes() {
                fmt.Println(node)
        }
}
