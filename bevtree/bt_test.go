package bevtree

import (
	"testing"
	"math/rand"
	"fmt"
	"time"
)

// 构建行为树
func createBevTree() IBTNode {

	// 创建行为节点
	step := NewActionStep(nil)   // 漫步
	run := NewActionRun(nil)     // 奔跑
	shout := NewActionShout(nil) // 叫喊
	slash := NewActionSlash(nil) // 挥砍

	// 构建逻辑节点
	// 逃跑，并行节点，一边跑，一边叫
	escape := NewActionEscape(nil)
	escape.AddNode(run)
	escape.AddNode(shout)

	// 战斗，序列节点，跑到敌人面前，然后挥砍
	fight := NewActionFight(nil)
	fight.AddNode(run)
	fight.AddNode(slash)

	// 根节点，选择节点，选择是执行逃跑，战斗，还是走步逻辑
	root := NewBTSelector(nil)
	root.AddNode(step)
	root.AddNode(escape)
	root.AddNode(fight)

	return root
}

// 一个简单机器人
type Robot struct {
	Name     string
	Ability  int
	behavior IBTNode
}

func TestBehavior_Tree(t *testing.T) {
	//  设置随机种子
	rand.Seed(time.Now().Unix())

	// 初始化机器人
	robot := Robot{
		Name:    "Hodor",
		Ability: rand.Intn(ABILITY_LIMIT),
	}

	fmt.Printf("I'm a robot name: %v, ability: %v\n", robot.Name, robot.Ability)

	// 保存机器人战力
	blackBoard := GetBlackboard()
	blackBoard.SetValueAsInt(HERO_ABILITY, robot.Ability)

	// 创建机器人行为树
	robot.behavior = createBevTree()

	// 循环执行100次，观察行为
	counter := 0
	for ; ;
	{
		robot.behavior.Execute()
		counter ++
		if counter > 100 {
			break
		}
		time.Sleep(time.Second * 5)
	}
}
