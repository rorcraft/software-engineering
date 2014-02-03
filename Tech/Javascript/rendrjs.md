### Rendr

Server side:

`server` 
* sets up `dataAdapter = RestAdapter`, `viewEngine`, `router`
* `expressApp = express()`, set `views`, `views engine`, `engine`
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

`BaseRouter`
* extends `Backbone.Events`
* `buildRoutes` - loads `routes.js` and push into an array. Each call `route`
* `route` - split pattern into controller, action. gets `getHandler`, stores `routeObj = [pattern, route, handler]` into _routes and trigger 'route:add' with `routeObj`. 

Shared:

`app` 
* extends Backbone.Model
* sets `templateAdapter`
* sets `fetcher`
* calls `postInitialize()`

`fetcher`
* knows `app`, `modelUtils`, `modelStore`, `collectionStore`, extends `Backbone.Events`
* `getModelOrCollecionForSpec`, really delegates to `modelUtils.getModel` or `modelUtils.getCollection`, spec is `{ model: { } }` or `{ collection: [] }`
* `_retrieve` - wrap specs into batchedRequests and calls `async.parallel`. For each spec - if !readFromCache, fetch from api, else look at cache, if needsfetch, `fetchFromAPI` and caches data. (server defaults !readFromCache, client defaults readFromCache)
* `_retrieveModel` - wraps `_retrieveModelData` callback with data from `fetcher.modelStore`
* `fetchFromApi` - calls model/collection.fetch with data, success and error callbacks.
* `fetch` - trigger `fetch:start`, calls `_retrieve`. This is called by View as it renders, the callback will combine the fetchedData with template.

Client side:

`app`
* sets `ClientRouter` - `{ app: this, entryPath: '', appViewClass: require(`client/app_view`), rootPath}`
* bootstrapData, start
