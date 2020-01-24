package controller
import(
	"net/http"
	"github.com/daleksprinter/catechdojo/model"
	"github.com/daleksprinter/catechdojo/repository"
	"encoding/json"
	"github.com/google/uuid"
	"fmt"
)

func CreateUserController(w http.ResponseWriter, r *http.Request){
	//parse posted user name
	user := model.UserCreateRequest{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		return 
	}

	//generate access token
	token := uuid.Must(uuid.NewRandom()).String()

	//save user to database
	err = repository.CreateUser(user.Name, token)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		return
	}

	//response
	resp := model.UserCreateResponse{}
	resp.Token = token
	json.NewEncoder(w).Encode(resp)

}

func GetUserController(w http.ResponseWriter, r *http.Request){
	token := r.Header.Get("x-token")

	user, err := repository.GetUser(token)

	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		return
	}

	resp := model.UserGetResponse{}
	resp.Name = user.Name

	json.NewEncoder(w).Encode(resp)
}

func UpdateUserController(w http.ResponseWriter, r *http.Request){
	token := r.Header.Get("x-token")

	user := model.UserUpdateRequest{}
	json.NewDecoder(r.Body).Decode(&user)

	err := repository.UpdateUser(token, user.Name)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		return 
	}
}