package sshutil

import (
	"bufio"
	"fmt"
	"github.com/wonderivan/logger"
	"io"
	"strings"
)




//Cmd is in host exec cmd
func (ss *SSH) Cmd(host string, cmd string) []string {
	logger.Info("[%s]exec cmd is : %s", host, cmd)
	session, err := ss.Connect(host)
	defer func() {
		if r := recover(); r != nil {
			logger.Error("[%s]Error create ssh session failed,%s", host, err)
		}
	}()
	if err != nil {
		panic(1)
	}
	defer session.Close()
	var stdReader io.Reader
	if stdReader ,err = session.StdoutPipe();err != nil {
		logger.Error("ssh session stdout pipe error:%s",err)
		panic(1)
	}
	reader := bufio.NewReader(stdReader)
	if err = session.Start(cmd);err!= nil {
		logger.Error("ssh session start CMD error:%s",err)
		panic(1)
	}
	var result []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		//line = strings.TrimSpace(line)
		result = append(result,strings.TrimSpace(line))
		fmt.Println(strings.TrimSpace(line))
	}
	err = session.Wait()
	defer func() {
		if r := recover(); r != nil {
			logger.Error("[%s]Error exec command failed: %s", host, err)
		}
	}()
	if err != nil {
		panic(1)
	}
	return result
}

//CmdToString is in host exec cmd and replace to spilt str
func (ss *SSH) CmdToString(host, cmd, spilt string) string {
	data := ss.Cmd(host, cmd)
	if len(data) >0 {
		str := strings.Join(data,spilt)
		logger.Debug("[%s]command %s result is: %s", host,cmd, str)
		return str
	}
	return ""
}
