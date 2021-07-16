CREATE DATABASE  IF NOT EXISTS `parking_lot` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;
USE `parking_lot`;

--
-- Table structure for table `parking_lot`
--

DROP TABLE IF EXISTS `parking_lot`;
CREATE TABLE `parking_lot` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `registration_number` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
  `colour` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
  `is_occupied` boolean DEFAULT FALSE,
  `updated_at` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
