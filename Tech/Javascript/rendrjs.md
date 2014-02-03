### Rendr

Server side:

`server` 
* sets up `dataAdapter = RestAdapter`, `viewEngine`, `router`
* `expressApp = `express()`, set `views`, `views engine`, `engine`
* wraps `handle` to expressApp.handle, configure a middleware in between. The middleware is a Rendr App, intercepts Api paths '/-/', sets up `req.rendrApp = new App(attrs, { req, entryPath, modelUtils})` 
* middleware sets `req.dataAdapter`, dereference dataAdapter after `res.end`
* `buildRoutes` - get routes from `router.buildRoutes` and attach to expressApp (only `get` method ?)
* adds `errorHandler` middleware

`router`
* inherits `BaseRouter`
* connects `express.router`
* `getHandler` - returns an Express middleware function. Sets up rendrApp context (route, app, redirectTo), calls `action` with handler function that sets up viewData and calls express' res.render which is already connected to the Rendr ViewEngine.
* `addExpressRoute` - on `route:add`, calls `_expressRouter.route('get', path, [])`, stores in `routesByPath`
* `match` - returns route definition from `routesByPath`

`viewEngine`
* `render` sets up layoutData, body, app, then `renderWithLayout`
* `getLayoutTemplate` - get '__layout' from templateAdapter
* `getViewHtml` - get View instance of BaseView + locals `.getHTML()`
* `getBootstrappedData` - setup modelOrCollection data from locals.
