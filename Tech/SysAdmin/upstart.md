http://upstart.ubuntu.com/cookbook/

_It handles asynchronicity by emitting events._
_Upstart emits "events" which services can register an interest in. When an event -- or combination of events -- is emitted that satisfies some service's requirements, Upstart will automatically start or stop that service. If multiple jobs have the same "start on" condition, Upstart will start those jobs ''in parallel''_

`/etc/init/my_service.conf`

```
#!upstart
description "My Service"

start on runlevel [2345]
stop on runlevel [016]

script
  RBENV_VERSION=jruby-1.7.6
  export RBENV_VERSION
  cd /data/my_service/current
  exec su -c '/usr/local/rbenv/shims/bundle exec puma -C /etc/user_service.rb >> /data/my_service/shared/log/puma.log 2>&1' <%= @user %>
end script

post-stop script
    kill -2 `cat /data/my_service/shared/pids/puma.pid`
end script
```

__Commands:__

`/sbin/init`

`/sbin/start job_name`

`/sbin/stop job_name`

`/sbin/status job_name`

---------

`/etc/init.d` System V init tools (SysVinit).
It contains the `init` program (the first process that is run when the kernel has finished initializing¹) as well as some infrastructure to start and stop services and configure them

Files in the directory are shell scripts that respond to start, stop, restart.
