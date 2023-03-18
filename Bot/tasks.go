package bot

import (
	TA "discord_bot/TasksAPI"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"google.golang.org/api/tasks/v1"
)

// tlget
func tasklistget() *discordgo.MessageSend {
	items, err := TA.InitUser().Tasklists.List().Do()

	if err != nil {
		panic(err)
	}
	var taskArray []string

	for _, i := range items.Items {
		taskArray = append(taskArray, fmt.Sprintf("%s :%s", i.Id, i.Title))
	}
	return &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title:       "TaskLists",
			Description: strings.Join(taskArray, "\n"),
			Color:       1752220,
		},
	}
}

// tget
func taskget(contents []string) *discordgo.MessageSend {
	if len(contents) == 1 {
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "Error",
				Description: "コマンドの構文が間違っています。詳細は!help tget\n```!tget -tl:<tasklist_id>```",
			},
		}
	}

	if include(contents, "-tl") == "" {
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "Error",
				Description: "コマンドの構文が間違っています。詳細は!help tget\n```-tl:<tasklist_id>``` ",
			},
		}
	}
	items, err := TA.InitUser().Tasks.List(include(contents, "-tl")).Do()

	if err != nil {
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       include(contents, "-tl"),
				Description: "タスクリストが見つかりませんでした。Tasklists not found",
				Color:       2002060,
			},
		}
	}

	if len(items.Items) > 0 {
		var taskArray []string
		for _, i := range items.Items {
			if i.Status != "completed" {
				taskArray = append(taskArray, fmt.Sprintf("%s :%s", i.Id, i.Title))
			}
		}

		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       include(contents, "-tl"),
				Description: strings.Join(taskArray, "\n"),
				Color:       1752220,
			},
		}

	} else {

		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       include(contents, "-tl"),
				Description: "タスクが見つかりませんでした。Tasks lists not found",
				Color:       2002060,
			},
		}
	}

}

// tdget
func taskdetailget(contents []string) *discordgo.MessageSend {
	if len(contents) <= 2 {
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "Error",
				Description: "コマンドの構文が間違っています。詳細は!help tdget\n```!tdget -tl:<tasklist_id> -ts:<task_id>```",
			},
		}
	}

	if include(contents, "-tl") == "" {
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "Error",
				Description: "コマンドの構文が間違っています。詳細は!help tdget\n```-tl:<tasklist_id>```",
			},
		}
	}
	if include(contents, "-ts") == "" {
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "Error",
				Description: "コマンドの構文が間違っています。詳細は!help tdget\n```-ts:<task_id>```",
			},
		}
	}

	items, err := TA.InitUser().Tasks.Get(include(contents, "-tl"), include(contents, "-ts")).Do()

	if err != nil {
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       include(contents, "-ts"),
				Description: "タスクが見つかりませんでした。Taskl not found",
				Color:       2002060,
			},
		}
	}

	return &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title:       items.Title,
			Description: items.Notes,
			Footer: &discordgo.MessageEmbedFooter{
				Text: items.Due,
			},
			Color: 1752220,
		},
	}
}

// tlpost
func tasklistinsert(contents []string) *discordgo.MessageSend {
	if len(contents) == 1 {
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "Error",
				Description: "コマンドの構文が間違っています。詳細は!help tlpost\n```!tlpost <tasklist_name>```",
			},
		}
	}

	newlist := tasks.TaskList{Title: contents[1]}
	items, err := TA.InitUser().Tasklists.Insert(&newlist).Do()

	if err != nil {
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "Error",
				Description: "タスクリストを追加できませんでした",
			},
		}
	}
	return &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title:       "Status:sccess",
			Description: fmt.Sprintf("%s : %s", items.Id, items.Title),
			Color:       1752220,
		},
	}
}

// tpost
func tasksinsert(contents []string) *discordgo.MessageSend {

	if len(contents) == 1 {
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "Error",
				Description: "コマンドの構文が間違っています。詳細は!help tpost\n```!tpost -tl:<tasklist_id> -ti:<title> option[-d:due -n:notes]```",
			},
		}
	}

	if include(contents, "-tl") == "" {
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "Error",
				Description: "コマンドの構文が間違っています。詳細は!help tpost\n-tl:<tasklist_id>",
			},
		}
	}

	if include(contents, "-ti") == "" {
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "Error",
				Description: "コマンドの構文が間違っています。詳細は!help tpost\n-ti:<title>",
			},
		}
	}
	task, _ := newtasks(contents)

	items, err := TA.InitUser().Tasks.Insert(include(contents, "-tl"), &task).Do()

	if err != nil {
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "Error",
				Description: "タスクリストの挿入に失敗しました",
			},
		}
	}

	return &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title:       "Status:success",
			Description: fmt.Sprintf("ID:%s\ntitle:%s\nnotes:%s\nDue:%s\nを追加しました", items.Id, items.Title, items.Notes, items.Due),
			Color:       1752220,
		},
	}
}

// tlput
func tasklistupdate(contents []string) *discordgo.MessageSend {
	if len(contents) == 1 {
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "Error",
				Description: "コマンドの構文が間違っています。詳細は!help tlput\n```!tlput -otl:<tasklist_id> -ntl:<new-name>```",
			},
		}
	}

	if include(contents, "-otl") == "" {
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "Error",
				Description: "コマンドの構文が間違っています。詳細は!help tlput\n-otl:<tasklist_id>",
			},
		}
	}

	if include(contents, "-ntl") == "" {
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "Error",
				Description: "コマンドの構文が間違っています。詳細は!help tlput\n-ntl:<new-name>",
			},
		}
	}
	newlist := tasks.TaskList{Title: include(contents, "-ntl")}
	items, err := TA.InitUser().Tasklists.Patch(include(contents, "-otl"), &newlist).Do()

	if err != nil {
		fmt.Println(err)
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "Error",
				Description: "タスクリリストの変更に失敗しました",
			},
		}
	}
	return &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title:       "Status:sccess",
			Description: fmt.Sprintf("%s : %s", items.Id, items.Title),
			Color:       1752220,
		},
	}
}

