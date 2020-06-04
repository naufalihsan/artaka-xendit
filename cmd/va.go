/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/xendit/xendit-go/client"
	"github.com/xendit/xendit-go/virtualaccount"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// vaCmd represents the va command
var vaCmd = &cobra.Command{
	Use:   "va",
	Short: "Virtual Account",
	Long: `Virtual Account adalah rekening tidak nyata (virtual).
Virtual Account itu sendiri berisikan nomor ID customer yang dibuat Bank (sesuai permintaan perusahaan)
untuk melakukan transaksi.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		secretKey := os.Getenv("SECRET_KEY")
		artakaClient := client.New(secretKey)

		banks, _ := cmd.Flags().GetBool("banks")

		if banks {
			resp, _ := artakaClient.VirtualAccount.GetAvailableBanks()
			fmt.Print(resp)
		}

		create, _ := cmd.Flags().GetBool("create")

		if create {

			closedPayment := true
			createVAFixData := virtualaccount.CreateFixedVAParams{
				ExternalID:     "va-" + time.Now().String(),
				BankCode:       "BCA",
				Name:           "Naufal Ihsan Pratama",
				IsClosed:       &closedPayment,
				ExpectedAmount: 99000,
				IsSingleUse:    &closedPayment,
			}

			resp, _ := artakaClient.VirtualAccount.CreateFixedVA(&createVAFixData)
			marshal, _ := json.Marshal(resp)
			fmt.Println(string(marshal))
		}

		simulate, _ := cmd.Flags().GetBool("simulate")

		if simulate {
			getVAFixData := virtualaccount.GetFixedVAParams{
				ID: "5ed9030128269e5eb097ef14",
			}

			resp, _ := artakaClient.VirtualAccount.GetFixedVA(&getVAFixData)
			marshal, _ := json.Marshal(resp)
			fmt.Println(string(marshal))

			if resp.Status == "ACTIVE" {

				type VACallback struct {
					ID            string `json:"id"`
					OwnerID       string `json:"owner_id"`
					ExternalID    string `json:"external_id"`
					AccountNumber string `json:"account_number"`
					MerchantCode  string `json:"merchant_code"`
					Amount        string `json:"amount"`
					BankCode      string `json:"bank_code"`
				}

				callbackURL := fmt.Sprintf(
					"https://api.xendit.co/callback_virtual_accounts/external_id=%s/simulate_payment", resp.ExternalID)
				callbackPayload, _ := json.Marshal(VACallback{
					ID:            resp.ID,
					OwnerID:       resp.OwnerID,
					ExternalID:    resp.ExternalID,
					AccountNumber: resp.AccountNumber,
					MerchantCode:  resp.MerchantCode,
					Amount:        "99000",
					BankCode:      resp.BankCode,
				})

				payload := bytes.NewReader(callbackPayload)

				req, _ := http.NewRequest("POST", callbackURL, payload)
				req.Header.Set("Authorization", os.Getenv("AUTHORIZATION"))
				req.Header.Set("Content-Type", "application/json")

				response, _ := http.DefaultClient.Do(req)
				defer response.Body.Close()

				var res map[string]interface{}

				body, _ := ioutil.ReadAll(response.Body)
				if err := json.Unmarshal(body, &res); err != nil {
					panic(err)
				}
				fmt.Println(res)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(vaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// vaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// vaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	vaCmd.PersistentFlags().BoolP("banks", "b", false, "Get Available Banks")
	vaCmd.PersistentFlags().BoolP("create", "c", false, "Create Virtual Account")
	vaCmd.PersistentFlags().BoolP("simulate", "s", false, "Simulate Payment Virtual Account")

}
