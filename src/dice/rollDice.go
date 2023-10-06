package dice

import (
    "math/rand"
)

// Individual functions for rolling each type of die
func RollD4() int {
    return rand.Intn(4) + 1
}

func RollD6() int {
    return rand.Intn(6) + 1
}

func RollD8() int {
    return rand.Intn(8) + 1
}

func RollD10() int {
    return rand.Intn(10) + 1
}

func RollD100() int {
    tens := rand.Intn(10) * 10  // This will give random tens place (0, 10, 20, ..., 90)
    return tens
}

func RollD12() int {
    return rand.Intn(12) + 1
}

func RollD20() int {
    return rand.Intn(20) + 1
}

type DiceRollResults struct {
    D4   int
    D6   int
    D8   int
    D10  int
    D100 int
    D12  int
    D20  int
}

// Parent function to roll all DnD dice
func RollAllDice() DiceRollResults {
    return DiceRollResults{
        D4:   RollD4(),
        D6:   RollD6(),
        D8:   RollD8(),
        D10:  RollD10(),
        D100: RollD100(),
        D12:  RollD12(),
        D20:  RollD20(),
    }
}