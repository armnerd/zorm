package zorm

type Options struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

type OptionFunc func(opts *Options)

func loadOptions(options ...OptionFunc) *Options {
	opts := new(Options)
	for _, option := range options {
		option(opts)
	}
	return opts
}

func WithOptions(options Options) OptionFunc {
	return func(opts *Options) {
		*opts = options
	}
}

func WithHost(Host string) OptionFunc {
	return func(opts *Options) {
		opts.Host = Host
	}
}

func WithPort(Port string) OptionFunc {
	return func(opts *Options) {
		opts.Port = Port
	}
}

func WithUser(User string) OptionFunc {
	return func(opts *Options) {
		opts.User = User
	}
}

func WithPass(Pass string) OptionFunc {
	return func(opts *Options) {
		opts.Pass = Pass
	}
}

func WithDatabase(Database string) OptionFunc {
	return func(opts *Options) {
		opts.Database = Database
	}
}
