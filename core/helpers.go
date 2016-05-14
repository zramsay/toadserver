package core

import (
	"bytes"
	"fmt"

	log "github.com/Sirupsen/logrus"
	//"github.com/eris-ltd/common/go/common"
	"github.com/eris-ltd/common/go/ipfs"
)

// TODO clean up any unused
// -> figure out soln for this func
func CacheHashAll(hash string) error {

	//TODO handle errors to prevent getting here...
	log.Warn("Pinning hash to your local IPFS node")

	pinned, err := ipfs.PinToIPFS(hash, bytes.NewBuffer([]byte{}))
	if err != nil {
		log.WithField("=>", fmt.Sprintf("%s", err)).Warn("")
		return fmt.Errorf("error pinning to local IPFS node: %v\n", err)
	}
	log.WithField("=>", pinned).Warn("Hash pinned to you local node")

	// IPaddrs, _ := getTheNames() -> use mindy to get ipAddrs
	/*IPaddrs := os.Getenv("TOADSERVER_IPFS_NODES")
	IPs := strings.Split(IPaddrs, ",")
	for _, ip := range IPs {
		endpoint := fmt.Sprintf("http://%s:11113/cacheHash/%s", ip, hash)
		log.WithField("=>", endpoint).Warn("Pinning hash to remote IPFS/toadserver node:")
		_, err := http.Post(endpoint, "", nil)
		if err != nil {
			log.WithField("=>", endpoint).Warn("error making post request to:")
			log.Error(err)
			continue
			//TODO return err?
		}
	}*/
	return nil
}