drop database if exists banking;

create database banking;

use banking;

create table `Customers` (
	`Id` int auto_increment primary key,
	`Name` varchar(40),
	`City` varchar(40),
	`ZipCode` varchar(10),
	`DateOfBirth` date,
	`Status` char
);

insert into `Customers` (`Name`, `City`, `ZipCode`, `DateOfBirth`, `Status`) 
values ('Keneth', 'Duran', '09883', str_to_date('02/12/2000', '%d/%m/%Y'), '1');

insert into `Customers` (`Name`, `City`, `ZipCode`, `DateOfBirth`, `Status`) 
values ('Camila', 'Guayaquil', '09884', str_to_date('27/12/1999', '%d/%m/%Y'), '1');

insert into `Customers` (`Name`, `City`, `ZipCode`, `DateOfBirth`, `Status`) 
values ('Maylin', 'Naranjal', '09863', str_to_date('12/12/1975', '%d/%m/%Y'), '1');

insert into `Customers` (`Name`, `City`, `ZipCode`, `DateOfBirth`, `Status`) 
values ('Juan', 'Sauces', '09887', str_to_date('01/12/2000', '%d/%m/%Y'), '0');

create table `Accounts` (
	`Id` int auto_increment primary key,
	`CustomerId` int,
	`OpeningDate` varchar(50),
	`Type` varchar(20),
	`Amount` float,
	`Status` char,
	foreign key (`CustomerId`) references Customers(`Id`)
);

insert into `Accounts` (`CustomerId`, `OpeningDate`, `Type`, `Amount`, `Status`)
values (2, '', 'checking', 7560.99, '1');

insert into `Accounts` (`CustomerId`, `OpeningDate`, `Type`, `Amount`, `Status`)
values (3, '', 'saving', 10000.59, '1');