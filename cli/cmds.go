package cli

func RegisterCommands() {
	RootCmd.AddCommand(GrayscaleCmd)
	// 他のサブコマンドもここで登録できます。
}
