-- MySQL dump 10.13  Distrib 8.0.20, for Win64 (x86_64)
--
-- Host: localhost    Database: qadatabase
-- ------------------------------------------------------
-- Server version	8.0.20

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `exam_test`
--

DROP TABLE IF EXISTS `exam_test`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `exam_test` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `date` datetime NOT NULL,
  `user_id` bigint NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `subject` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `number_of_questions` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `FK_UserIdExamTest_idx` (`user_id`),
  CONSTRAINT `FK_UserIdExamTest` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exam_test`
--

LOCK TABLES `exam_test` WRITE;
/*!40000 ALTER TABLE `exam_test` DISABLE KEYS */;
INSERT INTO `exam_test` VALUES (37,'2021-07-15 13:29:18',1,'TestFormat.docx','SWT391',30),(38,'2021-07-15 13:29:44',1,'TestFormat.docx','SWT391',30);
/*!40000 ALTER TABLE `exam_test` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `knowledge`
--

DROP TABLE IF EXISTS `knowledge`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `knowledge` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `date` datetime DEFAULT NULL,
  `user_id` bigint NOT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_UserId_idx` (`user_id`),
  CONSTRAINT `FK_UserId` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `knowledge`
--

LOCK TABLES `knowledge` WRITE;
/*!40000 ALTER TABLE `knowledge` DISABLE KEYS */;
INSERT INTO `knowledge` VALUES (20,'finaldtb.sql','2021-07-15 13:29:59',1),(21,'TestFormat.docx','2021-07-15 13:30:06',1),(22,'phaicoexamtest.sql','2021-07-15 13:30:14',1),(23,'Poster_GaManhHoach.png','2021-07-15 13:30:22',1),(24,'pngtree-gourmet-chicken-leg-cuisine-poster-image_208616.jpg','2021-07-15 13:30:32',1);
/*!40000 ALTER TABLE `knowledge` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `option`
--

DROP TABLE IF EXISTS `option`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `option` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `question_id_id` bigint NOT NULL,
  `key` char(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `content` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `paragraph` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  KEY `FK_QuestionId_idx` (`question_id_id`),
  CONSTRAINT `FK_QuestionId` FOREIGN KEY (`question_id_id`) REFERENCES `question` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=785 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `option`
--

LOCK TABLES `option` WRITE;
/*!40000 ALTER TABLE `option` DISABLE KEYS */;
INSERT INTO `option` VALUES (545,143,'A','i,ii,iii are true and iv is false',''),(546,143,'B','i,,iv are true and ii is false',''),(547,143,'C','i,ii are true and iii,iv are false',''),(548,143,'D','ii,iii,iv are true and i is false',''),(549,144,'A','i,ii,iv are true and iii is false',''),(550,144,'B','i,,iv are true and ii is false',''),(551,144,'C','i,ii are true and iii,iv are false',''),(552,144,'D','ii,iii,iv are true and i is false',''),(553,145,'A','Test Analysis and Design',''),(554,145,'B','Test Planning and control',''),(555,145,'C','Test Implementation and execution',''),(556,145,'D','Evaluating exit criteria and reporting',''),(557,146,'A','CLASS',''),(558,146,'B','cLASS',''),(559,146,'C','CLass',''),(560,146,'D','CLa01ss',''),(561,147,'A','£4800; £14000; £28000',''),(562,147,'B','£5200; £5500; £28000',''),(563,147,'C','£28001; £32000; £35000',''),(564,147,'D','£5800; £28000; £32000',''),(565,148,'A','Designed by persons who write the software under test',''),(566,148,'B','Designed by a person from a different section',''),(567,148,'C','Designed by a person from a different organization',''),(568,148,'D','Designed by another person',''),(569,149,'A','User Acceptance Test Cases',''),(570,149,'B','Integration Level Test Cases',''),(571,149,'C','Unit Level Test Cases',''),(572,149,'D','Program specifications',''),(573,150,'A','Options i,ii,iii,iv are true',''),(574,150,'B','ii is true and i,iii,iv are false',''),(575,150,'C','i,ii,iii are true and iv is false',''),(576,150,'D','iii is true and i,ii,iv are false.',''),(577,151,'A','Component testing',''),(578,151,'B','Non-functional system testing',''),(579,151,'C','User acceptance testing',''),(580,151,'D','Maintenance testing',''),(581,152,'A','Re Testing',''),(582,152,'B','Confirmation Testing',''),(583,152,'C','Regression Testing',''),(584,152,'D','Negative Testing',''),(585,153,'A','How much regression testing should be done',''),(586,153,'B','Exit Criteria',''),(587,153,'C','How many more test cases need to written',''),(588,153,'D','Different Tools to perform Regression Testing',''),(589,154,'A','testing that the system functions with other systems',''),(590,154,'B','testing that the components that comprise the system function together',''),(591,154,'C','testing the end to end functionality of the system as a whole',''),(592,154,'D','testing the system performs functions within specified response times',''),(593,155,'A','Inspection',''),(594,155,'B','Walkthrough',''),(595,155,'C','Technical Review',''),(596,155,'D','Formal Review',''),(597,156,'A','ii is True; i, iii, iv & v are False',''),(598,156,'B','i & v are True; ii, iii & iv are False',''),(599,156,'C','ii & iii are True; i, iv & v are False',''),(600,156,'D','ii, iii & iv are True; i & v are False',''),(601,157,'A','Explaining the objective',''),(602,157,'B','Fixing defects found typically done by author',''),(603,157,'C','Follow up',''),(604,157,'D','Individual Meeting preparations',''),(605,158,'A','i-d , ii-a , iii-c , iv-b',''),(606,158,'B','i-c , ii-d , iii-a , iv-b',''),(607,158,'C','i-b , ii-a , iii-d , iv-c',''),(608,158,'D','i-c , ii-a , iii-d , iv-b',''),(609,159,'A','Test Planning and Control',''),(610,159,'B','Test implementation and Execution',''),(611,159,'C','Requirement Analysis',''),(612,159,'D','Evaluating Exit criteria and reporting',''),(613,160,'A','State transition testing',''),(614,160,'B','LCSAJ (Linear Code Sequence and Jump)',''),(615,160,'C','syntax testing',''),(616,160,'D','boundary value analysis',''),(617,161,'A','ii,iii,iv are correct and i is incorrect',''),(618,161,'B','iii , i , iv is correct and ii is incorrect',''),(619,161,'C','i , iii , iv , ii is in correct',''),(620,161,'D','ii is correct',''),(621,162,'A','i , ii,iii,iv is correct',''),(622,162,'B','iii ,is correct I,,ii,iv are incorrect.',''),(623,162,'C','i ,ii, iii and iv are incorrect',''),(624,162,'D','iv, ii is correct',''),(625,163,'A','Specifications',''),(626,163,'B','Test Cases',''),(627,163,'C','Test Data',''),(628,163,'D','Test Design',''),(629,164,'A','Equivalence partitioning, Decision Table and Control flow are White box Testing Techniques.',''),(630,164,'B','Equivalence partitioning , Boundary Value Analysis , Data Flow are Black Box Testing Techniques',''),(631,164,'C','Equivalence partitioning, State Transition, Use Case Testing are black box Testing Techniques.',''),(632,164,'D','Equivalence Portioning , State Transition , Use Case Testing and Decision Table are White Box Testing\nTechniques',''),(633,165,'A','i & ii are true, iii, iv & v are false',''),(634,165,'B','ii, iii & iv are true, i & v are false',''),(635,165,'C','ii & iv are true, i, iii & v are false',''),(636,165,'D','ii is true, i, iii, iv & v are false',''),(637,166,'A','Independent testers are much more qualified than Developers',''),(638,166,'B','Independent testers see other and different defects and are unbiased.',''),(639,166,'C','Independent Testers cannot identify defects.',''),(640,166,'D','Independent Testers can test better than developers',''),(641,167,'A','Statement coverage is 2, Branch Coverage is 2',''),(642,167,'B','Statement coverage is 3 and branch coverage is 2',''),(643,167,'C','Statement coverage is 1 and branch coverage is 2',''),(644,167,'D','Statement Coverage is 4 and Branch coverage is 2',''),(645,168,'A','Statement coverage is 4',''),(646,168,'B','Statement coverage is 1',''),(647,168,'C','Statement coverage is 3',''),(648,168,'D','Statement Coverage is 2',''),(649,169,'A','Anomaly Report',''),(650,169,'B','Defect Report',''),(651,169,'C','Test Defect Report',''),(652,169,'D','Test Incident Report',''),(653,170,'A','i, ii, iii is true and iv is false',''),(654,170,'B','ii,iii,iv is true and i is false',''),(655,170,'C','i is true and ii,iii,iv are false',''),(656,170,'D','iii and iv is correct and i and ii are incorrect',''),(657,171,'A','i, ii are true and iii and iv are false',''),(658,171,'B','iii is true and i,ii, iv are false',''),(659,171,'C','ii ,iii is true and i,iv is false',''),(660,171,'D','iii and iv are true and i,ii are false',''),(661,172,'A','i , ii , iv are true and iii is false',''),(662,172,'B','i , ii , iii are true and iv is false',''),(663,172,'C','i , iii , iv are true and ii is false',''),(664,172,'D','All of above are true',''),(665,173,'A','i,ii,iii are true and iv is false',''),(666,173,'B','i,,iv are true and ii is false',''),(667,173,'C','i,ii are true and iii,iv are false',''),(668,173,'D','ii,iii,iv are true and i is false',''),(669,174,'A','i,ii,iv are true and iii is false',''),(670,174,'B','i,,iv are true and ii is false',''),(671,174,'C','i,ii are true and iii,iv are false',''),(672,174,'D','ii,iii,iv are true and i is false',''),(673,175,'A','Test Analysis and Design',''),(674,175,'B','Test Planning and control',''),(675,175,'C','Test Implementation and execution',''),(676,175,'D','Evaluating exit criteria and reporting',''),(677,176,'A','CLASS',''),(678,176,'B','cLASS',''),(679,176,'C','CLass',''),(680,176,'D','CLa01ss',''),(681,177,'A','£4800; £14000; £28000',''),(682,177,'B','£5200; £5500; £28000',''),(683,177,'C','£28001; £32000; £35000',''),(684,177,'D','£5800; £28000; £32000',''),(685,178,'A','Designed by persons who write the software under test',''),(686,178,'B','Designed by a person from a different section',''),(687,178,'C','Designed by a person from a different organization',''),(688,178,'D','Designed by another person',''),(689,179,'A','User Acceptance Test Cases',''),(690,179,'B','Integration Level Test Cases',''),(691,179,'C','Unit Level Test Cases',''),(692,179,'D','Program specifications',''),(693,180,'A','Options i,ii,iii,iv are true',''),(694,180,'B','ii is true and i,iii,iv are false',''),(695,180,'C','i,ii,iii are true and iv is false',''),(696,180,'D','iii is true and i,ii,iv are false.',''),(697,181,'A','Component testing',''),(698,181,'B','Non-functional system testing',''),(699,181,'C','User acceptance testing',''),(700,181,'D','Maintenance testing',''),(701,182,'A','Re Testing',''),(702,182,'B','Confirmation Testing',''),(703,182,'C','Regression Testing',''),(704,182,'D','Negative Testing',''),(705,183,'A','How much regression testing should be done',''),(706,183,'B','Exit Criteria',''),(707,183,'C','How many more test cases need to written',''),(708,183,'D','Different Tools to perform Regression Testing',''),(709,184,'A','testing that the system functions with other systems',''),(710,184,'B','testing that the components that comprise the system function together',''),(711,184,'C','testing the end to end functionality of the system as a whole',''),(712,184,'D','testing the system performs functions within specified response times',''),(713,185,'A','Inspection',''),(714,185,'B','Walkthrough',''),(715,185,'C','Technical Review',''),(716,185,'D','Formal Review',''),(717,186,'A','ii is True; i, iii, iv & v are False',''),(718,186,'B','i & v are True; ii, iii & iv are False',''),(719,186,'C','ii & iii are True; i, iv & v are False',''),(720,186,'D','ii, iii & iv are True; i & v are False',''),(721,187,'A','Explaining the objective',''),(722,187,'B','Fixing defects found typically done by author',''),(723,187,'C','Follow up',''),(724,187,'D','Individual Meeting preparations',''),(725,188,'A','i-d , ii-a , iii-c , iv-b',''),(726,188,'B','i-c , ii-d , iii-a , iv-b',''),(727,188,'C','i-b , ii-a , iii-d , iv-c',''),(728,188,'D','i-c , ii-a , iii-d , iv-b',''),(729,189,'A','Test Planning and Control',''),(730,189,'B','Test implementation and Execution',''),(731,189,'C','Requirement Analysis',''),(732,189,'D','Evaluating Exit criteria and reporting',''),(733,190,'A','State transition testing',''),(734,190,'B','LCSAJ (Linear Code Sequence and Jump)',''),(735,190,'C','syntax testing',''),(736,190,'D','boundary value analysis',''),(737,191,'A','ii,iii,iv are correct and i is incorrect',''),(738,191,'B','iii , i , iv is correct and ii is incorrect',''),(739,191,'C','i , iii , iv , ii is in correct',''),(740,191,'D','ii is correct',''),(741,192,'A','i , ii,iii,iv is correct',''),(742,192,'B','iii ,is correct I,,ii,iv are incorrect.',''),(743,192,'C','i ,ii, iii and iv are incorrect',''),(744,192,'D','iv, ii is correct',''),(745,193,'A','Specifications',''),(746,193,'B','Test Cases',''),(747,193,'C','Test Data',''),(748,193,'D','Test Design',''),(749,194,'A','Equivalence partitioning, Decision Table and Control flow are White box Testing Techniques.',''),(750,194,'B','Equivalence partitioning , Boundary Value Analysis , Data Flow are Black Box Testing Techniques',''),(751,194,'C','Equivalence partitioning, State Transition, Use Case Testing are black box Testing Techniques.',''),(752,194,'D','Equivalence Portioning , State Transition , Use Case Testing and Decision Table are White Box Testing\nTechniques',''),(753,195,'A','i & ii are true, iii, iv & v are false',''),(754,195,'B','ii, iii & iv are true, i & v are false',''),(755,195,'C','ii & iv are true, i, iii & v are false',''),(756,195,'D','ii is true, i, iii, iv & v are false',''),(757,196,'A','Independent testers are much more qualified than Developers',''),(758,196,'B','Independent testers see other and different defects and are unbiased.',''),(759,196,'C','Independent Testers cannot identify defects.',''),(760,196,'D','Independent Testers can test better than developers',''),(761,197,'A','Statement coverage is 2, Branch Coverage is 2',''),(762,197,'B','Statement coverage is 3 and branch coverage is 2',''),(763,197,'C','Statement coverage is 1 and branch coverage is 2',''),(764,197,'D','Statement Coverage is 4 and Branch coverage is 2',''),(765,198,'A','Statement coverage is 4',''),(766,198,'B','Statement coverage is 1',''),(767,198,'C','Statement coverage is 3',''),(768,198,'D','Statement Coverage is 2',''),(769,199,'A','Anomaly Report',''),(770,199,'B','Defect Report',''),(771,199,'C','Test Defect Report',''),(772,199,'D','Test Incident Report',''),(773,200,'A','i, ii, iii is true and iv is false',''),(774,200,'B','ii,iii,iv is true and i is false',''),(775,200,'C','i is true and ii,iii,iv are false',''),(776,200,'D','iii and iv is correct and i and ii are incorrect',''),(777,201,'A','i, ii are true and iii and iv are false',''),(778,201,'B','iii is true and i,ii, iv are false',''),(779,201,'C','ii ,iii is true and i,iv is false',''),(780,201,'D','iii and iv are true and i,ii are false',''),(781,202,'A','i , ii , iv are true and iii is false',''),(782,202,'B','i , ii , iii are true and iv is false',''),(783,202,'C','i , iii , iv are true and ii is false',''),(784,202,'D','All of above are true','');
/*!40000 ALTER TABLE `option` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `question`
--

DROP TABLE IF EXISTS `question`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `question` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `content` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `exam_test_id` bigint NOT NULL,
  `number` bigint NOT NULL,
  `mark` double NOT NULL,
  `answer` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_ExamTestId_idx` (`exam_test_id`),
  CONSTRAINT `FK_ExamTestId` FOREIGN KEY (`exam_test_id`) REFERENCES `exam_test` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=203 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `question`
--

LOCK TABLES `question` WRITE;
/*!40000 ALTER TABLE `question` DISABLE KEYS */;
INSERT INTO `question` VALUES (143,'Deciding how much testing is enough should take into account:\ni. Level of Risk including Technical and Business product and project risk\nii. Project constraints such as time and budget\niii. Size of Testing Team\niv. Size of the Development Team\n',37,1,0,'C'),(144,'Test planning has which of the following major tasks?\ni. Determining the scope and risks, and identifying the objectives of testing.\nii. Determining the test approach (techniques, test items, coverage, identifying and interfacing the teams\ninvolved in testing, testware)\niii. Reviewing the Test Basis (such as requirements, architecture, design, interface)\niv. Determining the exit criteria.',37,2,0,'A'),(145,'Evaluating testability of the requirements and system are a part of which phase:',37,3,0,'A'),(146,'One of the fields on a form contains a text box which accepts alphabets in lower or upper case.\nIdentify the invalid Equivalence class value.',37,4,0,'D'),(147,'In a system designed to work out the tax to be paid:\nAn employee has £4000 of salary tax free. The next £1500 is taxed at 10% The next £28000 is\ntaxed at 22% Any further amount is taxed at 40% Which of these groups of numbers would fall\ninto the same equivalence class?',37,5,0,'D'),(148,'Which of the following has highest level of independence in which test cases are',37,6,0,'C'),(149,'We use the output of the requirement analysis, the requirement  specification as the input for writing:',37,7,0,'A'),(150,'Validation involves which of the following\ni. Helps to check the Quality of the Built Product\nii. Helps to check that we have built the right product.\niii. Helps in developing the product\niv. Monitoring tool wastage and obsoleteness.',37,8,0,'B'),(151,'Which of the following uses Impact Analysis most?',37,9,0,'D'),(152,'Repeated Testing of an already tested program, after modification, to discover any defects introduced or uncovered as a result of the changes in the software being tested or in another related or unrelated software component:',37,10,0,'C'),(153,'Impact Analysis helps to decide',37,11,0,'A'),(154,'Functional system testing is:',37,12,0,'C'),(155,'Peer Reviews are also called as:',37,13,0,'C'),(156,'Consider the following statements:\ni. 100% statement coverage guarantees 100% branch coverage.\nii. 100% branch coverage guarantees 100% statement coverage.\niii. 100% branch coverage guarantees 100% decision coverage.\niv. 100% decision coverage guarantees 100% branch coverage.\nv. 100% statement coverage guarantees 100% decision coverage.',37,14,0,'D'),(157,'The Kick Off phase of a formal review includes the ',37,15,0,'A'),(158,'Match every stage of the software Development Life cycle with the Testing Life cycle:\ni. Hi-level design               a Unit tests\nii. Code                              b Acceptance tests\niii. Low-level design          c System tests\niv. Business requirements d Integration tests',37,16,0,'D'),(159,'Which of the following is not phase of the Fundamental Test Process?',37,17,0,'C'),(160,'Which of the following techniques is NOT a black box technique?',37,18,0,'B'),(161,'Success Factors for a review include:\ni. Each Review does not have a predefined objective\nii. Defects found are welcomed and expressed objectively\niii. Management supports a good review process.\niv. There is an emphasis on learning and process improvement.',37,19,0,'A'),(162,'Defects discovered by static analysis tools include:\ni. Variables that are never used.\nii. Security vulnerabilities.\niii. Programming Standard Violations\niv. Uncalled functions and procedures',37,20,0,'A'),(163,'Test Conditions are derived from:',37,21,0,'A'),(164,'Which of the following is true about White and Black Box Testing Technique:',37,22,0,'C'),(165,'Regression testing should be performed:\ni. every week\nii. after the software has changed\niii. as often as possible\niv. when the environment has changed\nv. when the project manager says',37,23,0,'C'),(166,'Benefits of Independent Testing',37,24,0,'B'),(167,'Minimum Tests Required for Statement Coverage and Branch Coverage :\nRead P\nRead Q\nIf p+q > 100 then\nPrint “Large”\nEnd if\nIf p > 50 then\nPrint “pLarge”\nEnd if',37,25,0,'C'),(168,'Minimum Test Required for Statement Coverage:\nDisc = 0\nOrder-qty = 0\nRead Order-qty\nIf Order-qty >=20 then\nDisc = 0.05\nIf Order-qty >=100 then\nDisc =0.1\nEnd if\nEnd if',37,26,0,'B'),(169,'The structure of an incident report is covered in the Standard for Software Test Documentation IEEE 829 and is called as:',37,27,0,'D'),(170,'Which of the following is the task of a Test Lead / Leader\ni. Interaction with the Test Tool Vendor to identify best ways to leverage test tool on the project.\nii. Write Test Summary Reports based on the information gathered during testing\niii. Decide what should be automated, to what degree and how.\niv. Create the Test Specifications',37,28,0,'A'),(171,'Features of White Box Testing Technique:\nWe use explicit knowledge of the internal workings of\nthe item being tested to select the test data.\nUses specific knowledge of programming code to examine outputs and assumes that the tester knows the\npath of logic in a unit or a program.\nChecking for the performance of the application\nAlso checks for functionality',37,29,0,'A'),(172,'Which of the following is a part of Test Closure Activities?\ni. Checking which planned deliverables\nhave been delivered\nii. Defect report analysis.\niii. Finalizing and archiving testware.\niv. Analyzing lessons.',37,30,0,'C'),(173,'Deciding how much testing is enough should take into account:\ni. Level of Risk including Technical and Business product and project risk\nii. Project constraints such as time and budget\niii. Size of Testing Team\niv. Size of the Development Team\n',38,1,0,'C'),(174,'Test planning has which of the following major tasks?\ni. Determining the scope and risks, and identifying the objectives of testing.\nii. Determining the test approach (techniques, test items, coverage, identifying and interfacing the teams\ninvolved in testing, testware)\niii. Reviewing the Test Basis (such as requirements, architecture, design, interface)\niv. Determining the exit criteria.',38,2,0,'A'),(175,'Evaluating testability of the requirements and system are a part of which phase:',38,3,0,'A'),(176,'One of the fields on a form contains a text box which accepts alphabets in lower or upper case.\nIdentify the invalid Equivalence class value.',38,4,0,'D'),(177,'In a system designed to work out the tax to be paid:\nAn employee has £4000 of salary tax free. The next £1500 is taxed at 10% The next £28000 is\ntaxed at 22% Any further amount is taxed at 40% Which of these groups of numbers would fall\ninto the same equivalence class?',38,5,0,'D'),(178,'Which of the following has highest level of independence in which test cases are',38,6,0,'C'),(179,'We use the output of the requirement analysis, the requirement  specification as the input for writing:',38,7,0,'A'),(180,'Validation involves which of the following\ni. Helps to check the Quality of the Built Product\nii. Helps to check that we have built the right product.\niii. Helps in developing the product\niv. Monitoring tool wastage and obsoleteness.',38,8,0,'B'),(181,'Which of the following uses Impact Analysis most?',38,9,0,'D'),(182,'Repeated Testing of an already tested program, after modification, to discover any defects introduced or uncovered as a result of the changes in the software being tested or in another related or unrelated software component:',38,10,0,'C'),(183,'Impact Analysis helps to decide',38,11,0,'A'),(184,'Functional system testing is:',38,12,0,'C'),(185,'Peer Reviews are also called as:',38,13,0,'C'),(186,'Consider the following statements:\ni. 100% statement coverage guarantees 100% branch coverage.\nii. 100% branch coverage guarantees 100% statement coverage.\niii. 100% branch coverage guarantees 100% decision coverage.\niv. 100% decision coverage guarantees 100% branch coverage.\nv. 100% statement coverage guarantees 100% decision coverage.',38,14,0,'D'),(187,'The Kick Off phase of a formal review includes the ',38,15,0,'A'),(188,'Match every stage of the software Development Life cycle with the Testing Life cycle:\ni. Hi-level design               a Unit tests\nii. Code                              b Acceptance tests\niii. Low-level design          c System tests\niv. Business requirements d Integration tests',38,16,0,'D'),(189,'Which of the following is not phase of the Fundamental Test Process?',38,17,0,'C'),(190,'Which of the following techniques is NOT a black box technique?',38,18,0,'B'),(191,'Success Factors for a review include:\ni. Each Review does not have a predefined objective\nii. Defects found are welcomed and expressed objectively\niii. Management supports a good review process.\niv. There is an emphasis on learning and process improvement.',38,19,0,'A'),(192,'Defects discovered by static analysis tools include:\ni. Variables that are never used.\nii. Security vulnerabilities.\niii. Programming Standard Violations\niv. Uncalled functions and procedures',38,20,0,'A'),(193,'Test Conditions are derived from:',38,21,0,'A'),(194,'Which of the following is true about White and Black Box Testing Technique:',38,22,0,'C'),(195,'Regression testing should be performed:\ni. every week\nii. after the software has changed\niii. as often as possible\niv. when the environment has changed\nv. when the project manager says',38,23,0,'C'),(196,'Benefits of Independent Testing',38,24,0,'B'),(197,'Minimum Tests Required for Statement Coverage and Branch Coverage :\nRead P\nRead Q\nIf p+q > 100 then\nPrint “Large”\nEnd if\nIf p > 50 then\nPrint “pLarge”\nEnd if',38,25,0,'C'),(198,'Minimum Test Required for Statement Coverage:\nDisc = 0\nOrder-qty = 0\nRead Order-qty\nIf Order-qty >=20 then\nDisc = 0.05\nIf Order-qty >=100 then\nDisc =0.1\nEnd if\nEnd if',38,26,0,'B'),(199,'The structure of an incident report is covered in the Standard for Software Test Documentation IEEE 829 and is called as:',38,27,0,'D'),(200,'Which of the following is the task of a Test Lead / Leader\ni. Interaction with the Test Tool Vendor to identify best ways to leverage test tool on the project.\nii. Write Test Summary Reports based on the information gathered during testing\niii. Decide what should be automated, to what degree and how.\niv. Create the Test Specifications',38,28,0,'A'),(201,'Features of White Box Testing Technique:\nWe use explicit knowledge of the internal workings of\nthe item being tested to select the test data.\nUses specific knowledge of programming code to examine outputs and assumes that the tester knows the\npath of logic in a unit or a program.\nChecking for the performance of the application\nAlso checks for functionality',38,29,0,'A'),(202,'Which of the following is a part of Test Closure Activities?\ni. Checking which planned deliverables\nhave been delivered\nii. Defect report analysis.\niii. Finalizing and archiving testware.\niv. Analyzing lessons.',38,30,0,'C');
/*!40000 ALTER TABLE `question` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `role` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'khailq','1234','admin'),(2,'binhtb','1234','user'),(3,'thienlh','1234','staff');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-07-15 20:40:20
