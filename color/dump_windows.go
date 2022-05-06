//go:build windows
// +build windows

package color

func Text(txt string, params ...StyleOrColor) string {
	return txt
}

func Info(txt string) string {
	return txt
}

func Warn(txt string) string {
	return txt
}

func Error(txt string) string {
	return txt
}

func Success(txt string) string {
	return txt
}

func Failure(txt string) string {
	return txt
}
