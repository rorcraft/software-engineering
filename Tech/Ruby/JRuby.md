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

# Speed up JRuby + Rails startup time

check - https://github.com/jruby/jruby/wiki/Improving-startup-time

`~/.rvm/hooks/after_use_jruby_opts`
```
#!/usr/bin/env bash

. "${rvm_path}/scripts/functions/hooks/jruby"

export PROJECT_JRUBY_OPTS="-Xcompile.invokedynamic=false -J-XX:+TieredCompilation -J-XX:TieredStopAtLevel=1 -J-noverify -Xcompile.mode=OFF"

if [[ "${rvm_ruby_string}" =~ "jruby" ]]
then
  jruby_options_append "${PROJECT_JRUBY_OPTS[@]}"
else
  jruby_options_remove "${PROJECT_JRUBY_OPTS[@]}"
  jruby_clean_project_options
fi
```

## One by one

`-Xcompile.invokedynamic=false` optimizations on bytecode instruction level dynamically.

`-J-XX:+TieredCompilation -J-XX:TieredStopAtLevel=1` same effect as client mode on 64bit JVM.

`-J-noverify` not verifying

`-Xcompile.mode=OFF` don't do JIT

## Install

`thor install:jruby_opts`

## Nailgun

> "Insanely Fast Java"

Preloads a JVM.

Not that much gain < 1 second

```
cd ~/.rvm/rubies/jruby-1.7.4/tool/nailgun
./configure
make install
```
```
$ time ruby bin/thor list
real  0m1.919s

$ ruby --ng-server &
NGServer started on all interfaces, port 2113.

$ time ruby --ng bin/thor list
real  0m1.331s
```

## Drip
Charles Nutter @headius Nov 20 2012
"Big JRuby startup time improvements using Drip…give it a shot, let us know how it works: https://github.com/flatland/drip"

```
  export JAVACMD=`which drip`
  export DRIP_INIT_CLASS=org.jruby.main.DripMain
  
  dripmain.rb
  # org.jruby.main.DripMain
  public static final String JRUBY_DRIP_PREBOOT_FILE = "./dripmain.rb";
```

Unfortunately buggy - https://github.com/flatland/drip/issues/51


## All together

0. BEFORE LAZY REQUIRE (Bundler.require)

```
$ time ruby bin/rspec  spec/models/tag_spec.rb

saved about ~3 secs.
```


1. DEFAULT

```
$ time ruby bin/rspec  spec/models/tag_spec.rb

real  0m23.325s
```

2. JRUBY_OPTS

```
$ export JRUBY_OPTS="-Xcompile.invokedynamic=false -J-XX:+TieredCompilation -J-XX:TieredStopAtLevel=1 -J-noverify -Xcompile.mode=OFF"
$ time ruby bin/rspec  spec/models/tag_spec.rb

real  0m10.726s
```

3. JRUBY_OPTS + NAILGUN + SPORK

```
$ ruby --ng-server &
$ ruby --ng bin/spork
$ time ruby --ng bin/rspec  spec/models/tag_spec.rb

real  0m3.453s
```

----
### Auto enable Nailgun

`~/.rvm/hooks/after_use_jruby`

```
#!/usr/bin/env bash

\. "${rvm_path}/scripts/functions/hooks/jruby"

if [[ "${rvm_ruby_string}" =~ "jruby" ]]
then
  jruby_ngserver_start
  jruby_options_append "--ng" "${PROJECT_JRUBY_OPTS[@]}"
else
  jruby_options_remove "--ng" "${PROJECT_JRUBY_OPTS[@]}"
  jruby_clean_project_options
fi
```

### Prefer binstub over bundle exec
```
bundle binstub rspec-core
```
`bundle exec` requires the executable script and call `exec` to replace the process. 
From the wiki "Avoid spawning sub-rubies".

# JRuby Internals

* http://marcelozambranav.blogspot.com/2011/01/jruby-hacking-guide.html
* http://realjenius.com/2009/09/16/distilling-jruby-method-dispatching-101/
* https://kenai.com/projects/jruby/pages/Internals
* http://www.infoq.com/news/2009/11/jruby-ir
* http://blog.headius.com/2010/05/kicking-jruby-performance-up-notch.html
* http://prezi.com/ebpcrtgmwgln/jruby-source-code-reading-guide/
* https://docs.google.com/spreadsheet/ccc?key=0AiZsKd8d4hSJdHVFYk1jcFVwY29sUWxpdFZyajVqdGc&hl=en#gid=0
* MRI Ruby hacking guide - http://rhg.rubyforge.org/
