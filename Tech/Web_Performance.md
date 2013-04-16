# Web Performance:
```
 ___________
|  Browser  |
| --------- |
|  Network  |
| --------- |
|  App      |
| --------- |
|  Storage  |
|___________|
```

# Browser

* Parse HTML, Javascript, CSS
* Render incrementally.

## Parse HTML

* Bytes -> Characters -> Tokens -> Nodes -> Dom
* Dom is constructed incrementally as bytes arrives on the wire.
* Sync Javascript will block render of page.
* Defer - download in background, execute in order before DomContentLoaded
* Async - download in background, execute when ready

