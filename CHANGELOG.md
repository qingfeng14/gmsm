# 更新日志

## 1.2 更新(Feb 20, 2019)

- [NEW] 实现PKCS#7签名及验签
- [FIX] SM2 签名及验签方法完全遵循标准GM/T 0003系列，兼容CFCA Java-SDK
- **破坏性更新，证书与之前版本不兼容**

## 1.1.1更新

- 新增以下函数支持用户其他信息
  - SignDigitToSignData 将签名所得的大数r和s转换为签名的格式
  - Sm2Sign     支持用户信息的签名
  - Sm2Verify   支持用户信息的验签

## 1.1.0更新

- 改进新能，具体提升如下

> 注:本次优化并不彻底，只是第一次尝试优化，后续有时间还会继续优化

```sh
    old:
        generate key:
            BenchmarkSM2-4          1000   2517147 ns/op 1156476 B/op   11273 allocs/op
        sign:
            BenchmarkSM2-4           300   6297498 ns/op 2321890 B/op   22653 allocs/op
        verify:
            BenchmarkSM2-4          2000   8557215 ns/op 3550626 B/op   34627 allocs/op
        encrypt:
            BenchmarkSM2-4          2000   8304840 ns/op 3483113 B/op   33967 allocs/op
        decrypt:
            BenchmarkSM2-4          2000   5726181 ns/op 2321728 B/op   22644 allocs/op
    new:
        generate key:
            BenchmarkSM2-4          5000    303656 ns/op    2791 B/op      41 allocs/op
        sign:
            BenchmarkSM2-4          2000    652465 ns/op    8828 B/op     133 allocs/op
        verify:
            BenchmarkSM2-4          1000   2004511 ns/op  122709 B/op    1738 allocs/op
        encrpyt:
            BenchmarkSM2-4          1000   1984419 ns/op  118560 B/op    1687 allocs/op
        decrypt:
            BenchmarkSM2-4          1000   1725001 ns/op  118331 B/op    1679 allocs/op
```

### 1.0.1 更新

- 添加全局的sbox改进sm4效率(by https://github.com/QwertyJack)

### 1.0 更新

- 添加以下oid
  - SM3WithSM2 1.2.156.10197.1.501
  - SHA1WithSM2 1.2.156.10197.1.502
  - SHA256WithSM2 1.2.156.10197.1.503

- x509生成的证书如今可以使用SM3作为hash算法

- 引入了以下hash算法
  - RIPEMD160
  - SHA3_256
  - SHA3_384
  - SHA3_512
  - SHA3_SM3

用户需要自己安装`golang.org/x/crypto`
