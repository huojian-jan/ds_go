package main

import (
	"fmt"
	"reflect"
)

type LinkedNode struct{
	Data int
	Next *LinkedNode
}


func insertNewData(lastNode *LinkedNode,data int) *LinkedNode{
	newnode:=new(LinkedNode)
	newnode.Data=data
	*lastNode.Next=*newnode
	return newnode
}

func traverseLinkedList(head *LinkedNode){
	for{
		fmt.Println(head.Data)
		if(head.Next==nil){
			break
		}
		head=head.Next
	}
}


func test1(){
	node_0:=new(LinkedNode)
	node_0.Data=0

	node_1:=new(LinkedNode)
	node_1.Data=1
	node_0.Next=node_1

	node_2:=new(LinkedNode)
	node_2.Data=2
	node_1.Next=node_2

	traverseLinkedList(node_0)
}

type Ring struct{
	Data int
	Pre *Ring
	Next *Ring
}


func insert(pre *Ring,data int){
	newNode:=new(Ring)
	newNode.Data=data
	nextNode:=pre.Next

	pre.Next=newNode
	nextNode.Pre=newNode

	newNode.Pre=pre
	newNode.Next=nextNode
}


func ring_test(){
head:=new(Ring)
head.Data=0
head.Pre=head
head.Next=head

insert(head,3)
insert(head,2)
insert(head,1)

fmt.Println(head.Data)
fmt.Println(head.Next.Data)
fmt.Println(head.Next.Next.Data)
fmt.Println(head.Next.Next.Next.Data)
fmt.Println(head.Next.Next.Next.Next.Data)

fmt.Println("===========================")

fmt.Println(head.Pre.Data)
fmt.Println(head.Pre.Pre.Data)
fmt.Println(head.Pre.Pre.Pre.Data)
fmt.Println(head.Pre.Pre.Pre.Pre.Data)
fmt.Println(head.Pre.Pre.Pre.Pre.Pre.Data)

}


func named_return_test(add1,add2 int)(sum int){
	sum=add1+add2
	return
}

func return_multiple_vals_test()(int,int){
left:=100
right:=200
return left,right
}


func varidic_numbers_test1(data ...int) int{
	fmt.Println(data)
	sum:=0
	for idx,val:=range data{
		fmt.Printf("idx=%d\t,val=%d\n",idx,val)
		sum+=val
	}
	return sum
}


func variadic_args_test(data ...int){
for _, val :=range  data{
	fmt.Println(val)
}	
}



func slice_test1(){
	data:=[5]int{1,2,3,4,5}
	fmt.Println(reflect.TypeOf(data))
}


func slice_test2(){
	slice:=make([]int,5,10)
	fmt.Println(slice)
}



func where(data []int,predicate func(int) bool) []int{
	result:=[]int{}
	for _,v:=range data{
		if predicate(v){
			result=append(result,v)
		}
	}
	return result
}


func where_test1(){
	predicate:=func(v int) bool{
		return v>5
	}

	data:=[]int{1,2,5,8,6,12,27,23}

	result:=where(data,predicate)

	fmt.Println(result)
}


type student struct{
	name string
	age int
	id int
}


func student_test1(){
	stu1:=student{"zhangsan",18,101}
	fmt.Printf("name in outside change name:%s\n",stu1.name)
	student_change_name(&stu1)
	fmt.Printf("name in outside change name:%s\n",stu1.name)
}

func student_change_name(stu *student){
	stu.name="lisi"
	fmt.Printf("name in change func:%s\n",stu.name)
}

type province struct{
	name string
}

type city struct{
	name string
}

type address struct{
	province province
	city city
}

type man struct{
	name string
	address address
}

func(p province) printProvince(){
fmt.Println(p.name)
}

func struct_func(){
	p:=province{
		name: "zhejiang",
	}
	p.printProvince()
}


func struct_composition_test1(){
}

func main(){

	struct_func()

	// where_test1()
	// data:=make([]int,10);
	// for i:=0;i<10;i++{
	// 	data[i]=i
	// }
	// variadic_args_test(data...)


// greeting :=func(){
// 	fmt.Println("hello friend")
// }

// fmt.Printf("%T\n",greeting)


//closure

// x:=100
// fmt.Println(x){
// 	fmt.println(x)

// 	y:="y var inside"
// 	fmt.println(y)
// }
// // fmt.println(y)
}
