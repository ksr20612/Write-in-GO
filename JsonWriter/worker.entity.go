package entities

import (
	"fmt"
)

type Worker struct {
	Worker_id     string `gorm:"column:worker_id"`  
	Worker_passwd string `gorm:"column:worker_passwd"`
	Worker_comms  string `gorm:"column:worker_comms"` 
	Worker_domain   string `gorm:"column:worker_domain"`  
	Worker_topic string `gorm:"column:worker_topic"`
	Worker_name string `gorm:"column:worker_name"`
	Worker_sex string `gorm:"column:worker_sex"`
	Worker_age string `gorm:"column:worker_age"`
	Worker_home string `gorm:"column:worker_home"`
	Worker_loc string `gorm:"column:worker_loc"`
	Worker_tool string `gorm:"column:worker_tool"`
	Worker_dialect string `gorm:"column:worker_dialect"`

}

func (user *Worker) TableName() string {
	return "worker"
}

func (user Worker) ToString() string {
	return fmt.Sprintf("worker_id: %s, worker_passwd: %s, worker_comms: %s , worker_domain: %s , worker_topic: %s , worker_name: %s , worker_sex: %s , worker_age: %s , worker_home: %s , worker_loc: %s , worker_tool: %s, worker_dialect: %s\n, Comm_id: %d, Move_comm : %s", user.Worker_id, user.Worker_passwd, user.Worker_comms, user.Worker_domain, user.Worker_topic, user.Worker_name, user.Worker_sex, user.Worker_age, user.Worker_home, user.Worker_loc, user.Worker_tool, user.Worker_dialect, user.Comm_id, user.Move_comm)
}
