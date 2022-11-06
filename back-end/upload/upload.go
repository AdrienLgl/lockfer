package upload

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lockfer/decrypt"
	"lockfer/encrypt"
	"lockfer/key_generator"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/go-chi/httprate"
	"github.com/gorilla/mux"
)

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

type TokResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type OKResp struct {
	Status  bool   `json:"status"`
	UUID    string `json:"uuid"`
	Message string `json:"message"`
}

type DecryptResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type Decrypt struct {
	Token string `json:"token"`
}

type Identifier struct {
	Id string `json:"id"`
}

func UploadMultiplesFiles(w http.ResponseWriter, r *http.Request) {

	AddLog("upload files", r.RemoteAddr)

	key := key_generator.GetKey()
	id, e := key_generator.CreateUUID()
	if e != nil {
		ErrorResponse(w, r, "Error during key generation, please retry")
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 800*1024*1024) // 100 Mb
	err := r.ParseMultipartForm(10)                        // grab the multipart form
	formdata := r.MultipartForm
	if formdata == nil || formdata.File["multiplefiles"] == nil {
		ErrorResponse(w, r, "Too large files, your total encryption should be less than 800MB")
		return
	}

	if err != nil {
		ErrorResponse(w, r, "Too large files, your total encryption should be less than 800MB")
		return
	}

	//get the *fileheaders
	files := formdata.File["multiplefiles"] // grab the filenames

	for i := range files { // loop through the files one by one
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		fileBytes, err := ioutil.ReadAll(file)
		status, err := encrypt.EncryptFile(fileBytes, key, files[i].Filename, id)
		if !status {
			ErrorResponse(w, r, err.Error())
			return
		}
	}
	token := key_generator.CreateToken(hex.EncodeToString(key), id)
	TokenResponse(w, r, token, "Your files has been encrypted successfully")
}

func DecryptFiles(w http.ResponseWriter, r *http.Request) {
	AddLog("decrypt files", r.RemoteAddr)
	var key Decrypt
	err := json.NewDecoder(r.Body).Decode(&key)
	if err != nil {
		ErrorResponse(w, r, "Please, provide a body request")
		return
	}
	resp, err := key_generator.DecryptToken(key.Token)
	if err != nil {
		TimeResponse(w, r, "The files you are looking for are no longer available (24h).")
		return
	}
	files, err := ioutil.ReadDir("files/encrypted/" + resp.Id)
	if err != nil {
		ErrorResponse(w, r, "Your id is invalid: not found")
		return
	}
	errM := os.MkdirAll("files/decrypted/"+resp.Id, os.ModePerm)
	if errM != nil {
		fmt.Println("Error during file decryption")
		ErrorResponse(w, r, "Error during file decryption")
		return
	}
	count := 0
	for _, file := range files {
		f, _ := ioutil.ReadFile("files/encrypted/" + resp.Id + "/" + file.Name())
		secret_key, _ := hex.DecodeString(resp.Key)
		d, err := decrypt.DecryptFile(f, file.Name(), secret_key)
		if err != nil {
			ErrorResponse(w, r, "Error during file decryption")
			return
		}
		if d != nil {
			c, err := createFile(d, file, resp.Id)
			if err != nil {
				ErrorResponse(w, r, "Error during file decryption")
				return
			}
			if c {
				count++
			}
		}
	}

	if count == len(files) {
		AddLog("zip creation...", "server")
		decrypt.CreateZip("files/decrypted/"+resp.Id, resp.Id)
		OkResponse(w, r, "Your files has been decrypted successfully", resp.Id)
	}
}

func createFile(file []byte, fileInfo os.FileInfo, id string) (bool, error) {
	AddLog("create file: "+fileInfo.Name()+" - "+id, "server")
	filename := fileNameWithoutExtSliceNotation(fileInfo.Name())
	err := ioutil.WriteFile("files/decrypted/"+id+"/"+filename, file, 0777)
	if err != nil {
		fmt.Println("Error during file creation...")
		return false, err
	}
	return true, nil
}

func fileNameWithoutExtSliceNotation(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}

func DownloadArchive(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		ErrorResponse(w, r, "Please, provide uuid parameter")
		return
	}
	ReturnFile(w, r, id)
}

func ReturnFile(writer http.ResponseWriter, req *http.Request, id string) {
	AddLog("download archive: "+id, req.RemoteAddr)
	writer.Header().Set("Content-type", "application/zip")
	http.ServeFile(writer, req, "files/decrypted/"+id+".zip")
	//delete file from server once it has been served
	err := os.RemoveAll("files/decrypted/" + id)
	if err != nil {
		println(err.Error())
	}
	defer os.Remove("files/decrypted/" + id + ".zip")
}

func TokenResponse(w http.ResponseWriter, r *http.Request, token string, message string) {
	AddLog(message, r.RemoteAddr)
	data := TokResponse{
		true,
		message,
		token,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func OkResponse(w http.ResponseWriter, r *http.Request, message string, id string) {
	AddLog(message, r.RemoteAddr)
	data := OKResp{
		true,
		id,
		message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func ErrorResponse(w http.ResponseWriter, r *http.Request, errorMessage string) {
	AddLog(errorMessage, r.RemoteAddr)
	data := Response{
		false,
		"",
		"Error: " + errorMessage,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(data)
}

func TimeResponse(w http.ResponseWriter, r *http.Request, errorMessage string) {
	AddLog(errorMessage, r.RemoteAddr)
	data := Response{
		false,
		"",
		"Error: " + errorMessage,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(data)
}

func SetupRoutes(wg *sync.WaitGroup) {
	AddLog("starting server...", "server")
	defer wg.Done()
	r := mux.NewRouter()
	// rate limiter (20 requests per minute)
	r.Use(httprate.LimitByIP(20, 1*time.Minute))
	// log middleware
	r.Use(logMiddleware)
	r.HandleFunc("/api/v1/upload", UploadMultiplesFiles)
	r.HandleFunc("/api/v1/decrypt", DecryptFiles)
	r.HandleFunc("/api/v1/download/{uuid}", DownloadArchive)
	// timeout middleware
	muxWithMiddlewares := http.TimeoutHandler(r, time.Minute*3, "Service Unavailable, timeout")
	fmt.Println("Server listening on http://127.0.0.1:8080")
	http.ListenAndServe(":8080", muxWithMiddlewares)
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		AddLog(r.Method, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

func AddLog(action string, ip string) {
	f, err := os.OpenFile("log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Printf("[%d] - %s - action: %s \n", time.Now(), ip, action)
}
