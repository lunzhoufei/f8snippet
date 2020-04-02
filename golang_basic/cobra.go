
// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "util",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		id = strings.Split(commaSepId, ",")

		// rawStdout, _ := syscall.Dup(syscall.Stdout)
		// rawStderr, _ := syscall.Dup(syscall.Stderr)
		syscall.Stdout, _ = syscall.Dup(syscall.Stdout)
		syscall.Stderr, _ = syscall.Dup(syscall.Stderr)
		// syscall.Close(1)
		srfCtx = srfs.NewContext(context.Background())
		initCallee()
		// os.Stdout = os.NewFile(uintptr(rawStdout), "mystandout")
		// os.Stderr = os.NewFile(uintptr(rawStderr), "mystanderr")
		syscall.Dup2(syscall.Stdout, 1)
		syscall.Dup2(syscall.Stderr, 2)

		os.Stdout = os.NewFile(uintptr(syscall.Stdout), "mystandout")
		os.Stderr = os.NewFile(uintptr(syscall.Stdin), "mystanderr")
		color.Output = os.Stdout
	},
}

