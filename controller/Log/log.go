package Log

import (
	"SWP490_G21_Backend/model/response"
	"SWP490_G21_Backend/utility"
	"bufio"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"time"
)

func StreamLogFile(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userName := claims["username"].(string)

	logFile, err := os.OpenFile("log/file.log", os.O_APPEND, 0666)
	if err != nil {
		utility.FileLog.Println(err)
		return c.JSON(http.StatusInternalServerError, response.Message{Message: utility.Error021OpenFileError})
	}
	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {
		}
	}(logFile)
	utility.FileLog.Println("Admin " + userName + " access log file")
	reader := bufio.NewReader(logFile)
	logBytes := make(chan []byte)
	errChan := make(chan error)
	go func() {
		var bs []byte
		endOfFile := false
		for {
			b, err := reader.ReadByte()
			if err != nil {
				if err == io.EOF {
					err := reader.UnreadByte()
					if err != nil {
						errChan <- err
						return
					}
					endOfFile = true
					time.Sleep(time.Millisecond * 500)
					continue
				}
				errChan <- err
				return
			}
			if !endOfFile {
				bs = append(bs, b)
				if reader.Buffered() <= 0 {
					tempBytes := bs
					bs = []byte{}
					logBytes <- tempBytes
				}
			} else {
				endOfFile = false
			}
		}
	}()
	for {
		select {
		case <-c.Request().Context().Done():
			utility.FileLog.Println("Admin " + userName + " finished reading log file")
			return c.JSON(http.StatusInternalServerError, response.Message{Message: utility.Error021OpenFileError})
		case out := <-logBytes:
			_, err := c.Response().Write(out)
			if err != nil {
				utility.FileLog.Println("Response write error: ", err.Error())
				return c.JSON(http.StatusInternalServerError, response.Message{Message: utility.Error064WriteError})
			}
			c.Response().Flush()
		case err = <-errChan:
			utility.FileLog.Println("Error reading log file: ", err.Error())
			return c.JSON(http.StatusInternalServerError, response.Message{Message: utility.Error063ReadLogError})
		}
	}
}

func readLog(logString chan string, reader *bufio.Reader) {

}
