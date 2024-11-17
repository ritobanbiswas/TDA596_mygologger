package mylogger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func MyLogger(logdirname string, logfilename string, logprefix string) (*log.Logger, *os.File, error) {
	/* Log file as in cloud env will never see the stdout*/
	flog, err := os.OpenFile(filepath.Join(logdirname, logfilename),
		os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf(`error opening Log file: %v`, err)
	}
	logPrefix := fmt.Sprintf("[Server: %s] ", logprefix)
	logger := log.New(flog, logPrefix, log.LstdFlags)
	return logger, flog, nil
}
