//
// This file is part of the quberok/event-dispatcher package.
//
// (c) Vladimir Podluzhnyy <quber.one@icloud.com>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.
//

package main

import (
	"math/rand"
)

// Our fake entity
type Order struct {
	number int
}

// newOrder returns a new order
func newOrder() Order {
	return Order{
		number: rand.New(rand.NewSource(99)).Int(),
	}
}
