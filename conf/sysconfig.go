package conf

type sysconfig struct {

	//thrift 服务ip
	ThriftHost string `json:"ThriftHost"`
	ThriftPort string `json:"ThriftPort"`

	//日志存储地址、级别
	LoggerPath  string `json:"LoggerPath"`
	LoggerLevel string `json:"LoggerLevel"`

	//日志中显示相关密文
	ShadeInLog bool `json:ShadeInLog`
}

const ProjectName = "MPDCDS_ProtocolConversion"
