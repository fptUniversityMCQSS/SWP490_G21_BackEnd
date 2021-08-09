package controller

import (
	"SWP490_G21_Backend/model"
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
	"log"
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

	var knows []*model.Knowledge
	var knowRs []*response.KnowledgResponse
	// Get a QuerySeter object. User is table name
	_, err := utility.DB.QueryTable("knowledge").OrderBy("-id").RelatedSel().All(&knows)

	//if has problem in connection
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "query error",
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
	log.Printf(userName + " get list knowledge")
	return c.JSON(http.StatusOK, knowRs)

}

func KnowledgeUpload(c echo.Context) error {
	// Read form fields

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userName := claims["username"].(string)
	userId := claims["userId"].(float64)
	IntUserId := int64(userId)

	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "file error",
		})
	}
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "open error",
		})
	}
	defer src.Close()

	user := &model.User{
		Id:       IntUserId,
		Username: userName,
	}

	know := &model.Knowledge{
		Name:   file.Filename,
		User:   user,
		Status: "Processing",
	}
	i, err := utility.DB.QueryTable("knowledge").PrepareInsert()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
		})
	}
	//fmt.Println(i)
	insert, err := i.Insert(know)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "Insert file error",
		})
	}
	//fmt.Println(insert)
	err1 := i.Close()
	if err1 != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "Close error",
		})
	}

	userPath := "knowledge/" + userName
	stringIdInsert := strconv.Itoa(int(insert))
	fileFolderPath := userPath + "/" + stringIdInsert
	err = os.MkdirAll(fileFolderPath, os.ModePerm)
	if err != nil {
		log.Print(err)
	}
	know.Path = fileFolderPath
	_, err = utility.DB.Update(know)

	filePath := fileFolderPath + "/" + file.Filename

	dst, err := os.Create(filePath)
	if err != nil {
		log.Print(err)
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "Copy file error",
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
	enc.Encode(knowledgeResponse)
	c.Response().Flush()

	extension := filepath.Ext(file.Filename)

	switch extension {

	case ".pdf":
		err := convertPdfToTxt(filePath, file.Filename, extension, fileFolderPath, insert)
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: "Can not read file pdf",
			})
		}

	case ".doc":
		err := converDocToText(fileFolderPath, file.Filename, extension, insert)
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: "Can not read file doc",
			})
		}

	case ".docx":
		err := convertDocxToText(filePath, file.Filename, extension, fileFolderPath, insert)
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: "Can not read file docx",
			})
		}

	case ".txt":
		err := modifyTxtFile(filePath, file.Filename, extension, fileFolderPath, insert)
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusInternalServerError, response.Message{
				Message: "Can not read file txt",
			})
		}
	default:
		src.Close()
		dst.Close()
		err := os.RemoveAll(fileFolderPath)
		if err != nil {
			log.Print(err)
			return err
		}
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "please check format file it must pdf, doc, docx or txt",
		})
	}

	know.Status = "Encoding"
	_, err = utility.DB.Update(know)
	knowledgeResponse = response.KnowledgResponse{
		Id:       insert,
		Name:     file.Filename,
		Date:     time.Now(),
		Username: user.Username,
		Status:   know.Status,
	}

	enc.Encode(knowledgeResponse)
	c.Response().Flush()

	//placeToSaveFileTxt := createFolderOfTxtFile(file.Filename,extension,fileFolderPath,insert)          //placeToSaveFileTxt get path of txt file
	//utility.SendFileRequest(utility.ConfigData.AIServer+"/knowledge", "POST", placeToSaveFileTxt)

	know.Status = "Ready"
	_, err = utility.DB.Update(know)
	knowledgeResponse = response.KnowledgResponse{
		Id:       insert,
		Name:     file.Filename,
		Date:     time.Now(),
		Username: user.Username,
		Status:   know.Status,
	}

	log.Printf(userName + " upload file : " + file.Filename)
	return c.JSON(http.StatusOK, knowledgeResponse)
}

func DownloadKnowledge(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userName := claims["username"].(string)

	knowledgeId := c.Param("id")
	intKnowledgeId, _ := strconv.ParseInt(knowledgeId, 10, 64)
	var knowledge model.Knowledge

	err := utility.DB.QueryTable("knowledge").Filter("id", intKnowledgeId).One(&knowledge)

	//if has problem in connection
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "query error",
		})
	}

	log.Printf(userName + " downloaded " + knowledge.Name)

	return c.Attachment(knowledge.Path+"/"+knowledge.Name, knowledge.Name)
}

