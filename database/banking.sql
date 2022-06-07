drop database if exists banking;

create database banking;

use banking;

create table `Customers` (
	`Id` int auto_increment primary key,
	`Name` varchar(40),
	`City` varchar(40),
	`ZipCode` varchar(10),
	`DateOfBirth` date,
	`Status` tinyint(1)
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
	`Amount` decimal(10, 2),
	`Status` tinyint(1),
	foreign key (`CustomerId`) references Customers(`Id`)
);

insert into `Accounts` (`CustomerId`, `OpeningDate`, `Type`, `Amount`, `Status`)
values (2, '', 'checking', 7560.99, '1');

insert into `Accounts` (`CustomerId`, `OpeningDate`, `Type`, `Amount`, `Status`)
values (3, '', 'saving', 7560.99, '1');

insert into `Accounts` (`CustomerId`, `OpeningDate`, `Type`, `Amount`, `Status`)
values (4, '', 'saving', 10000.59, '1');

create table `Transactions` (
	`Id` int auto_increment primary key,
	`AccountId` int,
	`Amount` float,
	`Type` varchar(20),
	`Date` datetime not null default current_timestamp,
	foreign key (`AccountId`) references Accounts(`Id`)
);

create table `Users` (
	`Username` varchar(25) not null primary key,
	`Password` varchar(25) not null,
	`Role` varchar(25) not null,
	`CustomerId` int default null,
	`CreatedOn` datetime not null default current_timestamp,
	foreign key (`CustomerId`) references Customers(`Id`)
);

insert into `Users` values
  ('root','root','admin', null, '2020-08-09 10:27:22'),
  ('camilucas','root','user', 2, '2020-08-09 10:27:22'),
  ('kenethrrizzo','root','admin', 1, '2020-08-09 10:27:25');
  
 
 select u.Username, u.CustomerId, u.Role, group_concat(a.Id) as AccountNumbers 
		from Users u
		left join Accounts a on a.Id = u.CustomerId
		where u.Username = '2001' and u.Password = 'abc123'
		group by a.CustomerId;