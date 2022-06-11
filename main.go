package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/afrizal423/sopibot/bot"
	"github.com/afrizal423/sopibot/pengguna"
	"github.com/joho/godotenv"
)

func Concurrently(numOfQuotes int, u pengguna.User) {
	wg := sync.WaitGroup{}

	for i := 0; i < numOfQuotes; i++ {
		wg.Add(1)
		go func(idx int) {
			inputx := os.Getenv("URLBARANG")
			item := bot.Fetch_item_from_url(inputx, u.Csrf_token)
			fmt.Println()
			selected_model := os.Getenv("MODEL")
			modelnum, _ := strconv.Atoi(selected_model)
			// fmt.Print(selected_model)
			if !item.Is_flash_sale {
				if item.Upcoming_flash_sale.Name == "" {
					// f, _ := strconv.ParseInt(item.Upcoming_flash_sale.Start_time, 10, 64)
					fmt.Println("Menunggu flash sale")
					goto FINISH
				} else {
					fmt.Println("flash sale telah terlewat")
					goto FINISH
				}
			} else {
				fmt.Println()
				fmt.Println("********************************")
				fmt.Println("Proses menambahkan ke keranjang belanja")
				fmt.Println("********************************")

				if os.Getenv("BELANJA") != "sukses" {
					item_id := bot.Add_to_card(item, modelnum, u.Csrf_token)

					fmt.Println("✅Barang sudah CHECKOUT✅")
					bot.Checkout(item_id, item, modelnum, u.Csrf_token)
					f, err := os.OpenFile(".env", os.O_APPEND|os.O_WRONLY, 0600)
					if err != nil {
						panic(err)
					}
					defer f.Close()

					if _, err = f.WriteString("BELANJA=sukses\n"); err != nil {
						panic(err)
					}
				} else {
					fmt.Println("✅Barang sudah PERNAH di CHECKOUT✅")
					goto FINISH
				}
			}

		FINISH:
			// fmt.Println("********************************")
			fmt.Println("Selesai")
			// fmt.Println("********************************")
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func main() {
	loop := 2
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	fmt.Println("Mengambil informasi user ...")
	user := pengguna.Login()
	if user.Phone != "" && user.Email != "" {
		Concurrently(loop, user)
	} else {
		panic("Session udah habis")
	}
}
