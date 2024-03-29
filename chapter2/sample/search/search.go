package search

import (
	"log"  // log 包提供打印日志信息到标准输出（stdout）、标准错误（stderr）或者自定义设备的功能。
	"sync" // sync 包提供同步 goroutine 的功能。
)

// 注册用于搜索的匹配器的映射,且声明为 Matcher 类型的映射（map），这个映射以 string 类型值作为键，Matcher类型值作为映射后的值。
var matchers = make(map[string]Matcher)

// Run 执行搜索逻辑
func Run(searchTerm string) {
	// 获取需要搜索的数据源列表
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个无缓冲的通道， 接收匹配后的结果
	results := make(chan *Result)

	// 构造一个waitGroup， 以便处理所有的数据源
	var waitGroup sync.WaitGroup

	// 设置需要等待处理
	// 每个数据源的goroutine的数量
	waitGroup.Add(len(feeds))

	// 为每个数据源启动一个goroutine来查找结果
	for _, feed := range feeds {
		// 获取一个匹配器用于查找
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]z
		}

		// 启动一个goroutine来监控是否所有的工作都做完了
		go func(matcher Matcher, feed *Feed) {
			Matcher(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// 启动一个goroutine来监控是否所有的工作都做完了
	go func() {
		// 等候所有任务完成
		waitGroup.Wait()

		// 用关闭通道的方式，通知Display函数
		// 可以退出程序了
		close(rersults)
	}()

	// 启动函数，显示最后返回的结果，并且
	// 在最后一个结果显示完后返回
	Display(results)
}
