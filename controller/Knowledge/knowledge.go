package Knowledge

import (
	"SWP490_G21_Backend/model/entity"
	"SWP490_G21_Backend/model/response"
	"SWP490_G21_Backend/utility"
	"encoding/json"
	_ "github.com/astaxie/beego/orm"
	"github.com/gen2brain/go-fitz"
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"github.com/golang-jwt/jwt"
	"github.com/guylaor/goword"
	"github.com/labstack/echo/v4"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func ListKnowledge(c echo.Context) error {

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userName := claims["username"].(string)

	//userId := claims["userId"].(float64)

	var knows []*entity.Knowledge
	var knowRs []*response.KnowledgResponse
	// Get a QuerySeter object. User is table name
	_, err := utility.DB.QueryTable("knowledge").OrderBy("-id").RelatedSel().All(&knows)

	//if has problem in connection
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error019ErrorQueryForGetAllKnowledge,
		})
	}
	//add selected data to knowledge_Res list
	for _, k := range knows {
		var knowR = new(response.KnowledgResponse)
		knowR.Id = k.Id
		knowR.Name = k.Name
		knowR.Date = k.Date
		knowR.Username = k.User.Username
		knowR.Status = k.Status
		knowRs = append(knowRs, knowR)
	}
	utility.FileLog.Println(userName + " get list knowledge")
	if knowRs == nil {
		knowRs = []*response.KnowledgResponse{}
	}
	return c.JSON(http.StatusOK, knowRs)

}

func KnowledgeUpload(c echo.Context) error {
	//WARNING: missing delete knowledge if uploading failed

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userName := claims["username"].(string)
	userId := claims["userId"].(float64)
	IntUserId := int64(userId)

	file, err := c.FormFile("file")
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error020FileError,
		})
	}
	src, err := file.Open()
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error021OpenFileError,
		})
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			utility.FileLog.Println(err)
			return
		}
	}(src)

	user := &entity.User{
		Id:       IntUserId,
		Username: userName,
	}

	know := &entity.Knowledge{
		Name:   file.Filename,
		User:   user,
		Status: "Processing",
	}
	i, err := utility.DB.QueryTable("knowledge").PrepareInsert()
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error023CantGetKnowledge,
		})
	}
	//fmt.Println(i)
	insert, err := i.Insert(know)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error024InsertKnowledgeError,
		})
	}
	//fmt.Println(insert)
	err1 := i.Close()
	if err1 != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error022CloseConnectionError,
		})
	}

	userPath := "knowledge/" + userName
	stringIdInsert := strconv.Itoa(int(insert))
	fileFolderPath := userPath + "/" + stringIdInsert
	err = os.MkdirAll(fileFolderPath, os.ModePerm)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error025CreateDirectoryError,
		})
	}
	know.Path = fileFolderPath
	_, err = utility.DB.Update(know)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error026UpdateKnowledgeFailed,
		})
	}

	filePath := fileFolderPath + "/" + file.Filename

	dst, err := os.Create(filePath)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error025CreateDirectoryError,
		})
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
			utility.FileLog.Println(err)
			return
		}
	}(dst)

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error027CopyFileError,
		})
	}

	knowledgeResponse := response.KnowledgResponse{
		Id:       insert,
		Name:     file.Filename,
		Date:     time.Now(),
		Username: user.Username,
		Status:   know.Status,
	}

	enc := json.NewEncoder(c.Response())
	err = enc.Encode(knowledgeResponse)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error028EncodingResponseError,
		})
	}
	c.Response().Flush()

	extension := filepath.Ext(file.Filename)

	var placeToSaveFileTxt string

	switch extension {

	case ".pdf":
		placeToSaveFileTxt, err = convertPdfToTxt(filePath, file.Filename, extension, fileFolderPath, insert)
		if err != nil {
			utility.FileLog.Println(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: utility.Error029ReadFilePdfError,
			})
		}

	case ".doc":
		placeToSaveFileTxt, err = convertDocToText(fileFolderPath, file.Filename, extension, insert)
		if err != nil {
			utility.FileLog.Println(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: utility.Error030ReadFileDocError,
			})
		}

	case ".docx":
		placeToSaveFileTxt, err = convertDocxToText(filePath, file.Filename, extension, fileFolderPath, insert)
		if err != nil {
			utility.FileLog.Println(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: utility.Error031ReadFileDocxError,
			})
		}

	case ".txt":
		placeToSaveFileTxt, err = modifyTxtFile(filePath, file.Filename, extension, fileFolderPath, insert)
		if err != nil {
			utility.FileLog.Println(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: utility.Error032ReadFileTxtError,
			})
		}
	default:
		err = src.Close()
		if err != nil {
			utility.FileLog.Println(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: utility.Error033CloseFileError,
			})
		}
		err = dst.Close()
		if err != nil {
			utility.FileLog.Println(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: utility.Error033CloseFileError,
			})
		}
		err = os.RemoveAll(fileFolderPath)
		if err != nil {
			utility.FileLog.Println(err)
			return err
		}
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error034CheckFormatFile,
		})
	}
	know.Status = "Encoding"
	_, err = utility.DB.Update(know)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error026UpdateKnowledgeFailed,
		})
	}
	knowledgeResponse = response.KnowledgResponse{
		Id:       insert,
		Name:     file.Filename,
		Date:     time.Now(),
		Username: user.Username,
		Status:   know.Status,
	}

	err = enc.Encode(knowledgeResponse)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error028EncodingResponseError,
		})
	}
	c.Response().Flush()

	err = utility.SendFileRequest(utility.ConfigData.AIServer+"/knowledge", "POST", placeToSaveFileTxt)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error035RequestModelAI,
		})
	}

	know.Status = "Ready"
	_, err = utility.DB.Update(know)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error026UpdateKnowledgeFailed,
		})
	}
	knowledgeResponse = response.KnowledgResponse{
		Id:       insert,
		Name:     file.Filename,
		Date:     time.Now(),
		Username: user.Username,
		Status:   know.Status,
	}
	utility.FileLog.Println(userName + " upload file : " + file.Filename)
	return c.JSON(http.StatusOK, knowledgeResponse)
}

