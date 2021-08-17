package utility

const (
	//user error
	Error001InvalidUser              = "Invalid user"
	Error002ErrorQueryForGetAllUsers = "Can't get all users"
	Error003UserExisted              = "User already exists"
	Error004CantGetTableUser         = "Can't get table user"
	Error005InsertUserError          = "Add user error"
	Error006UserNameModified         = "Username must not contains special characters and has length at least 8 characters"
	Error007CantGetUser              = "Can't get user with your id"
	Error008UserIdInvalid            = "User id is invalid"
	Error009DeleteUserFailed         = "Can't delete user"
	Error010RoleOfUserIsInvalid      = "Role of user is invalid"
	Error011UpdateUserFailed         = "Can't update user"
	Error012PasswordEmpty            = "Password is empty"
	Error013CreateTokenOfUserFailed  = "Can't create token of user"
	Error060CurrentPasswordInvalid   = "Current Password isn't correct"
	Error063RoleOfUserIsInvalid      = "Role of user is invalid"
	Error064PasswordOfUserIsInvalid  = "Password has at least 8 character"
	Error065EmailInvalid             = "Email has a form xxx@xxx.xxx"
	Error066PhoneInvalid             = "Phone must be 10 digit"
	Error067FullNameInvalid          = "Full Name has 8 to 30 characters"
	//exam error
	Error014ErrorQueryForGetAllExamTest  = "Can't get all exam tests with your id"
	Error015CantGetExamTest              = "Can't get exam with your id"
	Error040CantPrepareStatementExamTest = "Can't Prepare Statement ExamTest"
	Error042UpdateExamFailed             = "Can't update exam"
	Error041InsertExamFailed             = "Can't insert exam"
	Error061ExamIdInvalid                = "Exam id is invalid"
	Error062DeleteExamFailed             = "Can't delete exam"

	//questions error
	Error016ErrorQueryForGetAllQuestions = "Can't get all questions with your id"
	Error052InsertQuestionError          = "Add question error"
	Error053CantGetQuestion              = "Can't get question"
	Error056UpdateAnswerError            = "Update answer error"

	//options error
	Error017ErrorQueryForGetAllOptions = "Can't get all options with your id"
	Error053InsertOptionError          = "Add option error"
	Error054CantGetOption              = "Can't get option"

	//authority error
	Error018DontHavePermission = "You dont have permission to access this"
	Error057AccessDenied       = "Access Denied"

	//knowledge error
	Error019ErrorQueryForGetAllKnowledge = "Can't get all knowledge"
	Error023CantGetKnowledge             = "Can't get knowledge"
	Error024InsertKnowledgeError         = "Add knowledge error"
	Error026UpdateKnowledgeFailed        = "Can't update knowledge"
	Error036KnowIdInvalid                = "Knowledge id is invalid"
	Error037DeleteKnowledgeFailed        = "Can't delete knowledge"

	//file error
	Error020FileError              = "File not found please try again"
	Error021OpenFileError          = "Open file error"
	Error027CopyFileError          = "Copy file error"
	Error029ReadFilePdfError       = "Read file pdf error"
	Error030ReadFileDocError       = "Read file doc error"
	Error031ReadFileDocxError      = "Read file docx error"
	Error032ReadFileTxtError       = "Read file txt error"
	Error033CloseFileError         = "Close file error"
	Error034CheckFormatFile        = "Please check format file it must pdf, doc, docx or txt"
	Error038RemoveFileError        = "Remove file error"
	Error043CantParseFileSize      = "Can't parse file size"
	Error044CantGetFilePath        = "Can't get file path"
	Error047StatFileError          = "Stat file error"
	Error050ReadFileDocOrDocxError = "Read file doc or docx error"

	//connection error
	Error022CloseConnectionError = "Close connection error"

	//directory error
	Error025CreateDirectoryError = "Create directory error"

	//encoding response error
	Error028EncodingResponseError = "Encoding response error"

	//request model ai error
	Error035RequestModelAI                    = "Failed to request to AI server"
	Error039RequestDeleteKnowledgeFormModelAI = "Failed to delete knowledge form AI server"
	Error055CantGetResponseModelAI            = "Fail to receive response from AI server"

	//result
	Error045CantGetResult               = "Error for get result"
	Error046CantParseFileDocToDocx      = "Error for parse file from doc to docx"
	Error048ParseFileXmlError           = "Error for parse xml file"
	Error051ParseNumberOfQuestionError  = "Error for parse number of question"
	Error058ParseNumberOfQuestionsError = "Error for parse number of questions"
	Error059ReformatStringFalse         = "Missing ':' Subject or Number Of Question"
)
