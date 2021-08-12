package main

import (
	"fmt"
	"math/rand"
	"os"
)

type FoodMenuManager struct {
	addFoodMenu map[int64]*FoodMenu
}

type FoodMenu struct {
	id   int64
	name string
}

func newFoodMenuManager(allStudent map[int64]*FoodMenu) *FoodMenuManager {
	return &FoodMenuManager{
		addFoodMenu: allStudent,
	}
}
func newFoodMenu(id int64, name string) *FoodMenu {
	return &FoodMenu{
		id:   id,
		name: name,
	}
}

func (s FoodMenuManager) showFoodMenu() {
	for index, value := range s.addFoodMenu {
		fmt.Println("ID：", index, "  食物名称：", value.name)
	}
}

func (s FoodMenuManager) addFoodMenuNew(id int64, name string) {
	_, flag := s.addFoodMenu[id]
	if flag {
		fmt.Println("已存在，若想修改，请重新选择3：修改功能")
	} else {
		var menu *FoodMenu
		menu = &FoodMenu{
			id:   id,
			name: name,
		}
		fmt.Println("我要开始添加", menu)
		s.addFoodMenu[id] = menu
		fmt.Println("添加成功！")
		for index, value := range s.addFoodMenu {
			fmt.Println("ID：", index, "  食物名称：", value.name)
		}
	}
}

func (s FoodMenuManager) changeFoodMenu(id int64, name string) {
	value, flag := s.addFoodMenu[id]
	if !flag {
		fmt.Println("查无食物名称，请先添加信息，请重新选择2：添加功能")
	} else {
		value.name = name
		fmt.Println("修改成功")
		for index, value := range s.addFoodMenu {
			fmt.Println("ID：", index, "  食物名称：", value.name)
		}
	}
}

func (s FoodMenuManager) deleteFoodMenu(id int64) {
	delete(s.addFoodMenu, id)
	fmt.Println("删除成功")
	for index, value := range s.addFoodMenu {
		fmt.Println("ID：", index, "  食物名称：", value.name)
	}
}

func (s FoodMenuManager) randomFoodMenu() {
	r := rand.Intn(len(s.addFoodMenu))
	for index, value := range s.addFoodMenu {

		if r == 0 {
			fmt.Println("^-^运气真好！那就吃它吧^-^\t")
			fmt.Println("ID：", index, "  食物名称：", value.name)
		}
		r--
	}
}

func main() {
	allFoodMenuMap := make(map[int64]*FoodMenu, 48)
	allFoodMenuMap[1001] = newFoodMenu(1001, "螺蛳粉")
	allFoodMenuMap[1002] = newFoodMenu(1002, "炒河粉")
	allFoodMenuMap[1003] = newFoodMenu(1003, "肉丝炒饭")
	allFoodMenuMap[1004] = newFoodMenu(1004, "热干面")
	allFoodMenuMap[1005] = newFoodMenu(1005, "麻辣烫")
	allFoodMenuMap[1006] = newFoodMenu(1006, "卤肉饭")
	var foodMenuM = newFoodMenuManager(allFoodMenuMap)
	for {
		fmt.Print(`
	欢迎来到食物随机管理系统！
	1.查看所有食物信息
	2.新增食物信息
	3.修改食物信息
	4.删除食物信息
	5.筛选我要吃啥？
	6.退出
	请选择您需要的功能：
`)
		var tmp int8
		fmt.Scanln(&tmp)
		switch tmp {
		case 1:
			foodMenuM.showFoodMenu()
		case 2:
			var (
				id   int64
				name string
			)
			fmt.Print("请输入食物ID：")
			fmt.Scanln(&id)
			fmt.Print("请输入食物名称：")
			fmt.Scanln(&name)
			foodMenuM.addFoodMenuNew(id, name)
		case 3:
			var (
				id   int64
				name string
			)
			fmt.Print("请输入食物ID：")
			fmt.Scanln(&id)
			fmt.Print("请输入想要修改的食物名称：")
			fmt.Scanln(&name)
			foodMenuM.changeFoodMenu(id, name)
		case 4:
			var (
				id int64
			)
			fmt.Print("请输入删除食物的ID：")
			fmt.Scanln(&id)
			foodMenuM.deleteFoodMenu(id)
		case 5:
			fmt.Println()
			fmt.Print("^-^拼命计算中...")
			fmt.Println()
			foodMenuM.randomFoodMenu()
		case 6:
			os.Exit(0)
		default:
			fmt.Println("无效的输入，请重试")
		}
	}
}
