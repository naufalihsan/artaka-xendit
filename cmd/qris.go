/*
Copyright © 2020 NAUFAL IHSAN <naufal.ihsan21@gmail.com>

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
	"fmt"
	"github.com/joho/godotenv"
	"github.com/naufalihsan/artaka-xendit/qris"
	"github.com/skip2/go-qrcode"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// qrisCmd represents the qris command
var qrisCmd = &cobra.Command{
	Use:   "qris",
	Short: "Quick Response Code Indonesian Standard",
	Long: `Quick Response Code Indonesian Standard atau biasa disingkat QRIS
adalah penyatuan berbagai macam QR dari berbagai Penyelenggara Jasa Sistem Pembayaran (PJSP)
menggunakan QR Code.`,
	Run: func(cmd *cobra.Command, args []string) {

		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		secretKey := os.Getenv("SECRET_KEY")
		client := qris.New(secretKey)

		create, _ := cmd.Flags().GetBool("create")

		if create {
			createQRISData := qris.CreateQRISParams{
				ExternalID:  "c4n5ky-Stud1o",
				Type:        qris.QRISTypeDYNAMIC,
				CallbackURL: "https://naufalihsan.co.id/callback",
				Amount:      99000,
			}
			resp, _ := client.QRIS.CreateQRIS(&createQRISData)

			err := qrcode.WriteFile(resp.QRString, qrcode.Medium, 256, "qris.png")
			if err != nil {
				log.Fatal("Error generate file")
			}

		}

		get, _ := cmd.Flags().GetBool("get")

		if get {
			getQRISParam := qris.GetQRISParam{
				ExternalID: "c4n5ky-Stud1o",
			}
			resp, _ := client.QRIS.GetQRIS(&getQRISParam)
			fmt.Print(resp.QRString)
		}

		simulate, _ := cmd.Flags().GetBool("simulate")

		if simulate {
			paymentParam := qris.PaymentParam{
				ExternalID: "c4n5ky-Stud1o",
				Amount:     99000,
			}
			resp, _ := client.QRIS.SimulatePayment(&paymentParam)
			fmt.Print(resp.Status)
		}
	},
}

func init() {
	rootCmd.AddCommand(qrisCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// qrisCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// qrisCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	qrisCmd.PersistentFlags().BoolP("create", "c", false, "Generate QRIS Code")
	qrisCmd.PersistentFlags().BoolP("get", "g", false, "Get QRIS By External ID")
	qrisCmd.PersistentFlags().BoolP("simulate", "s", false, "Simulate Payment")
}
