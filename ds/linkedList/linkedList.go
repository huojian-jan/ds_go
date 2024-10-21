package main

import (
	"fmt"
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

func main(){
	ring_test()
}
