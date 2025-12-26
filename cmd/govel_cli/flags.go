package main

var (
	executeFlag string
	versionFlag bool
)

func initFlags() {
	rootCmd.PersistentFlags().StringVarP(&executeFlag, "create", "c", "", "create project govel. if project_name is not specified, ecs_govel name will be used")
	rootCmd.PersistentFlags().StringVarP(&executeFlag, "info", "i", "", "show project info. if project_name is not specified, ecs_govel name will be used")
	rootCmd.PersistentFlags().BoolVarP(&versionFlag, "version", "v", false, "version of govel_cli and ecs_govel")
}
