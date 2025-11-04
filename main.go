package main

import (
	"aws-iot-jobs-tui/config"
	"aws-iot-jobs-tui/helper"
)

func main() {
	helper.Must(config.Setup())
}
