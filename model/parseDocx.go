package model

import "encoding/xml"

type XMLDocument struct {
	XMLName xml.Name `xml:"document"`
	XMLBody XMLBody  `xml:"body"`
}
type XMLBody struct {
	XMLName     xml.Name     `xml:"body"`
	XMLBodyPs   []XMLBodyP   `xml:"p"`
	XMLBodyTbls []XMLBodyTbl `xml:"tbl"`
}

type XMLBodyP struct {
	XMLName   xml.Name    `xml:"p"`
	XMLBodyPr []XMLBodyPr `xml:"r"`
}
type XMLBodyPr struct {
	XMLName xml.Name `xml:"r"`
	Subject string   `xml:"t"`
}

type XMLBodyTbl struct {
	XMLName     xml.Name      `xml:"tbl"`
	XMLBodyTblR []XMLBodyTblR `xml:"tr"`
}
type XMLBodyTblR struct {
	XMLName      xml.Name       `xml:"tr"`
	XMLBodyTblRC []XMLBodyTblRC `xml:"tc"`
}
type XMLBodyTblRC struct {
	XMLName       xml.Name        `xml:"tc"`
	XMLBodyTblRCP []XMLBodyTblRCP `xml:"p"`
}
type XMLBodyTblRCP struct {
	XMLName        xml.Name         `xml:"p"`
	XMLBodyTblRCPR []XMLBodyTblRCPR `xml:"r"`
}
type XMLBodyTblRCPR struct {
	XMLName xml.Name `xml:"r"`
	Br      xml.Name `xml:"br"`
	QN      string   `xml:"t"`
}
