LiteMessage
===========

轻量级的消息编码格式，用于客户端服务器间数据交换。简单性是此协议格式的第一目的，为此不考虑客户端服务器端新旧协议的兼容性；能自动生成客户端/服务器端的序列号/反序列化代码，但手动写也非常简单。

数据类型
--------
- string
- bool/byte/int16/int32/int64
- float/double
- 嵌套自定义类型
- 以上各种类型的数组

消息定义示例
--------

	message VoUser{
	  string name
	  int64 exp
	  int32 level
	}
	
	message CSLogin = 12001 
	{
	  string name  = "" //用户名
	  string password
	}
	
	message SCLogin = 12002
	{
	  byte result
	  VoUser[] userInfo
	}

