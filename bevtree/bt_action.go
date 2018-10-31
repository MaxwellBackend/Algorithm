package bevtree

import (
	"fmt"
	"math/rand"
)

const (
	MONSTER_ABILITY = 1 // 怪物战力
	HERO_ABILITY    = 2 // 机器人战力
	STEP_COUNTER    = 3 // 步数计数器
)

const (
	ABILITY_LIMIT = 1000 // 战力上限
)


/*
	行为节点 闲置
 */
type ActionIdle struct {
	*BTNode
}

func NewActionIdle(parent IBTNode) *ActionIdle {
	return &ActionIdle{
		BTNode: NewBTNode(parent),
	}
}

func (a *ActionIdle) Enter() {
	fmt.Println("================================================")
	fmt.Println("action idle enter")
}

func (a *ActionIdle) PreCondition() bool {
	return true
}

func (a *ActionIdle) Execute() bool {

	a.Enter()
	fmt.Println("action idle excute")
	a.Exit()
	return true
}

func (a *ActionIdle) Exit() {
	fmt.Println("action idle exit")
	fmt.Println("================================================")
}

/*
	行为节点 漫步
 */
type ActionStep struct {
	*BTNode
}

func NewActionStep(parent IBTNode) *ActionStep {
	return &ActionStep{
		BTNode: NewBTNode(parent),
	}
}

func (a *ActionStep) Enter() {
	fmt.Println("================================================")
	fmt.Println("action step enter")
}

func (a *ActionStep) PreCondition() bool {
	blackBoard := GetBlackboard()
	ability, _ := blackBoard.GetValueAsInt(MONSTER_ABILITY)
	if ability > 0 {
		return false
	}
	return true
}

func (a *ActionStep) Execute() bool {

	a.Enter()
	fmt.Println("action step excute")

	blackBoard := GetBlackboard()

	// 检查步数
	step, _ := blackBoard.GetValueAsInt(STEP_COUNTER)

	nowStep := step + 1
	blackBoard.SetValueAsInt(STEP_COUNTER, nowStep)

	fmt.Printf("robot now step： %v\n", nowStep)

	// 随机小怪兽战力
	if rand.Float32() > 0.5 {
		ability := rand.Intn(ABILITY_LIMIT)
		blackBoard.SetValueAsInt(MONSTER_ABILITY, ability)
		fmt.Printf("find a monster with ability： %v\n", ability)
	} else {
		blackBoard.SetValueAsInt(MONSTER_ABILITY, 0)
		fmt.Printf("find nothing\n")
	}

	a.Exit()
	return true
}

func (a *ActionStep) Exit() {
	fmt.Println("action step exit")
	fmt.Println("================================================")

}

/*
	行为节点 奔跑
 */
type ActionRun struct {
	*BTNode
}

func NewActionRun(parent IBTNode) *ActionRun {
	return &ActionRun{
		BTNode: NewBTNode(parent),
	}
}

func (a *ActionRun) Enter() {
	fmt.Println("================================================")
	fmt.Println("action run enter")
}

func (a *ActionRun) PreCondition() bool {
	return true
}

func (a *ActionRun) Execute() bool {
	a.Enter()
	fmt.Println("action run excute")
	a.Exit()
	return true
}

func (a *ActionRun) Exit() {
	fmt.Println("action run exit")
	fmt.Println("================================================")
}

/*
	行为节点 挥刀进攻
 */
type ActionSlash struct {
	*BTNode
}

func NewActionSlash(parent IBTNode) *ActionSlash {
	return &ActionSlash{
		BTNode: NewBTNode(parent),
	}
}

func (a *ActionSlash) Enter() {
	fmt.Println("================================================")
	fmt.Println("action slash enter")
}

func (a *ActionSlash) Execute() bool {
	a.Enter()
	fmt.Println("action slash excute")
	a.Exit()
	return true
}

func (a *ActionSlash) Exit() {
	fmt.Println("action slash exit")
	fmt.Println("================================================")
}

/*
	行为节点 喊叫
 */
type ActionShout struct {
	*BTNode
}

func NewActionShout(parent IBTNode) *ActionShout {
	return &ActionShout{
		BTNode: NewBTNode(parent),
	}
}

func (a *ActionShout) Enter() {
	fmt.Println("================================================")
	fmt.Println("action shout enter")
}

func (a *ActionShout) Execute() bool {
	a.Enter()
	fmt.Println("action shout excute")
	a.Exit()
	return true
}

func (a *ActionShout) Exit() {
	fmt.Println("action shout exit")
	fmt.Println("================================================")
}

/*
	并行节点 逃跑
*/
type ActionEscape struct {
	*BTParallel
}

func NewActionEscape(parent IBTNode) *ActionEscape {
	return &ActionEscape{
		BTParallel: NewBtParallel(parent),
	}
}

func (a *ActionEscape) Enter() {
	fmt.Println("================================================")
	fmt.Println("action escape enter")
}

func (a *ActionEscape) PreCondition() bool {
	blackBoard := GetBlackboard()
	heroAbility, _ := blackBoard.GetValueAsInt(HERO_ABILITY)
	monsterAbility, _ := blackBoard.GetValueAsInt(MONSTER_ABILITY)
	canFight := heroAbility >= monsterAbility

	if !canFight {
		fmt.Printf("robot ability: %v, monster ability: %v cant fight: %v\n", heroAbility, monsterAbility, canFight)
		return true
	}
	return false
}

func (a *ActionEscape) Execute() bool {
	a.Enter()
	fmt.Println("action escape excute")
	a.BTParallel.Execute()
	a.Exit()

	blackBoard := GetBlackboard()
	blackBoard.SetValueAsInt(MONSTER_ABILITY, 0)
	return true
}

func (a *ActionEscape) Exit() {
	fmt.Println("action escape exit")
	fmt.Println("================================================")
}

/*
	序列节点 战斗
*/
type ActionFight struct {
	*BTSequence
}

func NewActionFight(parent IBTNode) *ActionFight {
	return &ActionFight{
		BTSequence: NewBtSequence(parent),
	}
}

func (a *ActionFight) Enter() {
	fmt.Println("================================================")
	fmt.Println("action fight enter")
}

func (a *ActionFight) PreCondition() bool {
	blackBoard := GetBlackboard()
	heroAbility, _ := blackBoard.GetValueAsInt(HERO_ABILITY)
	monsterAbility, _ := blackBoard.GetValueAsInt(MONSTER_ABILITY)
	canFight := heroAbility >= monsterAbility
	if canFight {
		fmt.Printf("robot ability: %v, monster ability: %v cant fight: %v\n", heroAbility, monsterAbility, canFight)
		return true
	}
	return false
}

func (a *ActionFight) Execute() bool {
	a.Enter()
	fmt.Println("action fight excute")
	a.BTSequence.Execute()
	a.Exit()

	blackBoard := GetBlackboard()
	blackBoard.SetValueAsInt(MONSTER_ABILITY, 0)

	return true
}

func (a *ActionFight) Exit() {
	fmt.Println("action fight exit")
	fmt.Println("================================================")
}
