CREATE DATABASE  IF NOT EXISTS `qadatabase` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `qadatabase`;
-- MySQL dump 10.13  Distrib 8.0.25, for Win64 (x86_64)
--
-- Host: localhost    Database: qadatabase
-- ------------------------------------------------------
-- Server version	8.0.25

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
                             `name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                             `subject` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
                             `number_of_questions` bigint NOT NULL DEFAULT '0',
                             PRIMARY KEY (`id`),
                             KEY `FK_UserIdExamTest_idx` (`user_id`),
                             CONSTRAINT `FK_UserIdExamTest` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exam_test`
--

LOCK TABLES `exam_test` WRITE;
/*!40000 ALTER TABLE `exam_test` DISABLE KEYS */;
INSERT INTO `exam_test` VALUES (1,'2021-05-26 00:00:00',1,'nameofuser1','',0),(2,'2021-05-25 00:00:00',2,'nameofuser2','',0),(3,'2021-05-24 00:00:00',3,'nameofuser3','',0),(4,'2021-05-26 00:00:00',4,'nameofuser4','',0),(5,'2021-05-27 00:00:00',5,'nameofuser5','',0),(6,'2021-05-26 00:00:00',5,'nameofuser5','',0),(31,'2021-07-10 16:24:52',2,'oke','',0);
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
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `knowledge`
--

LOCK TABLES `knowledge` WRITE;
/*!40000 ALTER TABLE `knowledge` DISABLE KEYS */;
INSERT INTO `knowledge` VALUES (1,'knowledge1',NULL,1),(2,'knowledge2',NULL,1),(3,'knowledge3',NULL,1),(4,'knowledge4',NULL,2),(5,'knowledge5',NULL,3),(6,'knowledge6',NULL,3),(7,'knowledge7',NULL,4),(8,'knowledge8',NULL,4),(9,'knowledge9',NULL,5),(10,'knowledge10',NULL,5),(11,'AS1.txt','2021-06-06 00:00:00',3),(12,'AS1.txt','2021-06-06 00:00:00',3),(13,'AS1.txt','2021-06-06 00:00:00',3),(14,'AS1.txt','2021-06-06 00:00:00',3),(15,'AS1.txt','2021-06-06 00:00:00',5),(16,'AS1.txt','2021-06-06 00:00:00',5),(17,'AS1.txt','2021-06-25 00:00:00',5),(18,'AS1.txt','2021-06-25 05:53:36',5),(19,'AS1.txt','2021-06-25 15:18:20',2);
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
) ENGINE=InnoDB AUTO_INCREMENT=305 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `option`
--

