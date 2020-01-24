package controller
import(
	"net/http"
	"github.com/daleksprinter/catechdojo/model"
	"github.com/daleksprinter/catechdojo/repository"
	"encoding/json"
	"github.com/google/uuid"
)

func CreateUserController(w http.ResponseWriter, r * http.Request){
	//parse posted user name
	user := model.UserCreateRequest{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}

	//generate access token
	token := uuid.Must(uuid.NewRandom()).String()

	//save user to database
	repository.CreateUser(user.Name, token)

	//response
	resp := model.UserCreateResponse{}
	resp.Token = token
	json.NewEncoder(w).Encode(resp)

}