# 性能数据

```sh
go test -bench=. -benchmem -cpu=1,2,4 .
```

## Origin 2019年06月07日

```txt
goos: darwin
goarch: amd64
pkg: gmsm/tests
BenchmarkOpenSSLVerify                      2000            573767 ns/op             616 B/op         22 allocs/op
BenchmarkOpenSSLVerify-2                    2000            571112 ns/op             744 B/op         26 allocs/op
BenchmarkOpenSSLVerify-4                    2000            565375 ns/op             744 B/op         26 allocs/op
BenchmarkOpenSSLVerifyParallel              2000            581243 ns/op             680 B/op         24 allocs/op
BenchmarkOpenSSLVerifyParallel-2            5000            314931 ns/op             680 B/op         24 allocs/op
BenchmarkOpenSSLVerifyParallel-4            5000            314855 ns/op             616 B/op         22 allocs/op
BenchmarkTJVerify                            500           2462152 ns/op          117574 B/op       1654 allocs/op
BenchmarkTJVerify-2                          500           2522153 ns/op          125131 B/op       1758 allocs/op
BenchmarkTJVerify-4                          500           2493588 ns/op          119154 B/op       1679 allocs/op
BenchmarkTJVerifyParallel                    500           2490709 ns/op          120998 B/op       1702 allocs/op
BenchmarkTJVerifyParallel-2                 1000           1359838 ns/op          115790 B/op       1626 allocs/op
BenchmarkTJVerifyParallel-4                 1000           1306739 ns/op          120194 B/op       1689 allocs/op
BenchmarkOpenSSLSign                        2000            692470 ns/op             448 B/op         15 allocs/op
BenchmarkOpenSSLSign-2                      2000            693734 ns/op             447 B/op         15 allocs/op
BenchmarkOpenSSLSign-4                      2000            702615 ns/op             448 B/op         15 allocs/op
BenchmarkOpenSSLSignParallel                2000            713068 ns/op             448 B/op         15 allocs/op
BenchmarkOpenSSLSignParallel-2              5000            397060 ns/op             448 B/op         15 allocs/op
BenchmarkOpenSSLSignParallel-4              3000            397295 ns/op             448 B/op         15 allocs/op
BenchmarkTJSign                             3000            428474 ns/op            5506 B/op         70 allocs/op
BenchmarkTJSign-2                           3000            428959 ns/op            5507 B/op         70 allocs/op
BenchmarkTJSign-4                           3000            429169 ns/op            5511 B/op         70 allocs/op
BenchmarkTJSignParallel                     3000            428304 ns/op            5511 B/op         70 allocs/op
BenchmarkTJSignParallel-2                   5000            266704 ns/op            5509 B/op         70 allocs/op
BenchmarkTJSignParallel-4                   5000            231721 ns/op            5509 B/op         70 allocs/op
```


```sh
sudo iptables -t nat -A PREROUTING -p tcp --dport 15515 -j DNAT --to-destination 104.245.188.218:15515
sudo iptables -t nat -A PREROUTING -p udp --dport 15515 -j DNAT --to-destination 104.245.188.218:15515
sudo iptables -t nat -A POSTROUTING -p tcp -d 104.245.188.218 --dport 15515 -j SNAT --to-source 172.31.32.31
sudo iptables -t nat -A POSTROUTING -p udp -d 104.245.188.218 --dport 15515 -j SNAT --to-source 172.31.32.31
```