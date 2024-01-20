package request_handler

import (
	socket "backend/appsocket"
	"backend/model/auth_token_data"
	"backend/model/common_models"
	"backend/redisdb"
	r2 "backend/utilities/cloudflareR2utils"
	"bytes"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

var filePool = sync.Pool{
	New: func() any { return new(bytes.Buffer) },
}

func SendMessageWithFile(w http.ResponseWriter, r *http.Request) (e error, statusCode int) {
	ctx := r.Context()
	authTokenData, containsAuthToken := ctx.Value("user").(auth_token_data.Model)
	if !containsAuthToken {
		return errors.New("Invalid"), 400
	}

	err := r.ParseMultipartForm(55 << 20)
	if err != nil {
		return err, 400
	}

	var (
		form = r.MultipartForm
	)

	if len(form.File) == 0 {
		return errors.New("No files"), 400
	}

	var (
		files                       []*multipart.FileHeader
		tempSlice                   []string
		toPersonID                  int64
		fileNameString, contentType string
		file                        multipart.File
		ok                          bool
		fileSizeInBytes             int64
		nowTime                     time.Time
	)

	if files, ok = form.File["files"]; !ok || len(files) == 0 {
		return errors.New("No files"), 400
	}
	err = nil

	pathBuilder := strings.Builder{}
	pathBuilder.Grow(35)
	pathBuilder.WriteString("Chats/AudioFiles/")
	if tempSlice, ok = form.Value["toPersonID"]; ok && len(tempSlice) > 0 {
		if tmpInt, parseErr := strconv.ParseInt(tempSlice[0], 10, 64); err != nil {
			return parseErr, 400
		} else {
			toPersonID = tmpInt
		}

	}

	fileBytes := filePool.Get().(*bytes.Buffer)
	defer func() {
		fileBytes.Reset()
		filePool.Put(fileBytes)
	}()
	formFile := files[0]

	fileBytes.Reset()
	fileSizeInBytes = formFile.Size >> 10
	nowTime = time.Now()
	pathBuilder.WriteString(fmt.Sprint(nowTime.UnixMilli()))
	pathBuilder.WriteString("_")
	pathBuilder.WriteString(formFile.Filename)
	fileNameString = pathBuilder.String()
	contentType = r2.GetContentType(formFile.Filename)

	if file, err = formFile.Open(); err != nil {
		_ = file.Close()
		return err, 400
	} else if _, err = io.Copy(fileBytes, file); err != nil {
		_ = file.Close()
		return err, 400
	}
	_ = file.Close()
	if err = r2.UploadBytes(fileBytes.Bytes(), fileNameString, contentType); err != nil {
		return err, 400
	}
	/*if strings.HasPrefix(contentType, "image") {
		if img, err = imaging.Decode(fileBytes); err == nil {
			size := img.Bounds().Size()
			thumb := imaging.Resize(img, int(float64(size.X)*0.15), int(float64(size.Y)*0.15), imaging.NearestNeighbor)
			switch contentType {
			case "image/png":
				err = png.Encode(thumbBytes, thumb)
			case "image/jpeg":
				err = jpeg.Encode(thumbBytes, thumb, nil)
			case "image/bmp":
				err = bmp.Encode(thumbBytes, thumb)
			case "image/gif":
				err = gif.Encode(thumbBytes, thumb, nil)
			case "image/tiff":
				err = tiff.Encode(thumbBytes, thumb, nil)
			}
			if err == nil {
				_ = r2.UploadBytes(thumbBytes.Bytes(), strings.Join([]string{
					pathBuilder.String()[1:],
					fileNameString,
				}, ""), s3BucketName, contentType)
			}

		}
	}*/
	redisM := common_models.MessageDBModel{
		MessageID:    0,
		Timestamp:    nowTime.UnixMilli(),
		MessageAudio: fileNameString,
		Reaction:     "",
	}
	if !redisdb.ChatExists(authTokenData.PersonID, toPersonID) {
		//if err = redisdb.CreateChat(authTokenData.PersonID, toPersonID); err != nil {
		//
		//}
		socket.Broadcast([]int64{authTokenData.PersonID, toPersonID, authTokenData.PersonID}, "new_chat", fmt.Sprintf("%d,%d", authTokenData.PersonID, toPersonID))
	}
	err = redisdb.AddMessage(authTokenData.PersonID, toPersonID, &redisM)
	if err != nil {
		return err, 400
	}
	m := common_models.SocketAndResponseModel{
		FromPersonID:           authTokenData.PersonID,
		ToPersonID:             toPersonID,
		MessageID:              redisM.MessageID,
		MessageFile:            fileNameString,
		MessageFileSizeInBytes: fileSizeInBytes,
		MessageTime:            nowTime.UnixMilli(),
	}
	socket.Broadcast([]int64{authTokenData.PersonID, toPersonID, authTokenData.PersonID}, "new_message", m)
	_, err = w.Write([]byte("success"))
	return err, 400
}

func AddReactionToMessage(w http.ResponseWriter, r *http.Request) (e error, statusCode int) {
	messageIDRaw := chi.URLParam(r, "messageID")
	queryParams := r.URL.Query()

	reaction := queryParams.Get("reaction")
	if reaction == "" {
		return errors.New("Invalid reaction"), 400
	}
	toPersonIDStr := queryParams.Get("toPersonID")
	if toPersonIDStr == "" {
		return errors.New("Invalid toPersonID"), 400
	}
	var err error
	var toPersonID int64
	if toPersonID, err = strconv.ParseInt(toPersonIDStr, 10, 64); err != nil {
		return err, 400
	}

	var messageIDString string
	if messageIDString, err = url.QueryUnescape(messageIDRaw); err != nil {
		messageIDString = messageIDRaw
	}
	var messageID int64
	if messageID, err = strconv.ParseInt(messageIDString, 10, 64); err != nil {
		return err, 400
	}
	authTokenData, containsAuthToken := r.Context().Value("user").(auth_token_data.Model)
	if !containsAuthToken {
		return errors.New("Invalid"), 400
	}
	if err = redisdb.AddReaction(authTokenData.PersonID, toPersonID, messageID, reaction); err != nil {
		return err, 400
	}
	socket.Broadcast([]int64{authTokenData.PersonID, toPersonID, authTokenData.PersonID}, "new_reaction", common_models.SocketAndResponseModel{
		MessageID: messageID,
		Reaction:  reaction,
	})
	_, err = w.Write([]byte("success"))
	return err, 400

}
