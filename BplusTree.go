package main

import "fmt"

type Bplustree map[int]node //定义 存储的时候用map存储

func NewBplustree()  *Bplustree  {
	b := Bplustree{}
	leaf := NewLeafNode(nil)
	r := NewinteriorNode(nil,leaf)
	leaf.parent=r
	b[-1] = r  //根节点
	b[0]=leaf
	return &b
	return nil
}

func (bpt *Bplustree)Root() node  {
	return (*bpt)[-1]  //返回根节点
}

//处理第一个节点
func (bpt *Bplustree)First()node  {
	return (*bpt)[0]
}

//统计数量
func (bpt *Bplustree)Count () int  {
	count := 0
	leaf := (*bpt)[0].(*LeafNode)
	for {
		count+=leaf.CountNum()
		if leaf.next == nil{
			break
		}

		leaf = leaf.next

	}
	return count
}

//查找
func (bpt *Bplustree)Values () []*LeafNode  {
	nodes := make([]*LeafNode,0) //开辟节点
	leaf := (*bpt)[0].(*LeafNode)
	for {
		nodes = append(nodes,leaf)  //数据节点叠加
		if leaf.next == nil{
			break
		}

		leaf = leaf.next

	}
	return nodes
}
//插入

func (bpt *Bplustree)Insert (key int,value string)  {
	//插入前  确定插入的位置 是否存在
	_ , oldindex , leaf := search((*bpt)[-1],key)
	p := leaf.Parent()  //保存父节点
	//插入叶子节点，判断是否分裂
	mid , nextleaf,bump := leaf.insert(key,value)
	if !bump{ //没有分裂的话  直接返回
		return
	}

	//分裂的节点 插入B+树
	(*bpt)[mid] = nextleaf

	var minnode node  //声明一个 中间节点变量
	minnode = leaf
	p.kcs[oldindex].child = leaf.next //设置父节点
	leaf.next.SetParent(p)  //分裂的节点 设置父节点
	interior , interiorP := p,p.Parent()  //获取中间节点和父节点
	//接下里是平衡的过程

	//迭代向上判断  是否平衡
	for{
		var oldindex int //保存老的索引
		var newinterior *interiorNode  //新的节点
		//判断是否到达根节点
		isRoot:= interiorP ==nil
		if !isRoot{
			oldindex, _ = interiorP.find(key)  //查找
		}
		//叶子节点分裂后的中间节点传入父节点中间节点 传入分裂节点
		mid , newinterior, bump = interior.insert(mid,minnode)
		if !bump{
			return
		}
		(*bpt)[newinterior.kcs[0].key] = newinterior //插入填充好了的 map
		if !isRoot{
			interiorP.kcs[oldindex].child=newinterior //没有到达根节点  直接插入父节点
			newinterior.SetParent(interiorP)
			minnode = interior

		}else {

			//更新节点
			(*bpt)[interior.kcs[0].key] =(*bpt)[-1] //备份根节点
			(*bpt)[-1] = NewinteriorNode(nil,newinterior)
			node := (*bpt)[-1].(*interiorNode)
			node.insert(mid,interior)
			(*bpt)[-1]= node
			newinterior.SetParent(node)

		}

		interior,interiorP=interiorP,interiorP.Parent()


	}



}

func (bpt *Bplustree)Search (key int) (string,bool)  {
	kv , _, _ := search((*bpt)[-1],key)
	if kv==nil{
		return "",false
	}else {
		return kv.value,true
	}

}

func search(n node , key int) (*kv,int,*LeafNode)  {
	curr := n
	oldindex := -1
	for {
		switch t:=curr.(type) {
		case *LeafNode:  //叶子节点搜索
			i,ok := t.find(key)
			if !ok{
				return nil,oldindex,t
			}else {
				return &t.kvs[i],oldindex,t
			}
		case *interiorNode:
			i , _ := t.find(key)
			curr = t.kcs[i].child  //中间节点查找
			oldindex = i
		default:
			panic("异常节点")
		}
	}
}

func main()  {

	bpt := NewBplustree()
	for i:=0 ; i<250000;i ++{
		bpt.Insert(i," ")
	}
	fmt.Println(bpt.Count())
	fmt.Println(bpt.Search(3))
}