// tput
func tasksupdate(contents []string) *discordgo.MessageSend {

	if len(contents) > 1 {
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "Error",
				Description: "コマンドの構文が間違っています。詳細は!help tput\n```!tpost -tl:<tasklist_id> -ts:<task_id> -ti:<title> option[ -d:due -n:notes ]```",
			},
		}
	}

	if include(contents, "-tl") == "" {
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "Error",
				Description: "コマンドの構文が間違っています。詳細は!help tput\n-tl:<tasklist_id>",
			},
		}
	}

	if include(contents, "-ts") == "" {
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "Error",
				Description: "コマンドの構文が間違っています。詳細は!help tput\n-ts:<task_id>",
			},
		}
	}

	task, _ := newtasks(contents)
	fmt.Println(include(contents, "-tl"), include(contents, "-ts"))
	items, err := TA.InitUser().Tasks.Patch(include(contents, "-tl"), include(contents, "-ts"), &task).Do()

	if err != nil {
		fmt.Println(err)
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "Error",
				Description: "タスクの変更に失敗しました",
			},
		}
	}
	return &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title:       "Status:success",
			Description: fmt.Sprintf("ID:%s\ntitle:%s\nnotes:%s\nDue:%s\nに変更しました", items.Id, items.Title, items.Notes, items.Due),
			Color:       1752220,
		},
	}
}

// tcomp
func taskcomplete(contents []string) *discordgo.MessageSend {

	if len(contents) < 4 {
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "Error",
				Description: "コマンドの構文が間違っています。詳細は!help tcomp\n```!tcomp -tl:<tasklist_id> -t:<task_id> option[ -ti:<title> -d:<due> -n:<notes> 'どれか1つ以上必須'] ```",
			},
		}
	}

	if include(contents, "-tl") == "" {
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "Error",
				Description: "コマンドの構文が間違っています。詳細は!help tcomp\n-tl:<tasklist_id>",
			},
		}
	}

	if include(contents, "-ts") == "" {
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "Error",
				Description: "コマンドの構文が間違っています。詳細は!help tcomp\n-t:<task_id>",
			},
		}
	}

	task := tasks.Task{
		Status: "completed",
	}

	items, err := TA.InitUser().Tasks.Patch(include(contents, "-tl"), include(contents, "-ts"), &task).Do()

	if err != nil {
		fmt.Println(err)
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "Error",
				Description: "タスクの変更に失敗しました",
			},
		}
	}

	return &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title:       "Status:success",
			Description: fmt.Sprintf("ID:%s\ntitle:%s\nnotes:%s\nDue:%s\nのタスクを完了しました。", items.Id, items.Title, items.Notes, items.Due),
			Color:       1752220,
		},
	}
}

// tldel
func tasklistdelete(contents []string) *discordgo.MessageSend {
	if len(contents) == 1 {
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "Error",
				Description: "コマンドの構文が間違っています。詳細は!help tldel\n```!tldel -tl:<tasklist_id>```",
			},
		}
	}

	if include(contents, "-tl") == "" {
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "Error",
				Description: "コマンドの構文が間違っています。詳細は!help tldel\n-tl:<tasklist_id>",
			},
		}
	}
	err := TA.InitUser().Tasklists.Delete(include(contents, "-tl")).Do()

	if err != nil {
		fmt.Println(err)
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "Error",
				Description: "タスクの削除に失敗しました",
			},
		}
	}

	return &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title:       "Status:success",
			Description: fmt.Sprintf("ID:%s\nを削除しました。", include(contents, "-tl")),
			Color:       1752220,
		},
	}
}

// tdel
func taskdelete(contents []string) *discordgo.MessageSend {
	if len(contents) == 1 {
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "Error",
				Description: "コマンドの構文が間違っています。詳細は!help tdel\n```!tdel -tl:<tasklist_id> -ts:<task_id>```",
			},
		}
	}

	if include(contents, "-tl") == "" {
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "Error",
				Description: "コマンドの構文が間違っています。詳細は!help tdel\n-tl:<tasklist_id>",
			},
		}
	}

	if include(contents, "-ts") == "" {
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "Error",
				Description: "コマンドの構文が間違っています。詳細は!help tdel\n-t:<task_id>",
			},
		}
	}
	err := TA.InitUser().Tasks.Delete(include(contents, "-tl"), include(contents, "-ts")).Do()

	if err != nil {
		fmt.Println(err)
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "Error",
				Description: "タスクの削除に失敗しました",
			},
		}
	}

	return &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title:       "Status:success",
			Description: fmt.Sprintf("ID:%s\nを削除しました。", include(contents, "-ts")),
			Color:       1752220,
		},
	}
}

func newtasks(contents []string) (tasks.Task, error) {
	var task tasks.Task
	if include(contents, "-ti") != "" {
		task = tasks.Task{Title: include(contents, "-ti")}
	}
	if include(contents, "-d") != "" {
		task = tasks.Task{Due: include(contents, "-d")}
	}
	if include(contents, "-n") != "" {
		task = tasks.Task{Notes: include(contents, "-n")}
	}

	return task, nil

}
func include(array []string, target string) string {
	for _, item := range array {
		if strings.Contains(item, target) {
			return strings.Replace(strings.Replace(item, target, "", -1), ":", "", 1)
		}
	}
	return ""
}
