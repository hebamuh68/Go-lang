DROP DATABASE ProteinData;

CREATE DATABASE ProteinData;

USE ProteinData;

CREATE TABLE Fasta(
	Fasta_id VARCHAR(100) PRIMARY KEY,
    Name VARCHAR(100) NOT NULL,
    Header MEDIUMTEXT NOT NULL,
    Sequence LONGTEXT NOT NULL
    );
    
CREATE TABLE MsSample(
	Sample_id VARCHAR(100) PRIMARY KEY,
    Name VARCHAR(100) NOT NULL,
    NumberOfSpectra INT NOT NULL
    );

CREATE TABLE Spectra(
	SpectrumIndex INT NOT NULL,
    MZML JSON,
    Sample_id VARCHAR(100) NOT NULL,
    FOREIGN KEY(Sample_id) REFERENCES MsSample(Sample_id) ON UPDATE CASCADE,
    PRIMARY KEY(SpectrumIndex, Sample_id)
);

SELECT * FROM MsSample;
SELECT * FROM Spectra;

DELETE FROM MsSample WHERE Sample_id = "MS:1000777";
DELETE FROM Spectra WHERE Sample_id = "MS:1000777";