package ai

import "fmt"

func ProcessPrompt(prompt Prompt) (string, error) {
	response := fmt.Sprintf("Processed prompt: %s", prompt.text)

	return response, nil
}
