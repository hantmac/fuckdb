package cmd

import (
	"os/exec"

	"github.com/sirupsen/logrus"
)

func openBrowser(url string) {
	cmd := exec.Command("cmd", "/C", "start", url)

	if err := cmd.Run(); err != nil {
		logrus.Errorln("open browser error, please open browser manually:", url)
	}
}
