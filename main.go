package main

func main() {
	GuildedRose()
}

type Item struct {
	Name    string
	SellIn  int
	Quality int
}

var Items = []Item{
	Item{"+5 Dexterity Vest", 10, 20},
	Item{"Aged Brie", 2, 0},
	Item{"Elixir of the Mongoose", 5, 7},
	Item{"Sulfuras, Hand of Ragnaros", 0, 80},
	Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
	Item{"Conjured Mana Cake", 3, 6},
}

func GuildedRose() {
	for _, item := range Items {
		if item.Name == "Aged Brie" || item.Name == "Backstage passes to a TAFKAL80ETC concert" {
			item.SellIn -= 1
			if item.Quality < 50 {
				item.Quality += 1
			}
			if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
				if item.SellIn <= 10 && item.SellIn > 5 && item.Quality < 50 {
					item.Quality += 1
				} else if item.SellIn <= 5 && item.Quality < 49 {
					item.Quality += 2
				} else if item.SellIn <= 5 {
					item.Quality += 1
				} else if item.SellIn < 0 {
					if item.Quality != 0 {
						item.Quality = 0
					}
				}
			}
		} else {
			if item.SellIn == 0 || item.SellIn < 0 { //after the sellin date
				item.SellIn -= 1
				if item.Quality >= 2 && item.Name != "Sulfuras, Hand of Ragnaros" { //quality never negative
					item.Quality -= 2
				} else if item.Quality >= 1 && item.Name != "Sulfuras, Hand of Ragnaros" { //quality never negative
					item.Quality -= 1
				}
			} else { //before the SellIn date
				item.SellIn -= 1
				if item.Quality >= 1 && item.Name != "Aged Brie" && item.Name != "Backstage passes to a TAFKAL80ETC concert" && item.Name != "Sulfuras, Hand of Ragnaros" {
					if item.Quality > 0 { //quality never negative
						item.Quality -= 1
					}
				} else if item.Name == "Conjured Mana Cake" {
					if item.Quality > 1 { //quality never negative
						item.Quality -= 2
					} else if item.Quality > 0 {
						item.Quality -= 1
					}
				}
			}
		}
	}
}