func DeleteKnowledge(c echo.Context) error {

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userName := claims["username"].(string)

	knowledgeId := c.Param("id")
	var knowledge model.Knowledge
	intKnowledgeId, _ := strconv.ParseInt(knowledgeId, 10, 64)
	err3 := utility.DB.QueryTable("knowledge").Filter("id", intKnowledgeId).One(&knowledge)
	if err3 != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "query error",
		})
	}
	_, err := utility.DB.QueryTable("knowledge").Filter("id", intKnowledgeId).Delete()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "query error",
		})
	}

	err2 := os.RemoveAll(knowledge.Path)

	if err2 != nil {
		return c.JSON(http.StatusInternalServerError, response.Message{
			Message: "remove file error",
		})
	}
	message := response.Message{
		Message: "Delete successfully",
	}
	log.Printf(userName + " delete file " + knowledge.Name)
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
		return nil, err
	}
	unknown, err := oleutil.CreateObject("Word.Application")
	if err != nil {
		return nil, err
	}
	word, err := unknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return nil, err
	}
	_, err = oleutil.PutProperty(word, "Visible", true)
	if err != nil {
		return nil, err
	}
	documents := oleutil.MustGetProperty(word, "Documents").ToIDispatch()
	defer documents.Release()
	document := oleutil.MustCallMethod(documents, "Open", inputDoc, false, true).ToIDispatch()
	defer document.Release()
	oleutil.MustCallMethod(document, "SaveAs", outPutDocx, 16).ToIDispatch()

	_, err = oleutil.CallMethod(word, "Quit")
	if err != nil {
		return nil, err
	}
	word.Release()
	ole.CoUninitialize()
	open, err := os.Open(outPutDocx)
	if err != nil {
		return nil, err
	}
	return open, nil
}
func parseTextFromDocorDocx(path string) (string, error) {
	text, err := goword.ParseText(path)
	return text, err
}

func createFolderOfTxtFile(fileName string, extension string, fileFolderPath string, id int64) string {
	stringId := strconv.Itoa(int(id))
	extensionNewFormat := strings.ReplaceAll(fileName, extension, stringId+".txt")
	placeToSaveFileTxt := fileFolderPath + "/" + extensionNewFormat
	return placeToSaveFileTxt
}
func convertPdfToTxt(filepath, fileFileName, extension, fileFolderPath string, insert int64) error {
	doc, err := fitz.New(filepath)
	if err != nil {
		log.Print(err)
		return err
	}
	defer func(doc *fitz.Document) {
		err := doc.Close()
		if err != nil {
			return
		}
	}(doc)
	placeToSaveFileTxt := createFolderOfTxtFile(fileFileName, extension, fileFolderPath, insert)
	f, err := os.Create(placeToSaveFileTxt)
	for n := 0; n < doc.NumPage(); n++ {
		text, err := doc.Text(n)
		if err != nil {
			panic(err)
		}
		f.WriteString(text)
	}
	return nil
}
func converDocToText(fileFolderPath, fileFileName, extension string, insert int64) error {
	dir, err := filepath.Abs(fileFolderPath)
	if err != nil {
		log.Fatal(err)
	}
	fileName, err := toDocx(dir, fileFileName)
	text, err := parseTextFromDocorDocx(fileName.Name())
	if err != nil {
		log.Print(err)
		return err
	}
	placeToSaveFileTxt := createFolderOfTxtFile(fileFileName, extension, fileFolderPath, insert)
	f, err := os.Create(placeToSaveFileTxt)
	_, err = f.WriteString(text)
	if err != nil {
		log.Print(err)
		return err
	}
	err2 := fileName.Close()
	if err2 != nil {
		log.Print(err)
		return err
	}
	os.Remove(fileName.Name())
	return nil
}

func convertDocxToText(filePath, fileFileName, extension, fileFolderPath string, insert int64) error {
	text, err := parseTextFromDocorDocx(filePath)
	if err != nil {
		log.Print(err)
		return err
	}
	placeToSaveFileTxt := createFolderOfTxtFile(fileFileName, extension, fileFolderPath, insert)
	f, err := os.Create(placeToSaveFileTxt)
	_, err = f.WriteString(text)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}
func modifyTxtFile(filePath, fileFileName, extension, fileFolderPath string, insert int64) error {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	text := string(content)
	placeToSaveFileTxt := createFolderOfTxtFile(fileFileName, extension, fileFolderPath, insert)
	f, err := os.Create(placeToSaveFileTxt)
	_, err = f.WriteString(text)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}
