Critical rendering path - https://www.udacity.com/course/ud884

### Compression

* http://www.html5rocks.com/en/tutorials/speed/img-compression/
* http://www.html5rocks.com/en/tutorials/speed/txt-compression/

### CSS

* https://github.com/giakki/uncss remove unused css.

### Javascript

* https://addyosmani.github.io/basket.js/ - cache js or css on browser
* http://alistapart.com/article/application-cache-is-a-douchebag - HTML5 application cache, html manifest
* https://www.igvita.com/2014/05/20/script-injected-async-scripts-considered-harmful/

<table style="margin-top:1em; text-align:left">
<tbody><tr>
<th width="50%"><code>&lt;script src="..."&gt;</code></th>
<th><code>&lt;script async src="..."&gt;</code></th>
</tr>
<tr>
<td class="red">Blocks DOM construction</td>
<td class="green">Does not block DOM construction</td>
</tr>
<tr>
<td class="red">Execution is blocked on CSSOM</td>
<td class="green">Execution is not blocked on CSSOM</td>
</tr>
<tr>
<td class="green">Preload scanner discoverable</td>
<td class="green">Preload scanner discoverable</td>
</tr>
<tr>
<td class="green">Ordered execution of scripts</td>
<td class="red">Unordered execution</td>
</tr>
<tr style="font-style:italic">
<td>Use where execution order matters, place these scripts at the bottom.</td>
<td>Can be placed anywhere, ideal for scripts that can tolerate out-of-order execution.</td>
</tr>
</tbody></table>

* https://developers.google.com/speed/docs/insights/mobile
* http://www.stevesouders.com/blog/2010/12/15/controljs-part-1/

### Asset Pipeline

* https://github.com/assetgraph/assetgraph
