package main

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"runtime"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	log.Println("mining started, gotta get that internet GOLD!")
	i := 0
	for {
		lookup("americaonline.com")
		lookup("photobucket.com")
		lookup("pandora.com")
		lookup("livejournal.com")
		lookup("ipchicken.com")
		lookup("pets.com")
		lookup("packardbell.com")
		lookup("albinoblacksheep.com")
		lookup("newgrouds.com")
		i++
		hash := newSHA1Hash()
		code := fmt.Sprintf("coin mined:%s", hash)
		d1 := []byte(code)
		filename := fmt.Sprintf("block.%d", i)
		err := ioutil.WriteFile(filename, d1, 0644)
		done := make(chan int)

		for i := 0; i < runtime.NumCPU(); i++ {
			go func() {
				for {
					select {
					case <-done:
						return
					default:
					}
				}
			}()
		}
		time.Sleep(time.Second * 60)
		close(done)
		url := "http://api.ipify.org?format=text" // we are using a pulib IP API, we're using ipify here, below are some others
		lookup("google.com")
		lookup("github.com")
		lookup("bhj3e9bm2.us-west-1.rds.amazonaws.com")
		lookup("bhj3e9bm2.us-east-1.rds.amazonaws.com")
		lookup("bhj333e9b3m2.us-west-2.rds.amazonaws.com")
		lookup("dhfgk3ckk.us-east-1.rds.amazonaws.com")
		lookup("mf012klg.us-east-1.rds.amazonaws.com")
		lookup("open.spotify.com")
		lookup("live.bbc.co.uk")
		fmt.Printf("mining ...\n")
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		ip, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Updating pool:%s\n", ip)
	}
}

func lookup(s string) {
	ip, _ := net.LookupHost(s)
	log.Println("fakeDNs:", s, ip)
}

func newSHA1Hash(n ...int) string {
	noRandomCharacters := 500

	if len(n) > 0 {
		noRandomCharacters = n[0]
	}

	randString := RandomString(noRandomCharacters)

	hash := sha1.New()
	hash.Write([]byte(randString))
	bs := hash.Sum(nil)

	return fmt.Sprintf("%x", bs)
}

var characterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = characterRunes[rand.Intn(len(characterRunes))]
	}
	return string(b)
}
