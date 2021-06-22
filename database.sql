CREATE DATABASE  IF NOT EXISTS `question_answer_db` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `question_answer_db`;
-- MySQL dump 10.13  Distrib 8.0.25, for Win64 (x86_64)
--
-- Host: localhost    Database: question_answer_db
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
-- Table structure for table `knowledge`
--

DROP TABLE IF EXISTS `knowledge`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `knowledge` (
                             `id` bigint NOT NULL AUTO_INCREMENT,
                             `name` varchar(255) NOT NULL DEFAULT '',
                             `date` datetime DEFAULT NULL,
                             `user_id` bigint NOT NULL,
                             PRIMARY KEY (`id`),
                             KEY `fk_userId_idx` (`user_id`),
                             CONSTRAINT `FK_UserIdKnowledge` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `knowledge`
--

LOCK TABLES `knowledge` WRITE;
/*!40000 ALTER TABLE `knowledge` DISABLE KEYS */;
INSERT INTO `knowledge` VALUES (1,'knowledge1',NULL,1),(2,'knowledge2',NULL,1),(3,'knowledge3',NULL,1),(4,'knowledge4',NULL,2),(5,'knowledge5',NULL,3),(6,'knowledge6',NULL,3),(7,'knowledge7',NULL,4),(8,'knowledge8',NULL,4),(9,'knowledge9',NULL,5),(10,'knowledge10',NULL,5);
/*!40000 ALTER TABLE `knowledge` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `option`
--

DROP TABLE IF EXISTS `option`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `option` (
                          `question_id` bigint NOT NULL,
                          `id` bigint NOT NULL AUTO_INCREMENT,
                          `key` varchar(10) NOT NULL DEFAULT '',
                          `content` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
                          `paragraph` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
                          PRIMARY KEY (`id`),
                          KEY `FK_QuestionId_idx` (`question_id`),
                          CONSTRAINT `FK_QuestionId` FOREIGN KEY (`question_id`) REFERENCES `question` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `option`
--

LOCK TABLES `option` WRITE;
/*!40000 ALTER TABLE `option` DISABLE KEYS */;
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
                            `content` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
                            `date` datetime NOT NULL,
                            `answer_id` bigint DEFAULT NULL,
                            `user_id` bigint NOT NULL,
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `answer_id_UNIQUE` (`answer_id`),
                            KEY `FK_AnswerId_idx` (`answer_id`),
                            KEY `FK_UserIdQuestion_idx` (`user_id`),
                            CONSTRAINT `FK_AnswerId` FOREIGN KEY (`answer_id`) REFERENCES `option` (`id`),
                            CONSTRAINT `FK_UserIdQuestion` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `question`
--

LOCK TABLES `question` WRITE;
/*!40000 ALTER TABLE `question` DISABLE KEYS */;
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
                        `username` varchar(255) NOT NULL DEFAULT '',
                        `password` varchar(255) NOT NULL DEFAULT '',
                        `role` varchar(255) NOT NULL DEFAULT '',
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'thienlh','1234','user'),(2,'khailq','1234','admin'),(3,'haokx','1234','user'),(4,'binhtb','1234','user'),(5,'minhpa','1234','user');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'question_answer_db'
--

--
-- Dumping routines for database 'question_answer_db'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-06-23  0:10:17
