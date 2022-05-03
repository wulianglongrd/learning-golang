package main

//type MyDuration = time.Duration

// Cannot define new methods on the non-local type 'time.Duration'
//func (m MyDuration) EasySet(a string) {
//
//}

// 解决以上问题可以使用`新的类型`替代`别名`：type MyDuration time.Duration
