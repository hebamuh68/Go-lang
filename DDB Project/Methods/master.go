package Methods

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// This function returns the slave url based on its slave id...
func Choose_Slave(slave_id int) (url string) {
	switch slave_id {
	case 1:
		return "http://127.0.0.1:8091/"
	case 2:
		return "http://127.0.0.1:8092/"
	case 3:
		return "http://127.0.0.1:8093/"
	case 4:
		return "http://127.0.0.1:8094/"

	}

	return
}

// This function download ta fasta file from its uniprot id and store the result in PyData dir...
func Get_New_Fasta(fasta_id string) {
	toolDir := "PyData"
	commandStr := fmt.Sprintf("python3 proteomics_pipeline.py --fasta_id %s", fasta_id)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute) // Set a timeout
	defer cancel()

	cmd := exec.CommandContext(ctx, "bash", "-c", commandStr)
	cmd.Dir = toolDir

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing Python script:", err)
		fmt.Println("Error output:", stderr.String())
		panic(err.Error())
	}
}


// This function makes the post request to the slave determined by its url, and the function takes the fasta data as json value
// to send it...
func Fasta_Router(slave_url string, json_value []byte) {
	_, err := http.Post(slave_url+"storeFasta", "application/json", bytes.NewBuffer(json_value))
	if err != nil {
		panic(err.Error())
	}
}

// The function sends fasta file to the slave that has the turn...
func Send_Fasta_To_Slave(slave_id int, id, name, header, sequence string) {
	data_to_slave := map[string]string{
		"fasta_id": id,
		"name":     name,
		"header":   header,
		"sequence": sequence,
	}

	json_value, _ := json.Marshal(&data_to_slave)

	url := Choose_Slave(slave_id)

	Fasta_Router(url, json_value)
}

// This function sends the fasta file requested by the client...
func Send_Fasta(slave_id int, fasta_id string, w http.ResponseWriter) {
	url := Choose_Slave(slave_id)

	fastaIDBytes := []byte(fasta_id)

	// Get the fasta file from slave and send it the client...
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(fastaIDBytes))
	if err != nil {
		panic(err)
	}

	// Connect then send the fasta file to the requesting client...
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "%s", string(body))
}

// The function takes the fasta id in bytes and slave id that stores the file then returns the respond returned from the slave...
func Get_Fasta_From_Slave(slave_id int, fasta_id []byte) *http.Response {
	url := Choose_Slave(slave_id)

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(fasta_id))
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	return resp
}

func Send_Thr_Channel(channel chan []Record, record_obj []Record) {
	channel <- record_obj
}

func Mapper_Router(slave_url string, request_body []byte, w http.ResponseWriter) []Record {
	// send the request to the slave...
	req, err := http.NewRequest("GET", slave_url, bytes.NewBuffer(request_body))
	if err != nil {
		panic(err)
	}

	slaveClient := &http.Client{}
	resp, err := slaveClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Recieve the results from the slave after sending the request...
	len := resp.ContentLength
	body := make([]byte, len)
	resp.Body.Read(body)

	var search_data []Record

	json.Unmarshal(body, &search_data)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return nil
	}

	return search_data
}

func Map_Data_To_Slave(slave_id int, sample_id string, start_index int, fasta_data Fasta, w http.ResponseWriter) {
	// Make the request object that will be send to the slave to start the peptide identification process...
	slaveRequestBody := RequestBody{
		MsSampleID: sample_id,
		FastaFile:  fasta_data,
		Index:      start_index,
	}
	// Convert it to json...
	requestBody, err := json.Marshal(slaveRequestBody)
	if err != nil {
		panic(err)
	}

	switch slave_id {
	case 1:
		url := Choose_Slave(slave_id) + "peptideSearch"
		// Get the data from slave after sending the request to it...
		search_data := Mapper_Router(url, requestBody, w)
		// Send the result into its specified channel
		Send_Thr_Channel(Slave_1, search_data)
	case 2:
		url := Choose_Slave(slave_id) + "peptideSearch"
		// Get the data from slave after sending the request to it...
		search_data := Mapper_Router(url, requestBody, w)
		// Send the result into its specified channel
		Send_Thr_Channel(Slave_2, search_data)
	case 3:
		url := Choose_Slave(slave_id) + "peptideSearch"
		// Get the data from slave after sending the request to it...
		search_data := Mapper_Router(url, requestBody, w)
		// Send the result into its specified channel
		Send_Thr_Channel(Slave_3, search_data)
	case 4:
		url := Choose_Slave(slave_id) + "peptideSearch"
		// Get the data from slave after sending the request to it...
		search_data := Mapper_Router(url, requestBody, w)
		// Send the result into its specified channel
		Send_Thr_Channel(Slave_4, search_data)
	}
}

