//
// This file is part of the quberok/event-dispatcher package.
//
// (c) Vladimir Podluzhnyy <quber.one@icloud.com>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.
//

package event

// An StoppableEvent whose processing may be interrupted when the event has been handled.
//
// A Dispatcher implementation MUST check to determine if an Event
// is marked as stopped after each listener is called. If it is then it should
// return immediately without calling any further Listeners.
type StoppableEvent interface {
	IsPropagationStopped() bool
}
