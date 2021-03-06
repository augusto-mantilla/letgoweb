-- MySQL Script generated by MySQL Workbench
-- seg 27 mai 2019 15:52:22 WEST
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';

-- -----------------------------------------------------
-- Schema letmefix
-- -----------------------------------------------------
DROP SCHEMA IF EXISTS `letmefix` ;

-- -----------------------------------------------------
-- Schema letmefix
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `letmefix` DEFAULT CHARACTER SET utf8 ;
USE `letmefix` ;

-- -----------------------------------------------------
-- Table `letmefix`.`address`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `letmefix`.`address` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `zip_code` VARCHAR(45) NOT NULL,
  `address` VARCHAR(255) NOT NULL,
  `latitude` DOUBLE NULL,
  `longitude` DOUBLE NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `letmefix`.`user`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `letmefix`.`user` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `nif` VARCHAR(45) NOT NULL,
  `phone_number` VARCHAR(45) NOT NULL,
  `email` VARCHAR(255) NOT NULL,
  `address_id` INT NOT NULL,
  `foto_url` VARCHAR(255) NULL,
  `admin` TINYINT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_User_Address1_idx` (`address_id` ASC),
  CONSTRAINT `fk_User_Address1`
    FOREIGN KEY (`address_id`)
    REFERENCES `letmefix`.`address` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `letmefix`.`professional`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `letmefix`.`professional` (
  `user_id` INT NOT NULL,
  `prof_area` ENUM('construction', 'carpentry', 'plumbing', 'electricity') NOT NULL,
  `company_name` VARCHAR(45) NOT NULL,
  `skills` VARCHAR(255) NULL,
  PRIMARY KEY (`user_id`),
  CONSTRAINT `fk_table1_User`
    FOREIGN KEY (`user_id`)
    REFERENCES `letmefix`.`user` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `letmefix`.`commercial_client`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `letmefix`.`commercial_client` (
  `user_id` INT NOT NULL,
  `company_name` VARCHAR(45) NOT NULL,
  `commercial_branch` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`user_id`),
  CONSTRAINT `fk_table1_User1`
    FOREIGN KEY (`user_id`)
    REFERENCES `letmefix`.`user` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `letmefix`.`public_request`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `letmefix`.`public_request` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `start` DATETIME NOT NULL,
  `end` DATETIME NOT NULL,
  `prof_area` ENUM('construction', 'carpentry', 'plumbing', 'electricity') NOT NULL,
  `user_id` INT NOT NULL,
  `address_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_public_request_user1_idx` (`user_id` ASC),
  INDEX `fk_public_request_address1_idx` (`address_id` ASC),
  CONSTRAINT `fk_public_request_user1`
    FOREIGN KEY (`user_id`)
    REFERENCES `letmefix`.`user` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_public_request_address1`
    FOREIGN KEY (`address_id`)
    REFERENCES `letmefix`.`address` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `letmefix`.`proposal`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `letmefix`.`proposal` (
  `id` INT NOT NULL,
  `start` DATETIME NOT NULL,
  `end` DATETIME NOT NULL,
  `cost` DOUBLE NOT NULL,
  `cost_of_cancel` DOUBLE NOT NULL,
  `equipment` DOUBLE NULL,
  `public_request_id` INT NOT NULL,
  `professional_user_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_proposal_public_request1_idx` (`public_request_id` ASC),
  INDEX `fk_proposal_professional1_idx` (`professional_user_id` ASC),
  CONSTRAINT `fk_proposal_public_request1`
    FOREIGN KEY (`public_request_id`)
    REFERENCES `letmefix`.`public_request` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_proposal_professional1`
    FOREIGN KEY (`professional_user_id`)
    REFERENCES `letmefix`.`professional` (`user_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `letmefix`.`job`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `letmefix`.`job` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `start` DATETIME NOT NULL,
  `end` DATETIME NOT NULL,
  `address_id` INT NOT NULL,
  `professional_user_id` INT NOT NULL,
  `proposal_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_Job_Address1_idx` (`address_id` ASC),
  INDEX `fk_job_professional1_idx` (`professional_user_id` ASC),
  INDEX `fk_job_proposal1_idx` (`proposal_id` ASC),
  CONSTRAINT `fk_Job_Address1`
    FOREIGN KEY (`address_id`)
    REFERENCES `letmefix`.`address` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_job_professional1`
    FOREIGN KEY (`professional_user_id`)
    REFERENCES `letmefix`.`professional` (`user_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_job_proposal1`
    FOREIGN KEY (`proposal_id`)
    REFERENCES `letmefix`.`proposal` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `letmefix`.`score`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `letmefix`.`score` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `value` VARCHAR(45) NOT NULL,
  `comment` VARCHAR(255) NOT NULL,
  `description` VARCHAR(255) NOT NULL,
  `user_id` INT NOT NULL,
  `job_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_classification_user1_idx` (`user_id` ASC),
  INDEX `fk_score_job1_idx` (`job_id` ASC),
  CONSTRAINT `fk_classification_user1`
    FOREIGN KEY (`user_id`)
    REFERENCES `letmefix`.`user` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_score_job1`
    FOREIGN KEY (`job_id`)
    REFERENCES `letmefix`.`job` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `letmefix`.`execution_request`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `letmefix`.`execution_request` (
  `id` INT NOT NULL,
  `description_exe` VARCHAR(5000) NOT NULL,
  `user_id` INT NOT NULL,
  `address_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_execution_request_user1_idx` (`user_id` ASC),
  INDEX `fk_execution_request_address1_idx` (`address_id` ASC),
  CONSTRAINT `fk_execution_request_user1`
    FOREIGN KEY (`user_id`)
    REFERENCES `letmefix`.`user` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_execution_request_address1`
    FOREIGN KEY (`address_id`)
    REFERENCES `letmefix`.`address` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `letmefix`.`direct_request`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `letmefix`.`direct_request` (
  `id` INT NOT NULL,
  `start` DATETIME NOT NULL,
  `end` DATETIME NOT NULL,
  `user_id` INT NOT NULL,
  `professional_user_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_private_request_user1_idx` (`user_id` ASC),
  INDEX `fk_private_request_professional1_idx` (`professional_user_id` ASC),
  CONSTRAINT `fk_private_request_user1`
    FOREIGN KEY (`user_id`)
    REFERENCES `letmefix`.`user` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_private_request_professional1`
    FOREIGN KEY (`professional_user_id`)
    REFERENCES `letmefix`.`professional` (`user_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `letmefix`.`images`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `letmefix`.`images` (
  `id` INT NOT NULL,
  `image_url` VARCHAR(255) NOT NULL,
  `private_request_id` INT NULL,
  `public_request_id` INT NULL,
  `address_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_images_private_request1_idx` (`private_request_id` ASC),
  INDEX `fk_images_public_request1_idx` (`public_request_id` ASC),
  INDEX `fk_images_address1_idx` (`address_id` ASC),
  CONSTRAINT `fk_images_private_request1`
    FOREIGN KEY (`private_request_id`)
    REFERENCES `letmefix`.`direct_request` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_images_public_request1`
    FOREIGN KEY (`public_request_id`)
    REFERENCES `letmefix`.`public_request` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_images_address1`
    FOREIGN KEY (`address_id`)
    REFERENCES `letmefix`.`address` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `letmefix`.`example_service`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `letmefix`.`example_service` (
  `id` INT NOT NULL,
  `description` VARCHAR(255) NOT NULL,
  `price` DOUBLE NOT NULL,
  `professional_user_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_example_service_professional1_idx` (`professional_user_id` ASC),
  CONSTRAINT `fk_example_service_professional1`
    FOREIGN KEY (`professional_user_id`)
    REFERENCES `letmefix`.`professional` (`user_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `letmefix`.`message`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `letmefix`.`message` (
  `id` INT NOT NULL,
  `date` DATETIME NOT NULL,
  `user_id` INT NOT NULL,
  `professional_user_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_message_user1_idx` (`user_id` ASC),
  INDEX `fk_message_professional1_idx` (`professional_user_id` ASC),
  CONSTRAINT `fk_message_user1`
    FOREIGN KEY (`user_id`)
    REFERENCES `letmefix`.`user` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_message_professional1`
    FOREIGN KEY (`professional_user_id`)
    REFERENCES `letmefix`.`professional` (`user_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
