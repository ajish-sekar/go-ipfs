package auth

import (
	"fmt"
	"net/http"
	"strings"
	"bytes"
	"strconv"
	"crypto/rand"
	"encoding/base64"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	contracts "github.com/ipfs/go-ipfs/contracts"
)

type authMiddleware struct{}

var randVals  = make(map[string]string)

func Handler() *authMiddleware {
	return &authMiddleware{}
}

func unauthorized(w http.ResponseWriter, headers map[string][]string) {
	for k, v := range headers {
		w.Header()[k] = v
	}
	w.WriteHeader(http.StatusForbidden)
	fmt.Fprintf(w, "Unauthorized User\n")
}

func sendRandString(w http.ResponseWriter, s string, headers map[string][]string) {
	for k, v := range headers {
		w.Header()[k] = v
	}
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprintf(w,s)
}


func GenerateRandomBytes(n int) ([]byte, error) {
    b := make([]byte, n)
    _, err := rand.Read(b)
    // Note that err == nil only if we read len(b) bytes.
    if err != nil {
        return nil, err
    }

    return b, nil
}

func GenerateRandomString(s int) (string, error) {
    b, err := GenerateRandomBytes(s)
    return base64.URLEncoding.EncodeToString(b), err
}


func (_ *authMiddleware) Handle(w http.ResponseWriter, r *http.Request, headers map[string][]string) bool {

	if(r.Method != http.MethodGet){
		return true
	}

	host := r.Host

	randString, ok := randVals[host]

	if !ok {
		randVal,err := GenerateRandomString(32)

		if err!= nil {
			fmt.Println("Random Generator Failed")
			unauthorized(w,headers)
			return false
		}
		randVals[host] = randVal
		sendRandString(w,randVal,headers)
		return false
	}

	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		delete(randVals,host)
		unauthorized(w,headers)
		return false
	}


	urlPaths := strings.Split(r.URL.Path,"/")

	pathLen := len(urlPaths)

	if urlPaths[pathLen-2] != "ipfs"{
		return true
	}

	fileHash := urlPaths[pathLen-1]	

	fmt.Printf("Auth: %s\n",authHeader)
	authValues := strings.Split(authHeader,",")
	fmt.Println("Split:",authValues)
	fileAddressString := ""
	userAddressString := ""
	msgHash := ""
	signature := ""

	for _, element := range authValues {
		tokens := strings.Split(element," ")
		fmt.Printf("%s -> %s\n",tokens[0],tokens[1])
		if(tokens[0] == "SCCF"){
			fileAddressString = tokens[1]
		}else if(tokens[0] == "SCCU"){
			userAddressString = tokens[1]
		}else if(tokens[0] == "SCCH"){
			msgHash = tokens[1]
		}else if(tokens[0] == "SCCS"){
			signature = tokens[1]
		}
	}

	if(fileAddressString == "" || userAddressString == "" || msgHash == "" || signature == ""){
		fmt.Println("Missing credentials")
		delete(randVals,host)
		unauthorized(w,headers)
		return false
	}

	// 2 for the 2 /
	messageLength := len(fileAddressString) + len(userAddressString) + len(randString) + 2

	data := []byte("\x19Ethereum Signed Message:\n" + strconv.Itoa(messageLength) + fileAddressString + "/" + userAddressString + "/" + randString)

	hash := crypto.Keccak256Hash(data)

	hashBytes, err := hexutil.Decode(msgHash)

	if err != nil {
		fmt.Println("Cannot decode hash ",err)
		delete(randVals,host)
		unauthorized(w,headers)
		return false
	}

	hashMatches := bytes.Equal(hashBytes,hash.Bytes())

	if !hashMatches {
		fmt.Println("Invalid Hash")
		delete(randVals,host)
		unauthorized(w,headers)
		return false
	}

	fileAddress := common.HexToAddress(fileAddressString)
	userAddress := common.HexToAddress(userAddressString)


	

	sigBytes, err := hexutil.Decode(signature)

	if err != nil {
		fmt.Println("Cannot decode signature ",err)
		delete(randVals,host)
		unauthorized(w,headers)
		return false
	}

	sigPublicKeyECDSA, err := crypto.SigToPub(hashBytes, sigBytes)

	if err != nil {
		fmt.Println("Cannot recover key ",err)
		delete(randVals,host)
		unauthorized(w,headers)
		return false
	}

	sigAddress := crypto.PubkeyToAddress(*sigPublicKeyECDSA)

	fmt.Println("SigAddress",sigAddress.Hex())
	fmt.Println("EthAddress",userAddress.Hex())

	sigAddressBytes := sigAddress.Bytes()
	userAddressBytes := userAddress.Bytes();


	matches := bytes.Equal(sigAddressBytes,userAddressBytes)

	if !matches {
		fmt.Println("Invalid Signature")
		delete(randVals,host)
		unauthorized(w,headers)
		return false
	}




	client,err := ethclient.Dial("http://localhost:7545")

	if err != nil {
		fmt.Println("Unable to connect to client ",err)
		delete(randVals,host)
		unauthorized(w,headers)
		return false
	}

	instance, err := contracts.NewFile(fileAddress,client)

	if err!=nil {
		fmt.Println("Unable to load file contract ",err)
		delete(randVals,host)
		unauthorized(w,headers)
		return false
	}

	isHashMatch, err := instance.IsFileHash(nil,fileHash)

	if err!=nil {
		fmt.Println("Unable to access contract ",err)
		delete(randVals,host)
		unauthorized(w,headers)
		return false
	}

	if !isHashMatch{
		fmt.Println("File hash does not match")
		delete(randVals,host)
		unauthorized(w,headers)
		return false
	}

	allow, err := instance.HasAccess(nil,userAddress)

	if err!=nil {
		fmt.Println("Unable to access contract ",err)
		delete(randVals,host)
		unauthorized(w,headers)
		return false
	}

	if !allow{
		delete(randVals,host)
		unauthorized(w,headers)
		return false
	}



	

	// authToken := strings.SplitN(authHeader, " ", 2)
	// if len(authToken) < 2 {
	// 	unauthorized(w)
	// 	return false
	// }

	// validate token here. token can be an access token,
	// bearer, or jwt. commonly, the value of `authType`
	// would be "Bearer", or "JWT".
	// authType, token := authToken[0], authToken[1]

	delete(randVals,host)
	return true
}