http://backbonejs.org/docs/backbone.html

* version 1.1.0
* requires underscore.js
* Backbone.$ = root.jQuery || root.Zepto || root.ender || root.$;
* legacy support - emulateHTTP (fake patch, put, delete), emulateJSON (fake ajax with form post)

### Backbone.Events

* mixed into Model, Collection, View and Backbone itself.
* `on` (bind), `once`, `off` (unbind), `trigger` (also trigger listening on 'all'), `stopListening`
* eventsApi - parses events
* triggerEvents - custom dispatch function - calling .call instead of .apply for args <= 3.

### Backbone.Model

* idAttribute - defaults to 'id'
* empty `initialize` function
* sync
* `get`, `escape`, `has`, `set`, `set`
* `set` - chunk of the logic - triggering 'change' events, handle previous attributes, keep track of changes. (`silent` - won't trigger change events)
* `unset`, `clear`, `hasChanged`, `changedAttributes`, `previous`, `previousAttributes` (clones)
* `fetch` - calls sync to 'read' data from server and handle changes.
* `save` - create/update ajax sync with server, handles validation and waiting states.
* `destroy` - trigger destroy and sync delete.
* `url`, `parse` (json data to model), `clone`, `isNew`, `isValid`
* mix in `keys`, `values`, `pairs`, `invert`, `pick`, `omit` from underscore on attributes.

### Backbone.Collection

* `options.comparator` - maintains order of the model objects
* `model` - defines collection of model types
* `initialize` - empty function interface
* `toJSON` - maps the collection of model.toJSON
* `sync` - proxy to Backbone.sync
* `add`, `remove`
* `set` - add, remove and merge with the list of models. Trigger add, remove, change events on the model.
* `reset` - replace models without firing add, remove events.
* `push`, `pop`, `unshift`, `slice`, `get` (by id, cid, obj), `at`, `where` match attributes (return first or all), `findWhere` == where(attrs, true)
* `sort`, `pluck`
* `fetch` - sync 'read' from server - reset or set.
* `create` - new model and save to server, waits for server to confirm.
* `parse`, `clone`
* also a bunch of list/array methods mixed in from underscore.

### Backbone.View

* represents a logical chunk of UI in the DOM
* `tagName` - 'div'
* `$` - scoped to $el
* `render` - empty interface method, should build HTML on $el and return this
* `remove` - remove from DOM and stopListening to events
* `setElement` - changes $el, undelegateEvents and re-delegateEvents
* `delegateEvents` - binds events to specific elements. - does not work for 'focus', 'blur' and 'change', 'submit', 'reset' on IE.
* `undelegateEvents` - unbinds all events

### Backbone.sync

* an abstraction over plain ajax calls
* handles the emulateHTTP and emulateJSON
* for advanced features, this need to be overridden. (batching, websocket, xml)

### Backbone.Router

* Routers map faux-URLs to actions, and fire events when routes are matched.
* `initialize` empty interface method
* `route` - binds a route to a function
* `navigate` - proxy to Backbone.history

### Backbone.History

* cross-browser history management, based on either pushState and real URLs, or onhashchange and URL fragments. 
* `start` - begin the history tracking, figure out configration / strategy
* `stop` - Disable Backbone.history, perhaps temporarily. Not useful in a real app, but possibly useful for unit testing Routers
* `checkUrl` - check for url changes, if changed, calls loadUrl
* `loadUrl` - if valid route, calls the callback
* `navigate` - loadUrl and add to history, `replace: true` to not add to history.

