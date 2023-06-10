DROP DATABASE MetaData;

CREATE DATABASE MetaData;

USE MetaData;

CREATE TABLE Fasta(
	Fasta_id VARCHAR(100) PRIMARY KEY,
    Name VARCHAR(100) NOT NULL,
    Slave INT
    );
    
CREATE TABLE MsSample(
	Sample_id VARCHAR(100) PRIMARY KEY,
    Name VARCHAR(100) NOT NULL,
    NumberOfSpectra INT NOT NULL
);

CREATE TABLE Level1(
	Level1_id INT PRIMARY KEY AUTO_INCREMENT,
    Slave INT,
    FromSpectra INT NOT NULL,
    ToSpectra INT NOT NULL,
    Sample_id VARCHAR(100),
    FOREIGN KEY(Sample_id) REFERENCES MsSample(Sample_id) ON DELETE SET NULL ON UPDATE CASCADE
);

CREATE TABLE Level2(
	Level2_id INT PRIMARY KEY AUTO_INCREMENT,
    Slave INT,
    FromSpectra INT NOT NULL,
    ToSpectra INT NOT NULL,
    Sample_id VARCHAR(100),
    FOREIGN KEY(Sample_id) REFERENCES MsSample(Sample_id) ON DELETE SET NULL ON UPDATE CASCADE
);

SELECT * FROM MsSample;
SELECT * FROM Level1;
SELECT * FROM Level2;
SELECT * FROM Fasta;

DELETE FROM MsSample WHERE Sample_id = "MS:1000777";
DELETE FROM Level1 WHERE Sample_id = "MS:1000777";
DELETE FROM Level2 WHERE Sample_id = "MS:1000777";