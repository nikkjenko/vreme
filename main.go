package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/nikkjenko/vreme/vreme"
	"github.com/urfave/cli/v3"
)

func main() {

	cmd := &cli.Command{
		Name:  "vreme",
		Usage: "Dodajanje in spremljanje vremenskih meritev",
		Flags: []cli.Flag{
			&cli.Float64Flag{
				Name:  "minTemp",
				Value: -50.0,
				Usage: "najnižja dovoljena temperatura meritve v °C",
			},
			&cli.Float64Flag{
				Name:  "maxTemp",
				Value: 50.0,
				Usage: "najvišja dovoljena temperatura meritve v °C",
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Value: 3,
				Usage: "najmanjše število meritev za izračun statistike",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			minTemp := c.Float64("minTemp")
			maxTemp := c.Float64("maxTemp")
			stMeritev := c.Int("stMeritev")

			vreme.DodajPostajo("LJ01", "Prva", "Ljubljana")
			vreme.DodajPostajo("LJ02", "Druga", "Ljubljana")
			vreme.DodajPostajo("LJ03", "Tretja", "Ljubljana")

			vreme.DodajMeritev("LJ01", float64(15.2), minTemp, maxTemp)
			vreme.DodajMeritev("LJ01", float64(13), minTemp, maxTemp)
			vreme.DodajMeritev("LJ01", float64(15), minTemp, maxTemp)
			vreme.DodajMeritev("LJ02", float64(13.01), minTemp, maxTemp)
			vreme.DodajMeritev("LJ02", float64(14), minTemp, maxTemp)
			vreme.DodajMeritev("LJ02", float64(20), minTemp, maxTemp)
			vreme.DodajMeritev("LJ03", float64(24.17), minTemp, maxTemp)
			vreme.DodajMeritev("LJ03", float64(25), minTemp, maxTemp)
			vreme.DodajMeritev("LJ03", float64(26), minTemp, maxTemp)

			vreme.DodajMeritev("LJ01", 15.2, minTemp, maxTemp)
			vreme.DodajMeritev("LJ01", 13.0, minTemp, maxTemp)
			vreme.DodajMeritev("LJ01", 15.0, minTemp, maxTemp)

			vreme.DodajMeritev("LJ02", 20.12, minTemp, maxTemp)
			vreme.DodajMeritev("LJ04", 20.12, minTemp, maxTemp)
			vreme.DodajMeritev("LJ02", 61.0, minTemp, maxTemp)

			vreme.DodajMeritev("LJ03", 24.17, minTemp, maxTemp)
			vreme.DodajMeritev("LJ03", 25.0, minTemp, maxTemp)
			vreme.DodajMeritev("LJ03", 26.0, minTemp, maxTemp)

			fmt.Println()
			vreme.IzpisMeritev()
			fmt.Println()
			vreme.IzpisPorocila(stMeritev)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}