LOCK TABLES `option` WRITE;
/*!40000 ALTER TABLE `option` DISABLE KEYS */;
INSERT INTO `option` VALUES (1,1,'A','Bernard Arnault & Family','paragraph'),(2,1,'B','Jeff Bezos','paragraph'),(3,1,'C','Elon Musk','paragraph'),(4,1,'D','Bill Gates','paragraph'),(5,2,'A','Qatar','paragraph'),(6,2,'B','Ireland','paragraph'),(7,2,'C','	Singapore','paragraph'),(8,2,'D','Luxembourg','paragraph'),(9,3,'A','194 countries','paragraph'),(10,3,'B','195 countries','paragraph'),(11,3,'C','196 countries','paragraph'),(12,3,'D','200  countries','paragraph'),(13,4,'A','Monitoring progress and test coverage.','paragraph'),(14,4,'B','Measuring and analyzing results.','paragraph'),(15,4,'C','Scheduling test analysis and design tasks','paragraph'),(16,4,'D','Initiating corrective actions.','paragraph'),(17,5,'A','Failure of third party vendor','paragraph'),(18,5,'B','Training issues','paragraph'),(19,5,'C','Problems requirements definition','paragraph'),(20,5,'D','Poor software functionality','paragraph'),(21,6,'A','Customers and users','paragraph'),(22,6,'B','Developers and designers','paragraph'),(23,6,'C','Business and systems analysts','paragraph'),(24,6,'D','System and acceptance testers','paragraph'),(25,7,'A','Developers','paragraph'),(26,7,'B','Analysts','paragraph'),(27,7,'C','Testers','paragraph'),(28,7,'D','Incident Managers','paragraph'),(29,8,'A','Test case specification','paragraph'),(30,8,'B','Test design specification.','paragraph'),(31,8,'C','Test procedure specification.','paragraph'),(32,8,'D','Test results.','paragraph'),(33,9,'A','To enhance the security of the system','paragraph'),(34,9,'B','To prevent the endless loops in code.','paragraph'),(35,9,'C','To swerve as an alternative or \"Plan-B\"','paragraph'),(36,9,'D','To define when to stop testing','paragraph'),(37,10,'A','Ensuring proper environment setup','paragraph'),(38,10,'B','Writing a test summary report','paragraph'),(39,10,'C','Assessing the need for additional tests','paragraph'),(40,10,'D','Finalizing and archiving testware.','paragraph'),(41,11,'A','Testing performed by potential customers at the developers location.','paragraph'),(42,11,'B','Testing performed by potential customers at their own locations.','paragraph'),(43,11,'C','Testing performed by product developers at the customer\'s location.','paragraph'),(44,11,'D','Testing performed by product developers at their own locations.','paragraph'),(45,12,'A','Usability defects found by customers','paragraph'),(46,12,'B','Defects in infrequently used functionality','paragraph'),(47,12,'C','Defects that were detected early','paragraph'),(48,12,'D','Minor defects that were found by users','paragraph'),(49,13,'A','Implementation and execution.','paragraph'),(50,13,'B','Planning and control.','paragraph'),(51,13,'C','Analysis and design.','paragraph'),(52,13,'D','Test closure.','paragraph'),(53,14,'A','During test planning.','paragraph'),(54,14,'B','During test analysis.','paragraph'),(55,14,'C','During test execution.','paragraph'),(56,14,'D','When evaluating exit criteria','paragraph'),(57,15,'A','Damaged reputation','paragraph'),(58,15,'B','Lack of methodology','paragraph'),(59,15,'C','Inadequate training','paragraph'),(60,15,'D','Regulatory compliance','paragraph'),(61,16,'A','It is cheaper than designing tests during the test phases.','paragraph'),(62,16,'B','It helps prevent defects from being introduced into the code.','paragraph'),(63,16,'C','Tests designed early are more effective than tests designed later.','paragraph'),(64,16,'D',' It saves time during the testing phases when testers are busy.','paragraph'),(65,17,'A','To define when a test level is complete','paragraph'),(66,17,'B','To determine when a test has completed.','paragraph'),(67,17,'C','To identify when a software system should be retired','paragraph'),(68,17,'D','To determine whether a test has passed.','paragraph'),(69,18,'A','Data driven testing technique','paragraph'),(70,18,'B','Experience-based technique','paragraph'),(71,18,'C','White-box techniqueD. Structure-based technique','paragraph'),(72,18,'D','Analysis','paragraph'),(73,19,'A','Tool support for performance and monitoring.','paragraph'),(74,19,'B','Tool support for static testing.','paragraph'),(75,19,'C','Tool support for test execution and logging.','paragraph'),(76,19,'D','Tool support for the management of testing and tests.','paragraph'),(77,20,'A','Supporting reviews','paragraph'),(78,20,'B','Validating models of the software.','paragraph'),(79,20,'C','Testing code executed in a special test harness.','paragraph'),(80,20,'D','Enforcement of coding standards.','paragraph'),(81,21,'A','use IT for competitive advantage.','paragraph'),(82,21,'B','buy what customers want from designers.','paragraph'),(83,21,'C','provide products to a niche market.','paragraph'),(84,21,'D','provide services at a lower cost.','paragraph'),(177,58,'A','10 nam','Ngo Ba Kha - Kha Banh , dang choi do bi Cong an bat phat tu 10 nam'),(178,58,'B','9 nam','Ngo Ba Kha - Kha Banh , dang choi do bi Cong an bat phat tu 10 nam'),(179,58,'C','8 nam','Ngo Ba Kha - Kha Banh , dang choi do bi Cong an bat phat tu 10 nam'),(180,58,'D','khong bi di tu','Ngo Ba Kha - Kha Banh , dang choi do bi Cong an bat phat tu 10 nam'),(181,59,'A','20 t','Ngo Ba Kha - Kha Banh , dang choi do bi Cong an bat phat tu 10 nam'),(182,59,'B','21 t','Ngo Ba Kha - Kha Banh , dang choi do bi Cong an bat phat tu 10 nam'),(183,59,'C','22 t','Ngo Ba Kha - Kha Banh , dang choi do bi Cong an bat phat tu 10 nam'),(184,59,'D','23 t','Ngo Ba Kha - Kha Banh , dang choi do bi Cong an bat phat tu 10 nam');
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
  `answer` char(1) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_ExamTestId_idx` (`exam_test_id`),
  CONSTRAINT `FK_ExamTestId` FOREIGN KEY (`exam_test_id`) REFERENCES `exam_test` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=83 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `question`
--

LOCK TABLES `question` WRITE;
/*!40000 ALTER TABLE `question` DISABLE KEYS */;
INSERT INTO `question` VALUES (1,'who is the best richest man in the world ?',1,1,1,'A'),(2,'what is the richest country in the world',1,2,1,'0'),(3,'how many countries in the world',1,3,1,'0'),(4,'Which of the following is a MAJOR task of test planning?',1,4,1,'0'),(5,'Which is a potential product risk factor?',2,1,1,'0'),(6,'Who typically use static analysis tools?',2,2,1,'0'),(7,'Who would USUALLY perform debugging activities',2,3,1,'0'),(8,'Which of the following defines the expected results of a test?',2,4,1,'0'),(9,'In software testing what is the main purpose of exit criteria',3,1,1,'0'),(10,'Which of the following is a KEY test closure task?',3,2,1,'0'),(11,'What is beta testing?',3,3,1,'0'),(12,'Which defects are OFTEN much cheaper to remove?',3,4,1,'0'),(13,'Which activity in the fundamental test process creates test suites for efficient test execution?',4,1,1,'0'),(14,'When should configuration management procedures be implemented?',4,2,1,'0'),(15,'Which of the problems below BEST characterize a result of software failure?',4,3,1,'0'),(16,'What is the MAIN benefit of designing tests early in the life cycle?',4,4,1,'0'),(17,'What is the purpose of exit criteria?',5,1,1,'0'),(18,'Which test design technique relies heavily on prior thorough knowledge of the system?',5,2,1,'0'),(19,'With which of the following categories is a test comparator tool USUALLY associated?',5,3,1,'0'),(20,'For which of the following would a static analysis tool be MOST useful?',5,4,1,'0'),(21,'A key success factor for Net-A-Porter is the ability to',6,1,1,'0'),(58,'Kha Banh di tu bao nhieu nam',31,1,1,'A'),(59,'Kha Banh bao nhieu tuoi',31,2,1,'B');
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
INSERT INTO `user` VALUES (1,'thienlh','1234','user'),(2,'khailq','1234','admin'),(3,'haokx','1234','user'),(4,'binhtb','1234','user'),(5,'minhpa','12345','user');
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

-- Dump completed on 2021-07-14 10:33:33
