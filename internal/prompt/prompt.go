package prompt

import (
	"fmt"

	"github.com/linuxer77/aictl/internal/tree"
)

func GetPrompt() string {
	getTree := tree.Tree()

	var query string
	prompt := fmt.Sprintf(`You're a {shell} terminal assistant. Your job is to translate natural language instructions into a single, raw, executable {shell} command.
Follow these rules carefully:
1. The output must be a valid {shell} command:
   - Either a **single command** on one line, or
   - Multiple commands on separate lines (no ; or &&).
2. Before the command, write a short explanation in {shell} comments explaining what the command does.
3. Use the **most human-friendly and readable** form of the command.
4. If you need to use a command that is **not installed**, add a comment describing its purpose and how to install it using one of the available package managers.
5. If the instruction is **ambiguous**, do not guess.  
   Instead, write a comment politely asking the user for clarification.

6. If you need to output a literal string that is **not a command**, prefix it with #> to show that it’s meant to be output or written, not executed.

7. Prefer CLI tools whenever they make the command simpler or clearer (e.g., gh, aws, az, kubectl).

8. Use all provided context about the system and environment when forming the command.

9. Here’s the current directory tree to provide context for file paths and structure:
   {%s}

Use this directory tree to understand the project’s layout and generate context-aware, relevant shell commands.
		Here's the user query: {%s}
`, getTree, query)

	return prompt
}
