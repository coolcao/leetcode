package main

import "fmt"

/*
小明家住A城市，他有事要到B城市办事，从A城市到B城市中间有n个城市，
每个城市到下一个城市都会有两种方式，乘坐火车或者乘坐汽车，
输入一个二维数组，其中每个元素三个值，分别表示该城市到下一城市乘坐火车的时间，乘坐汽车的时间，以及从火车站到汽车站的时间。
假设，小明家到A城市的火车站以及到汽车站的时间都是30分钟，那么，求小明从家到B城市所需的最少时间。

输入是一个二维数组：
[
    [30,40,10],
    [35,20,15],
    [38,32,8]
]
表示从小明所在的A城市到B城市有3个中间城市，
从A到C1乘坐火车需30分钟，乘坐汽车需40分钟，A城市从火车站到汽车站需要10分钟（反过来从汽车站到火车站也是10分钟），
从C1到C2乘坐火车需要35分钟，乘坐汽车需要20分钟，C1城市从火车站到汽车站需要15分钟，
从C2到B城市乘坐火车需要38分钟，乘坐汽车需要32分钟，C2城市火车站到汽车站需要8分中。
*/

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func minTime(times [][]int) int {
	citys := len(times)
	if citys == 0 {
		return 0
	}
	// minTrain,minBus分别标识坐火车和坐汽车到达某一城市所需的最少时间
	minTrain, minBus := times[0][0], times[0][1]
	// minTime标识到达某一城市所需的最少时间
	minTime := min(minTrain, minBus)

	for i := 1; i < citys; i++ {
		time := times[i]

		// timeTrain,timeBus在这里其实就是“下一个”minTrain，minBus。
		// 由于这里的计算会重复用到上一个minTrain，因此只能重新定义一个变量来标识
		timeTrain := min(minTrain+time[0], minBus+time[2]+time[0])
		timeBus := min(minBus+time[1], minTrain+time[2]+time[1])

		minTrain, minBus = timeTrain, timeBus

		minTime = min(minTrain, minBus)
	}

	// 最后时间要加30分钟，因为小明从家到A城市的火车站或汽车站都需要30分钟
	return minTime + 30
}

func main() {
	times := [][]int{
		{30, 45, 10},
		{41, 37, 11},
		{48, 51, 19},
	}
	min := minTime(times)
	fmt.Printf("%d\n", min)
}
