package customUtil

type LoggerUtil struct {
	requestSqlLog []string //sql日志
	requestCustomLog []string //自定义日志

}

func NewLoggerUtil() *LoggerUtil {
	return &LoggerUtil{}
}

//写入sql日志记录
func (this *LoggerUtil) InfoSqlLog(log string)  {
	this.requestSqlLog = append(this.requestSqlLog,log)
}

//获取sql日志
func (this *LoggerUtil) GetSqlLog() []string{
	return this.requestSqlLog
}

//写入自定义日志
func (this *LoggerUtil) Info(log string)  {
	this.requestCustomLog = append(this.requestCustomLog, log)
}

//获取自定义的日志
func (this *LoggerUtil) GetCustomLog() []string{
	 return this.requestCustomLog
}