func DownloadKnowledge(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userName := claims["username"].(string)

	knowledgeId := c.Param("id")
	intKnowledgeId, err := strconv.ParseInt(knowledgeId, 10, 64)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error036KnowIdInvalid,
		})
	}
	var knowledge entity.Knowledge

	err = utility.DB.QueryTable("knowledge").Filter("id", intKnowledgeId).One(&knowledge)

	//if has problem in connection
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error023CantGetKnowledge,
		})
	}
	utility.FileLog.Println(userName + " downloaded " + knowledge.Name)

	return c.Attachment(knowledge.Path+"/"+knowledge.Name, knowledge.Name)
}

func DeleteKnowledge(c echo.Context) error {

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userName := claims["username"].(string)
	role := claims["role"].(string)

	knowledgeId := c.Param("id")
	var knowledge entity.Knowledge
	intKnowledgeId, err := strconv.ParseInt(knowledgeId, 10, 64)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error036KnowIdInvalid,
		})
	}
	err3 := utility.DB.QueryTable("knowledge").Filter("id", intKnowledgeId).One(&knowledge)
	if err3 != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error023CantGetKnowledge,
		})
	}
	if role != "admin" && knowledge.User.Username != userName {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error018DontHavePermission,
		})
	}
	_, err = utility.DB.QueryTable("knowledge").Filter("id", intKnowledgeId).Delete()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error037DeleteKnowledgeFailed,
		})
	}
	err2 := os.RemoveAll(knowledge.Path)
	if err2 != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error038RemoveFileError,
		})
	}
	err = utility.DeleteKnowledge(utility.ConfigData.AIServer, "DELETE", knowledge.Name)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: utility.Error039RequestDeleteKnowledgeFormModelAI,
		})
	}
	message := response.Message{
		Message: "Delete successfully",
	}
	utility.FileLog.Println(userName + " delete file " + knowledge.Name)
	return c.JSON(http.StatusOK, message)
}

func toDocx(folderPath string, fileName string) (*os.File, error) {
	inputDoc := folderPath + "/" + fileName
	splitPath := strings.Split(fileName, ".")

	outPutDocx := ""
	for i := 0; i < len(splitPath)-1; i++ {
		outPutDocx += splitPath[i]
	}
	outPutDocx = outPutDocx + ".docx"

	outPutDocx = folderPath + "/" + outPutDocx

	err := ole.CoInitialize(0)
	if err != nil {
		utility.FileLog.Println(err)
		return nil, err
	}
	unknown, err := oleutil.CreateObject("Word.Application")
	if err != nil {
		utility.FileLog.Println(err)
		return nil, err
	}
	word, err := unknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		utility.FileLog.Println(err)
		return nil, err
	}
	_, err = oleutil.PutProperty(word, "Visible", true)
	if err != nil {
		utility.FileLog.Println(err)
		return nil, err
	}
	documents := oleutil.MustGetProperty(word, "Documents").ToIDispatch()
	defer documents.Release()
	document := oleutil.MustCallMethod(documents, "Open", inputDoc, false, true).ToIDispatch()
	defer document.Release()
	oleutil.MustCallMethod(document, "SaveAs", outPutDocx, 16).ToIDispatch()

	_, err = oleutil.CallMethod(word, "Quit")
	if err != nil {
		utility.FileLog.Println(err)
		return nil, err
	}
	word.Release()
	ole.CoUninitialize()
	open, err := os.Open(outPutDocx)
	if err != nil {
		utility.FileLog.Println(err)
		return nil, err
	}
	return open, nil
}
func parseTextFromDocorDocx(path string) (string, error) {
	text, err := goword.ParseText(path)
	return text, err
}

