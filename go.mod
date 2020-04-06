module github.com/cuisongliu/sshcmd

go 1.13

require (
	github.com/pkg/sftp v1.10.1
	github.com/spf13/cobra v0.0.5
	github.com/wonderivan/logger v1.0.0
	golang.org/x/crypto v0.0.0-20190820162420-60c769a6c586
	golang.org/x/sys v0.0.0-20200223170610-d5e6a3e2c0ae // indirect
)

replace github.com/wonderivan/logger => github.com/fanux/sealos/pkg/logger v0.0.0-20200406033522-73f3ef41977c
