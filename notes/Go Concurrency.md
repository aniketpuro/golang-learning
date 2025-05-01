#learnMore
Go Concurrency is made available by what's called `go-routines` , basically when a function is preceded with the `go` keyword, it runs in a `go-routine`, think of go-routine as a thread (_**though they're different…**)__._** go-routines is one of the most important features of Go that makes it and its concurrency model special.

For data synchronization you can use mutex Locks, WaitGroups, and Atomic operations, however.. **It's recommended** to use Go Channels for data synchronization, though using the sync package (using mutex, locks, atomics, and WaitGroups) is also usable if it make sense for your use case.
