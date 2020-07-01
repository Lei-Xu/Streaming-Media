package main

import (

	"io"
	"io/ioutil"
	"net/http"


	"github.com/julienschmidt/httprouter"
)

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	vl := VIDEO_DIR + vid

	video, err := os.Open(vl)

	if err != nil {
		log.Printf("Read file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}

	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)

	defer video.Close()
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "File is too large")
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	data, err := iotil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal server error")
	}

	fn := p.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_DIR + fn, data, 0666)

	if err != nil {
		log.Printf("Read file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Uploaded successfully")
}
