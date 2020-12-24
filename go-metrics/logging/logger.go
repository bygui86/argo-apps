package logging

import "github.com/sirupsen/logrus"

// Log exposed
var Log = logrus.New()

// init - init logger wrapper automatically on package import
func init() {
	if Log == nil {
		_ = logrus.New()
	}
}
