package main

//题目9:用两个栈实现一个队列
/*
  用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，
  分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )
  type CQueue struct {

  }


  func Constructor() CQueue {

  }


  func (this *CQueue) AppendTail(value int)  {

  }


  func (this *CQueue) DeleteHead() int {

  }

 */

//自定义栈
type CStack struct {
	data []int //数据域
	top  int   //栈顶指针
	num  int   //当前栈元素数
}

func NewCStack() *CStack {
	return &CStack{
		data:make([]int, 0, 10000),
		top : -1,
		num : 0,
	}
}


//入栈
func (cs *CStack) push(data int) {
	cs.data = append(cs.data, data)
	cs.top++
	cs.num++
}

//出栈
func (cs *CStack) pop() int {
	if cs.isEmpty() {
		return -1
	}

	elem := cs.data[cs.top]
	cs.data = cs.data[:cs.top]
	cs.num--
	cs.top--
	return elem
}

//判断栈空
func (cs *CStack) isEmpty() bool {
	return cs.top == -1
}

func (cs *CStack) getTop() int {
	if cs.isEmpty() {
		return -1
	}

	return cs.data[cs.top]
}

type CQueue struct {
	stack1 *CStack   //栈1始终入栈
	stack2 *CStack   //栈2主要是辅助栈1
}

//构造一个队列
func Constructor() CQueue {
	return CQueue{
		stack1:NewCStack(),
		stack2:NewCStack(),
	}

}

//入队列
func (this *CQueue) AppendTail(value int)  {
	this.stack1.push(value)
}

//出队列
func (this *CQueue) DeleteHead() int {
	if this.stack1.isEmpty() {
		return -1
	}

	if this.stack1.num == 1 {
		return this.stack1.pop()
	}

	//栈1元素多于1个,需要借助栈2
	deleteItem := -1
	//先将stack1元素转移到stack2
	transNum := this.stack1.num  //这里注意要转移数据的次数
	for i := 0; i < transNum; i++ {
		this.stack2.push(this.stack1.pop())
	}

	//要出队列的元素
	deleteItem = this.stack2.pop()
	//再恢复
	transNum = this.stack2.num
	for i := 0; i < transNum; i++ {
		this.stack1.push(this.stack2.pop())
	}

	return deleteItem
}

//上面的DeleteHead还可以再优化下
//每次出队列的时候,先看下stack2是否为空,不为空就直接出栈stack2的元素
func (this *CQueue) DeleteHead2() int {
	if !this.stack2.isEmpty() {
		return this.stack2.pop()
	}

	if this.stack1.num == 1 {
		return this.stack1.pop()
	}

	//stack2为空,转移
	transNum := this.stack1.num
	for i := 0; i < transNum; i++ {
		this.stack2.push(this.stack1.pop())
	}

	//弹出stack2的栈顶即可
	return this.stack2.pop()
	return -1

}

/*
func main() {
	queue := Constructor()
	queue.AppendTail(1)
	queue.AppendTail(2)
	queue.AppendTail(3)
	param1 := queue.DeleteHead2()
	fmt.Printf("first delte head:%v\n", param1)
	queue.AppendTail(4)
	queue.AppendTail(5)
	param2 := queue.DeleteHead2()
	fmt.Printf("second delte head:%v\n", param2)
	param3 := queue.DeleteHead2()
	fmt.Printf("third delte head:%v\n", param3)
	param4 := queue.DeleteHead2()
	fmt.Printf("forth delte head:%v\n", param4)
	param5 := queue.DeleteHead2()
	fmt.Printf("fifth delte head:%v\n", param5)

	param6 := queue.DeleteHead2()
	fmt.Printf("sixth delte head:%v\n", param6)
}
*/
