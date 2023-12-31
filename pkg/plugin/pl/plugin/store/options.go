package store

// GetOpts - iterate the inbound Options and return a struct
func GetOpts(opt ...Option) options {
	opts := getDefaultOptions()
	for _, o := range opt {
		o(&opts)
	}
	return opts
}

// Option - how Options are passed as arguments.
type Option func(*options)

// options = how options are represented
type options struct {
	withName        string
	withDescription string
	withPublicID    string
}

func getDefaultOptions() options {
	return options{
		withDescription: "",
		withName:        "",
	}
}

// WithDescription provides an optional description.
func WithDescription(desc string) Option {
	return func(o *options) {
		o.withDescription = desc
	}
}

// WithName provides an optional name.
func WithName(name string) Option {
	return func(o *options) {
		o.withName = name
	}
}

// WithPublicID provides an optional specific public ID
func WithPublicID(with string) Option {
	return func(o *options) {
		o.withPublicID = with
	}
}
