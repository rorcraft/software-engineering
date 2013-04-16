A reverse proxy setup (e.g. nginx forwarding HTTP requests to Starman) has the following advantages:

* things are a bit easier to debug, since you can easily hit directly the backend server;

* if you need to scale your backend server, you can easily use something like pound/haproxy between the frontend (static-serving) HTTP and your backends (Zope is often deployed like that);

* it can be a nice sidekick if you are also using some kind of outward-facing, caching, reverse proxy (like Varnish or Squid) since it allows to bypass it very easily.

However, it has the following downsides:

* the backend server has to figure out the real originating IP, since all it will see is the frontend server address (generally localhost); there is almost an easy way to find out the client IP address in the HTTP headers, but that's something extra to figure out;

* the backend server does not generally know the orignal "Host:" HTTP header, and therefore, cannot automatically generated an absolute URL to a local resource; Zope addresses this with special URLs to embed the original protocol, host and port in the request to the backend, but it's something you don't have to do with FastCGI/Plack/...;

* the frontend cannot automatically spawn backend processes, like it could do with FastCGI for instance.

## FastCGI - the forgotten treasure
http://www.nongnu.org/fastcgi/#AEN147
