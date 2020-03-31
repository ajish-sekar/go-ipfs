package auth

import (
	"fmt"
	"net/http"
	"strings"
	"reflect"
	"bytes"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	contracts "github.com/ipfs/go-ipfs/contracts"
)

type authMiddleware struct{}

func Handler() *authMiddleware {
	return &authMiddleware{}
}

func unauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprintf(w, "Unauthorized User\n")
}

func (_ *authMiddleware) Handle(w http.ResponseWriter, r *http.Request) bool {

	if(r.Method != http.MethodGet){
		return true
	}

	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		unauthorized(w)
		return false
	}

	fmt.Printf("Type: %s\n",reflect.TypeOf(authHeader))
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
		unauthorized(w)
		return false
	}

	fileAddress := common.HexToAddress(fileAddressString)
	userAddress := common.HexToAddress(userAddressString)


	hashBytes, err := hexutil.Decode(msgHash)

	if err != nil {
		fmt.Println("Cannot decode hash ",err)
		unauthorized(w)
		return false
	}

	sigBytes, err := hexutil.Decode(signature)

	if err != nil {
		fmt.Println("Cannot decode signature ",err)
		unauthorized(w)
		return false
	}

	sigPublicKeyECDSA, err := crypto.SigToPub(hashBytes, sigBytes)

	if err != nil {
		fmt.Println("Cannot recover key ",err)
		unauthorized(w)
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
		unauthorized(w)
		return false
	}




	client,err := ethclient.Dial("http://localhost:7545")

	if err != nil {
		fmt.Println("Unable to connect to client ",err)
		unauthorized(w)
		return false
	}

	instance, err := contracts.NewFile(fileAddress,client)

	if err!=nil {
		fmt.Println("Unable to load file contract ",err)
		unauthorized(w)
		return false
	}

	allow, err := instance.HasAccess(nil,userAddress)

	if err!=nil {
		fmt.Println("Unable to access contract ",err)
		unauthorized(w)
		return false
	}

	if !allow{
		unauthorized(w)
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

	return true
}