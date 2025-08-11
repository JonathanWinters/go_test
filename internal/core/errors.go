package core

type Error int

const (
	RECTANGULAR Error = iota
	DIMENSIONS
	VALUES
)

var errorName = map[Error]string{
	RECTANGULAR: "The map is not rectangular.",
	DIMENSIONS:  "The dimensions are larger than 100 in one or both dimensions",
	VALUES:      "The values do not fit the allowable values of 0-4",
}

func (err Error) String() string {
	return errorName[err]
}

type ObfuscatedError struct {
	opaque string
}

func GetObfuscatedError(err Error) ObfuscatedError {

	return ObfuscatedError{
		opaque: errorName[err],
	}
}
