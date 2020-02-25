package taillog

import (
	"github.com/hpcloud/tail"
)

var (
	Tails *tail.Tail
)

func InitTaillog(logfile string) (err error) {
	conf := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}

	Tails, err = tail.TailFile(logfile, conf)
	return
}
