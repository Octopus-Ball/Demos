// 定时任务调度器

package main

import (
	"container/list"
	"fmt"
	"time"

	"github.com/gorhill/cronexpr"
)

// NewCronDispatch 实例化CronDispatch的工厂方法
func NewCronDispatch() (cd *CronDispatch) {
	cd = new(CronDispatch)
	cd.ReadyCronJobCh = make(chan CronJob, 50)
	cd.DispatchRate = time.Second
	return
}

// CronJob 任务
type CronJob struct {
	jobName    string
	cronString string
	expr       *cronexpr.Expression
}

// CronDispatch 定时任务调度器
type CronDispatch struct {
	// 用来传输就绪的任务给消费者
	ReadyCronJobCh chan CronJob
	// 调度检查频率
	DispatchRate time.Duration
}

// initCronJon 实例化一个定时任务
func (c *CronDispatch) initCronJon(cronSring string, jobName string) (cronjob *CronJob, err error) {
	cronjob = new(CronJob)
	cronjob.jobName = jobName
	cronjob.cronString = cronSring
	cronjob.expr, err = cronexpr.Parse(cronSring)
	return
}


// AddCronJob 注册一个定时任务
func (c *CronDispatch) AddCronJob(cronSring string, jobName string) {
	if cronjob, err := c.initCronJon(cronSring, jobName); err != nil {
		fmt.Printf("%v\n", err)
	} else {
		c.JobList.PushBack(cronjob)
	}

}

// Dispatch 循环检查就绪任务
func (c *CronDispatch) Dispatch() {
	for {
		now := time.Now()                   // 当前时间
		nextTime := now.Add(c.DispatchRate) // 下次调度的时间
		// 遍历任务列表
		for node := c.JobList.Front(); node != nil; node = node.Next() {
			cronjob := node.Value.(*CronJob)
			cronjobNext := cronjob.expr.Next(now) // 该任务下次改执行的时间
			// 判断任务是否会在下次调度前到期
			if cronjobNext.Before(nextTime) || cronjobNext.Equal(nextTime) {
				go c.Worker(cronjob)
			}
		}
		time.Sleep(c.DispatchRate)
	}
}

// Worker 任务执行者
func (c *CronDispatch) Worker(cronjob *CronJob) {
	now := time.Now()
	runTime := cronjob.expr.Next(now) // 该任务下次执行的时间
	time.AfterFunc(runTime.Sub(now), func() {
		fmt.Println(cronjob.jobName, now)
	})
}

//00:26:45

func main() {
	cronDispatch := NewCronDispatch()
	cronDispatch.AddCronJob("*/5 * * * * * *", "5_job")
	cronDispatch.AddCronJob("*/1 * * * * * *", "1_job")
	cronDispatch.Dispatch()

	time.Sleep(time.Second * 20)
}
