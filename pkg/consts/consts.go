package consts

const (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"
)

const (
	FmtStarting      = "starting"
	FmtCannotStart   = "can't start"
	FmtErrOnStarting = "error on starting"
	FmtStopping      = "stopping"
	FmtCannotStop    = "can't stop"
	FmtErrOnStopping = "error on stopping"
)

const (
	UserRoleID = iota + 1
	AdminRoleID
)

const (
	Admin = "admin"
	User  = "user"
)
