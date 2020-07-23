# computeexample
Example of using vron/compute to run openGL compute kernels on CPU in go. See source code..

Mostly an example of how to use it, not a particularly good or interesting example.

Note that this example is **linux only** for the time being (since it uses docker for building,
if you build locally win and osx is also supported.)

Bench results for this shader (see the source, not fully representative):

    goos: darwin
    goarch: amd64
    BenchmarkGo128x128-4                         572           2183993 ns/op
    BenchmarkShader128x128-4                     534           2210871 ns/op
    BenchmarkGo2048x2048-4                         2         545968129 ns/op
    BenchmarkShader2048x2048-4                     2         557022096 ns/op
    BenchmarkShader2048x2048Parallel-4             6         186456526 ns/op
    PASS