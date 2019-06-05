package timer

import (
	"context"
	"log"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type Schedule struct {
	//时间精度
	accuracy  int64
	ticker    *time.Ticker
	cancel    context.CancelFunc
	tasklist  map[int64][]string
	lock      *sync.RWMutex
	taskqueue taskBucket
	//uuid对应一个处理函数
	taskFuncs map[string]func()
}

type taskFunc struct {
	handler func()
	id      string
}

func NewSchedule(duration time.Duration) *Schedule {
	ctx, cancel := context.WithCancel(context.Background())
	s := &Schedule{
		ticker:    time.NewTicker(duration),
		cancel:    cancel,
		tasklist:  make(map[int64][]string),
		taskqueue: newTaskHeap(),
		lock:      &sync.RWMutex{},
		accuracy:  int64(duration) / int64(time.Second),
		taskFuncs: make(map[string]func()),
	}
	go s.run(ctx)
	return s
}

func (s *Schedule) Close() {
	s.cancel()
	//销毁所有已经分配的空间
	s.ticker.Stop()
	s.tasklist = nil
	s.taskqueue = nil
	s.taskFuncs = nil
}

func (s *Schedule) run(ctx context.Context) {
	for {
		select {
		case <-s.ticker.C:
			//找到执行时间在一分钟以内的任务，设置timer执行
			logrus.Infoln("ticker ticking, now:", time.Now().Unix())
			logrus.Infoln("tash queue:", s.taskqueue)
			for {
				ok, stamp := s.taskqueue.Peek()
				if ok && (stamp-time.Now().Unix() <= s.accuracy) {
					stamp = s.taskqueue.PopTask()
					logrus.Infoln("发现任务 stamp:", stamp, s.tasklist[stamp])
					//拉起这段时间内的所有任务.
					s.lock.RLock()
					for _, taskId := range s.tasklist[stamp] {
						logrus.Infoln("task id is:", taskId)
						handler, ok := s.taskFuncs[taskId]
						if !ok {
							//说明此任务被删除,跳过这个任务id
							continue
						}
						startHandler(ctx, time.Unix(stamp, 0).Sub(time.Now()), handler)
						//既然所有任务都被拉起来了，删除这些任务
						delete(s.tasklist, stamp)
						//任务id被拉起来了，删除这个任务
						delete(s.taskFuncs, taskId)
					}
					s.lock.RUnlock()
				} else {
					break
				}
			}
		case <-ctx.Done():
			logrus.Infoln("scheduler 退出")
		}
	}
}

func startHandler(ctx context.Context, d time.Duration, handler func()) {
	timer := time.NewTimer(d)
	go func() {
		select {
		case <-timer.C:
			handler()
			return
		case <-ctx.Done():
			logrus.Infoln("handler exit.")
			return
		}
	}()
}

//添加到相应队列中
func (s *Schedule) AddScheduleTask(execAt time.Time, taskHandler func()) (id string) {
	logrus.Infoln("添加定时任务:", execAt)
	stamp := execAt.Unix()
	log.Println("stamp is:", stamp)
	uid, _ := uuid.NewV4()

	s.lock.Lock()
	//可以减少内存分配
	//任务的执行时间内有没有其他待执行的任务
	taskids, ok := s.tasklist[stamp]
	if ok {
		//当前时间已经有任务了，直接放在桶中
		taskids = append(taskids, uid.String())
	} else {
		//新建一个桶，需要对桶进行排序
		s.taskqueue.Insert(stamp)
		logrus.Infoln("task queue:", s.taskqueue)
		taskids = make([]string, 0)
		taskids = append(taskids, uid.String())
		s.tasklist[stamp] = taskids
	}
	s.taskFuncs[uid.String()] = taskHandler
	s.lock.Unlock()
	return uid.String()
}

//删除一个任务
func (s *Schedule) CancelTask(id string) error {
	delete(s.taskFuncs, id)
	return nil
}
