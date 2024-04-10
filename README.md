# sending slog to VECTOR


https://vector.dev/


In the example, 
- Vector is acting as the server listeing to socket
- https://vector.dev/docs/reference/configuration/sources/http_server/
- https://vector.dev/docs/reference/configuration/sources/socket/

- running localhost:



 - 1 install vector
curl --proto '=https' --tlsv1.2 -sSfL https://sh.vector.dev | bash



- 2 confirm vector is installed
vector --version



- 3 run Vector with config file


//vector --config /etc/vector/vector.yaml
vector --config /vconfig/vector02.yaml





## Current Error
when I run this, the socket returns.

```
ERROR source{component_kind="source" component_id=my_source_id component_type=socket}: vector::internal_events::socket: Error receiving data. error=connection refused error_code="socket_receive" error_type="reader_failed" stage="receiving" mode=tcp internal_log_rate_limit=true
```




main2
Not sure how to pass sLog to a handler,  but my first guess is this.. and it works.
passing net.Conn to slog.NewJSONHandler as io.Writer
But I still prefer the original way I did it.
