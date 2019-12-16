package method

import "github.com/snail/util"


func  (args *Args)GetMsgId()byte{
	return args.MsgId
}

func  (args *Args)GetMethodName()string{
	return args.Name
}

func (args *Args) ToInt64() int64 {
	return util.BytesToInt64(args.Data)
}

func (args *Args) ToString()  string{
	return string(args.Data)
}

func (args *Args) ToMap() (map[string]interface{},error) {
	return util.BytesToMap(args.Data)
}
