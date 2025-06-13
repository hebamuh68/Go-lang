# Distributed Database System for Protein Sequence Analysis

A distributed system built in Go for efficient protein sequence analysis and peptide identification. This system implements a master-slave architecture to handle large-scale biological data processing.

## Project Overview

This system is designed to handle protein sequence (FASTA) files and mass spectrometry data (mzML) for peptide identification. It uses a distributed architecture with one master node and multiple slave nodes to process and analyze biological data efficiently.

### Key Features

- Distributed FASTA file storage and retrieval
- Parallel peptide identification processing
- RESTful API endpoints for data access
- Efficient data routing and load balancing
- CSV output for search results

## System Architecture
![Screenshot from 2023-06-10 16-48-50](https://github.com/hebamuh68/Go-lang/assets/69214737/45d4e3ff-6ca7-4b42-aa5c-de982a349d75)

![Screenshot from 2023-06-10 16-48-58](https://github.com/hebamuh68/Go-lang/assets/69214737/cfa75277-e6df-48e4-b281-bd4db8e1f3da)

The system consists of the following components:

- **Master Node**: Handles client requests, coordinates slave nodes, and manages data distribution
- **Slave Nodes**: Process and store FASTA files and perform peptide identification
- **Client Interface**: REST API for interacting with the system

## Prerequisites

- Go 1.20 or higher
- MySQL database
- Network connectivity between master and slave nodes

## Dependencies

- github.com/gin-gonic/gin v1.9.0
- github.com/go-sql-driver/mysql v1.7.0

## Project Structure

```
.
├── Methods/           # Shared methods and utilities
├── PyData/           # Python data processing scripts
├── SQL/              # Database schemas and queries
├── client1/          # Client implementation
├── master/           # Master node implementation
├── slave1/           # Slave node 1
├── slave2/           # Slave node 2
├── slave3/           # Slave node 3
├── slave4/           # Slave node 4
├── go.mod            # Go module definition
└── go.sum            # Go module checksums
```

## API Endpoints

### Search FASTA
- **Endpoint**: `/`
- **Method**: POST
- **Feature**: "Search_fasta"
- **Parameters**: 
  - FastaID: Identifier for the FASTA file

### Process mzML
- **Endpoint**: `/`
- **Method**: POST
- **Feature**: "Mzml"
- **Parameters**:
  - MsSampleID: Mass spectrometry sample identifier
  - FastaID: FASTA file identifier

## Setup and Installation

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Configure the database connection
4. Start the master node:
   ```bash
   cd master
   go run master.go
   ```
5. Start the slave nodes:
   ```bash
   cd slave1
   go run slave.go
   # Repeat for other slave nodes
   ```

## Usage

1. Send a POST request to the master node (localhost:8080) with the appropriate parameters
2. For FASTA search, the system will:
   - Check if the file exists in the distributed storage
   - Download and store if not found
   - Return the file to the client
3. For mzML processing, the system will:
   - Distribute the processing across slave nodes
   - Combine results
   - Return a CSV file with the search results

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
