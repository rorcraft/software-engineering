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
* `_retrieve` - wrap specs into batchedRequests and calls `async.parallel`. For each spec - if !readFromCache, `fetchFromApi`, else look at cache, if needsfetch, `fetchFromAPI` and caches data. (server defaults !readFromCache, client defaults readFromCache)
* `_retrieveModel` - wraps `_retrieveModelData` callback with data from `fetcher.modelStore`, `_retrieveModelData` - returns model or fetchFromApi.
* `fetchFromApi` - calls model/collection.fetch with data, success and error callbacks. (model/collection.fetch is a Backbone method)
* `fetch` - trigger `fetch:start`, calls `_retrieve`. This is called by controllers.  
```javascript
  index: function(params, callback) {
    this.app.set('title', 'Repos');

    var spec = {
      collection: {collection: 'Repos', params: params}
    };
    this.app.fetch(spec, function(err, result) {
      callback(err, result);
    });
  },
```
the callback will combine the fetchedData with template.

`model` 
* extends `Backbone.Model`, `syncer`
* sets app (rendrApp)
* on 'change', store
* `store` - `app.fetcher.modelStore.set(this)`

`collection`
* extends `Backbone.Collection`, `syncer`
* sets app, model. wraps `fetch`, `store` - `this.app.fetcher.collectionStore.set(this)`

`syncer`
* `serverSync` - calls `req.dataAdapter.request`, attach resp `{ body, status}` to `success` or `error` callbacks.
* `clientSync` - calls `Backbone.sync(method, model, options)`

`modelStore`, `collectionStore`
* extends `memoryStore` - key/value with expiration
* key - 'modelName:id' , collection is more complicated.

`view`
* extends `Backbone.View`
* initialize with `parseOptions` and then call `postInitialize`
* `parseOptions` - sets app, parentView, model (model_name, model_id), collection - on `this` as well as `options`.
* `getTemplateData` - model/collection.toJSON
* `getTemplate` - delegates to `templateAdapter`
* `getInnerHtml`, `getHtml` - renders template with data and add html attributes.
* `render` - attach html to dom and calls `_postRender` 

All Together:
```javascript
// high level psuedo code:
controller {
  action: function(params, callback)  {
    var spec = {
      collection: {collection: 'Repos', params: params}
    };
    this.app.fetch(spec, function(err, result) {
      callback(err, result);
    });
  }
}

// buildRoutes
Router.route('controller/action') {
  handler = function(req, res) { /* wraps action() with viewData */ }
  routes.push(['controller/action', Route, handler]) 
}


// server/server
this.handle
  this.configure // lazy configure
    this.expressApp.use // overwrite res.end middleware
    this.expressApp.use // attach req.rendrApp + appData, apiPath, entryPath, modelUtils
    // add custom middleware
    this.expressApp.use // apiPath, apiProxy
    buildRoutes // attach wrappedhandler to expressApp
      // for each controller/action, wrap action into wrapperHandler
      // builds an array of [pattern, route, wrappedhandler]
      // for each route
        this.expressApp.get(pattern, wrappedHandler)
    this.expressApp.use // errorHandler
  this.expressApp.handle // normal handle that calls a Router wrapped handler. 
    handler(req, res) {
      context = { route, app, redirectTo() }
      context.action(params, fetch_callback(err, viewPath, locals) {
        // this.app.fetch(spec, function(err, result))
          fetch_callback(err, result) {
            // res.render(viewPath, viewData, function(err, html) {})
            ViewEngine.render(viewPath, viewData, callback(err, html) {}) {
              layoutData = { body, appData, bootstrappedData, _app }
              body = new View(locals).getHtml();
              html = templateFn(layoutData)
              // callback(err, html) 
                res.end(html);
            
```

Client side:

`app`
* sets `ClientRouter` - `{ app: this, entryPath: '', appViewClass: require('client/app_view'), rootPath}`
* bootstrapData, start

`view`
* `attach` - `hydrate` - `preRender(), postRender(), fetchLazy, trigger('attach')`

`app_view`
* extend `BaseView` (Backbone.View)
* attaches to `<body>` and `#content`
* interceptClicks and manage history, calls `app.router.redirectTo`

`router`
* wraps around Backbone.Router, extends BaseRouter
* on `route:add` calls addBackboneRoute
* on initialize, calls appView.render, buildRoutes and postInitialize
* `addBackboneRoute` - Backbone.Router.route(pattern, "controller:action", handler)
* `getHandler` - attach MainView on first render, or call actionCall(params, getRenderCallback)
* `navigate` - delegates to Backbone.Router
* `getRenderCallback` - BaseView.getView from route and render with locals.
