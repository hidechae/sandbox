package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx := context.Background()

	// 途中でエラーが起きたときのキャンセル処理をするためにWithContextを使う
	// キャンセルしないなら、 `new(errgroup.Group)` で良い
	eg, ctx := errgroup.WithContext(ctx)

	// 並列数を設定する
	eg.SetLimit(3)

	for i := 1; i <= 10; i++ {
		i := i
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				// エラーが出た後は他のgoroutineはキャンセルされる
				return ctx.Err()
			default:
				// 正常時の処理
				fmt.Println(i)
				time.Sleep(1 * time.Second)

				// 5のときだけエラーを返す
				if i == 5 {
					// エラーを返すと、他のgoroutineはキャンセルされる
					return fmt.Errorf("error: %d", i)
				}
				return nil
			}
		})
	}

	// 処理が終わるのを待つ
	err := eg.Wait()
	if err != nil {
		panic(err)
	}
}
