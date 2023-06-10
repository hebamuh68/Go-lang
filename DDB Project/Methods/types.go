package Methods

import "database/sql"

var Data_db *sql.DB

type Fasta struct {
	Fasta_id string `json: "fasta_id"`
	Name     string `json: "name"`
	Header   string `json: "header"`
	Sequence string `json: "sequence"`
}

type FastaMetaData struct {
	//Client_url string `json: "client_url"`
	Fasta_id string `json: "fasta_id"`
}

type SpectraMetaData struct {
	From int `json: "from"`
	To   int `json: "to"`
}

type MsSample struct {
	Sample_id         string `json: "sample_id"`
	Name              string `json: "name"`
	Number_of_spectra int    `json: "number_of_spectra"`
}

type SpectrumData struct {
	Spectrum_index int
	Spectrum       MzML
	Sample_id      string
}

type MzML struct {
	To    int      `json: "to"`
	From  int      `json: "from"`
	MZML  []string `json: "mzml"`
	Level int      `json: "level"`
}

type Spectrum struct {
	Index int       `json: "index"`
	MZ    []float32 `json: "mz"`
	I     []float32 `json: "i"`
	RT    float32   `json: "rt"`
	Level int       `json: "level"`
}

// ====================================================== Master
var Data_dir = "../PyData"

type MetaFasta struct {
	Fasta_id string `json: "fasta_id"`
	Name     string `json: "name"`
	Slave    int    `json: "slave"`
}

type Level1 struct {
	Level1_id    int `json: "level1_id"`
	Slave        int `json: "slave"`
	From_spectra int `json: "from_spectra"`
	To_spectra   int `json: "to_spectra"`
}

type Level2 struct {
	Level2_id    int `json: "level2_id"`
	Slave        int `json: "slave"`
	From_spectra int `json: "from_spectra"`
	To_spectra   int `json: "to_spectra"`
}

type RequestBody struct {
	MsSampleID     string `json:"msSampleId"`
	FastaID        string `json:"fastaId"`
	FastaFile      Fasta
	SampleMetaData []Level2 `json:"sampleMetaData"`
	Index          int      `json:"index"`
	Feature        string   `json:"feature"`
}

// ====================================================== Meta data
var Meta_db *sql.DB

var Slave_1 = make(chan []Record)
var Slave_2 = make(chan []Record)
var Slave_3 = make(chan []Record)
var Slave_4 = make(chan []Record)

type SlaveCount struct {
	count int
}

type Record struct {
	Sample          string
	Spectrum_index  int
	Precursor_mass  float64
	Intensity_sum   float64
	Intensity_mean  float64
	RT              float64
	Matched_peptide string
}
