package main

import "fmt"

func main() {
	user, err := getInvitedUserWithEmail(ctx, email)
	if err != nil {
		// 呼び出し先で発生したエラーをラップし、付与情報を付与して呼び出し元に返却
		fmt.Errorf("fail to get invited user with email(%s): %w", email, err)
	}
}
