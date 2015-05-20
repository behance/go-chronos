package chronos

type Config struct {
	/* the url for chronos */
	URL string
	/* switch on debugging */
	Debug bool
	/* the timeout for requests */
	RequestTimeout int
}

func NewDefaultConfig() Config {
	return Config{
		URL:            "http://127.0.0.1:4400",
		Debug:          false,
		RequestTimeout: 5}
}
