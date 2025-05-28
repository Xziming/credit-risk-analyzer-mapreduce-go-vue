package main

import (
    "encoding/csv"
    "os"
    "strconv"
    "gorm.io/gorm"

    "credit/db"
    "credit/models"
)

func ImportFullClientsFromCSV(db *gorm.DB, filePath string) error {
    file, err := os.Open(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return err
    }

    for i, row := range records[1:] {
        var c models.ClientsInfo

        c.ClientID, _ = strconv.Atoi(row[0])
        c.LimitBal, _ = strconv.ParseFloat(row[1], 64)
        c.Sex, _ = strconv.Atoi(row[2])
        c.Education, _ = strconv.Atoi(row[3])
        c.Marriage, _ = strconv.Atoi(row[4])
        c.Age, _ = strconv.Atoi(row[5])

        c.Pay0, _ = strconv.Atoi(row[6])
        c.Pay2, _ = strconv.Atoi(row[7])
        c.Pay3, _ = strconv.Atoi(row[8])
        c.Pay4, _ = strconv.Atoi(row[9])
        c.Pay5, _ = strconv.Atoi(row[10])
        c.Pay6, _ = strconv.Atoi(row[11])

        c.BillAmt1, _ = strconv.ParseFloat(row[12], 64)
        c.BillAmt2, _ = strconv.ParseFloat(row[13], 64)
        c.BillAmt3, _ = strconv.ParseFloat(row[14], 64)
        c.BillAmt4, _ = strconv.ParseFloat(row[15], 64)
        c.BillAmt5, _ = strconv.ParseFloat(row[16], 64)
        c.BillAmt6, _ = strconv.ParseFloat(row[17], 64)

        c.PayAmt1, _ = strconv.ParseFloat(row[18], 64)
        c.PayAmt2, _ = strconv.ParseFloat(row[19], 64)
        c.PayAmt3, _ = strconv.ParseFloat(row[20], 64)
        c.PayAmt4, _ = strconv.ParseFloat(row[21], 64)
        c.PayAmt5, _ = strconv.ParseFloat(row[22], 64)
        c.PayAmt6, _ = strconv.ParseFloat(row[23], 64)

        c.Defaulted, _ = strconv.Atoi(row[24])

        if err := db.Create(&c).Error; err != nil {
            return err
        }

        if i%1000 == 0 {
            println("Imported:", i)
        }
    }

    println("Import finished.")
    return nil
}

func main() {
    db.Init()
    db.DB.AutoMigrate(&models.ClientsInfo{})
    err := ImportFullClientsFromCSV(db.DB, "full_clients_info.csv")
    if err != nil {
        panic(err)
    }
}

