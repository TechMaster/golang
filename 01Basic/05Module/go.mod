module github.com/techmaster.vn/app

go 1.16

replace github.com/TechMaster/greeting => ./greeting

replace github.com/TechMaster/foo => ./foo

require (
	github.com/TechMaster/foo v1.0.0
	github.com/TechMaster/greeting v1.0.0
	github.com/TechMaster/mygomodule v0.0.0-20210526162113-c243dee7fcfd
)
