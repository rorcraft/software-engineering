http://www.igvita.com/2013/01/15/faster-websites-crash-course-on-web-performance/

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

## Request lifecycle

http://blog.catchpoint.com/2010/09/17/anatomyhttp/

* DNS lookup
* tcp handshake
* first packet
* stream HTML
* browser parse html
* more HTTP requests
* TBD...

# Browser

[Blink’s Rendering Pipeline](https://docs.google.com/a/change.org/document/d/1wYNK2q_8vQuhVSWyUHZMVPGELzI0CYJ07gTPWP1V1us/pub)

* Parse HTML, Javascript, CSS
* Render incrementally.

## Parse HTML

* Bytes -> Characters -> Tokens -> Nodes -> Dom
* Dom is constructed incrementally as bytes arrives on the wire.
* Sync Javascript will block render of page.
* Defer - download in background, execute in order before DomContentLoaded
* Async - download in background, execute when ready

