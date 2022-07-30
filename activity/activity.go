package activity

import (
	"context"
	"fmt"
)

func PrintActivity1(ctx context.Context, student interface{}) error {
	fmt.Println("this is PrintActivity1 func. student:", student)
	return nil
}

func PrintActivity2(ctx context.Context, student interface{}) error {
	fmt.Println("this is PrintActivity2 func. student:", student)
	return nil
}
