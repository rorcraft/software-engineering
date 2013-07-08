* JRuby 1.7.4
* Oracle Java JDK 1.7

### Environment Options

```
JRUBY_OPTS=-Xcompile.invokedynamic=true

```

### ruby-debug
https://kenai.com/projects/jruby/pages/UsingTheJRubyDebugger
```
require 'rubygems'
require 'ruby-debug'
debugger
=> ruby --debug bundle exec rspec
* make sure to pass --debug flag, or "next" will always act as if it were the "step" command.
```

### Speed up startup time

* Drip - https://github.com/flatland/drip/wiki/JRuby
