# OpenAI Helper

This is a command-line interface (CLI) for interacting with OpenAI's GPT-3.5 Turbo and GPT-4 models.

## Prerequisites

Please create `~/.config/prompts.json` file:

```json
{
  "rewrite": "Correct any grammar errors and make is simple for non-native English speakers to understand"
}
```

## Usage

The CLI provides two flags:

- `--gpt35`: Use this flag to interact with the GPT-3.5 Turbo model.
- `--gpt4`: Use this flag to interact with the GPT-4 model.

The CLI also provides a `rewrite` command, which rewrites text in a way you prefer.

Here's an example of how to use the CLI:

```bash
$ openai-helper --gpt35 rewrite 
```

will open Vim and allow you to add the text you want to rewrite. Once you save and exit Vim, the CLI will send your text to the selected model and return the rewritten text.
