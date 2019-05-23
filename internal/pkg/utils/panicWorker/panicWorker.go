package panicWorker

import "github.com/sirupsen/logrus"

func PanicWorker(job func()) {
	defer func() {
		if err := recover(); err != nil {
			logrus.WithField("err", err).Error("OOOOOPA PANIC in worker, recovered")
		}
	}()

	job()
}
