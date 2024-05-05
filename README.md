## Checklist

### Data Storage
- [x] set up a SQLite database
    - [x] open SQLite DB
    - [x] add task
    - [x] delete task
    - [x] edit task
    - [x] get tasks

### Making a CLI with [Cobra][cobra]
- [s] add CLI
    - [s] add task
    - [s] delete task
    - [s] edit task
    - [s] get tasks

### Add a little... *Je ne sais quoi*
- [s] print to table layout with [Lip Gloss][lipgloss]
- [s] print to Kaban layout with [Lip Gloss][lipgloss]

## Project layout

`db.go` - here we create our custom `task` struct and our data layer.

`main.go` - our main file handles our initial setup including opening a
database and setting the data path for our application.

`cmds.go` - this is where we do all of our Cobra commands and setup for our CLI.

[lipgloss]: https://github.com/charmbracelet/lipgloss
[cobra]: https://github.com/spf13/cobra
