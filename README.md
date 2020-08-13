## GO PORT SCANNER

scan tcp/udp network ports to check status [open/unavailable/in-use]

> I wrote this to mainly benchmark Go's concurrency.

#### Benchmark test :watch:

```txt
--> scanning 65000 ports in a single process
Duration: 3.6s <avg>

--> scanning with multiple processes [100 goroutines ==> 1 goroutine/650 ports]
Duration: 580ms <avg>
```

Usage :memo:
```shell
$ go run main.go --protocol tcp // scan all ports

$ go run main.go --port 5432 --protocol udp // scan single port
```