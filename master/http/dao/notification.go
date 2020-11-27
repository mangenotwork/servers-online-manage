package dao

import (
	"github.com/mangenotwork/servers-online-manage/master/http/models"
)

type NotificationDao struct {
	models.Notifincation
}


func (this *NotificationDao) PendingCount() int64{
	return 10
}