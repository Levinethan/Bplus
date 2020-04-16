package main

const  (
	//叶子节点的最大存储数量  保证2^n -1
	MaxKV = 511   //数据量 = max的平方

	//中间节点数量
	MaxKC = 1023
)

//接口设计
type node interface {
	find (key int)(int , bool) //查找key
	Parent() *interiorNode  //返回父节点


	SetParent (*interiorNode)  //设置父亲节点
	full() bool  //判断是否满了
	CountNum ()int  //统计元素数量

}
