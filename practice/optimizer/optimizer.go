package optimizer

import (
	"fmt"
	"practice/pizza"
	"sort"
)

type likesDislikes struct {
	Likes    int
	Dislikes int
}

type orderedIngredient struct {
	Name string
	Val  int
}

func Solve(clients []pizza.Client) []string {
	ingredients := buildIndexOfIngridients(clients)
	byLike, byDislike := order(ingredients)
	pizza := optimize(byLike, byDislike, ingredients)
	return pizza
}

func optimize(like []orderedIngredient, dislike []orderedIngredient, ing map[string]likesDislikes) []string {
	il, id := 0, 0
	pizza := make(map[string]struct{})
	banned := make(map[string]struct{})
	for il < len(like) && id < len(dislike) {
		// todo: >= vs >
		if like[il].Val >= dislike[id].Val {
			_, ok := banned[like[il].Name]
			if ok {
				il++
				continue
			}

			pizza[like[il].Name] = struct{}{}
			il++
		} else {
			_, ok := pizza[dislike[id].Name]
			if ok {
				id++
				continue
			}

			banned[dislike[id].Name] = struct{}{}
			id++
		}
	}

	if id >= len(dislike) {
		for k := range ing {
			_, inPizza := pizza[k]
			_, inBan := banned[k]
			if !inPizza && !inBan {
				fmt.Println("heuristic worked")
				pizza[k] = struct{}{}
			}
		}
	}

	return flatten(pizza)
}

func flatten(p map[string]struct{}) []string {
	slice := make([]string, 0, len(p))
	for k := range p {
		slice = append(slice, k)
	}

	return slice
}

func order(ingredients map[string]likesDislikes) ([]orderedIngredient, []orderedIngredient) {
	byLike := make([]orderedIngredient, 0, len(ingredients))
	byDislike := make([]orderedIngredient, 0, len(ingredients))

	for k, v := range ingredients {
		byLike = append(byLike, orderedIngredient{
			Name: k,
			Val:  v.Likes,
		})

		byDislike = append(byDislike, orderedIngredient{
			Name: k,
			Val:  v.Dislikes,
		})
	}

	sort.SliceStable(byLike, func(i, j int) bool {
		return byLike[i].Val > byLike[j].Val
	})
	sort.SliceStable(byDislike, func(i, j int) bool {
		return byDislike[i].Val > byDislike[j].Val
	})

	return byLike, byDislike
}

func buildIndexOfIngridients(clients []pizza.Client) map[string]likesDislikes {
	index := make(map[string]likesDislikes)

	for _, c := range clients {
		for _, ingredient := range c.Likes {
			x, ok := index[ingredient]
			if !ok {
				index[ingredient] = likesDislikes{
					Likes:    1,
					Dislikes: 0,
				}
				continue
			}

			x.Likes++
			index[ingredient] = x
		}

		for _, ingredient := range c.Dislikes {
			x, ok := index[ingredient]
			if !ok {
				index[ingredient] = likesDislikes{
					Likes:    0,
					Dislikes: 1,
				}
				continue
			}

			x.Dislikes++
			index[ingredient] = x
		}
	}

	return index
}
