package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "fibonachii",
	Short: "Fibonachii is fast and efficient fibonacci calculator",
	Run: func(cmd *cobra.Command, args []string) {
		// cmd.Help()

		// fmt.Println(worker.Work(10))

		// producer.Produce()
		// consumer.Consume()

		// fib := types.Fibonacci{
		// 	N: 10,
		// }

		// fmt.Println(producer.Produce(fib))
	},
	Version: "0.2.4-alpha.0",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetDefault("KafkaAddress", "kafka")
	viper.SetDefault("KafkaTopic", "fibonachii")
	viper.SetDefault("KafkaGroupID", "fibonachii")
	viper.SetDefault("KafkaPartition", 0)

	viper.SetDefault("RedisAddress", "redis:6379")
	viper.SetDefault("RedisPassword", "")
	viper.SetDefault("RedisDB", 0)
}
