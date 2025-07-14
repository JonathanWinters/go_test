package types

type HttpMethod int

const (
	PUT HttpMethod = iota
	GET
	POST
)

var requestName = map[HttpMethod]string{
	PUT:  "PUT",
	GET:  "GET",
	POST: "POST",
}

func (method HttpMethod) String() string {
	return requestName[method]
}
