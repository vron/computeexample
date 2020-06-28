# computeexample
Example of using vron/compute to run openCL kernels on CPU in go. See source code..

Note that this example is **linux only** for the time being.

Bench results for this shader:

    goos: linux
    goarch: amd64
    BenchmarkGo128x128-4                 696           1679334 ns/op
    BenchmarkShader128x128-4             696           1717018 ns/op
    BenchmarkGo2048x2048-4                 3         428402654 ns/op
    BenchmarkShader2048x2048-4             3         435269681 ns/op
