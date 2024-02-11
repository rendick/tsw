package cmd

import (
	"log"
	"os"
)

func LogsWriter() {
	_, err := os.Stat("/var/log/tsw/")
	if os.IsNotExist(err) {
		MkdirLog := os.Mkdir("/var/log/tsw/", 0750)
		if MkdirLog != nil && !os.IsExist(MkdirLog) {
			log.Fatal(MkdirLog)
		}
	}

	OpenLog, err := os.OpenFile("/var/log/tsw/tsw.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer OpenLog.Close()

	_, WriteLog := OpenLog.WriteString(Log + "\n\n")
	if WriteLog != nil {
		log.Fatal(err)
	}
}

// sudo chmod 777 /var/log/pem
