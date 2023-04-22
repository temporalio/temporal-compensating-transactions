package app

import (
	"context"
	"fmt"
)

func GetBowl(ctx context.Context) error {
	fmt.Println("Getting bowl")
	return nil
}

func PutBowlAway(ctx context.Context) error {
	fmt.Println("Putting bowl away")
	return nil
}

func AddCereal(ctx context.Context) error {
	fmt.Println("Adding cereal")
	return nil
}

func PutCerealBackInBox(ctx context.Context) error {
	fmt.Println("Putting cereal back in box")
	return nil
}

func AddMilk(ctx context.Context) error {
	fmt.Println("Adding milk")
	return nil
}
