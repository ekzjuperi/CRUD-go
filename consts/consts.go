package consts

const (
	Dump    = "dump"
	Dumpint = "dumpint"
	Text    = `
	db [string]string
	create key value
	read key
	update key value
	delete key
	exist key

	db [string]int
	createint key value
	readint key
	updateint key value
	deleteint key
	existint key
	sum
	avg
	med
	gt value
	lt value
	eq value
	count
	Press q to quit`
	Create    = "create"
	Read      = "read"
	Update    = "update"
	Delete    = "delete"
	Exist     = "exist"
	Quit      = "q"
	Sum       = "sum"
	Avg       = "avg"
	Gt        = "gt"
	Lt        = "lt"
	Eq        = "eq"
	Count     = "count"
	Med       = "med"
	Createint = "createint"
	Updateint = "updateint"
	Invalid   = "Invalid command"
	Readint   = "readint"
	Deleteint = "deleteint"
	Existint  = "existint"
)
