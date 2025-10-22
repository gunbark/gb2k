package cmd

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// rollCmd represents the roll command
var rollCmd = &cobra.Command{
	Use:   "roll",
	Short: "roll yer dice",
	Long: `roll dice separated by spaces or commas, eg:
		
	roll 2d20, 1d4
	roll 10d100

dice groups MUST be one NdM format, limited to N<10, M<100 
will display successes as 💥
will display critical failures as 💀`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rolling..")
		rollGroup(args)
	},
}

func init() {
	rootCmd.AddCommand(rollCmd)
	rand.Seed(time.Now().UnixNano())
}

func rollGroup(dice []string) {
	dice = validateDice(dice)
	var successes, failures int
	re := regexp.MustCompile(`^(\d)d((\d{1,2})|100)$`)
	for i, die := range dice {
		time.Sleep(500 * time.Millisecond)
		die = strings.TrimSuffix(die, ",")
		die = strings.ToLower(die)
		if re.MatchString(die) {
			fmt.Println(".. dice group " + die + ":")
			s, f := roll(die)
			successes += s
			failures += f
		} else {
			fmt.Println(".. dice group "+strconv.Itoa(i+1)+" was not formatted correctly:", die)
		}
	}
	printSuccess(successes)
	printFailure(failures)
}

func roll(die string) (int, int) {
	var successes, failures int
	numDice, _ := strconv.Atoi(strings.Split(die, "d")[0])
	diceSize, _ := strconv.Atoi(strings.Split(die, "d")[1])
	for i := 0; i < numDice; i++ {
		d := rand.Intn(diceSize) + 1
		fmt.Println("\t#" + strconv.Itoa(i+1) + " - d" + strconv.Itoa(diceSize) + " .. " + strconv.Itoa(d))
		if d >= 10 {
			successes += 2
		} else if d >= 6 {
			successes++
		} else if d == 1 {
			failures++
		}
	}
	return successes, failures
}

func validateDice(args []string) []string {
	var newArgs []string
	for _, arg := range args {
		arg = strings.TrimPrefix(arg, ",")
		arg = strings.TrimSuffix(arg, ",")
		for _, i := range strings.Split(arg, ",") {
			newArgs = append(newArgs, strings.ToLower(i))
		}
	}
	return newArgs
}

func printSuccess(successes int) {
	fmt.Println("successes:")
	var s string
	for i := 0; i < successes; i++ {
		s += "💥"
	}
	fmt.Println("\t" + s)
}

func printFailure(failures int) {
	fmt.Println("failures:")
	var s string
	for i := 0; i < failures; i++ {
		s += "💀"
	}
	fmt.Println("\t" + s)
}
