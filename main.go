package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Todo 単一のTODOアイテムwを表す構造体
type Todo struct {
	ID        int
	Title     string
	Completed bool
	CreatedAt time.Time
}

// TodoManagerはTODOアイテムを管理するための構造体
type TodoManager struct {
	todos  []Todo
	nextId int
}

func NewTodoManager() *TodoManager {
	return &TodoManager{
		todos:  []Todo{},
		nextId: 1,
	}
}

// Addは新しいTODOを追加する
func (m *TodoManager) Add(title string) Todo {
	todo := Todo{
		ID:        m.nextId,
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}
	m.todos = append(m.todos, todo)
	m.nextId++
	return todo
}

// GetAll はすべてのTODOを取得する
func (tm *TodoManager) GetAll() []Todo {
	return tm.todos
}

// GetByIDはIDに一致するTODOを取得する
func (m *TodoManager) GetByID(id int) (Todo, bool) {
	for _, todo := range m.todos {
		if todo.ID == id {
			return todo, true
		}
	}
	return Todo{}, false
}

// updateはIDに一致するTODOを更新する
func (m *TodoManager) Update(id int, title string, completed bool) (Todo, bool) {
	for i, todo := range m.todos {
		if todo.ID == id {
			m.todos[i].Title = title
			m.todos[i].Completed = completed
			return m.todos[i], true
		}
	}
	return Todo{}, false
}

// DeleteはIDに一致するTODOを削除する
func (m *TodoManager) Delete(id int) bool {
	for i, todo := range m.todos {
		if todo.ID == id {
			m.todos = append(m.todos[:i], m.todos[i+1:]...)
			return true
		}
	}
	return false
}

// ToggleCompletedはIDに一致するTODOのCompletedをトグルする
func (m *TodoManager) ToggleCompleted(id int) (Todo, bool) {
	for i, todo := range m.todos {
		if todo.ID == id {
			m.todos[i].Completed = !m.todos[i].Completed
			return m.todos[i], true
		}
	}
	return Todo{}, false
}

func main() {
	todoManager := NewTodoManager()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n===== TODOアプリ =====")
		fmt.Println("1. TODOの追加")
		fmt.Println("2. TODOの一覧表示")
		fmt.Println("3. TODOの更新")
		fmt.Println("4. TODOの削除")
		fmt.Println("5. TODOの完了/未完了切り替え")
		fmt.Println("0. 終了")
		fmt.Print("選択してください: ")

		scanner.Scan()
		choice := scanner.Text()
		switch choice {
		case "1":
			fmt.Print("TODOのタイトルを入力してください: ")
			scanner.Scan()
			title := scanner.Text()
			todo := todoManager.Add(title)
			fmt.Printf("TODOを追加しました: ID=%d, タイトル=\"%v\"\n", todo.ID, todo.Title)

		case "2":
			todos := todoManager.GetAll()
			if len(todos) == 0 {
				fmt.Println("TODOはありません")
				continue
			}
			fmt.Println("===== TODO一覧 =====")
			for _, todo := range todos {
				status := "[ ]"
				if todo.Completed {
					status = "[x]"
				}
				fmt.Printf("%d: %s %s (作成日時: %s)\n", todo.ID, status, todo.Title, todo.CreatedAt.Format("2006-01-02 15:04:05"))
			}

		case "3":
			fmt.Print("更新するTODOのIDを入力してください: ")
			scanner.Scan()
			idStr := scanner.Text()
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("IDは数値で入力してください")
				continue
			}

			todo, ok := todoManager.GetByID(id)
			if !ok {
				fmt.Println("IDに一致するTODOが見つかりません")
				continue
			}

			fmt.Printf("現在のタイトル: %s\n", todo.Title)
			fmt.Print("新しいタイトルを入力してください: ")
			scanner.Scan()
			newTitle := scanner.Text()
			if newTitle == "" {
				newTitle = todo.Title
			}

			completed := todo.Completed
			fmt.Printf("完了状態 (現在: %t) [y/n]: ", todo.Completed)
			scanner.Scan()
			completedStr := strings.ToLower(scanner.Text())
			if completedStr == "y" {
				completed = true
			} else if completedStr == "n" {
				completed = false
			}

			updatedTodo, _ := todoManager.Update(id, newTitle, completed)
			fmt.Printf("更新しました: ID=%d, タイトル=\"%s\", 完了=%t\n", updatedTodo.ID, updatedTodo.Title, updatedTodo.Completed)

		case "4":
			fmt.Print("削除するTODOのIDを入力してください: ")
			scanner.Scan()
			idStr := scanner.Text()
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("IDは数値で入力してください")
				continue
			}

			if todoManager.Delete(id) {
				fmt.Printf("ID=%dのTODOを削除しました\n", id)
			} else {
				fmt.Println("指定されたIDのTODOが見つかりません")
			}

		case "5":
			fmt.Print("完了/未完了にするTODOのIDを入力してください: ")
			scanner.Scan()
			idStr := scanner.Text()
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("IDは数値で入力してください")
				continue
			}

			todo, ok := todoManager.ToggleCompleted(id)
			if ok {
				status := "未完了"
				if todo.Completed {
					status = "完了"
				}
				fmt.Printf("ID=%dのTODOを%sに設定しました\n", id, status)
			} else {
				fmt.Println("指定されたIDのTODOが見つかりません")
			}

		case "0":
			fmt.Println("終了します")
			return

		default:
			fmt.Println("無効な選択肢です")
		}
	}
}