func Post_Ms_Sample_V1(ms_sample MsSample) {
	json_value, _ := json.Marshal(&ms_sample)

	_, err := http.Post("http://127.0.0.1:8091/storeSample", "application/json", bytes.NewBuffer(json_value))
	if err != nil {
		panic(err.Error())
	}
}

func Post_Ms_Sample_V2(ms_sample MsSample) {
	json_value, _ := json.Marshal(&ms_sample)

	slave_urls := []string{"http://127.0.0.1:8091/storeSample", "http://127.0.0.1:8092/storeSample", "http://127.0.0.1:8093/storeSample", "http://127.0.0.1:8094/storeSample"}
	for _, url := range slave_urls {
		_, err := http.Post(url, "application/json", bytes.NewBuffer(json_value))
		if err != nil {
			panic(err.Error())
		}
	}
}

// The function sends the leve2 data returned from python to the slaves...
func Send_Level_2(sample_id string) {
	lev_twos := Get_Level_2(sample_id)

	files, _ := os.ReadDir(Data_dir)
	
	for _, file := range files{
		if strings.Contains(file.Name(), ".json") && strings.Contains(file.Name(), "level_2"){
			file_path := Data_dir + "/" + file.Name()
			mzml_data := Read_Json(file_path)
			for _, lev_2_obj := range lev_twos{
				for _, mzml_obj := range mzml_data{
					if lev_2_obj.From_spectra == mzml_obj.From{
						json_value, _ := json.Marshal(&mzml_obj)
						slave_url := Choose_Slave(lev_2_obj.Slave)
						slave_url = fmt.Sprintf("%sstoreSpectra/%s", slave_url, sample_id)
						_, err := http.Post(slave_url, "application/json", bytes.NewBuffer(json_value))
						if err != nil {
							panic(err.Error())
						}

					}
				}
				
			}
		}
	}

}

func Post_Spectra_To_Slaves(sample_id string) {

	Send_Level_2(sample_id)

}

func Populate_Spectra_Meta_Data(file_path string) {
	handle, err := os.Open(file_path)

	if err != nil {
		panic(err)
	}
	defer handle.Close()

	content := bufio.NewScanner(handle)

	var ms_sample MsSample
	var level_1 []Level1
	var level_2 []Level2

	content.Scan()
	ms_sample.Name = content.Text()
	content.Scan()
	ms_sample.Sample_id = content.Text()
	content.Scan()
	ms_sample.Number_of_spectra, _ = strconv.Atoi(content.Text())

	for content.Scan() {

		if content.Text() == "Level:1" {
			for i := 4; i > 0; i-- {
				content.Scan()
				ranges := strings.Split(content.Text(), ",")
				var lev_1 Level1
				fmt.Println(ranges)
				lev_1.From_spectra, _ = strconv.Atoi(ranges[0])
				lev_1.To_spectra, _ = strconv.Atoi(ranges[1])
				lev_1.Slave = i
				level_1 = append(level_1, lev_1)
			}

		}

		if content.Text() == "Level:2" {
			for i := 1; i <= 4; i++ {
				content.Scan()
				ranges := strings.Split(content.Text(), ",")
				var lev_2 Level2
				lev_2.From_spectra, _ = strconv.Atoi(ranges[0])
				lev_2.To_spectra, _ = strconv.Atoi(ranges[1])
				lev_2.Slave = i
				level_2 = append(level_2, lev_2)
			}
		}
	}

	ms_sample.Insert_Sample()
	for _, lev1 := range level_1 {
		lev1.Insert_Level_1(ms_sample.Sample_id)
	}

	for _, lev2 := range level_2 {
		lev2.Insert_Level_2(ms_sample.Sample_id)
	}

}

