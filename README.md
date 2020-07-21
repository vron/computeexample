# computeexample
Example of using vron/compute to run openGL compute kernels on CPU in go. See source code..

Note that this example is **linux only** for the time being (since it uses docker for building,
if you build locally win and osx is also supported.)

Bench results for this shader:

    goos: windows
    goarch: amd64
    BenchmarkGo128x128-24                        735           1600027 ns/op
    BenchmarkShader128x128-24                    714           1684173 ns/op
    BenchmarkGo2048x2048-24                        3         411007067 ns/op
    BenchmarkShader2048x2048-24                    3         429666800 ns/op
    BenchmarkShader2048x2048Parallel-24            9         122333233 ns/op
