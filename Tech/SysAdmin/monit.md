http://mmonit.com/monit/documentation/monit.html

`/etc/monit/monitrc`
`ls /etc/monit/conf.d/`

```
check process my_service
  with pidfile /data/my_service/shared/pids/puma.pid
  start program = "/sbin/start my_service" with timeout 150 seconds
  stop program = "/sbin/stop my_service"
```
