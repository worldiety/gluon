package gluon

func DeferClose(src error, dst *error) {
	if src != nil && *dst != nil {
		*dst = src
	}
}
