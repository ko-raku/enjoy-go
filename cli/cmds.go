package cli

func RegisterCommands() {
	RootCmd.AddCommand(GrayscaleCmd)
	RootCmd.AddCommand(TextFromImageCmd)
}
