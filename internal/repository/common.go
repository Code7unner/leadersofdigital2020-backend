package repository

type (
	// Printf is a (fmt.Printf|log.Printf|testing.T.Logf)-like function.
	Printf func(format string, args ...interface{})
)
