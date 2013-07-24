package main

import "fmt"
import "strings"
import "os"
import "os/exec"
import "bufio"

func init() {
}

func main() {
    var err error
    const AYS_PROCEED string = "Are you sure you want to run this?  Type 'yes' to proceed: "

    // Pull all the command arguments off the stack and reassemble them into the final command we're going to run.
    commandToRun := strings.TrimSpace(strings.Join(os.Args[1:], " "))
    if len(commandToRun) == 0 {
        fmt.Println("[ays] You must specify a command to run!")
        os.Exit(1)
    }

    // We have our command, so spit it back out to the user and ask them to confirm running it.
    fmt.Printf("[ays] Command entered: `%s`\n", commandToRun)
    fmt.Print(AYS_PROCEED)

    bufferedStdin := bufio.NewReader(os.Stdin)
    for {
        // Get their answer and check it.
        answer, err := bufferedStdin.ReadString('\n')
        if err != nil {
            fmt.Printf("[ays] Caught an error while reading user input: %s", err.Error())
            os.Exit(1)
        }

        if strings.TrimSpace(answer) != "yes" {
            fmt.Print(AYS_PROCEED)
            continue
        }

        // They answered yes, so let's go run the command.
        break
    }

    // Now run the command!
    fmt.Printf("[ays] Running command `%s`...\n", commandToRun)

    commandPieces := strings.Split(commandToRun, " ")
    if len(commandPieces) == 0 {
        fmt.Printf("[ays] Couldn't get the pieces of the command!  Input string: %s\n", commandToRun)
        os.Exit(1)
    }

    command := exec.Command(commandPieces[0], commandPieces[1:]...)
    command.Stdin = os.Stdin
    command.Stdout = os.Stdout
    command.Stderr = os.Stderr

    err = command.Run()
    if err != nil {
        fmt.Printf("[ays] Caught an error while running '%s': %s", commandToRun, err.Error())
        os.Exit(1)
    }
}
