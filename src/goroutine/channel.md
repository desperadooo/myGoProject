# channel和wait group使用遇到的问题
## 1. len(channel)长度会随着channel数据读出而改变
```
ch := make(chan int)
ch <- 1
ch <- 2
for i := 0; i < len(ch); i++ {
    fmt.Println(<-ch)
}
```
上述代码执行的时候只会打印出1，因为在第一次读取出1之后，len(ch)就变成了1，从而跳出for循环
正确的做法应该是
```
ch := make(chan int)
ch <- 1
ch <- 2
length := len(ch)
for i := 0; i < length; i++ {
    fmt.Println(<-ch)
}
```
或者
```
ch := make(chan int)
ch <- 1
ch <- 2
for i := range ch {
    fmt.Println(i)
}
```

## 2. 一个channel被关闭后，依然可以正常从channel读取数据

## 3. wait group踩坑
```
func main() {
	//case 3 扇入
	ch := make(chan int, 5)
	ch1 := make(chan int, 5)
	ch2 := make(chan int, 5)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	go func() {
		for i := 5; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		for i := 10; i < 15; i++ {
			ch2 <- i
		}
		close(ch2)
	}()
	for val := range merge(ch, ch1, ch2) {
		fmt.Printf("get number from out: %d\n", val)
	}
}

func merge(channels ...<- chan int) <- chan int {
    var wg sync.WaitGroup
    out := make(chan int)
    output := func(ch <- chan int) {
    for val := range ch {
    out <- val
    fmt.Printf("put number on out: %d\n", val)
    }
    wg.Done()
    }
    wg.Add(len(channels))
    for _, ch := range channels {
    go output(ch)
    }
    wg.Wait()
    close(out)
    //  go func() {
    //	wg.Wait()
    //	close(out)
    //}()
    return out
}
```
上述代码的目的是将三个channel的值合并到一个channel中去处理，但是如果我们在merge函数中就执行
wg.Wait()方法的话，会导致死锁
原因是我们创建的是一个无缓冲的channel，在merge函数中Wait会导致多个写入，从而死锁
解决的方法有两种
第一种是创建一个带缓冲的channel 
```
out := make(chan int, 15)
```
第二种是将wg.Wait放在协程中去处理
```
go func() {
	wg.Wait()
	close(out)
}
```