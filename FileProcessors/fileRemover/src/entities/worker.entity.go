package entities

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