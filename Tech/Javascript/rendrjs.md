### Rendr

Server side:

`server` 
* sets up `dataAdapter = RestAdapter`, `viewEngine`, `router`
* `expressApp = `express()`, set `views`, `views engine`, `engine`
* wraps `handle` to expressApp.handle, configure a middleware in between
* middleware sets `req.dataAdapter`, dereference dataAdapter after `res.end`
* `buildRoutes` - get routes from `router.buildRoutes` and attach to expressApp (only `get` method ?)

`router`

`viewEngine`
* `render` sets up layoutData, body, app, then `renderWithLayout`
* `getLayoutTemplate` - get '__layout' from templateAdapter
* `getViewHtml` - get View instance of BaseView + locals `.getHTML()`
* `getBootstrappedData` - setup modelOrCollection data from locals.
