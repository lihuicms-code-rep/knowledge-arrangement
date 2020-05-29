package main
import "fmt"

//非阻塞通道操作

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <- messages:
		fmt.Printf("msg:%s\n", msg)
	default:                                //不加default就会死锁
		fmt.Println("no msg recv")
	}


	select {
	case messages <- "hi":
		fmt.Println("sent msg")
	default:
		fmt.Println("no msg sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}