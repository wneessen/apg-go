// SPDX-FileCopyrightText: 2021-2024 Winni Neessen <wn@neessen.dev>
//
// SPDX-License-Identifier: MIT

package main

import (
	"fmt"

	"github.com/wneessen/apg-go"
)

func main() {
	config := apg.NewConfig(apg.WithAlgorithm(apg.AlgoRandom),
		apg.WithModeMask(apg.ModeNumeric),
		apg.WithFixedLength(6))
	generator := apg.New(config)
	password, err := generator.Generate()
	if err != nil {
		panic(err)
	}
	fmt.Println("Your PIN:", password)
}
