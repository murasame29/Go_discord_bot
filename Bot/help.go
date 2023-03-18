package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func helpcommand(contents []string) *discordgo.MessageSend {
	if len(contents) == 1 {
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title: "コマンド一覧",
				Description: fmt.Sprintf(`
				タスクリストの取得　: !tlget
				タスクの取得　　　　: !tget -tl
				タスクの詳細を取得　: !tdget -tl -ts
				タスクリストの追加　: !tlpost <新しいタスクリスト名>
				タスクの追加　　　　: !tpost -tl -ti オプション[-d -n]
				タスクリストの更新　: !tlput -otl -ntl
				タスクの更新　　　　: !tput -tl -ts オプション[-ti -d -n ※どれか１項目は必須]
				タスクの完了　　　　: !tcomp -tl -ts
				タスクリストの削除　: !tldel -tl
				タスクの削除　　　　: !tdel -tl -ts
				`),
			},
		}
	}
	switch contents[1] {
	case "tlget":
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title: "!tlget",
				Description: fmt.Sprintf(`
				概要　　 : タスクリストをすべて取得します。
				コマンド : !tlget
				`),
			},
		}
	case "tget":
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title: "!tget",
				Description: fmt.Sprintf(`
				概要　　 : 指定したタスクリスト内のタスクをすべて取得します
				コマンド : !tdget -tl
				引数　　 : -tl:<タスクリストID>
				`),
			},
		}
	case "tdget":
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title: "!tdget",
				Description: fmt.Sprintf(`
				概要　　 : 指定したタスクの詳細を表示します。
				コマンド : !tget -tl -ts
				引数　　 : -tl:<タスクリストID>
				　　　　 : -ts:<タスクID>
				`),
			},
		}
	case "tlpost":
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title: "!tlpost",
				Description: fmt.Sprintf(`
				概要　　 : 新しいタスクリストを追加
				コマンド : !tlpost <新しいタスクリスト名>
				`),
			},
		}
	case "tpost":
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title: "!tpost",
				Description: fmt.Sprintf(`
				概要　　 : 指定したタスクリストに新しいタスクを追加する
				コマンド : !tpost -tl -ti オプション[-d:期限 -n:ノート]
				引数　　 : -tl:<タスクリストID>
				　　　　 : -ti:<タイトル>
				　　　　 : -d:期限
				　　　　 : -n:ノート
				`),
			},
		}
	case "tlput":
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title: "!tlput",
				Description: fmt.Sprintf(`
				概要　　 : 指定したタスクリストの名前を変更する
				コマンド : !tlput -otl -ntl
				引数　　 : -otl:<タスクリストID>
				　　　　 : -ntl:<新しいタスクリスト名>
				`),
			},
		}
	case "tput":
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title: "!tput",
				Description: fmt.Sprintf(`
				概要　　 : 指定したタスクの内容を変更する
				コマンド : !tput -tl -ts オプション[-ti -d -n ※どれか１項目は必須]
				引数　　 : -tl:<タスクリストID>
				　　　　 : -ts:<タスクID>
				　　　　 : -ti:<タイトル>
				　　　　 : -d:<期限>
				　　　　 : -n:<ノート>
				`),
			},
		}
	case "tcomp":
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title: "!tcomp",
				Description: fmt.Sprintf(`
				概要　　 : 指定したタスクを完了にする。
				コマンド : !tcomp -tl -ts
				引数　　 : -tl:<タスクリストID>
				　　　　 : -ts:<タスクID>
				`),
			},
		}
	case "tldel":
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title: "!tldel",
				Description: fmt.Sprintf(`
				概要　　 : 指定したタスクを完了にする。
				コマンド : !tldel -tl
				引数　　 : -tl:<タスクリストID>
				`),
			},
		}
	case "tdel":
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title: "!tdel",
				Description: fmt.Sprintf(`
				概要　　 : 指定したタスクを完了にする。
				コマンド : !tdel -tl -ts
				引数　　 : -tl:<タスクリストID>
				　　　　 : -ts:<タスクID>
				`),
			},
		}
	default:
		return &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "コマンド一覧",
				Description: fmt.Sprintf(`そんなコマンドないで:)`),
			},
		}
	}
}
