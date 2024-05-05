## Checklist

### Data Storage
- [x] set up a SQLite database
    - [x] open SQLite DB
    - [x] add task
    - [x] delete task
    - [x] edit task
    - [x] get tasks

### Making a CLI with [Cobra][cobra]
- [x] add CLI
    - [x] add task
    - [x] delete task
    - [x] edit task
    - [x] get tasks

### Add a little... *Je ne sais quoi*
- [x] print to table layout with [Lip Gloss][lipgloss]
- [x] print to Kaban layout with [Lip Gloss][lipgloss]

## Project layout

`db/db.go` - here we create our custom `task` struct and our data layer.

`utils/utils.go` - utility functions handle our setup including opening a
database and setting the data path for our application.

`cmd/` - this is where we do all of our Cobra commands and setup for our CLI.

[lipgloss]: https://github.com/charmbracelet/lipgloss
[cobra]: https://github.com/spf13/cobra
