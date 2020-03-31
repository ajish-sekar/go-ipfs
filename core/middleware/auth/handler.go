package auth

import (
	"fmt"
	"net/http"
	"strings"
	"reflect"
	"github.com/ethereum/go-ethereum/common"
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
	userKeyString := ""

	for _, element := range authValues {
		tokens := strings.Split(element," ")
		fmt.Printf("%s -> %s\n",tokens[0],tokens[1])
		if(tokens[0] == "SCCF"){
			fileAddressString = tokens[1]
		}else if(tokens[0] == "SCCU"){
			userKeyString = tokens[1]
		}
	}

	client,err := ethclient.Dial("http://localhost:7545")

	if err != nil {
		fmt.Println("Unable to connect to client ",err)
		unauthorized(w)
		return false
	}

	fileAddress := common.HexToAddress(fileAddressString)
	userAddress := common.HexToAddress(userKeyString)

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