package panicWorker

import "github.com/sirupsen/logrus"

func PanicWorker(job func()) {
	defer func() {
		if err := recover(); err != nil {
			//log.Println("OOOOOPA PANIC recovered", err)
			logrus.WithField("err", err).Error("OOOOOPA PANIC in game, recovered")
		}
	}()

	job()
}
