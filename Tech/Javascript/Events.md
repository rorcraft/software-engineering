DOM level 2 Event model

* `addEventListener`, `removeEventListener`
* if IE<=8 `attachEvent`, `detachEvent`

Events propagate from event target up to the root of the DOM during the _bubble phase_ and then they traverse down the tree back to the target during the _capture phase_.

IE doesn't pass context the function.
```javascript
this.addEvent = function(elem, type, fn) {
  var bound = function() {
    return fn.apply(elem, arguments);
  };
  elem.attachEvent("on" + type, bound);
  return bound;
};
```

Event Model
* `event.target`, IE `event.srcElement`
* `relatedTarget` - `mouseover`, `mouseout`, IE `toElement`, `fromElement`
* `preventDefault`, IE `returnValue = false` - prevents default browser action.
* `stopPropagation`, IE `cancelBubble = true` - stop events from bubbling up.
* `pageX`, `pageY` - IE `clientX/Y` + `scrollTop/Left` + `clientTop/Left` - position of mouse relative to the whole document.
* `which`, IE `charCode` or `keyCode` - keyboard event.
* `button` (0,1,2) - IE ` (1,2,4) - mouse button clicked

