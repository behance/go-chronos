package chronos

// A Config defines a client configuration
type Config struct {
	/* the url for chronos */
	URL string
	/* switch on debugging */
	Debug bool
	/* the timeout for requests */
	RequestTimeout int
	/* string to prefix api paths with */
	APIPrefix string
}

// NewDefaultConfig returns a default configuration.
// Helpful for local testing/development.
func NewDefaultConfig() Config {
	return Config{
		URL:            "http://127.0.0.1:4400",
		Debug:          false,
		RequestTimeout: 5,
		APIPrefix:      "",
	}
}
