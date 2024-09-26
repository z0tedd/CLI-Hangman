package config

type Options struct {
	MaxAttempts int
}

func GetOptions(difficulty int) *Options {
	switch difficulty {
	case 1:
		return &Options{MaxAttempts: 12}
	case 2:
		return &Options{MaxAttempts: 9}
	case 3:
		return &Options{MaxAttempts: 7}
	default:
		return &Options{MaxAttempts: 7}
	}
}
