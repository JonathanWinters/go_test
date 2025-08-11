module github.com/JonathanWinters/go_test

go 1.24.5

replace github.com/luckypointlabs/gsu => ../gsu

require (
	github.com/go-redis/redis/v9 v9.0.0-rc.2
	github.com/google/uuid v1.6.0
	github.com/lib/pq v1.10.9
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
)
