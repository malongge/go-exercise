package container

type SingleList struct {
	head *Element
	tail *Element
	size int
}

type Element struct {
	value interface{}
	next  *Element
}



// 通过多个值，直接创建一个单链表
func NewSingleList(value ...interface{}) *SingleList {
	list := new(SingleList)
	if len(value) > 0 {
		list.Append(value...)
	}

	return list
}

func (list *SingleList) Swap(index1 int, index2 int) {

}

// 往链表中添加节点，可以一次添加多个
func (list *SingleList) Append(value ...interface{}) {
	for _, each := range value {
		newElement := &Element{each, list.tail}
		if list.size == 0 {
			list.head = newElement
			list.tail = newElement
		} else {
			list.tail.next = newElement
			list.tail = newElement
		}
		list.size += 1
	}
}

// 获取链表上，某个索引位置的值
func (list *SingleList) GetValue(index int) (findValue interface{}, findState bool) {

	if !list.withinRange(index) {
		return nil, false
	}

	e := list.head
	for i := 0; i != index; i, e = i+1, e.next {
	}
	return e.value, true

}


func (list *SingleList) swapValue(index1 int, index2 int) {

}

// 判断索引是否超出了链表的范围
func (list *SingleList) withinRange(index int) bool {
	if index >= 0 && index < list.size {
		return true
	}
	return false
}

// 获得链表的长度
func (list *SingleList) Size() int {
	return list.size
}

// 判断链表是否为空
func (list *SingleList) Empty() bool {
	return !(list.size > 0)
}


// 移除索引位置的元素， 超出访问的索引，将不起作用
func (list *SingleList) Remove(index int){
	if !list.withinRange(index){
		return
	}

	if list.size == 1 && index == 0{
		// 只有一个元素的情况下特殊处理
		list.head = nil
		list.tail = nil
		list.size = 0
		return
	}

	e := list.head
	var preElement *Element
	for i:=0; i !=index; i, e = i+1, e.next{
		preElement = e
	}
	// 如果删除的元素是第一个
	if e == list.head{
		list.head = e.next
	}
	// 如果删除的元素是最后一个
	if e == list.tail {
		list.tail = preElement
	}
	// 删除的是非首个元素的处理
	if preElement != nil{
		preElement.next = e.next
	}
	e = nil // 删除该元素
	list.size -= 1
}
