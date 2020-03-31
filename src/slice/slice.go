package main

import "fmt"

func main()  {
	//基于数组得到切片
	a := [5]int{55,56,57,58,59}
	b := a[1:4]
	fmt.Println(b)
	fmt.Printf("%T\n",b)

	//切片再次切片
	c := b[0:len(b)]
	fmt.Printf("%T\n",c)
	d := make([]int,5,10)
	fmt.Printf("%T\n",d)
	fmt.Println(d)
	fmt.Printf("%d\n%d",len(d),cap(d))
	var e []int
	var f = []int{}
	if e == nil {
		fmt.Println("\na是一个nil")
	}
	fmt.Println(e,len(e),cap(e),f)
	//切片是一个赋值拷贝
	 g := make([]int ,3)//[0 0 0]
	 h := g
	 g[0] = 100
	 fmt.Println(g,h)
	 //切片的遍历
	for i:=0;i<len(a) ;i++  {
		fmt.Println(a[i])
	}
	for _,j :=range a{
		fmt.Println(j)
	}
	//切片的扩容
	var slice_a [] int
	//切片要初始化后才能使用
	for i:=0;i<10 ;i++  {
		slice_a = append(slice_a,i)
	}
	//一次性添加多个元素
	slice_a = append(slice_a,10,11,21,232,321)
	fmt.Println(slice_a)
	//两个切片合并
	slice_b := []int {1,2,3,4,5,6}
	slice_a = append(slice_a,slice_b...)
	fmt.Println(slice_a)
	//copy赋值切片,copy的时候，只会填充len（需要拷贝到的切片）的长度
	sliceD := make([]int,20,20)
	copy(sliceD,slice_a)
	fmt.Println(sliceD)
	fmt.Println(slice_a)

}
