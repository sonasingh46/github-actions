/*
Copyright (c) 2020-present Author(sonasingh46@gmail.com). All Rights Reserved.
See LICENSE.txt for license information.
*/

package main

import (
	"fmt"
	"github.com/sonasingh46/github-actions/pkg/foo"
)

func main() {

	number := 3
	if foo.IsEven(number) {
		fmt.Println("%d is even", number)
	} else {
		fmt.Println("%d is odd", number)

	}
}
