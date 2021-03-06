package main

import (
	"errors"
	"flag"
	"fmt"
)

var tileDefinitions = []string{
	"1m",
	"2m",
	"3m",
	"4m",
	"5m",
	"6m",
	"7m",
	"8m",
	"9m",
	"1s",
	"2s",
	"3s",
	"4s",
	"5s",
	"6s",
	"7s",
	"8s",
	"9s",
	"1p",
	"2p",
	"3p",
	"4p",
	"5p",
	"6p",
	"7p",
	"8p",
	"9p",
	"E",
	"S",
	"W",
	"N",
	"H",
	"R",
	"C",
}

type Tile struct {
	name     string
	tileType string // bamboo, circle, character or honours
}

func NewTiles(raw string) ([]*Tile, error) {
	var tiles = []*Tile{} // 手牌 最終的に13個の配列になる
	var targetCharacters string = raw
	for i := 0; i < 13; i++ {
		// 残り0文字ならエラー出す
		if len(targetCharacters) == 0 {
			return nil, errors.New("与えられた文字列が短すぎます")
		}

		var newTile *Tile = nil
		// 字牌
		for _, item := range []string{"C", "H", "R", "E", "W", "N", "S"} {
			if targetCharacters[:1] == item {
				newTile = &Tile{
					name:     item,
					tileType: "honours",
				}
				break
			}
		}
		if newTile != nil {
			tiles = append(tiles, newTile)
			targetCharacters = targetCharacters[1:] // 字牌1枚分を落とす
			continue
		}

		// 残り1文字以下ならエラー出す
		if len(targetCharacters) < 2 {
			return nil, errors.New("与えられた文字列が短すぎます")
		}

		// マンズ
		for _, item := range []string{"1m", "2m", "3m", "4m", "5m", "6m", "7m", "8m", "9m"} {
			if targetCharacters[:2] == item {
				newTile = &Tile{
					name:     item,
					tileType: "character",
				}
				break
			}
		}
		if newTile != nil {
			tiles = append(tiles, newTile)
			targetCharacters = targetCharacters[2:] // 数牌1枚分を落とす
			continue
		}

		// ピンズ
		for _, item := range []string{"1p", "2p", "3p", "4p", "5p", "6p", "7p", "8p", "9p"} {
			if targetCharacters[:2] == item {
				newTile = &Tile{
					name:     item,
					tileType: "circle",
				}
				break
			}
		}
		if newTile != nil {
			tiles = append(tiles, newTile)
			targetCharacters = targetCharacters[2:] // 数牌1枚分を落とす
			continue
		}

		// ソーズ
		for _, item := range []string{"1s", "2s", "3s", "4s", "5s", "6s", "7s", "8s", "9s"} {
			if targetCharacters[:2] == item {
				newTile = &Tile{
					name:     item,
					tileType: "bamboo",
				}
				break
			}
		}
		if newTile != nil {
			tiles = append(tiles, newTile)
			targetCharacters = targetCharacters[2:] // 数牌1枚分を落とす
			continue
		}

		// どれにも当てはまらなかったらその時点でerror
		return nil, errors.New("指定された文字列が不正です" + targetCharacters[0:1])
	}

	if len(targetCharacters) > 0 {
		return nil, errors.New("指定された文字列が長すぎます")
	}

	return tiles, nil
}

func (t *Tile) getType() (string, error) {
	for _, name := range []string{"C", "H", "R", "E", "W", "N", "S"} {
		if name == t.name {
			return "honor", nil
		}
	}

	typeString := t.name[1:]

	switch typeString {
	case "m":
		return "manzu", nil
	case "s":
		return "souzu", nil
	case "p":
		return "pinzu", nil
	default:
		return "", errors.New("不正なtypeです")
	}
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("引数を指定してください.")
		return
	}
	var tiles []*Tile
	var err error
	tiles, err = NewTiles(flag.Args()[0])

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	tileType := tiles[0].tileType

	fmt.Println(tileType)
	fmt.Println("tileType")
}
