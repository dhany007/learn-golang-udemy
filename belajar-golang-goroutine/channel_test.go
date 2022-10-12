package belajargolanggoroutines

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	// membuat channel
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		// mengirimkan isi channel
		channel <- "Dhany Putra Aritonang"
		fmt.Println("Selesai mengirimkan data")
	}()

	// menerima data dari channel
	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Dhany Putra Aritonang"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	close(channel)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Dhany Putra Aritonang"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
	close(channel)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Dhany"
		channel <- "Putra"
		channel <- "Aritonang"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}

	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}

}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu data ...")
		}
		if counter == 2 {
			break
		}
	}

}

func TestRaceConditionChannel(t *testing.T) {
	x := 0

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				x = x + 1
			}
		}()
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Counter", x) // x pasti bukan 100000, karena ada variabel yang disharing
}

// menghindari race condition
func TestSyncMutexChannel(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock() // dilock dulu, kemudian menunggu sampai lock nya lepas
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Counter", x) // x = 100000
}

type BankAccount struct {
	RWMutex sync.RWMutex // lock unlock read write
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock() // ngelock write
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock() // unlock write
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock() // ngelock read
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestSyncRWMutexChannel(t *testing.T) {
	account := BankAccount{}

	for i := 1; i <= 100; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println("Final Balance", account.GetBalance())
}

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock", user2.Name)
	user2.Change(-amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Dhany",
		Balance: 2000,
	}

	user2 := UserBalance{
		Name:    "Budi",
		Balance: 1000,
	}

	go Transfer(&user1, &user2, 100)
	go Transfer(&user2, &user1, 200)

	time.Sleep(5 * time.Second)

	fmt.Println("User 1", user1.Name, " Balance", user1.Balance)
	fmt.Println("User 2", user2.Name, " Balance", user2.Balance)
}