func Send_Inf_TO_Slave(slave_id int, fasta_id string) {
	url := fmt.Sprintf("http://localhost:809%d", slave_id)

	requestData := map[string]string{
		"fasta_id": fasta_id,
	}

	requestDataBytes, err := json.Marshal(requestData)
	if err != nil {
		panic(err)
	}

	_, err = http.Post(url, "application/json", bytes.NewBuffer(requestDataBytes))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Fasta ID %s sent to slave %d\n", fasta_id, slave_id)
}

func Handle_Get_Fasta(fasta_id string) {
	meta_fasta := Get_Fasta_Id(fasta_id)

	if meta_fasta.Fasta_id == "" {
		Get_New_Fasta(fasta_id)
		meta_fasta.Fasta_id = fasta_id
		files, _ := os.ReadDir(Data_dir)
		for _, file := range files {
			if strings.Contains(file.Name(), ".fasta") {
				file_path := Data_dir + "/" + file.Name()
				_, name, header, sequence := Read_Fasta(file_path)
				meta_fasta.Name = name
				meta_fasta.Slave = The_Next_Slave_Fasta()
				meta_fasta.Insert_Fasta_Id()
				Send_Fasta_To_Slave(meta_fasta.Slave, meta_fasta.Fasta_id, name, header, sequence)
				err := os.Remove(file_path)
				if err != nil {
					panic(err.Error())
				}
				return
			}
		}
	} else {
		Send_Inf_TO_Slave(meta_fasta.Slave, meta_fasta.Fasta_id)
	}
}

func Insert_F_filePath(filepath string) {
	id, name, header, seq := Read_Fasta(filepath)

	var obj MetaFasta
	obj.Fasta_id = id
	obj.Name = name
	obj.Slave = The_Next_Slave_Fasta()
	obj.Insert_Fasta_Id()

	var obj2 Fasta
	obj2.Fasta_id = id
	obj2.Name = name
	obj2.Header = header
	obj2.Sequence = seq
	obj2.Insert_Fasta_Data()
}

// This function choose the slave that should store the next fasta file...
func The_Next_Slave_Fasta() int {
	query_statement_1 := "SELECT COUNT(*) FROM MetaData.Fasta WHERE Slave = 1"
	query_statement_2 := "SELECT COUNT(*) FROM MetaData.Fasta WHERE Slave = 2"
	query_statement_3 := "SELECT COUNT(*) FROM MetaData.Fasta WHERE Slave = 3"
	query_statement_4 := "SELECT COUNT(*) FROM MetaData.Fasta WHERE Slave = 4"

	slave_1_count, err := Meta_db.Query(query_statement_1)
	if err != nil {
		fmt.Println("Not Executed...")
		panic(err.Error())
	}

	slave_2_count, err := Meta_db.Query(query_statement_2)
	if err != nil {
		fmt.Println("Not Executed...")
		panic(err.Error())
	}

	slave_3_count, err := Meta_db.Query(query_statement_3)
	if err != nil {
		fmt.Println("Not Executed...")
		panic(err.Error())
	}

	slave_4_count, err := Meta_db.Query(query_statement_4)
	if err != nil {
		fmt.Println("Not Executed...")
		panic(err.Error())
	}

	var results_arr [4]SlaveCount

	for slave_1_count.Next() {
		_ = slave_1_count.Scan(&results_arr[0].count)
	}

	for slave_2_count.Next() {
		_ = slave_2_count.Scan(&results_arr[1].count)
	}

	for slave_3_count.Next() {
		_ = slave_3_count.Scan(&results_arr[2].count)
	}

	for slave_4_count.Next() {
		_ = slave_4_count.Scan(&results_arr[3].count)
	}

	minmum_count := math.Min(math.Min(float64(results_arr[0].count), float64(results_arr[2].count)), math.Min(float64(results_arr[2].count), float64(results_arr[3].count)))
	for i, _ := range results_arr {
		if results_arr[i].count == int(minmum_count) {
			return i + 1
		}
	}
	return 0
}
