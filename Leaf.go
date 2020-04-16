package main

import "sort"

type kv struct {
	key int
	value string
}

type kvs [MaxKC]kv
type LeafNode struct {
	kvs  kvs   //数据
	count int  //数量
	next *LeafNode  //下一个节点
	parent *interiorNode  //父节点
}

func (l *LeafNode) full() bool {
	return l.count == MaxKV
}

func (kvs *kvs)Len() int  {
	return len(kvs)
}

func (kvs *kvs)Swap(i , j int)  {
	kvs[i],kvs[j] = kvs[j],kvs[i]
}

func (kvs *kvs)Less (i , j int) bool  {
	if kvs[i].key==0{  //中间节点第一个元素是空的
		return false
	}
	if kvs[j].key==0{
		return true
	}


	return kvs[i].key < kvs[j].key
}

func NewLeafNode(parent *interiorNode) *LeafNode  {   //创建叶子节点
	return &LeafNode{
		parent: parent,
	}
}

func (l *LeafNode)find (key int) (int , bool)  {
	myfunc := func(i int) bool {
		return l.kvs[i].key >= key
	}

	//myfunc是个函数 用来对比大小
	i := sort.Search(l.count,myfunc)
	if i <l.count && l.kvs[i].key == key{
		return i,true
	}
	return i,false
}

func(l *LeafNode) isFull() bool  {
	return l.count == MaxKV //判断是否满了
}
func (l *LeafNode)Parent() *interiorNode  {
	return l.parent  //判断是否满了
}

func (l *LeafNode)SetParent(p *interiorNode)  {
	l.parent = p
}

func (l *LeafNode)CountNum() int  {
	return l.count
}

//初始化数组， 数组每个元素初始化
func (l *LeafNode)InitArray (num int) {
	for i := num; i < len(l.kvs);i++{
		l.kvs[i]= kv{}
	}
}

func (l *LeafNode)insert(key int,value string) (int , *LeafNode,bool)  {
	i , ok := l.find(key)
	if ok{
		l.kvs[i].value = value  //key已经存在  更新value
		return 0,nil,false
	}
	if !l.isFull(){
		copy(l.kvs[i+1:],l.kvs[i:l.count])  //如果没有满 找到i  把i和中间的数据拷贝
		l.kvs[i].key = key  //数组的删除需要整体往后移动
		l.kvs[i].value=value
		l.count++
		return 0,nil,false
	}else {
		next := l.split()  //分裂叶子节点
		if key < next.kvs[0].key{
			l.insert(key,value)
		}else {
			next.insert(key,value)
		}
		return next.kvs[0].key,next,true
	}

}

func (l *LeafNode)split () *LeafNode  {
	next := NewLeafNode(nil) //新建一个右边节点
	copy(next.kvs[0:],l.kvs[l.count/2+1:])
	l.InitArray(l.count/2+1)  //后半部数据清空
	next.count = MaxKV-l.count/2-1  //下一个节点的数量
	next.next = l.next
	l.count = l.count/2+1 //取得中间节点
	l.next = next
	return next
}

