package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const seconds = 10
const delay time.Duration = seconds * 1000 * time.Millisecond

var longSeparator = strings.Repeat("------", 25)

const initialBalance float32 = 300.00

var accountMovements = [...]float32{
	40.50,
	-20.72,
	89.01,
	60.22,
	-96.02,
	7.42,
	73.79,
	96.24,
	76.81,
	-32.98,
	24.76,
	17.48,
	-11.5,
	94.34,
	43.06,
	56.37,
	-29.41,
	-51.32,
	-53.99,
	-41.77,
	-33.04,
	-49.18,
	84.89,
	-49.91,
	34.41,
	-8.27,
	-83.34,
	-84.29,
	1.71,
	-75.72,
	-73.64,
	-97,
	-94.03,
	-15.48,
	0.95,
	18.02,
	70.39,
	-6.15,
	-80.81,
	11,
	-3.26,
	81.66,
	58.06,
	96.79,
	-97.54,
	-97.97,
	45.18,
	4.75,
	-84.8,
	75.03,
	31.05,
	-83.56,
	15.4,
	-43.92,
	16.62,
	66.13,
	38.67,
	-80.34,
	21.61,
	42.01,
	-94.07,
}

// Global variables subject to the effect of a race condition
var accountBalance = initialBalance
var position int = 0

func main() {

	rand.Seed(time.Now().UnixNano())

	var strSlice00 = make([]string, 0, 100)
	var strSlice01 = make([]string, 0, 100)
	var strSlice02 = make([]string, 0, 100)

	showRaceConditionDefinition()
	pauseAndPrompt()

	showInitialInfo()
	pauseAndPrompt()

	showStage01Info()
	pauseAndPrompt()

	fmt.Println("=========================================================")
	fmt.Printf("Initial Balance [%.2f].\n", accountBalance)
	fmt.Println("=========================================================")

	applyAccountMovements("MAIN Goroutine", len(accountMovements), &strSlice00, false)

	fmt.Println("=========================================================")
	fmt.Printf("Resulting Balance [%.2f].\n", accountBalance)
	fmt.Println("=========================================================")

	pauseAndPrompt()
	showStage02Info()
	pauseAndPrompt()

	// Resetting values to its initial state
	accountBalance = initialBalance
	position = 0

	fmt.Println("\n=========================================================")
	fmt.Printf("RESETTING Position[%d] and Initial Account Balance[%.2f]\n", position, accountBalance)
	fmt.Println("=========================================================\n")

	fmt.Println("=========================================================")
	fmt.Printf("Initial Balance [%.2f].\n", accountBalance)
	fmt.Println("=========================================================")

	go applyAccountMovements("Goroutine #01", len(accountMovements), &strSlice01, true)
	go applyAccountMovements("Goroutine #02", len(accountMovements), &strSlice02, true)

	fmt.Printf("Main Thread is currently Awaiting [%d] seconds to let all goroutines to safely end ...\n", seconds)
	time.Sleep(delay)

	fmt.Println("=========================================================")
	fmt.Printf("Resulting Balance [%.2f].\n", accountBalance)
	fmt.Println("=========================================================")

	pauseAndPrompt()
	showStage03Info()
	pauseAndPrompt()

	printHistoryLines("Goroutine #01", &strSlice01)
	pauseAndPrompt()
	printHistoryLines("Goroutine #02", &strSlice02)

	fmt.Println("=========================================================")
	fmt.Printf("Resulting Balance [%.2f].\n", accountBalance)
	fmt.Println("=========================================================")
	pauseAndPrompt()

	fmt.Println("GOODBYE !!!")
}

func applyAccountMovements(threadName string, totalNumberOfAccountMovements int, history *[]string, isRandomDelayActive bool) {
	threadBalance := accountBalance // Initial value at start and required to keep scope

	*history = append(*history, longSeparator)
	logLine := fmt.Sprintf(
		"\t\t--> Entering Goroutine [%s]. Theoretical initial Account Balance[%.2f] (%s)",
		threadName,
		threadBalance,
		time.Now().Format("2006-01-02 15:04:05.000000000"))
	fmt.Printf("\n%s\n\n", logLine)
	*history = append(*history, logLine)
	*history = append(*history, longSeparator)

	for position < totalNumberOfAccountMovements {
		threadBalance = accountBalance // May have been changed by other concurrent thread
		accountMovement := float32(0.0)
		// as position maybe changed by anoter thread
		// it is convenient test before again
		if position < totalNumberOfAccountMovements {
			accountMovement = accountMovements[position]
			threadUpdatedBalance := threadBalance + accountMovement

			logLine = fmt.Sprintf("[%s].\tTheroretical Current Balance[%.2f].\tAccount Movement[#%d][%.2f].\tUpdated Balance[%.2f] (%s)",
				threadName,
				threadBalance,
				position,
				accountMovement,
				threadUpdatedBalance,
				time.Now().Format("2006-01-02 15:04:05.000000000"))
			fmt.Printf("%s\n", logLine)
			*history = append(*history, logLine)

			// Update shared accessible global variable with
			// the "theoretically" proper updated value
			accountBalance = threadUpdatedBalance
			threadBalance = threadUpdatedBalance
			position++
			if isRandomDelayActive {
				time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond) // to obtain a more random interleaving
			}
		} else {
			break
		}
	}
	*history = append(*history, longSeparator)
	logLine = fmt.Sprintf(
		"\t\t--> Exiting from Goroutine [%s]. Theoretical resulting Account Balance[%.2f] (%s)",
		threadName,
		threadBalance,
		time.Now().Format("2006-01-02 15:04:05.000000000"))
	fmt.Printf("\n%s\n\n", logLine)
	*history = append(*history, logLine)
	*history = append(*history, longSeparator)
}

