DROP TABLE IF EXISTS Student;
DROP TABLE IF EXISTS Speciality;
DROP TABLE IF EXISTS Dormitory;

CREATE TABLE Dormitory (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
    );

CREATE TABLE Speciality (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
    );

CREATE TABLE Student (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    dormitoryId INT NOT NULL,
    specialityId INT NOT NULL,
      FOREIGN KEY(dormitoryId) REFERENCES Dormitory ("id"),
      FOREIGN KEY(specialityId) REFERENCES Speciality("id")
    );

INSERT INTO Dormitory (name)VALUES('Dormitory 1');
INSERT INTO Dormitory (name)VALUES('Dormitory 2');
INSERT INTO Dormitory (name)VALUES('Dormitory 3');

INSERT INTO Speciality (name)VALUES('biology');
INSERT INTO Speciality (name)VALUES('computerScience');
INSERT INTO Speciality (name)VALUES('literature');

INSERT INTO Student (name,dormitoryId,specialityId)VALUES('Pasha',1,1);
INSERT INTO Student (name,dormitoryId,specialityId)VALUES('Dasha',1,2);
INSERT INTO Student (name,dormitoryId,specialityId)VALUES('Sasha',2,1);
INSERT INTO Student (name,dormitoryId,specialityId)VALUES('Masha',3,3);
