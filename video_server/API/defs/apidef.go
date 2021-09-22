package defs

//requsets
type UserCreadential struct{
	Username string 'json:"user_name"'
	Pwd string 'json:"pwd"'
}

//Go里面的原生方法处理json的一种方式
//是在struct打Tag的方式，Tag的名称是user_name ，Tag的属性是json
//在序列化和反序列化json的时候会自动转换成我们想要的东西
//{
//	user_name: xxx,
//	pwd: xxx
//}