func printHistoryLines(threadName string, history *[]string) {
	fmt.Println("\n========================================================================")
	fmt.Printf("Showing now the event's history registered at Goroutine [%s]\n", threadName)
	fmt.Println("========================================================================")
	for _, line := range *history {
		fmt.Println(line)
	}
}

func showRaceConditionDefinition() {
	fmt.Println("\n============================================================================================================")
	fmt.Println("WELCOME !!!")
	fmt.Println("This program aims to show an example of the appearance of a RACE CONDITION. What does it means?")
	fmt.Println("\nRACE CONDITION DEFINITION:")
	fmt.Println("A 'RACE CONDITION' (a.k.a 'RACE HAZARD') in software systems is the condition where the system's substantive ")
	fmt.Println("behavior is dependent on the sequence or timing of interleaving actions.")
	fmt.Println("Race conditions arise in software when an application depends on the sequence or timing of processes or threads")
	fmt.Println("for it to operate properly:")
	fmt.Println("  ==> Interleaving is non-deterministic, consequently the outcome depends on non-deterministic ordering.")
	fmt.Println("  ==> Races occur due to communication (access to shared data/variables) among concurrently running threads")
	fmt.Println("============================================================================================================\n")
}

func showInitialInfo() {
	fmt.Println("\n============================================================================================================")
	fmt.Println("LET'S SHOW THAT WITH AN EXAMPLE !!!")
	fmt.Println("This program aims to show an example of the emergence of a RACE CONDITION when more than one thread is being")
	fmt.Println("run concurrently and having simultaneous read and write access to a zone of shared (variable) data.")
	fmt.Println("In the proposed example, a situation is simulated where there is an initial balance of an account and also a")
	fmt.Println("list of pending transactions to be applied on. The process to be followed is to increase the position in the")
	fmt.Println("list (cursor), take the value of the correspondig movement in the list and update the current account balance")
	fmt.Println("with its value. When the entire list of pending movements has been traversed, the resulting final balance is")
	fmt.Println("to be shown.")
	fmt.Println("If you carry out the test to execute this program successively, you will be able to observe that when the")
	fmt.Println("RACE CONDITION (badly managed) occurs, different results will be obtained each time: ")
	fmt.Println("--> UNDETERMINISTIC OUTCOME")
	fmt.Println("============================================================================================================\n")
}

func showStage01Info() {
	fmt.Println("\n============================================================================================================")
	fmt.Println("STAGE #01:")
	fmt.Println("First it will be shown a completely correct behaviour when a single thread is used (the main Goroutine).")
	fmt.Println("Once applied all movements the resulting balance should be [120.30].")
	fmt.Println("============================================================================================================\n")
}

func showStage02Info() {
	fmt.Println("\n============================================================================================================")
	fmt.Println("STAGE #02:")
	fmt.Println("Now we will reset both values: the Cursor (position) of the list of pending movements and the value of")
	fmt.Println("the initial Balance to their starting values")
	fmt.Println("Then two threads (Goroutines) will be launched and execute concurrently creating a RACE CONDITION as:")
	fmt.Println("1. The two threads will access concurrently trying to increase the shared value of the cursor of the")
	fmt.Println("   movement list and updating the shared value of the balance with the corresponding movement of the list")
	fmt.Println("2. Running this way final will drive to a totally different final Balance result each time")
	fmt.Println(" ")
	fmt.Println("Note: now the main thread will be waiting sleeping during a reasonable time to permit all other threads")
	fmt.Println("to succesfully conclude its tasks. DISCLAIMER: There are better way to sync. threads but they have not")
	fmt.Println("shown/explained at this point in the course.")
	fmt.Println("============================================================================================================\n")
}

func showStage03Info() {
	fmt.Println("\n============================================================================================================")
	fmt.Println("STAGE #03:")
	fmt.Println("Each thread has been writing in a log its activity showing at every moment in time the values that 'see'")
	fmt.Println("and update the shared variables: Cursor and Current balance.")
	fmt.Println("Now it's time to see how they were behaving individually.")
	fmt.Println("============================================================================================================\n")
}

func pauseAndPrompt() {
	in := bufio.NewScanner(os.Stdin)
	fmt.Print("Press ENTER to continue.")
	in.Scan()
}
