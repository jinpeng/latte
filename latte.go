package main

import (
	"os"
  "github.com/codegangsta/cli"
)

func train() {
  
}

func test() {
  
}

func device_query() {
  
}

func benchmark() {
  
}

func main() {
  app := cli.NewApp()
  app.Name = "latte"
  app.Usage = "train/test/device_query/time Deep Neural Network."
  app.Author = "Dong Jinpeng"
  app.Email = "jinpeng.dong@gmail.com"
  app.Commands = []cli.Command{
    {
      Name:      "train",
      ShortName: "r",
      Usage:     "train or finetune a model",
      Action: func(c *cli.Context) {
        println("train or finetune a model: ", c.Args().First())
        train()
      },
    },
    {
      Name:      "test",
      ShortName: "t",
      Usage:     "score a model",
      Action: func(c *cli.Context) {
        println("score a model: ", c.Args().First())
        test()
      },
    },
    {
      Name:      "device_query",
      ShortName: "d",
      Usage:     "show GPU diagnostic information",
      Action: func(c *cli.Context) {
        println("show GPU diagnostic information: ", c.Args().First())
        device_query()
      },
    },
    {
      Name:      "benchmark",
      ShortName: "b",
      Usage:     "benchmark model execution time",
      Action: func(c *cli.Context) {
        println("benchmark model execution time: ", c.Args().First())
        benchmark()
      },
    },
  }

  app.Run(os.Args)
}