package main

//type MyDuration = time.Duration

// Cannot define new methods on the non-local type 'time.Duration'
//func (m MyDuration) EasySet(a string) {
//
//}

// 解决以上问题可以使用新的类型替代别名：type MyDuration time.Duration
