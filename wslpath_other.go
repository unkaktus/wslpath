// +build !linux

package wslpath

func FromWindows(p string) (string, error) {
	return p, nil
}

func ToWindows(p string) (string, error) {
	return p, nil
}
