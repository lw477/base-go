package main

import (
	"errors"
	"fmt"
)

// for循环中的return、break、continue

type Order struct {
	OrderID int
	Amount  float64
	Status  string // "pending", "invalid", "cancel_batch", "fatal_error"
}

func ProcessOrders(orders []Order) (int, error) {
	processedCount := 0

	for _, order := range orders {
		// 1.使用continue，跳过无效的订单，继续处理下一个
		if order.Status == "invalid" || order.Amount <= 0 {
			fmt.Printf("【跳过】订单%d 状态无效或者金额异常\n", order.OrderID)
			continue
		}

		// 2. 使用break，遇到了批量取消指令，停止后续订单的处理，但函数正常结算
		if order.Status == "cancel_batch" {
			fmt.Printf("【中断】收到批量取消的指令（订单%d），停止后续处理! \n", order.OrderID)
			break
		}

		// 3. 使用return, 遇到严重的系统错误，直接中断整个函数并返回错误
		if order.Status == "fatal_error" {
			fmt.Printf("[报错] 订单 %d 触发严重错误，程序直接退出！\n", order.OrderID)
			return processedCount, errors.New("数据库连接中断，无法继续处理") // 立即结束整个函数
		}

		fmt.Printf("[成功] 订单 %d 处理成功，金额：%.2f\n", order.OrderID, order.Amount)
		processedCount++
		fmt.Println(processedCount)
	}
	// 只有当循环正常结束，或者被 break 中断时，才会执行到这里
	return processedCount, nil
}

func main() {
	orders := []Order{
		{OrderID: 101, Amount: 100.0, Status: "pending"},      // 正常处理
		{OrderID: 102, Amount: -500.9, Status: "invalid"},     // 跳过
		{OrderID: 103, Amount: 200.0, Status: "pending"},      // 正常处理
		{OrderID: 104, Amount: 100.0, Status: "cancel_batch"}, // 会触发 break
		{OrderID: 105, Amount: 300.0, Status: "pending"},      // 不会触发，因为前面break了
	}

	fmt.Println("===开始处理订单数据===")
	count, err := ProcessOrders(orders)

	if err != nil {
		fmt.Println("处理失败，错误信息:", err)
	} else {
		fmt.Printf("处理完成，共成功处理 %d 个订单\n", count)
	}
}

// 总结： return：直接结束当前函数，break：跳过当前循环，continue：继续下一次循环
