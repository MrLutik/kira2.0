package maintenance

import "errors"

var (
	ErrNotSelectFlag      = errors.New("not select flag")
	ErrOnlyOneFlagAllowed = errors.New("only one flag at a time is allowed")
	ErrGettingFlagError   = errors.New("error while getting flag")
)
