package main

import (
	"fmt"
	"log"
	"os"

	openai "github.com/sashabaranov/go-openai"
	"github.com/urfave/cli"
)

var appVersion = "0.1.0"
var appName = "openai-helper"
var defaultModel = openai.GPT3Dot5Turbo

func main() {

	app := &cli.App{
		Name:     appName,
		HelpName: appName,
		Usage:    "Cli to ask OpenAI questions",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:     "gpt35",
				Usage:    "Use GPT-3.5x Turbo",
				Required: false,
			},
			&cli.BoolFlag{
				Name:     "gpt4",
				Usage:    "Use GPT-4",
				Required: false,
			},
		},
		Commands: []cli.Command{
			{
				Name:  "rewrite",
				Usage: "rewrite text in a way You like",
				Flags: []cli.Flag{},
				Action: func(c *cli.Context) error {
					if c.Bool("gpt35") {
						defaultModel = openai.GPT3Dot5Turbo
					} else if c.Bool("gpt4") {
						defaultModel = openai.GPT4TurboPreview
					}

					prompts, errPrompts := loadPrompts("")
					if errPrompts != nil {
						fmt.Printf("Cannot load prompts: %v\n", errPrompts)
						return nil
					}
					fmt.Printf("Rewriting: %s\n", prompts.Rewrite)

					content, errContent := getContentViaVim()
					if errContent != nil {
						fmt.Printf("Cannot get content: %v\n", errContent)
						return nil
					}
					fmt.Printf("Content: %s\n", content)

					response, errResponse := sendOpenAIRequest(prompts.Rewrite, content, defaultModel)
					if errResponse != nil {
						fmt.Printf("Response error: %v\n", errResponse)
						return nil
					}
					fmt.Printf("\n\n%s\n\n", response)
					return nil
				},
			},
			{
				Name:  "version",
				Usage: "show version and exit",
				Flags: []cli.Flag{},
				Action: func(c *cli.Context) error {
					fmt.Printf("%s version %s\n", appName, appVersion)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
