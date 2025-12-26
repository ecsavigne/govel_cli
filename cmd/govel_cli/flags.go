package main

var (
	executeFlag string
)

func initFlags() {
	rootCmd.PersistentFlags().StringVarP(&executeFlag, "create", "c", "", "create project govel. if project_name is not specified, ecs_govel name will be used")
	rootCmd.PersistentFlags().StringVarP(&executeFlag, "info", "i", "", "show project info. if project_name is not specified, ecs_govel name will be used")
}
