package main

import "sort"

//存储结构
type kc struct {
	key int
	child node  //接口类型 node
}

type kcs [MaxKC+1]kc

func (k *kcs) Len() int {
	panic("implement me")
}

func (k *kcs) Less(i, j int) bool {
	panic("implement me")
}

func (k *kcs) Swap(i, j int) {
	panic("implement me")
}

//数组 ，存储一个数组

//中间节点数据结构
type interiorNode struct {
	kcs   kcs //数组存储结构
	count int  //存储元素的数量
	parent *interiorNode  //指向父节点
}

func (in *interiorNode) full() bool {
	return in.count == MaxKC
}

func(in *interiorNode) isFull() bool  {
	return in.count == MaxKC //判断是否满了
}
func (in *interiorNode)Parent() *interiorNode  {
	return in.parent  //判断是否满了
}

func (in *interiorNode)SetParent(p *interiorNode)  {
	in.parent = p
}

func (in *interiorNode)CountNum() int  {
	return in.count
}

//初始化数组， 数组每个元素初始化
func (in *interiorNode)InitArray (num int) {
	for i := num; i < len(in.kcs);i++{
		in.kcs[i]= kc{}
	}
}

func (in *interiorNode)find(key int) (int,bool)  {
	myfunc := func(i int) bool {
		return in.kcs[i].key >= key
	}
	//myfunc是个函数 用来对比大小
	i := sort.Search(in.count-1,myfunc)
	return i,false
}

func NewinteriorNode(p *interiorNode,largestChild node) *interiorNode  {
	in := &interiorNode{parent:p,count:1}
	if largestChild!=nil{
		in.kcs[0].child=largestChild
	}
	return in
}

func(in *interiorNode)insert (key int ,child node) (int ,*interiorNode,bool)  {
	//确定位置
	i , _ := in.find(key)
	if !in.isFull(){
		//没有满 插入  满了分裂
		copy(in.kcs[i+1:],in.kcs[i:in.count])

		in.kcs[i].key = key

		in.kcs[i].child = child

		child.SetParent(in)

		in.count++

		return 0,nil,false
	}else {
		in.kcs[MaxKC].key = key  //预留一个值
		in.kcs[MaxKC].child = child  //存储到最后
		child.SetParent(in) //设置父节点
		next, midkey :=in.split()
		return midkey,next,true
	}
}

func (in *interiorNode)split() (*interiorNode,int)  {
	//处理节点分裂  ，使节点插入正确位置
	sort.Sort(&in.kcs)

	midIndex := MaxKC /2
	midchild := in.kcs[midIndex].child
	midkey := in.kcs[midIndex].key

	//新建一个中间节点
	next := NewinteriorNode(nil,nil)
	copy(next.kcs[0:],in.kcs[midIndex+1:])
	in.InitArray(midIndex+1) //数据初始化
	next.count = MaxKC-midIndex //下一个节点数量
	for i:=0 ;i<next.count;i++{
		next.kcs[i].child.SetParent(next)  //新开辟节点的每个叶子的节点的祖先 设置为next
	}
	in.count = midIndex+1  //更新数量
	in.kcs[in.count-1].key = 0  //设置为0  预留一个
	in.kcs[in.count-1].child=midchild  //设置中间节点
	midchild.SetParent(in)  //设置父亲节点
	return next,midkey
}