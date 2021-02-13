//
// This file is part of the quberok/event-dispatcher package.
//
// (c) Vladimir Podluzhnyy <quber.one@icloud.com>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.
//

package event

// Event is the base struct for structs containing event data.
//
// This struct contains no event data. It is used by events that do not pass
// state information to an event handler when an event is raised.
//
// You can call the method stopPropagation() to abort the execution of
// further listeners in your event listener.
type Event struct {
	propagationStopped bool
}

// IsPropagationStopped Is propagation stopped?
//
// This will typically only be used by the Dispatcher to determine if the
// previous listener halted propagation.
//
// True if the Event is complete and no further listeners should be called.
// False to continue calling listeners.
func (e *Event) IsPropagationStopped() bool {
	return e.propagationStopped
}

// StopPropagation stops the propagation of the event to further event listeners.
//
// If multiple event listeners are connected to the same event, no
// further event listener will be triggered once any trigger calls
// StopPropagation().
func (e *Event) StopPropagation() {
	e.propagationStopped = true
}
