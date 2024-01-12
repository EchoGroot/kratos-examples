package biz

import (
	"context"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
)

type BackupDbUsecase struct {
	sync.Mutex
	dbAVersion int32
	dbBVersion int32
}

func NewBackupDbUsecase() *BackupDbUsecase {
	return &BackupDbUsecase{
		dbAVersion: 0,
		dbBVersion: 0,
	}
}

func (cu *BackupDbUsecase) BackupDbA(ctx context.Context) error {
	cu.dbAVersion++
	log.Infof("backup db a, version: %d", cu.dbAVersion)
	return nil
}

func (cu *BackupDbUsecase) BackupDbBWithLock(ctx context.Context) error {
	if cu.TryLock() {
		defer cu.Unlock()
		return cu.BackupDbB(ctx)
	}
	return nil
}

func (cu *BackupDbUsecase) BackupDbB(ctx context.Context) error {
	cu.dbBVersion++
	log.Infof("backup db b, version: %d", cu.dbBVersion)
	return nil
}
