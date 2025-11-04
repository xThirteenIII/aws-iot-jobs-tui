package main

import (
	"aws-iot-jobs-tui/config"
	"aws-iot-jobs-tui/helper"
	"aws-iot-jobs-tui/logger"
)

func main() {
	helper.Must(config.Setup())
	helper.Must(logger.Setup())
}
