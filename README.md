# adventofcode2023
Advent of Code 2023 Solutions

Current Coverage Result: [![codecov](https://codecov.io/gh/sfmiked/adventofcode2023/graph/badge.svg?token=63MRY6V67B)](https://codecov.io/gh/sfmiked/adventofcode2023)

## setup steps

### command line
```
# create go.mod file
go mod init github.com/miketzian/adventofcode2023

# add testify dependency
go get github.com/stretchr/testify
```

### devcontainer

1. In vscode - Add Dev Container Configuration Files
2. Choose 'Go' option
3. No options selected

### github 

1. In actions, choose [https://docs.github.com/en/actions/automating-builds-and-tests/building-and-testing-go](go action)
2. In test workflow step, add coverage option (by default no coverage report)
3. Login to codecov*, get the secret, create the repository secret
4. Add the codecov workflow action

*note! codecov.io wants to know our email address. They're not allowed to know this, so I had to create another dummy github account with an email address they're allowed to know. 