//
// This file is part of the quberok/event-dispatcher package.
//
// (c) Vladimir Podluzhnyy <quber.one@icloud.com>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.
//

package subscriber

// A Listener is similar to EventSubscriber, but listener doesn't know itself what events it is interested in.
// Using listeners separately are more flexible because your app can enable or disable each of them conditionally depending on some configuration value.
type Listener struct {
	Callable func(event interface{})
	Priority int
}

// An EventSubscriber knows itself what events it is interested in.
type EventSubscriber interface {
	// Returns a map of event names this subscriber wants to listen to.
	//
	// The code must not depend on runtime state.
	// All logic depending on runtime state must be put into the individual methods handling the events.
	//
	// The first string argument contains the event name to listen to
	GetSubscribedEvents() map[string][]Listener
}
