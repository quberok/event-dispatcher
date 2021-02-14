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
	"github.com/quberok/event-dispatcher/pkg/subscriber"
	"log"
)

// Create OrderSubscriber to hook into our order events.
type OrderSubscriber struct {
	logger log.Logger
}

// GetSubscribedEvents returns a map of event names this subscriber wants to listen to.
func (s *OrderSubscriber) GetSubscribedEvents() map[string][]subscriber.Listener {
	return map[string][]subscriber.Listener{
		OrderDeliveredEvent: {
			{Callback: s.log("Order № %d delivered successfully"), Priority: 0},
		},
		OrderCreatedEvent: {
			{Callback: s.log("Order № %d created successfully"), Priority: 0},
		},
		OrderPaidEvent: {
			{Callback: s.log("Order № %d paid successfully"), Priority: 0},
			{Callback: s.onPaid(), Priority: 10},
		},
	}
}

// Encapsulate the log function into our OrderSubscriber
func (s *OrderSubscriber) log(message string) func(event interface{}) {
	return func(event interface{}) {
		var oe orderEvent
		var ok bool

		if oe, ok = event.(orderEvent); !ok {
			return
		}

		s.logger.Printf(message, oe.order.number)
	}
}

// onPaid implementation some actions on paid
func (s *OrderSubscriber) onPaid() func(event interface{}) {
	return func(event interface{}) {
		var ope orderPaidEvent
		var ok bool

		if ope, ok = event.(orderPaidEvent); !ok {
			return
		}

		message := "Order numder: %$1d, details: %$2s, payment system: %$3s"

		s.logger.Printf(message, ope.order.number, ope.details, ope.paymentSystem)

		// or you can send sms to client, for example.
	}
}
