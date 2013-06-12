# HTTP Caching

1. Cache by expiration time.
1. Conditional Get by versioning.


## Cache by expiration time.
_Tells the client and intermediary cache what to cache._

HTTP 1.0:
```
Expires: Thu, 01 Dec 1994 16:00:00 GMT
```
Gives the date/time after which the response is considered stale. _Always returns an absolute date._

HTTP 1.1:
```
Cache-Control: max-age=3600
```
_Uses relative times in seconds, cf (s)max-age_

```
    Cache-Control   = "Cache-Control" ":" 1#cache-directive
    cache-directive = cache-request-directive
         | cache-response-directive
    cache-request-directive =
           "no-cache"                          ; Section 14.9.1
         | "no-store"                          ; Section 14.9.2
         | "max-age" "=" delta-seconds         ; Section 14.9.3, 14.9.4
         | "max-stale" [ "=" delta-seconds ]   ; Section 14.9.3
         | "min-fresh" "=" delta-seconds       ; Section 14.9.3
         | "no-transform"                      ; Section 14.9.5
         | "only-if-cached"                    ; Section 14.9.4
         | cache-extension                     ; Section 14.9.6
     cache-response-directive =
           "public"                               ; Section 14.9.1
         | "private" [ "=" <"> 1#field-name <"> ] ; Section 14.9.1
         | "no-cache" [ "=" <"> 1#field-name <"> ]; Section 14.9.1
         | "no-store"                             ; Section 14.9.2
         | "no-transform"                         ; Section 14.9.5
         | "must-revalidate"                      ; Section 14.9.4
         | "proxy-revalidate"                     ; Section 14.9.4
         | "max-age" "=" delta-seconds            ; Section 14.9.3
         | "s-maxage" "=" delta-seconds           ; Section 14.9.3
         | cache-extension                        ; Section 14.9.6
    cache-extension = token [ "=" ( token | quoted-string ) 
```

Request a fresh copy.
```
Cache-Control: no-cache
```
> Pragma: no-cache
The Pragma request header is a legacy header and should no longer be used.

```max-age:``` Specifies the period in seconds during which the cache will be considered fresh. On a request - how long has it already been cached. On a response, how long it can be cached for in the future.

__no-cache VS max-age=0__
* no-cache: force reload
* max-age=0: force revalidation. _Caches can avoid redownloading data, but they’re not sure if it’s still fresh_. See Conditional Get.

__Expires VS Cache-Control__
> When both Cache-Control and Expires are present, Cache-Control takes precedence.
> If a response includes both an Expires header and a max-age directive, the max-age directive overrides the Expires header, even if the Expires header is more restrictive

Response specific:
* ```public, private``` whether it is cacheable on intermediary proxies
* ```must-revalidate, proxy-revalidate``` Tells client to revalidate even though they can cache.
* ```no-store``` Not cacheable.

# Conditional Get by versioning.

* timestamp ```Last-Modified``` (HTTP 1.0)
* checksum  ```ETag``` (HTTP 1.1)

> Usually use both Last-Modified and ETag together.

__If-Modified-Since -> Last-Modified__
* Client stores the last modified at timestamp.

```
# Request
If-Modified-Since: Wed, 01 Sep 2004 13:24:52 GMT

# Response
# If not modified since the time specified, no data is returned.
HTTP/1.1 304 Not Modified
Last-Modified: Wed, 01 Sep 2004 13:24:52 GMT

# If modified since the time specified, no data is returned.
HTTP/1.1 200 OK
Last-Modified: Wed, 01 Sep 2004 13:24:52 GMT
(data)
```

Caveat - sub-second updates may never be realized.

__ETag -> If-None-Match__
* Clients stores the ETag
* ETag usually hash of important information on data that might represent change.

```
# Request
If-None-Match: "1edec-3e3073913b100"

# Response
# If not modified since the time specified, no data is returned.
HTTP/1.1 304 Not Modified
ETag: "1edec-3e3073913b100"

# If modified since the time specified, no data is returned.
HTTP/1.1 200 OK
ETag: "abcdef-3e3073913b100"
(data)
```

Caveat:
* gzip file format has a timestamp in it that means the ETag will change every time you re-compress it.

## Example 

Time@0s - Request from Client with empty cache 
```
GET / HTTP/1.1
Host: localhost
```

```
# Server Response
HTTP/1.1 200 OK
Cache-Control: max-age=10
ETag: "1edec-3e3073913b100"
Last-Modified: Wed, 01 Sep 2004 13:24:52 GMT
(other headers .e.g content-length, content-type)
(data)
```

* Client caches - content, etag, last-modified - expires in: 10 seconds.

Time@3s - Client can use stored cache, makes no request.

Time@10s - Client sends ETag and Last-Modified to revalidate.
```
GET / HTTP/1.1
Host: localhost
Cache-Control: max-age=0
If-None-Match: "1edec-3e3073913b100"
Last-Modified-Since: Wed, 01 Sep 2004 13:24:52 GMT
```

* Server load data, if ETag == "1edec-3e3073913b100" && updated_at <= Wed, 01 Sep 2004 13:24:52 GMT

```
HTTP/1.1 304 Not Modified
Cache-Control: max-age=10
ETag: "1edec-3e3073913b100"
Last-Modified: Wed, 01 Sep 2004 13:24:52 GMT
```



Reference:
* http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html
* https://www.varnish-software.com/static/book/HTTP.html
* http://en.wikipedia.org/wiki/List_of_HTTP_header_fields
* http://palizine.plynt.com/issues/2008Jul/cache-control-attributes/
* http://www.mnot.net/blog/2007/08/07/etags
