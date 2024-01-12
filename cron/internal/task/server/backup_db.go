package server

import (
	"context"
	"sync"

	"github.com/EchoGroot/kratos-examples/pkg/lang/syncx"

	"github.com/EchoGroot/kratos-examples/cron/internal/task/biz"
	"github.com/EchoGroot/kratos-examples/cron/internal/task/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
)

type BackupDbServer struct {
	crontab         *cron.Cron
	cancelFunc      context.CancelFunc
	wg              sync.WaitGroup
	backupDbACron   string
	backupDbBCron   string
	backupDbUsecase *biz.BackupDbUsecase
}

func NewBackupDbServer(
	c *conf.Bootstrap,
	backupDbUsecase *biz.BackupDbUsecase,
) *BackupDbServer {
	return &BackupDbServer{
		crontab:         cron.New(cron.WithSeconds()),
		wg:              sync.WaitGroup{},
		backupDbACron:   c.App.BackupDbACron,
		backupDbBCron:   c.App.BackupDbBCron,
		backupDbUsecase: backupDbUsecase,
	}
}

func (r *BackupDbServer) Start(ctx context.Context) error {
	log.Info("start crontab")

	withCancelCtx, cancelFunc := context.WithCancel(context.Background())
	r.cancelFunc = cancelFunc

	// 添加定时任务backupDbA
	_, err := r.crontab.AddFunc(r.backupDbACron, func() {
		r.wg.Add(1)
		defer func() {
			r.wg.Done()
		}()

		if err := r.backupDbUsecase.BackupDbA(withCancelCtx); err != nil {
			log.Errorf("backup db a error: %+v", err)
			return
		}
	})
	if err != nil {
		return errors.Wrap(err, "add backup db a task to crontab error")
	}

	// 添加定时任务backupDbB
	_, err = r.crontab.AddFunc(r.backupDbBCron, func() {
		r.wg.Add(1)
		defer func() {
			r.wg.Done()
		}()

		// 避免任务执行时间超过 cron 调度时间间隔，导致并发执行。
		if err := r.backupDbUsecase.BackupDbBWithLock(withCancelCtx); err != nil {
			log.Errorf("backup db b error: %+v", err)
			return
		}
	})
	if err != nil {
		return errors.Wrap(err, "add backup db b task to crontab error")
	}

	r.crontab.Start()
	return nil
}

func (r *BackupDbServer) Stop(ctx context.Context) error {
	log.Info("stop crontab")
	r.cancelFunc()

	wgCtx, wgCancelFunc := context.WithCancel(context.Background())
	syncx.Go(func() {
		r.wg.Wait()
		<-r.crontab.Stop().Done()
		wgCancelFunc()
	})

	select {
	case <-ctx.Done():
		return errors.New("stop crontab timeout")
	case <-wgCtx.Done():
		return nil
	}
}