func createFolderOfTxtFile(fileName string, extension string, fileFolderPath string, id int64) string {
	//stringId := strconv.Itoa(int(id))
	extensionNewFormat := strings.ReplaceAll(fileName, extension, ".txt")
	placeToSaveFileTxt := fileFolderPath + "/" + extensionNewFormat
	return placeToSaveFileTxt
}
func convertPdfToTxt(filepath, fileFileName, extension, fileFolderPath string, insert int64) (string, error) {
	doc, err := fitz.New(filepath)
	if err != nil {
		utility.FileLog.Println(err)
		return err.Error(), err
	}
	defer func(doc *fitz.Document) {
		err := doc.Close()
		if err != nil {
			return
		}
	}(doc)
	placeToSaveFileTxt := createFolderOfTxtFile(fileFileName, extension, fileFolderPath, insert)
	f, err := os.Create(placeToSaveFileTxt)
	if err != nil {
		utility.FileLog.Println(err)
		return err.Error(), err
	}
	for n := 0; n < doc.NumPage(); n++ {
		text, err := doc.Text(n)
		if err != nil {
			utility.FileLog.Println(err)
			return err.Error(), err
		}
		_, err = f.WriteString(text)
		if err != nil {
			utility.FileLog.Println(err)
			return err.Error(), err
		}
	}
	return placeToSaveFileTxt, nil
}
func convertDocToText(fileFolderPath, fileFileName, extension string, insert int64) (string, error) {
	dir, err := filepath.Abs(fileFolderPath)
	if err != nil {
		utility.FileLog.Println(err)
		return err.Error(), err
	}
	fileName, err := toDocx(dir, fileFileName)
	if err != nil {
		utility.FileLog.Println(err)
		return err.Error(), err
	}
	text, err := parseTextFromDocorDocx(fileName.Name())
	if err != nil {
		utility.FileLog.Println(err)
		return err.Error(), err
	}
	placeToSaveFileTxt := createFolderOfTxtFile(fileFileName, extension, fileFolderPath, insert)
	f, err := os.Create(placeToSaveFileTxt)
	if err != nil {
		utility.FileLog.Println(err)
		return err.Error(), err
	}
	_, err = f.WriteString(text)
	if err != nil {
		utility.FileLog.Println(err)
		return err.Error(), err
	}
	err = fileName.Close()
	if err != nil {
		utility.FileLog.Println(err)
		return err.Error(), err
	}
	err = os.Remove(fileName.Name())
	if err != nil {
		utility.FileLog.Println(err)
		return err.Error(), err
	}
	return placeToSaveFileTxt, nil
}

func convertDocxToText(filePath, fileFileName, extension, fileFolderPath string, insert int64) (string, error) {
	text, err := parseTextFromDocorDocx(filePath)
	if err != nil {
		utility.FileLog.Println(err)
		return err.Error(), err
	}
	placeToSaveFileTxt := createFolderOfTxtFile(fileFileName, extension, fileFolderPath, insert)
	f, err := os.Create(placeToSaveFileTxt)
	if err != nil {
		utility.FileLog.Println(err)
		return err.Error(), err
	}
	_, err = f.WriteString(text)
	if err != nil {
		utility.FileLog.Println(err)
		return err.Error(), err
	}
	return placeToSaveFileTxt, nil
}
func modifyTxtFile(filePath, fileFileName, extension, fileFolderPath string, insert int64) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		utility.FileLog.Println(err)
		return err.Error(), err
	}
	text := string(content)
	placeToSaveFileTxt := createFolderOfTxtFile(fileFileName, extension, fileFolderPath, insert)
	f, err := os.Create(placeToSaveFileTxt)
	if err != nil {
		utility.FileLog.Println(err)
		return err.Error(), err
	}
	_, err = f.WriteString(text)
	if err != nil {
		utility.FileLog.Println(err)
		return err.Error(), err
	}
	return placeToSaveFileTxt, nil
}
