### Express 
A web framework built on `connect` middleware framework.

`express = require(`express`); express();`
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
* `app.render`


### Connect
A http handler framework with middleware API and prebuilt handlers such as - cookieParser, csrf, query etc under `/middleware`
http://www.senchalabs.org/connect/

`app.use = function(route, fn)`
* fn can be an sub-app, `.handle` is wrapped as fn. Pushes to a stack.
* global middlewares are matched to `/` route.

`app.handle = function(req, res, out)`
* sets up `req` and `path`, execute matching middleware stack by calling next()
* inner function `function next(err)` - calls next handler in the stack or return out


merge
mixin
proto
Route
Router
req
res
   
