package main

import (
	"fmt"
	"log"
	"net"

	"github.com/briand787b/rfs/scratch/secrets"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	db *sqlx.DB
)

type ipPrac struct {
	ID      int
	IPBytes []byte

	LocalIP     net.IP
	NetworkCIDR *net.IPMask
}

func (ip *ipPrac) initialize() {
	ip.LocalIP = net.ParseIP(string(ip.IPBytes))

	if ip.NetworkCIDR == nil {
		return
	}

	// DO LATER
	// ip.NetworkCIDR = net.ParseCIDR()
}

func main() {
	m, err := secrets.PromptSecrets("user", "dbname", "password")
	if err != nil {
		log.Fatal(err)
	}

	db = sqlx.MustConnect("postgres",
		fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable",
			m["user"],
			m["dbname"],
			m["password"],
		),
	)

	fmt.Println("connected to db!")
	defer db.Close()

	r := db.QueryRowx(`
		SELECT
			id AS ID,
			local_ip AS IPBytes,
			network_cider AS NetworkCIDR
		FROM
			ip_prac
		WHERE id = 1;`,
	)

	if err := r.Err(); err != nil {
		log.Fatal(err)
	}

	var ipp ipPrac
	if err := r.StructScan(&ipp); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ipPrac: %+v\n", ipp)
	fmt.Printf("ipp's ip: %s\n", ipp.LocalIP)

	fmt.Printf("initializing...")

	fmt.Printf("ipPrac: %+v\n", ipp)
	fmt.Printf("ipp's ip: %s\n", ipp.LocalIP)

	// bs, err := ipp.LocalIP.MarshalText()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("ipp's ip: %s\n", bs)

	// ipp.LocalIP = []byte("192.168.1.1")
	// fmt.Printf("ipp's ip: %s\n", []byte(ipp.LocalIP))
	// bs, err := ipp.LocalIP.MarshalText()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("ipp's ip: %s\n", bs)
}
