# sm2fiat

[![ci](https://github.com/emmansun/sm2fiat/actions/workflows/ci.yml/badge.svg)](https://github.com/emmansun/sm2fiat/actions/workflows/ci.yml)

POC generate sm2 p256 curve with fiat-crypto &amp; addchain


**SM2 elliptic curve pure golang implementation Benchmark**

    goos: windows
    goarch: amd64
    pkg: github.com/emmansun/sm2fiat
    cpu: Intel(R) Core(TM) i5-9500 CPU @ 3.00GHz
    BenchmarkScalarBaseMult/SM2P256-6         	    3514	    346616 ns/op	    8408 B/op	     276 allocs/op
    BenchmarkScalarBaseMult/OldSM2P256-6      	    1150	   1045606 ns/op	     619 B/op	      13 allocs/op

    goos: windows
    goarch: amd64
    pkg: github.com/emmansun/sm2fiat
    cpu: Intel(R) Core(TM) i5-9500 CPU @ 3.00GHz
    BenchmarkScalarMult/SM2P256-6         	    1015	   1171729 ns/op	   10352 B/op	     339 allocs/op
    BenchmarkScalarMult/OldSM2P256-6      	     446	   2659460 ns/op	    1005 B/op	      17 allocs/op

    goos: windows
    goarch: amd64
    pkg: github.com/emmansun/sm2fiat
    cpu: Intel(R) Core(TM) i5-9500 CPU @ 3.00GHz
    BenchmarkMarshalUnmarshal/SM2P256/Uncompressed-6         	  317323	      3648 ns/op	     536 B/op	      12 allocs/op
    BenchmarkMarshalUnmarshal/SM2P256/Compressed-6           	   13086	     91704 ns/op	    3961 B/op	     137 allocs/op
    BenchmarkMarshalUnmarshal/OldSM2P256/Uncompressed-6      	  372056	      3265 ns/op	     896 B/op	      12 allocs/op
    BenchmarkMarshalUnmarshal/OldSM2P256/Compressed-6        	   13580	     89549 ns/op	    4289 B/op	     133 allocs/op
