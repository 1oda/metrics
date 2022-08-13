package core

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"metrics/utils/config"
	logs "metrics/utils/log"
	"strconv"
)

const (
	Mysql          = "mysql"
	MysqlChoose    = "default"
	OrmDebug       = false
	DefaultMaxConn = 100
	DefaultMaxIdle = 50
	DefaultForce   = false
	DefaultVerbose = false
)

// Process 进程信息
type Process struct {
	ID         int    `orm:"column(id);auto"`
	PID        int    `orm:"column(pid);int(11);" description:"pid"`
	User       string `orm:"column(user);size(200);" description:"user"`
	Cmd        string `orm:"column(cmd);size(2000);" description:"cmd"`
	CmdLine    string `orm:"column(cmd_line);size(2000)" description:"cmd_line"`
	UpdateTime int    `orm:"column(update_time);int(11)" description:"update time"`
}

// MonitorData 进程信息
type MonitorData struct {
	ID         int    `orm:"column(id);auto"`
	ProcessID  string `orm:"column(process_id);int(11);" description:"进程ID"`
	CPUUsage   string `orm:"column(cpu_usage);float64" description:"CPU 利用率"`
	MEMUsage   string `orm:"column(mem_usage);float64" description:"内存 利用率"`
	UpdateTime int    `orm:"column(update_time);int(11)" description:"update time"`
}

func InitDB() {

	defaultHost := config.AppConfBO.MysqlHost
	defaultPort := config.AppConfBO.MysqlPort
	defaultUser := config.AppConfBO.MysqlUser
	defaultPass := config.AppConfBO.MysqlPass
	defaultDatabase := config.AppConfBO.MysqlDatabase

	maxIdle, maxConn := DefaultMaxIdle, DefaultMaxConn
	if config.AppConfBO.MysqlMaxIdle != "" {
		maxIdle, err := strconv.Atoi(config.AppConfBO.MysqlMaxIdle)
		if err != nil {
			logs.Error(err)
		}

		if maxIdle < 1 {
			err = errors.New("wrong max idle number")
			logs.Error(err)
			panic(err)
		}
	}

	if config.AppConfBO.MysqlMaxConn != "" {
		maxConn, err := strconv.Atoi(config.AppConfBO.MysqlMaxConn)
		if err != nil {
			logs.Error(err)
		}

		if maxConn < 1 {
			err = errors.New("wrong max conn number")
			logs.Error(err)
			panic(err)
		}
	}

	err := orm.RegisterDriver(Mysql, orm.DRMySQL)
	if err != nil {
		logs.Error(err.Error())
		panic(err)
	}
	defaultDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", defaultUser, defaultPass, defaultHost, defaultPort, defaultDatabase)
	err = orm.RegisterDataBase(MysqlChoose, Mysql, defaultDsn, maxIdle, maxConn)
	if err != nil {
		logs.Error(err.Error())
		panic(err)
	}

	// 开启调试模式
	orm.Debug = OrmDebug
	if config.AppConfBO.Env == "dev" {
		orm.Debug = true
	}
	orm.RegisterModel(new(Process))
	orm.RegisterModel(new(MonitorData))
	err = orm.RunSyncdb(MysqlChoose, DefaultForce, DefaultVerbose)
	if err != nil {
		logs.Error(err.Error())
		panic(err)
	}
}

// 进程信息

func (m *Process) ProcessTableName() string {
	return "Process"
}

func (m *Process) ProcessTableEngine() string {
	return "INNODB DEFAULT CHARSET=utf8"
}

func (m *Process) ProcessInsert() int64 {
	id, err := orm.NewOrm().Insert(m)
	if err != nil {
		return 0
	}
	return id
}

func (m *Process) ProcessInsertMulti(mul []Process) error {
	if _, err := orm.NewOrm().InsertMulti(len(mul), mul); err != nil {
		return err
	}
	return nil
}

func (m *Process) ProcessRead(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Process) ProcessUpdate(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Process) ProcessDelete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Process) ProcessQuery() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

// 监控信息管理

func (m *MonitorData) MonitorDataTableName() string {
	return "MonitorData"
}

func (m *MonitorData) MonitorDataTableEngine() string {
	return "INNODB DEFAULT CHARSET=utf8"
}

func (m *MonitorData) MonitorDataInsert() int64 {
	id, err := orm.NewOrm().Insert(m)
	if err != nil {
		return 0
	}
	return id
}

func (m *MonitorData) MonitorDataInsertMulti(mul []MonitorData) error {
	if _, err := orm.NewOrm().InsertMulti(len(mul), mul); err != nil {
		return err
	}
	return nil
}

func (m *MonitorData) MonitorDataRead(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *MonitorData) MonitorDataUpdate(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *MonitorData) MonitorDataDelete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *MonitorData) MonitorDataQuery() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
