package main

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"io"
	"io/ioutil"

	"golang.org/x/net/html/charset"
)

type ListePDV struct {
	XMLName xml.Name `xml:"pdv_liste"`
	PDVs []PDV `xml:"pdv"`
}


type PDV struct {
	Id string `xml:"id,attr" bson:"id"`
	Latitude float64 `xml:"latitude,attr" bson:"latitude"`
	Longitude float64 `xml:"longitude,attr" bson:"longitude"`
	Cp string `xml:"cp,attr" bson:"cp"`
	Pop string `xml:"pop,attr" bson:"pop"`
	Adresse string `xml:"adresse" bson:"adresse"`
	Ville string `xml:"ville" bson:"ville"`
	Horaires Horaires `xml:"horaires" bson:"horaires"`
	Services []string `xml:"services>service" bson:"services"`
	Prix []Prix `xml:"prix" bson:"prix"`
}

type Horaires struct {
	Automate_24_24 string `xml:"automate_24_24,attr" bson:"automate_24_24"`
	Jours []Jour `xml:"jour" bson:"jours"`
}

type Jour struct {
	Id string `xml:"id,attr" bson:"id"`
	Nom string `xml:"nom,attr" bson:"nom"`
	Ferme string `xml:"ferme,attr" bson:"ferme"`
}


type Prix struct {
	Id string `xml:"id,attr" bson:"id"`
	Nom string `xml:"nom,attr" bson:"nom"`
	Maj string `xml:"maj,attr" bson:"maj"`
	Valeur float64 `xml:"valeur,attr" bson:"valeur"`
}

func Parse_XML(b []byte, xml_data *ListePDV) error {
	r := bytes.NewReader(b)
	d := xml.NewDecoder(r)
	d.CharsetReader = charset.NewReaderLabel
	err := d.Decode(&xml_data)
	if err != nil {
		return err
	}
	return nil
}

func Unzip(read_closer io.ReadCloser, unzipped *[]byte) error {
	body, err := ioutil.ReadAll(read_closer)
	if err != nil {
		return err
	}

	zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		return err
	}
	f, err := zipReader.File[0].Open()
	if err != nil {
		return err
	}
	defer f.Close()

	*unzipped, err = ioutil.ReadAll(f)
	return nil
}