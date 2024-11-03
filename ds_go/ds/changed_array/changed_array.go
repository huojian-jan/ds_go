package main

import "unsafe"

type slice struct{
	array unsafe.Pointer
	len int
	cap int
}



	