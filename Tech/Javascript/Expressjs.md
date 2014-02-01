### Express 
A web framework built on `connect` middleware framework.

`express = require('express'); express();`
under the hood
`createApplication()`
* extends `connect` with `application`
* attach `request`, `response` from base classes.
* `application.init()`

`application`
* configure defaults - implicitly use query middleware.
* sets up `mount` event to delegate req, res etc.
* proxies `app.use`, 
* Router
* wire up view template `app.engine('jade', require('jade').__express)`
* `app.param = function(name, fn)` - middleware to process param and load data into request.
* `app.set`, `app.get` - set/get in `settings` key value store.
* `app.render = function(name, options, fn)` - gets the view by name `view.render(options, fn)`
* `this._router[method].apply(this._router, arguments)` - loop through the methods and store the callbacks as instances of Route.

`request` - abstractions for HTTP request
- `req.get('Content-Type') // "text/plain"`
- `req.accepts('image/png') // undefined if not accepted`
- `acceptsCharset`, `acceptsLanguage` ... `range`, `param`, `is`, `secure`, `ip`, `host` etc

`response` - abstractions for HTTP response
- `status`, `links`
- `res.send = function(body)` - sets content-type, content-length, calls `end`
- `json` - call `send` with json as string.
- `sendfile` - streams and callbacks on done, `download` - wrapper of `sendfile`
- `res.format = function(obj)` - response to first accepts from request.
- `header`, `attachment`, `cookie`, `location`, `redirect`
- `res.render = function(view, options, fn)` - calls `app.render` with merged locals.

`view` - lightweight wrapper on top templateEngines
- `lookup = function(path)` => `if (exists(path)) return path;`
- `render = function(options, fn)` => `this.engine(this.path, options, fn)`, options = template locals
 
`router` - manage routes like `.route`, `.match`
- `dispatch` - invoke route's callbacks
- Route - object storing { path:, callback: }

### Connect
A http handler framework with middleware API and prebuilt handlers such as - cookieParser, csrf, query etc under `/middleware`
http://www.senchalabs.org/connect/

`app.use = function(route, fn)`
* fn can be an sub-app, `.handle` is wrapped as fn. Pushes to a stack.
* global middlewares are matched to `/` route.

`app.handle = function(req, res, out)`
* sets up `req` and `path`, execute matching middleware stack by calling next()
* inner function `function next(err)` - calls next handler in the stack or return out
   
