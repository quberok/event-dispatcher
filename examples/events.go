//
// This file is part of the quberok/event-dispatcher package.
//
// (c) Vladimir Podluzhnyy <quber.one@icloud.com>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.
//

package main

import "github.com/quberok/event-dispatcher/pkg/event"

// Our events
const (
	ORDER_CREATED_EVENT   = "order.created"
	ORDER_PAID_EVENT      = "order.confirmed"
	ORDER_DELIVERED_EVENT = "order.delivered"
)

// orderEvent any data related to order place here
type orderEvent struct {
	event.Event
	order interface{}
}

// orderPaidEvent extends orderEvent to store some payment details
type orderPaidEvent struct {
	orderEvent
	paymentSystem string
	details       string
}
