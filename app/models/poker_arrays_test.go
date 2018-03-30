package models

import (
	"fmt"
	"testing"
)

func TestPokerArr_SortForPlay(t *testing.T) {
	pokers := []Poker{
		{Value: 6, HuaSe: 2},
		{Value: 5, HuaSe: 3},
		{Value: 6, HuaSe: 1},
		{Value: 7, HuaSe: 4},
		{Value: 5, HuaSe: 1},
		{Value: 5, HuaSe: 2},
		{Value: 6, HuaSe: 3},
		{Value: 3, HuaSe: 1},
	}
	PokerArr(pokers).SortForPlay()

	pokers2 := []Poker{
		{Value: 7, HuaSe: 2},
		{Value: 7, HuaSe: 1},
	}
	PokerArr(pokers2).SortForPlay()

}

func TestPokerArr_GetType(t *testing.T) {
	pokers := []Poker{
		{Value: 6, HuaSe: 2},
		{Value: 5, HuaSe: 3},
		{Value: 6, HuaSe: 1},
		{Value: 7, HuaSe: 4},
		{Value: 5, HuaSe: 1},
		{Value: 5, HuaSe: 2},
		{Value: 6, HuaSe: 3},
		{Value: 3, HuaSe: 1},
	}
	fmt.Println(PokerArr(pokers).GetType().TypeInt == PokersTypeForSerialTrebleWithSingle)

	pokers2 := []Poker{
		{Value: 3},
		{Value: 6},
		{Value: 2},
		{Value: 5},
		{Value: 4},
	}
	fmt.Println(PokerArr(pokers2).GetType().TypeInt == PokersTypeForSerialSingle)
	pokers3 := []Poker{
		{Value: 3},
		{Value: 3},
		{Value: 6},
		{Value: 6},
		{Value: 2},
		{Value: 2},
		{Value: 5},
		{Value: 5},
		{Value: 5},
		{Value: 4},
		{Value: 4},
	}
	fmt.Println(PokerArr(pokers3).GetType().TypeInt)

}
