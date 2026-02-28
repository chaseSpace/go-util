package goutil

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Listening 监听退出信号并执行优雅关闭
func Listening(timeout time.Duration, onClose ...func(ctx context.Context)) {
	// 创建一个通道来接收退出信号
	quit := make(chan os.Signal, 1)

	// 监听指定的信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// 等待接收退出信号
	sig := <-quit

	log.Printf("Received signal: %s, start graceful shutdown...", sig)

	// 创建带有超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	for _, f := range onClose {
		f(ctx)
	}
}
