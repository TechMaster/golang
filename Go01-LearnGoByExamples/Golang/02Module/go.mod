module github.com/techmaster.vn/app

go 1.16

replace github.com/TechMaster/greeting => ./greeting
replace github.com/TechMaster/foo => ./foo
require (
	github.com/TechMaster/greeting v1.0.0
  github.com/TechMaster/foo v1.0.0
)